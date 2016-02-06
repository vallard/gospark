# Dockerfile for Spark webserver
FROM golang
ADD . /go/src/github.com/vallard/gospark/
WORKDIR /go/src/github.com/vallard/gospark
RUN go get
EXPOSE 7999
CMD ["go", "run", "main.go"]
