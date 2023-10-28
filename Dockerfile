FROM --platform=$TARGETPLATFORM golang:alpine3.18 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server -ldflags "-s -w" ./cmd/server/main.go

FROM --platform=$TARGETPLATFORM scratch
COPY --from=builder ./app/server ./server
CMD ["./server"]