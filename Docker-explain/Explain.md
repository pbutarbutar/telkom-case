FROM golang
ADD . /go/src/github.com/telkomdev/indihome/backend
WORKDIR /go/src/github.com/telkomdev/indihome
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
LISTEN 80


#Should changed use go mod like below

FROM golang
ADD . /go/src/github.com/telkomdev/indihome/backend
WORKDIR /go/src/github.com/telkomdev/indihome/backend
COPY go.mod ./
RUN go install .
EXPOSE 8081
CMD [ "/go/bin/backend" ]