FROM golang:1.18-bullseye

WORKDIR /app

COPY go.mod ./

#COPY go.sum ./

RUN go mod download

COPY main.go ./

RUN go build -o main

CMD ["/app/main"] 
