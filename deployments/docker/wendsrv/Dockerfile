FROM golang:1.22.1

WORKDIR /go/src/wendover

RUN mkdir .git
RUN mkdir -p cmd/wendsrv
RUN mkdir -p deployments/migrations
RUN mkdir internal
RUN mkdir pkg
RUN mkdir tools
COPY .git/                      .git/
COPY cmd/wendsrv/               cmd/wendsrv/
COPY deployments/migrations/    deployments/migrations/
COPY internal/                  internal/
COPY pkg/                       pkg/
COPY tools/                     tools/
COPY go.*                       .
COPY Makefile                   .

EXPOSE 8080

RUN make install_server

CMD ["wendsrv", "run"]