FROM golang:1.22.1

WORKDIR /go/src/wendover

RUN mkdir .git
RUN mkdir -p cmd/wendsrv
RUN mkdir internal
RUN mkdir pkg
COPY .git/          .git/
COPY cmd/wendsrv/   cmd/wendsrv/
COPY internal/      internal/
COPY pkg/           pkg/
COPY go.*           .
COPY Makefile       .

RUN make bin/wendsrv

CMD ["bash"]