# OrderKuota - RESTful API dengan Golang dan Vue.js

Proyek yang mengimplementasikan RESTful API menggunakan **Golang** dan **MySQL**, dengan fitur **autentikasi menggunakan Bearer Token**. API ini memungkinkan operasi CRUD (Create, Read, Update, Delete) pada resource **pengguna (user)**.  

Selain itu, proyek ini juga menyediakan **UI dengan Vue.js** untuk mempermudah interaksi dengan API. API ini dikembangkan sebagai bagian dari tugas teknikal test perusahaan Oke Digital (OrderKuota).

## 🎯 Fitur Utama  
✅ **CRUD User** - API dapat membuat, membaca, memperbarui, dan menghapus pengguna.  
✅ **Autentikasi Bearer Token** - Keamanan API menggunakan JWT (JSON Web Token).  
✅ **Format JSON** - API menerima dan mengirim data dalam format JSON.  
✅ **Database MySQL** - Penyimpanan data menggunakan MySQL.  
✅ **Dokumentasi API dengan Postman** - Termasuk contoh request & response positif & negatif.  
✅ **Goroutines** - Optimalisasi performa dengan concurrency di routes.  
✅ **Database Migration** - Menggunakan migrasi database untuk kemudahan setup.  
✅ **Error Handling** - Penanganan error yang jelas dan informatif menggunakan log.  

---

## 🏗️ Struktur Proyek  

```plaintext
orderkuota/
│── backend/            # Kode backend (Golang)
│   ├── adapters/       # Adapter untuk repository & service
│   ├── config/         # Konfigurasi aplikasi
│   ├── domain/         # Model domain
│   ├── dto/            # Data Transfer Objects
│   ├── log/            # Logging system
│   ├── middleware/     # Middleware (Auth, Logging, dll)
│   ├── migrations/     # Database migrations
│   ├── ports/          # Interface untuk dependency injection
│   ├── routes/         # Routing API
│   ├── utils/          # Utility functions
│   ├── .env            # Environment variables
│   ├── go.mod          # Go module dependencies
│   ├── go.sum          # Dependency checksum
│   ├── main.go         # Entry point aplikasi backend
│
│── frontend/           # Kode frontend (Vue.js)
│   ├── node_modules/   # Dependencies Vue.js
│   ├── public/         # Static files
│   ├── src/            # Kode utama frontend
│   ├── package.json    # Dependencies frontend
│   ├── vue.config.js   # Konfigurasi Vue.js
│
│── README.md           # Dokumentasi proyek
│── orderkuota.postman_collection.json # Dokumentasi API di Postman

# ⚙️ Teknologi yang Digunakan

## Backend (Golang)
- Gorilla Mux - Router untuk Golang
- Golang JWT - Autentikasi menggunakan JWT
- MySQL Driver - Koneksi ke database MySQL
- Viper - Manajemen konfigurasi
- Zerolog - Logging system
- Godotenv - Load konfigurasi dari file .env
- UUID Generator - Untuk ID unik pengguna

## Frontend (Vue.js)
- Vue 3 - Framework frontend
- Vue Router - Manajemen routing
- Vuex - Manajemen state
- Axios - HTTP client untuk komunikasi dengan backend
- Bootstrap 5 - Styling UI

# 🚀 Cara Menjalankan Proyek

## 1️⃣ Setup Backend

### 🛠 Prasyarat
- Install Golang (versi 1.23.4 atau terbaru)
- Install MySQL dan buat database orderkuota
- Buat file `.env` di folder backend/ berdasarkan contoh berikut:

```plaintext
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=password
DB_NAME=orderkuota

JWT_SECRET=mysecretkey
```

### 🚀 Jalankan Backend
```sh
cd backend
go mod tidy
go run main.go
```
API akan berjalan di [http://localhost:8080](http://localhost:8080)

## 2️⃣ Setup Frontend

### 🛠 Prasyarat
- Install Node.js (versi terbaru)
- Install Vue CLI (jika belum ada)

```sh
npm install -g @vue/cli
```

### 🚀 Jalankan Frontend
```sh
cd frontend
npm install
npm run serve
```
UI akan berjalan di [http://localhost:8081](http://localhost:8081)

# 📌 Endpoint API

## 🔐 Authentication

| Method | Endpoint             | Deskripsi                             |
|--------|----------------------|---------------------------------------|
| POST   | /api/auth/login      | Login user dan mendapatkan token      |
| POST   | /api/auth/register   | Registrasi user baru                  |

## 👤 User Management

| Method | Endpoint             | Deskripsi                             |
|--------|----------------------|---------------------------------------|
| GET    | /api/users           | Mendapatkan semua user                |
| GET    | /api/users/:id      | Mendapatkan user berdasarkan ID       |
| POST   | /api/users           | Menambahkan user baru                  |
| PUT    | /api/users/:id      | Memperbarui data user                 |
| DELETE | /api/users/:id      | Menghapus user berdasarkan ID          |

## 🔍 Dokumentasi API
Postman Collection tersedia di file `orderkuota.postman_collection.json`. Import ke Postman untuk mencoba semua API yang tersedia.

# 🌟 Additional Point
- ✅ Deployment di GitHub → Repo OrderKuota
- ✅ Frontend & Backend sudah berjalan sepenuhnya
- ✅ Penggunaan Goroutines untuk performa lebih baik
- ✅ Database Migration sudah diterapkan
- ✅ Error Handling sudah optimal

# 🎯 Catatan
Jika ada kendala saat menjalankan proyek, pastikan:
- MySQL sudah berjalan dan kredensial di `.env` benar.
- Golang dan Node.js sudah terinstall dengan versi terbaru.
- Dependencies backend dan frontend sudah diinstall (go mod tidy & npm install).

Terima kasih! 🚀🔥