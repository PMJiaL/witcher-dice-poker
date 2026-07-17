FROM golang:1.26-alpine AS build-stage
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main

FROM busybox:1.38.0-musl AS utilities

FROM scratch AS production-stage
WORKDIR /app
COPY --from=build-stage /build/main .
COPY --from=utilities /bin/busybox /bin/busybox
ENTRYPOINT ["./main"]
EXPOSE 2007
