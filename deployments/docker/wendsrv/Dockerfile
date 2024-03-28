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
COPY internal/                  internal/
COPY pkg/                       pkg/
COPY tools/                     tools/
COPY go.*                       .
COPY Makefile                   .

EXPOSE 8080

RUN make lambda

FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /go/src/wendover/bin/wendsrv-migrate-lambda ./wendsrv-migrate-lambda

ENV WENDOVER_CONFIG_DIRECTORY=DOZFAC
ENV WENDOVER_AWS_REGION=us-west-2

ENTRYPOINT ["./wendsrv-migrate-lambda"]