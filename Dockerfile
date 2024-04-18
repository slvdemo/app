FROM alpine AS build-env
ARG TARGETARCH
WORKDIR /build
COPY ./*.zip .
RUN unzip $TARGETARCH.zip

FROM cgr.dev/chainguard/static:latest
COPY --from=build-env /build/app /app
ENTRYPOINT ["/app"]