FROM alpine:latest

COPY sooty-tern /

COPY configs /configs

ARG SOOTY_TERN_ENV

ENV sooty_tern_env $SOOTY_TERN_ENV

WORKDIR /

EXPOSE 8080

ENTRYPOINT ["/sooty-tern"]
