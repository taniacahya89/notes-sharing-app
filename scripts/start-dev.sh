#!/bin/bash

# Development startup script for Notes Sharing App

echo "🚀 Starting Notes Sharing App (Development Mode)"
echo "================================================"

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker first."
    exit 1
fi

# Start PostgreSQL
echo ""
echo "📦 Starting PostgreSQL..."
docker run -d \
  --name notes-postgres-dev \
  -e POSTGRES_USER=notesuser \
  -e POSTGRES_PASSWORD=notespass \
  -e POSTGRES_DB=notes_db \
  -p 5432:5432 \
  postgres:15-alpine

# Wait for PostgreSQL
echo "⏳ Waiting for PostgreSQL to be ready..."
sleep 5

# Start Backend
echo ""
echo "🔧 Starting Backend (Golang)..."
cd ../backend
export DATABASE_URL="postgres://notesuser:notespass@localhost:5432/notes_db?sslmode=disable"
export JWT_SECRET="mysupersecretkey123"
export PORT="8080"
go run cmd/main.go &
BACKEND_PID=$!

# Wait for backend
sleep 3

# Start Frontend
echo ""
echo "🎨 Starting Frontend (Next.js)..."
cd ../frontend
export NEXT_PUBLIC_API_URL="http://localhost:8080"
npm run dev &
FRONTEND_PID=$!

echo ""
echo "✅ All services started!"
echo ""
echo "📍 Access points:"
echo "   Frontend: http://localhost:3000"
echo "   Backend:  http://localhost:8080"
echo "   Database: localhost:5432"
echo ""
echo "Press Ctrl+C to stop all services"

# Wait for Ctrl+C
trap "echo ''; echo '🛑 Stopping services...'; kill $BACKEND_PID $FRONTEND_PID; docker stop notes-postgres-dev; docker rm notes-postgres-dev; echo '✅ All services stopped'; exit" INT
wait