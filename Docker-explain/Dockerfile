FROM golang

ADD . /go/src/github.com/pbutarbutar/panganesia.id

WORKDIR /go/src/github.com/pbutarbutar/panganesia.id

COPY go.mod ./

RUN go install .

EXPOSE 8081
CMD [ "/go/bin/panganesia.id" ]