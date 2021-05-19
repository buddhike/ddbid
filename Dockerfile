FROM golang:1.16.3-alpine3.13

RUN mkdir -p /home/ddbid
WORKDIR /home/ddbid
COPY ./*.go /home/ddbid/
COPY ./go.mod /home/ddbid/

ENV  GOPROXY=https://goproxy.io,direct

RUN go get && \
    go build -o build/ddbid .

FROM alpine:3.13
COPY --from=0 /home/ddbid/build/ddbid /usr/local/bin/

EXPOSE 8001

ENTRYPOINT ["ddbid"]
