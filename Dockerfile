FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY data/*.go ./data/


RUN go build -o /planetor-reborn

EXPOSE 8080

CMD [ "/planetor-reborn" ]