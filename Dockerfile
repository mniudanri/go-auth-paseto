FROM golang:1.20

RUN mkdir build

ADD . /build

WORKDIR /build

RUN go build -o go-auth-paseto .
EXPOSE 8080

CMD ["/build/go-auth-paseto"]