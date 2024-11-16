FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o ci-cd ./main.go

FROM scratch

COPY --from=builder /build/ci-cd /

ENTRYPOINT ["/ci-cd"]