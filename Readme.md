# 🛒 E-Commerce API - Go Backend Documentation

Dokumentasi ini menjelaskan struktur folder dan arsitektur proyek yang menggunakan pola **Clean Architecture** dengan pendekatan **Modular Monolith**.

---

## 🚀 1. Command (cmd)
Titik masuk (*entry point*) utama aplikasi untuk proses kompilasi.

* **`cmd/server/main.go`**: Jantung dari aplikasi. Bertugas melakukan *bootstrapping*, memuat konfigurasi, inisialisasi koneksi database, merakit *Dependency Injection* dari setiap modul, dan menjalankan server HTTP (Echo).

---

## 🧠 2. Internal
Area terproteksi yang berisi logika bisnis inti. Kode di sini tidak dapat diimpor oleh paket di luar proyek ini.

* **`internal/middlewares/`**: Lapisan pencegat (*interceptor*) HTTP untuk menangani Autentikasi (JWT) dan Otorisasi (Admin Check).
* **`internal/models/`**: Definisi skema database (GORM Structs) dan relasi antar tabel (Products, Orders, Users, dll).
* **`internal/modules/`**: Implementasi fitur bisnis yang dibagi berdasarkan domain:
    * **`auth/`**: Registrasi, Login, dan OAuth (Google/Github).
    * **`cart/`**: Logika keranjang belanja.
    * **`inventory/`**: Manajemen stok barang.
    * **`order/`**: Alur transaksi dan histori pesanan.
    * **`payment/`**: Integrasi payment gateway dan callback.
    * **`products/`**: Katalog, kategori, atribut, dan varian produk.
    * **`promotion/`**: Sistem diskon dan promo.
    * **`review/`**: Rating dan ulasan pelanggan.
    * **`shops/`**: Manajemen profil toko/seller.

> **Struktur Standar Modul:**
> * `delivery/`: DTO (Data Transfer Object) untuk Request/Response.
> * `handler/`: Controller HTTP (Echo).
> * `services/`: Pusat logika bisnis (*Business Logic*).
> * `repository/`: Abstraksi akses database (Query GORM).
> * `provider/`: Registrasi modul & Dependency Injection.

---

## 📦 3. Package (pkg)
Pustaka bersama (*shared library*) dan utilitas infrastruktur yang bersifat agnostik terhadap logika bisnis.

* **`pkg/bootstrapper/`**: Inisialisasi server HTTP dan *graceful shutdown*.
* **`pkg/common/`**: Standarisasi konstanta (Enum), penanganan error kustom, dan format response API.
* **`pkg/config/`**: Manajemen variabel lingkungan (`.env`) ke dalam struct Go.
* **`pkg/database/`**: Pengaturan koneksi DB, Auto-Migration, dan **Transaction Manager (tx.go)**.
* **`pkg/logger/`**: Sistem pencatatan aktivitas dan audit sistem.
* **`pkg/middlewares/`**: Middleware global seperti CORS, Header security, dan Logger request.
* **`pkg/shared/`**: Helper pihak ketiga seperti Supabase Storage client dan utilitas upload file.

---

## 🛠️ Tech Stack
* **Language:** Go (Golang) v1.26
* **Framework:** Echo v5
* **ORM:** GORM (PostgreSQL)
* **Validation:** Go-Playground Validator v10
* **Authentication:** JWT (JSON Web Token)
* **Storage:** Supabase Storage

---