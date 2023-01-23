FROM golang:1.16-alpine

WORKDIR /controller-mock

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./controller-mock

EXPOSE 8080

CMD [ "./controller-mock" ]