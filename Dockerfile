FROM golang:alpine
WORKDIR /app
COPY . .

RUN go get && go build -o bin .
ENTRYPOINT ["/app/bin"]