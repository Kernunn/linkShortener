FROM golang:1.16-alpine

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