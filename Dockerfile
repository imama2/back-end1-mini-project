# Gunakan image golang:alpine sebagai base image
FROM golang:latest

EXPOSE 8080:8080

# Set working directory di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum ke dalam container
COPY go.mod go.sum ./

# Download dependensi Go
RUN go mod download

# Salin kode sumber aplikasi ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o myapp

# Menjalankan aplikasi saat container dijalankan
ENTRYPOINT go run .
#CMD ["./myapp"]
