FROM alpine:latest AS builder
RUN apk add --no-cache curl tar

ENV CRANE_VERSION=v0.21.7

RUN curl -L https://github.com/google/go-containerregistry/releases/download/${CRANE_VERSION}/go-containerregistry_Linux_x86_64.tar.gz -o crane.tar.gz && \
    tar -zxvf crane.tar.gz crane && \
    chmod +x crane


FROM alpine:latest
RUN apk add --no-cache bash


COPY --from=builder /crane /usr/local/bin/crane


COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
