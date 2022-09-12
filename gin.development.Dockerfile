FROM golang:1.19-alpine

RUN apk update && apk upgrade # && apk add --no-cache bash git openssh

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

# RUN go install -mod=mod github.com/cosmtrek/air

# RUN go install -mod=mod github.com/codegangsta/gin

ENTRYPOINT CompileDaemon -polling -log-prefix=false --build="go build -tags development main.go" --command=./main

# COPY . .

# CMD ["air"]

#ENTRYPOINT gin run main.go