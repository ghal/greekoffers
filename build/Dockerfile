FROM golang:1.13-alpine as builder
WORKDIR /src
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ARG version=dev
ENV GO111MODULE=on

RUN go build -mod=readonly -a -ldflags "-X main.version=$version" -o /bin/greekoffers ./cmd/greekoffers/main.go

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/greekoffers .
COPY --from=builder /src/build/entrypoint.sh .

ENTRYPOINT ["./entrypoint.sh"]
