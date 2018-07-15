FROM golang:1.10
RUN go get -d github.com/Houndie/dates
WORKDIR /go/src/github.com/Houndie/dates
RUN ./build.sh
CMD ./dates
