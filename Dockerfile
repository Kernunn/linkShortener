FROM golang:1.16-alpine

### Install postgresql ###
RUN apk update && apk add postgresql
RUN mkdir /run/postgresql
RUN chown postgres:postgres /run/postgresql/

USER postgres
RUN mkdir /var/lib/postgresql/data
RUN chmod 0700 /var/lib/postgresql/data

RUN initdb -D /var/lib/postgresql/data
RUN echo "host all all 0.0.0.0/0 md5" >> /var/lib/postgresql/data/pg_hba.conf
RUN echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf

RUN pg_ctl start -D /var/lib/postgresql/data

### Build app ###
USER root
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN mkdir proto
COPY proto/* ./proto/

RUN mkdir shortener
COPY shortener/* ./shortener/

COPY main.go ./

RUN go build -o /linkShortener

EXPOSE 8080

CMD ["/linkShortener"]
