FROM golang:latest
WORKDIR /go/src/github.com/nirmalvp/amadeusgo/
RUN go get -d -v github.com/parnurzeal/gorequest
COPY api/ api/
RUN go install ./api
CMD [/bin/bash]
