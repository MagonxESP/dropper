FROM golang:1.19-alpine as build

WORKDIR /build

RUN apk add --no-cache make

COPY . .

RUN make build

FROM golang:1.19-alpine

WORKDIR /app
ENV GIN_MODE=release
ENV DROPPER_HTTP_PORT=8080

COPY --from=build /build/build/dropper dropper

EXPOSE $DROPPER_HTTP_PORT

CMD ["./dropper"]