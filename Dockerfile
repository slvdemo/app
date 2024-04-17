FROM alpine:latest
ARG TARGETARCH
WORKDIR /ws
RUN apk update && apk add curl
COPY ./dist/app_linux_${TARGETARCH}*/app .
USER 65532:65532
ENTRYPOINT ["/ws/app"]