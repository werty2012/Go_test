FROM golang:1.19

WORKDIR /app

COPY main.go .
COPY server.go .
#COPY lib_pq.go .
#RUN go mod init app
#RUN go build
EXPOSE 8000

CMD ["go", "run", "main.go", "server.go"]