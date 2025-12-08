FROM golang:1.25
LABEL authors="dalinarkholin"


WORKDIR /app

COPY . .

RUN go build -o app

ENTRYPOINT ["./app"]