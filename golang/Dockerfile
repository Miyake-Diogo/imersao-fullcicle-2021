FROM golang:1.17


WORKDIR /go/src/

RUN apt-get update && apt-get install build-essential librdkafka-dev -y

# gambs pro container ficar executando
CMD ["tail", "-f", "/dev/null"]