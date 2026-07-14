# Registry Sync Action

Registry Sync 是一款专为 GitHub Actions 设计的容器镜像同步工具。基于 Go 语言原生开发，无需依赖 Docker 环境，提供高效、轻量级的多架构镜像同步能力。

## 核心特性

- **高性能执行**：采用纯 Go 静态编译二进制文件，消除对 Docker Daemon 的依赖，大幅提升启动与执行效率。
- **灵活的输入源**：支持单镜像地址、文本清单文件以及 Workflow 内联配置三种数据源接入方式。
- **多架构支持**：原生支持指定目标架构（如 `linux/amd64`, `linux/arm64`）的镜像拉取与推送。
- **可视化报告**：任务结束后，自动于 Workflow 的 Job Summary 面板生成格式化的同步状态报告。

## 输入参数 (Inputs)

> **注意**：目标仓库（`dst`）为必填项。此外，必须且仅需提供一种源镜像参数（`src`、`images`、`file` 其中之一）。

| 参数名 | 必填 | 默认值 | 描述 |
| :--- | :---: | :--- | :--- |
| **`dst`** | **是** | - | **目标仓库 (Destination)**：镜像同步的目标 Registry 地址。 |
| **`src`** | 否* | - | **单体模式**：单个源镜像地址（例如：`nginx:alpine`）。 |
| **`images`** | 否* | - | **内联模式**：在 Workflow 中直接声明的多行源镜像列表。 |
| **`file`** | 否* | - | **清单模式**：代码库中包含源镜像列表的文件路径（例如：`./images.txt`）。 |
| **`platform`**| 否 | `linux/amd64` | **目标架构**：指定拉取与同步的镜像系统架构。 |
| **`concurrency`**| 否 | `4` | **并发线程数**：执行镜像同步的并发数量，建议 4~6 个。 |
| **`retries`** | 否 | `3` | **重试次数**：网络请求失败时的最大重试次数。 |

*\*注：`src`、`images`、`file` 在 YAML 层面配置为非必填（required: false），但在实际执行逻辑中，必须提供其中之一作为同步数据源。*


## 使用范例

### 1. 单镜像同步
适用于仅需同步单一镜像的场景。

```yaml
steps:
  - name: Sync single image
    uses: yangtudou/registry-sync@main
    with:
      src: nginx:1.24-alpine
      dst: 目标仓库
```

### 2. 内联多镜像同步
适用于镜像数量较少，便于直接在 Workflow 文件中声明和维护的场景。

```yaml
steps:
  - name: Sync inline images
    uses: yangtudou/registry-sync@main
    with:
      dst: 目标仓库
      concurrency: '6'
      images: |
        ubuntu:22.04
        redis:7.0-alpine
        golang:1.21
```

### 3. 基于文件清单的批量同步
适用于大规模镜像同步，或镜像清单由其他上游流程自动生成的场景。

```yaml
steps:
  - name: Sync images from file
    uses: yangtudou/registry-sync@main
    with:
      file: 来源文件
      dst: 目标仓库
```
