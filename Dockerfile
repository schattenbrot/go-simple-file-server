#
# Build go app
#
FROM golang:1.22.1 AS builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o api ./cmd/api

# 
# Run app
# 
FROM golang:1.22.1

WORKDIR /run

ENV CORS false

COPY --from=builder /go/src/api ./

EXPOSE 8080

ENTRYPOINT [ "./api" ]