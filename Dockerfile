FROM golang:1.16.3-alpine

WORKDIR /go/src/qt
COPY . .
RUN go build

CMD ["./quincy-term"]