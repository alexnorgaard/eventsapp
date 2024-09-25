FROM golang:1.23

WORKDIR /api

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY --chmod=755 . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /echo-docker

CMD ["/echo-docker"]
