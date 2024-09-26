FROM golang:1.23

WORKDIR /api

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /echo-docker ./cmd/api/main.go
RUN chmod 755 /echo-docker

EXPOSE 8080
EXPOSE 5432

CMD ["/echo-docker"]
