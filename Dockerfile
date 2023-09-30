FROM golang:alpine
WORKDIR /qr-application
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o qrCodeApp
CMD ["./qrCodeApp"]
