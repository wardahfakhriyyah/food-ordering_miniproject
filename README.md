# Food Ordering Application
Ini adalah proyek aplikasi pemesanan makanan (Food Ordering Application) menggunakan bahasa pemrograman Golang dan RESTful API.

## Deskripsi
Aplikasi ini memungkinkan pengguna untuk melakukan pemesanan makanan melalui sebuah RESTful API. Pengguna dapat memilih makanan yang tersedia pada menu, menentukan jumlah makanan yang ingin dipesan, serta melakukan pembayaran melalui sistem pembayaran online yang tersedia pada aplikasi.

## Teknologi yang Digunakan
- Golang
- RESTful API
- MongoDB
- Docker

## Cara Menjalankan Aplikasi
1. Pertama-tama, pastikan Anda sudah menginstal Golang, MongoDB, dan Docker di komputer Anda.
2. Clone repositori ini ke komputer Anda.
3. Buka terminal dan navigasi ke direktori aplikasi ini.
4. Jalankan perintah `go mod tidy` untuk mengunduh dan menginstal dependensi proyek yang diperlukan.
5. Buat file konfigurasi untuk mengatur koneksi ke database MySQL dan menyimpannya dengan nama `.env` pada direktori utama proyek.
   Contoh isi file `.env`:

   ```
   DB_HOST=localhost
   DB_PORT=27017
   DB_USER=username
   DB_PASSWORD=password
   DB_NAME=food_ordering
   ```
   
   Pastikan informasi koneksi yang Anda masukkan sesuai dengan konfigurasi MySQL di komputer Anda.
   
6. Jalankan perintah `go run main.go` pada terminal untuk menjalankan aplikasi.
7. Buka web browser dan masukkan alamat `http://localhost:8000` untuk mengakses API dan melakukan pemesanan makanan.


## Docker
Anda juga dapat menjalankan aplikasi ini menggunakan Docker. Untuk melakukan itu, ikuti langkah-langkah berikut:

1. Pastikan Docker sudah terinstal di komputer Anda.
2. Buka terminal dan navigasi ke direktori aplikasi ini.
3. Jalankan perintah `docker build -t food_ordering .` untuk membangun image Docker.
4. Setelah image berhasil dibangun, jalankan perintah `docker run -p 8080:8000 food_ordering` untuk menjalankan aplikasi.
5. Buka web browser dan masukkan alamat `http://localhost:8080` untuk mengakses API dan melakukan pemesanan makanan.

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan lakukan fork repositori ini dan buat pull request dengan perubahan yang Anda ajukan. 

## Kontak

wardahfakhrdn@gmail.com
