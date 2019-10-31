FROM alpine:latest

COPY sooty-tern /

COPY configs /configs

ENV sooty_tern_env dev

WORKDIR /

EXPOSE 8080

ENTRYPOINT ["/sooty-tern"]
