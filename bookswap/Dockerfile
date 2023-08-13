FROM golang:1.21.0-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . .
RUN go mod download
RUN go build -o myapp ./cmd/main.go
EXPOSE ${BOOKSWAP_PORT}
CMD [ "./myapp" ]