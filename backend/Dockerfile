FROM golang:1.22-alpine
WORKDIR /go/src/github.com/poonchub/sa-67-example
COPY backend .
RUN apk add build-base
RUN go mod tidy
RUN go build -o main .
EXPOSE 8000
CMD ["./main"]