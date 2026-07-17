FROM golang:1.23-alpine AS build-stage
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main

FROM scratch AS production-stage
WORKDIR /app
COPY --from=build-stage /build/main .
ENTRYPOINT ["./main"]
EXPOSE 2007