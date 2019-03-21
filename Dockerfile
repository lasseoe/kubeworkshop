FROM golang
WORKDIR /go/src/kubeworkshop
ADD . /go/src/kubeworkshop
RUN go install kubeworkshop
ENTRYPOINT /go/bin/kubeworkshop
