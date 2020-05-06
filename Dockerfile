FROM golang:1.14.1

RUN mkdir -p /go/src/github.com/dinopuguh/kawulo-webservice/

WORKDIR /go/src/github.com/dinopuguh/kawulo-webservice/

COPY . .

RUN go build -o webservice main.go

EXPOSE 9090

CMD /go/src/github.com/dinopuguh/kawulo-webservice/webservice