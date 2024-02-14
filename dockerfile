FROM golang:1.21.6
WORKDIR /app
COPY CryptScrap ./CryptScrap
COPY StockScrap ./StockScrap
COPY router ./router
COPY middleware ./middleware
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY . /app
ENV RUNNING_IN_DOCKER true
RUN go build -o main
EXPOSE 8080
CMD ["./main"]
