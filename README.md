# 📝 Notes Sharing App

Aplikasi berbagi catatan dengan fitur autentikasi JWT, CRUD notes, dan upload gambar.

## 🏗️ Struktur Proyek

```
notes-sharing-app/
├── backend/
│   ├── cmd/
│   │   └── main.go
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── database.go
│   ├── handlers/
│   │   ├── auth.go
│   │   └── notes.go
│   ├── middleware/
│   │   └── auth.go
│   ├── models/
│   │   ├── user.go
│   │   └── note.go
│   ├── routes/
│   │   └── routes.go
│   ├── utils/
│   │   ├── jwt.go
│   │   └── logger.go
│   ├── uploads/
│   ├── logs/
│   ├── .env
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
├── frontend/
│   ├── app/
│   │   ├── login/
│   │   ├── register/
│   │   ├── notes/
│   │   ├── layout.tsx
│   │   └── page.tsx
│   ├── components/
│   ├── lib/
│   │   └── api.ts
│   ├── .env.local
│   ├── package.json
│   ├── tsconfig.json
│   ├── next.config.js
│   └── Dockerfile
├── docker/
│   └── docker-compose.yml
└── README.md
```

## 🚀 Cara Menjalankan Proyek

### Prerequisites
- Docker & Docker Compose terinstall
- Node.js 18+ (untuk development)
- Go 1.21+ (untuk development)

### Langkah-langkah:

1. **Clone atau extract proyek ini**
   ```bash
   cd notes-sharing-app
   ```

2. **Setup Environment Variables**
   
   Backend (`backend/.env`):
   ```env
   DATABASE_URL=postgres://notesuser:notespass@postgres:5432/notes_db?sslmode=disable
   JWT_SECRET=mysupersecretkey123
   PORT=8080
   ```

   Frontend (`frontend/.env.local`):
   ```env
   NEXT_PUBLIC_API_URL=http://localhost:8080
   ```

3. **Jalankan dengan Docker Compose**
   ```bash
   cd docker
   docker-compose up --build
   ```

   Tunggu hingga semua service berjalan:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - PostgreSQL: localhost:5432

4. **Akses Aplikasi**
   - Buka browser ke `http://localhost:3000`
   - Register akun baru
   - Login dan mulai membuat notes!

## 📌 API Endpoints

### Authentication
- `POST /api/auth/register` - Registrasi user baru
- `POST /api/auth/login` - Login user

### Notes (Requires JWT Token)
- `GET /api/notes` - Ambil semua notes milik user
- `POST /api/notes` - Buat note baru
- `GET /api/notes/:id` - Ambil note by ID
- `PUT /api/notes/:id` - Update note
- `DELETE /api/notes/:id` - Hapus note
- `POST /api/notes/:id/upload` - Upload gambar untuk note

## 🛠️ Development

### Menjalankan Backend Saja
```bash
cd backend
go mod download
go run cmd/main.go
```

### Menjalankan Frontend Saja
```bash
cd frontend
npm install
npm run dev
```

### Menjalankan PostgreSQL Saja
```bash
docker run -d \
  --name postgres-notes \
  -e POSTGRES_USER=notesuser \
  -e POSTGRES_PASSWORD=notespass \
  -e POSTGRES_DB=notes_db \
  -p 5432:5432 \
  postgres:15-alpine
```

## 🧪 Testing

### Test API dengan cURL

**Register:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"password123"}'
```

**Login:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}'
```

**Get Notes (gunakan token dari login):**
```bash
curl -X GET http://localhost:8080/api/notes \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 📂 Fitur Utama

- ✅ User Authentication dengan JWT
- ✅ CRUD Operations untuk Notes
- ✅ Upload Gambar untuk Notes
- ✅ Logging System (backend/logs/app.log)
- ✅ Dockerized (Frontend, Backend, PostgreSQL)
- ✅ Responsive UI dengan Next.js
- ✅ Secure password hashing (bcrypt)
- ✅ Protected routes dengan middleware

## 🔐 Security

- Password di-hash menggunakan bcrypt
- JWT token untuk autentikasi
- Middleware untuk protect routes
- CORS enabled untuk frontend

## 📝 Database Schema

### Users Table
```sql
id          SERIAL PRIMARY KEY
name        VARCHAR(255)
email       VARCHAR(255) UNIQUE
password    VARCHAR(255)
created_at  TIMESTAMP
```

### Notes Table
```sql
id          SERIAL PRIMARY KEY
user_id     INTEGER (FK to users)
title       VARCHAR(255)
content     TEXT
image_url   VARCHAR(500)
created_at  TIMESTAMP
updated_at  TIMESTAMP
```

## 🐛 Troubleshooting

**Port sudah digunakan:**
```bash
# Stop container yang berjalan
docker-compose down

# Atau ubah port di docker-compose.yml
```

**Database connection error:**
- Pastikan PostgreSQL container sudah running
- Check environment variables
- Tunggu beberapa detik untuk database initialization

**Frontend tidak bisa connect ke backend:**
- Pastikan `NEXT_PUBLIC_API_URL` benar
- Check CORS settings di backend

**Go dependencies error:**
```bash
cd backend
go mod tidy
go mod download
```

**Next.js build error:**
```bash
cd frontend
rm -rf .next node_modules
npm install
npm run build
```

## 📚 Tech Stack

- **Backend**: Go 1.21, Fiber v2, GORM, JWT
- **Frontend**: Next.js 14, TypeScript, TailwindCSS
- **Database**: PostgreSQL 15
- **DevOps**: Docker, Docker Compose

📸 Screenshots

![Register Page](screenshots/register.png)
![Login Page](screenshots/login.png)
![Notes Page](screenshots/notes.png)
![Add Note Page](screenshots/add-note.png)
![Delete Note Page](screenshots/delete-note.png)
![Add Image Page](screenshots/add-image.png)

