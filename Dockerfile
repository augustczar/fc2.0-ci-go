FROM golang:1.21.1

WORKDIR /app

RUN go mod init main

COPY . .

RUN go build -o math

CMD [ "./math" ]