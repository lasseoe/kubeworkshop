FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o kubeworkshop .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/kubeworkshop /app/
WORKDIR /app
CMD ["/app/kubeworkshop"]
