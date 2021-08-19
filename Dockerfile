FROM golang:1.16-alpine

### Install postgresql ###
RUN apk update && apk add postgresql
RUN mkdir /run/postgresql
RUN chown postgres:postgres /run/postgresql/

USER postgres
RUN mkdir /var/lib/postgresql/data
RUN chmod 0700 /var/lib/postgresql/data

RUN initdb -D /var/lib/postgresql/data

### Build app ###ps
USER root
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

ADD cmd ./cmd/
ADD internal ./internal/
ADD proto ./proto/
RUN go build -o /linkShortener cmd/main.go

ADD start.sh /
RUN chmod +x /start.sh

USER postgres

EXPOSE 8080

CMD ["/start.sh"]
