FROM golang:latest as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app

#--------------
FROM alpine:latest
RUN apk --no-cache add ca-certificates
#ENV MONGODB_URI="mongodb+srv://<username>:<password>@<cluster>/?retryWrites=true&w=majority"
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]