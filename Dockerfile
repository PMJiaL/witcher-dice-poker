FROM golang:1.23-alpine
COPY . .
RUN go build
CMD ["./witcher-dice-poker"]
EXPOSE 2007
