# Rest-API Checkout

## Langkah-langkah Pengambilan Proyek

1. **Clone Repository:**
    - Buka terminal.
    - Gunakan perintah `git clone` untuk mengunduh repository:

      ```bash
      git clone https://github.com/refaldyrk/checkout-api.git
      ```

2. **Pindah ke Direktori Proyek:**
    - Masuk ke direktori proyek:

      ```bash
      cd checkout-api
      ```

3. **Konfigurasi Database:**
    - Pastikan konfigurasi database sesuai dengan kebutuhan Anda.
    - Edit file konfigurasi atau lingkungan seperti `.env` dengan kredensial dan pengaturan database yang benar.

4. **Install Dependencies:**
    - Pastikan Anda memiliki Go dan Go Modules terinstal.
    - Jalankan perintah untuk mengunduh dan menginstal dependencies:

      ```bash
      go mod tidy
      ```

5. **Jalankan Aplikasi:**
    - Jalankan aplikasi Golang Anda:

      ```bash
      go run .
      ```

6. **Akses Dokumentasi API Swagger:**
    - Setelah aplikasi berjalan, akses dokumentasi API Swagger pada URL berikut: [http://localhost:port/swagger/index.html](http://localhost:port/swagger/index.html)

## Uji Coba API di Swagger

1. Pastikan API berhasil diuji menggunakan Swagger untuk setiap operasi CRUD.
2. Berikan hasil uji coba yang mencakup response code dan body.

---

## Swagger
go to `/swagger/index.html` route path to see API documentation via Swagger.

## What i can do?
- Customer can view product list by product category
- Customer can add product to shopping cart
- Customers can see a list of products that have been added to the shopping cart
- Customer can delete product list in shopping cart
- Customers can checkout and make payment transactions
- Login and register customers

## Pakai Docker? 
```bash
docker compose up -d
```
