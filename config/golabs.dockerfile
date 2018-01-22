FROM golang
MAINTAINER Yuri Souza
RUN mkdir -p "$GOPATH/src/github.com/yurisouza/labs-go-with-docker"
COPY . $GOPATH/src/github.com/yurisouza/labs-go-with-docker
WORKDIR $GOPATH/src/github.com/yurisouza/labs-go-with-docker
RUN go install
ENTRYPOINT $GOPATH/bin/labs-go-with-docker
EXPOSE 3000