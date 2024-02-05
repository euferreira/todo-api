FROM golang:1.19

WORKDIR /go/src
ENV path="/go/bin:${PATH}"

RUN apt-get update

CMD ["tail", "-f", "/dev/null"]