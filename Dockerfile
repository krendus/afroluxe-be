FROM golang:1.18-alpine
WORKDIR /afroluxe-be
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v .
CMD ["./afroluxe-be"]