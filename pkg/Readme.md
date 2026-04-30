# 📦 Package (pkg) Directory Documentation

Folder `pkg` berisi pustaka bersama (*shared library*), utilitas inti, dan konfigurasi infrastruktur yang bersifat agnostik terhadap logika bisnis. Semua modul di dalam folder `internal` bergantung pada paket-paket di sini.

---

### 📂 `bootstrapper/`
**Fungsi Utama:** Inisialisasi & Lifecycle Server.
* Mengelola proses *startup* aplikasi dan server HTTP (Echo).
* Mengatur konfigurasi dasar server seperti *timeout*, *port*, dan *graceful shutdown*.

### 📂 `common/`
**Fungsi Utama:** Standarisasi Data & Kontrak API.
* **Constants:** Lokasi pusat untuk `enum` (Status Order, Role User, dll) agar konsisten di seluruh aplikasi.
* **Customs:** Penanganan *error* kustom untuk validasi struct dan pemetaan pesan *error* yang ramah bagi Frontend.
* **Response:** Template JSON standar untuk *Success*, *Error*, dan *Pagination* agar format API seragam.

### 📂 `config/`
**Fungsi Utama:** Manajemen Konfigurasi.
* Membaca dan memetakan variabel lingkungan (`.env` atau *Environment Variables*) ke dalam struct Go.
* Menyediakan konfigurasi terpusat (DB URL, JWT Secret, Key pihak ketiga) untuk di-inject ke modul lain.

### 📂 `database/`
**Fungsi Utama:** Abstraksi & Integritas Data.
* Mengelola koneksi *database pool* menggunakan GORM.
* Menangani **Auto-Migration** untuk sinkronisasi skema database dengan model Go.
* Menyediakan **Transaction Manager (`tx.go`)** untuk memastikan operasi database yang kompleks bersifat *atomic* (semua sukses atau semua batal).

### 📂 `logger/`
**Fungsi Utama:** Pencatatan Aktivitas Sistem.
* Menyediakan utilitas logging terpusat untuk memantau perilaku aplikasi (*Behavior Tracking*).
* Membantu proses *debugging* dan audit di lingkungan produksi (Production).

### 📂 `middlewares/`
**Fungsi Utama:** Keamanan & Interceptor Request.
* Berisi filter fungsional seperti kebijakan **CORS**, validasi keamanan **Header**, pembatasan **Host**, serta logging otomatis untuk setiap HTTP Request yang masuk.

### 📂 `shared/`
**Fungsi Utama:** Integrasi Pihak Ketiga & Helper Umum.
* Berisi *client* untuk layanan eksternal seperti **Supabase Storage** (Cloud Bucket).
* Menyediakan fungsi pembantu (*Helper*) untuk pemrosesan file, seperti filter tipe file dan batasan ukuran *upload*.

---