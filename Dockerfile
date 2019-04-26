# build binary
FROM golang:1.12 AS build
WORKDIR /build/blackbox
COPY . /build/blackbox
RUN CGO_ENABLED=0 go build \
    -o /out/blackbox \
    github.com/wenkaler/blackbox/cmd/blackbox

# copy to alpine image
FROM alpine:3.8
WORKDIR /app
COPY --from=build /out/blackbox /app
RUN apk add --no-cache tzdata
ENV TZ Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
CMD ["/app/blackbox"]
