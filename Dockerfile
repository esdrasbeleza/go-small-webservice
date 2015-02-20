FROM golang:1.4-onbuild
ADD . /go/src/github.com/esdrasbeleza/go-small-webservice
RUN go install github.com/esdrasbeleza/go-small-webservice
#ENTRYPOINT /go/bin/go-small-webservice
EXPOSE 8000

