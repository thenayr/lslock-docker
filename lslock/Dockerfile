FROM alpine:3.1
RUN mkdir -p /go
RUN apk add --update \
            go \
            git \
            && export GOPATH=/go
ENV GOPATH /go
WORKDIR /go
RUN mkdir -p /go/src/github.com/thenayr/lslock
RUN go get -u github.com/theckman/go-flock
RUN go get -u github.com/Pallinder/go-randomdata

#RUN go get github.com/franela/goreq
ADD . /go/src/github.com/thenayr/lslock
RUN go install github.com/thenayr/lslock
#RUN go build -o /bin/docker-viz 
ENTRYPOINT ["/go/bin/lslock"]
CMD ["-d", "/tmp/lslock-test"]
