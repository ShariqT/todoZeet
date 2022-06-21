FROM golang:1.15

WORKDIR /app

COPY . .

RUN go mod download

RUN go build

EXPOSE 3000

ENV DATABASE=localhost

CMD ./app