FROM golang:1.17

WORKDIR /app

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download

COPY . .

RUN go get .

CMD [ "go", "run", "." ]

EXPOSE 3000