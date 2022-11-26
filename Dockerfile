FROM golang:1.19

WORKDIR /app

COPY main.go .
COPY server.go .
RUN go mod init appdocker
RUN cd /app & go build
EXPOSE 8000

CMD ["/app/appdocker"]