FROM golang:1-stretch

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install

VOLUME /go/src/app
VOLUME /go/bin

CMD ["go-wrapper", "run"]