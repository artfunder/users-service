FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git gcc

WORKDIR /app

COPY . .

RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server

# Final Image
FROM scratch

WORKDIR /bin

COPY --from=builder /app/server .

CMD [ "/bin/server" ]