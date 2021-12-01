FROM golang:1.17.2-buster AS go_build

WORKDIR /artsign

COPY go.mod go.sum /artsign/
RUN go mod download

COPY . /artsign
RUN make

FROM debian:buster-slim

WORKDIR /artsign
RUN apt update && apt install -y ca-certificates
COPY ./entry-point.sh ./entry-point.sh
COPY --from=go_build /artsign/bin/ ./bin/

RUN chmod 755 ./entry-point.sh
ENTRYPOINT [ "./entry-point.sh" ]
