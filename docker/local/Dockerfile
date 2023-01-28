FROM golang:1.18-alpine AS builder

WORKDIR /metaplayer-chain

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o metaplayer-chain main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /metaplayer-chain

COPY --from=builder /metaplayer-chain/metaplayer-chain ./
COPY ./docker/local/metaplayer-chain.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478 5001

ENTRYPOINT ["./metaplayer-chain.sh"]
