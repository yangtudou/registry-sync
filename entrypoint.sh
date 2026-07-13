#!/usr/bin/env bash

# 严格模式：遇到错误立即停止
set -Eeuo pipefail

# --- 映射 GitHub Actions 输入参数 ---
# 注意：GHA 会自动将 inputs 中的变量转换为大写并添加 INPUT_ 前缀
readonly SRC="${INPUT_SRC}"
readonly DST="${INPUT_DST}"
readonly MAX_CONCURRENT="${INPUT_MAX_CONCURRENT:-4}"
readonly PLATFORM="${INPUT_PLATFORM:-}"

# 构建 platform 参数
if [[ -n "$PLATFORM" ]]; then
    readonly PLATFORM_FLAG="--platform $PLATFORM"
else
    readonly PLATFORM_FLAG=""
fi

# --- 批量同步逻辑函数 ---
run_batch_sync() {
    local list_file=$1
    # 提取所有镜像，去重并过滤空行/注释
    mapfile -t IMAGES < <(grep -Ev '^[[:space:]]*(#|$)' "$list_file" | awk '!seen[$0]++')
    
    local total=${#IMAGES[@]}
    echo "🚀 开始批量同步任务，总数: $total"

    for ((i=0; i<total; i++)); do
        local img="${IMAGES[i]}"
        local task_num=$((i + 1))
        
        # 核心逻辑：保留命名空间
        # 去掉协议前缀 (例如 daemon://)，拼接目标
        local clean_name="${img#*://}"
        local target="${DST}/${clean_name}"

        (
            echo "[$task_num/$total] 正在同步: $img -> $target"
            # 使用 eval 或直接展开变量，确保 platform 参数正确生效
            # 这里将 PLATFORM_FLAG 作为字符串展开，因为它是可选的
            if crane copy ${PLATFORM_FLAG} "$img" "$target"; then
                echo "[$task_num/$total] ✅ 成功: $img"
            else
                echo "[$task_num/$total] ❌ 失败: $img"
                exit 1
            fi
        ) &

        # 并发限制
        while [[ $(jobs -rp | wc -l) -ge "$MAX_CONCURRENT" ]]; do
            wait -n
        done
    done
    wait
    echo "🎉 批量同步任务全部完成。"
}

# --- 逻辑分支 ---

# 1. 模式 A: 文件路径 (存在该文件则按文件处理)
if [[ -f "$SRC" ]]; then
    echo "检测到文件模式: $SRC"
    run_batch_sync "$SRC"

# 2. 模式 B: 多行文本列表 (包含换行符)
elif [[ "$SRC" == *$'\n'* ]]; then
    echo "检测到多行文本模式"
    tmp_list=$(mktemp)
    trap 'rm -f "$tmp_list"' EXIT
    printf "%s\n" "$SRC" > "$tmp_list"
    run_batch_sync "$tmp_list"

# 3. 模式 C: 单体同步
else
    echo "检测到单体模式: $SRC -> $DST"
    crane copy ${PLATFORM_FLAG} "$SRC" "$DST"
fi
