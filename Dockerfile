FROM golang:1.22.3 AS build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y make

RUN go build main.go

FROM ubuntu:22.04 AS runner

WORKDIR /app

RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*

COPY --from=build /build/main ./main

CMD ["/app/main" ]
