FROM golang:1.24-bookworm AS build
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -trimpath -o /forged ./cmd/forged
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends ffmpeg ca-certificates && rm -rf /var/lib/apt/lists/* && useradd -r -u 10001 forge
COPY --from=build /forged /usr/local/bin/forged
RUN mkdir /data && chown forge /data
USER forge
ENV FORGE_ROOT=/data FORGE_ADDR=0.0.0.0:8734
EXPOSE 8734
ENTRYPOINT ["/usr/local/bin/forged"]
