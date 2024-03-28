FROM golang:1.22.1 as build

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

ENV WENDOVER_CONFIG_DIRECTORY=DOZFAC

EXPOSE 8080

RUN make bin/wendsrv

CMD ["wendsrv", "run"]

FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /go/src/wendover/bin/wendsrv ./wendsrv
COPY --from=build /go/src/wendover/deployments/migrations ./migrations

CMD ["wendsrv", "run"]