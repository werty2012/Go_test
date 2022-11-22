FROM golang:1.19

WORKDIR /app

COPY main.go .
COPY server.go .

EXPOSE 8000

CMD ["go", "run", "main.go", "server.go"]