# 🧠 Internal Directory Documentation

Folder `internal` adalah area terproteksi yang berisi logika bisnis inti aplikasi. Kode di sini tidak dapat diimpor oleh paket di luar proyek ini, menjaga integritas dan keamanan sistem.

---

### 📂 `middlewares/`
**Definisi:** Lapisan pencegat (*interceptor*) permintaan HTTP.
* Berisi logika autentikasi (JWT), otorisasi (Role Admin/User), dan validasi akses sebelum permintaan mencapai handler utama.

### 📂 `models/`
**Definisi:** Representasi struktur data dan skema database.
* Berisi struct GORM yang mendefinisikan tabel, relasi antar tabel (Has-Many, Belongs-To), dan anotasi database (Unique, Index, Foreign Key).

### 📂 `modules/`
**Definisi:** Implementasi fitur bisnis yang terfragmentasi secara modular. Setiap sub-folder di dalamnya mewakili satu domain bisnis (Contoh: `auth`, `order`, `product`) dengan struktur internal sebagai berikut:

* **`delivery/`**: Berisi DTO (*Data Transfer Object*) untuk standarisasi format Request dan Response API.
* **`handler/`**: Layer entri poin HTTP (Echo) yang bertugas melakukan binding data dan mengirimkan HTTP Response.
* **`services/`**: Layer pusat logika bisnis (*Business Logic*). Tempat di mana aturan main aplikasi dijalankan.
* **`repository/`**: Layer abstraksi database. Berisi query (SQL/GORM) untuk interaksi langsung dengan penyimpanan data.
* **`provider/`**: Layer registrasi dan *Dependency Injection*. Tempat menyatukan Repository, Service, dan Handler serta mendaftarkan rute API.
* **`utils/`**: Utilitas spesifik yang hanya digunakan di dalam modul tersebut (misal: Helper JWT di modul Auth).

---

## 🧭 Daftar Modul Bisnis

1.  **`auth/`**: Manajemen kredensial, registrasi, login, dan integrasi OAuth (Google/Github).
2.  **`cart/`**: Logika pengelolaan keranjang belanja user.
3.  **`inventory/`**: Manajemen stok dan ketersediaan barang.
4.  **`order/`**: Pengelolaan alur transaksi, status pesanan, dan histori belanja.
5.  **`payment/`**: Integrasi gerbang pembayaran (*payment gateway*) dan penanganan callback.
6.  **`products/`**: Katalog produk, kategori, atribut, serta manajemen varian produk.
7.  **`promotion/`**: Sistem diskon, kode promo, dan kampanye pemasaran.
8.  **`review/`**: Sistem ulasan, rating pelanggan, dan social proof.
9.  **`shops/`**: Manajemen profil toko dan entitas penjual (*seller*).