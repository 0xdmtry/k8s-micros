FROM golang:1.19 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o server

FROM debian:buster-slim
RUN apt-get update -y && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /server

CMD ["/server"]