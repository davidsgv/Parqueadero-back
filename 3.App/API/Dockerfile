FROM golang:1.20.2-bullseye

WORKDIR /usr/src/app

COPY . .
#RUN go mod download && go mod verify

#RUN go build -v -o /usr/local/bin/app ./...

CMD ["go", "run", "main.go"]