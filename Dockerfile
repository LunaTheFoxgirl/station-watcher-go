FROM golang:1-stretch

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]