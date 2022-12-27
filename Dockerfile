FROM golang:1.19-alpine as build

WORKDIR /build

RUN apk add --no-cache make

COPY . .

RUN make build

FROM golang:1.19-alpine

WORKDIR /app
ENV GIN_MODE=release

COPY --from=build /build/build/dropper dropper

CMD ["./dropper"]