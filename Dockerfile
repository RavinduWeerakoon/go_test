FROM golang:alpine

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o main hello.go
RUN adduser -S -D -H -h /app appuser
USER appuser

CMD ["./main"]