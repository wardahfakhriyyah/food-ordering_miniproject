# Image yang digunakan
FROM golang:1.16

# Mengatur direktori kerja
WORKDIR /app

# Menyalin file go.mod dan go.sum ke dalam container
COPY go.mod .
COPY go.sum .

# Menyalin seluruh file dari direktori saat ini ke direktori /app di dalam container
COPY . .

# Menjalankan perintah untuk mengunduh dan menginstal dependensi yang diperlukan
RUN go mod download

# Build binary file
RUN go build -o main .

EXPOSE 8000

# Menjalankan binary file saat container dijalankan
CMD ["./main"]
