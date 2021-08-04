FROM golang:1.16 AS builder

RUN mkdir -p /build

WORKDIR /build
ADD . .
RUN cd core/cmd && go build

FROM ubuntu:18.04 AS test
WORKDIR /server
RUN apt update && apt install -y git ca-certificates && update-ca-certificates
COPY --from=builder /build/core/cmd/cmd .

CMD [ "/server/cmd" ]