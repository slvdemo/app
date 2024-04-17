FROM golang AS build-env
WORKDIR /build
COPY . /build
RUN go build -a -tags 'osusergo netgo static_build' -ldflags '-w -extldflags "-static"' -o app

FROM alpine:latest
RUN apk update && apk add curl
WORKDIR /ws
COPY --from=build-env /build/app /ws/
ENTRYPOINT ["/ws/app"]