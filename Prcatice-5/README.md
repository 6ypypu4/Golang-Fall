#  Инструкция по запуску Practice 4 

## 1️ Запуск базы данных PostgreSQL
```bash (из корня проекта)
docker-compose up -d
docker ps
```

---

## 2️ Применение миграций
```bash
migrate -path ./internal/db/migrations -database "postgres://postgres:postgres@localhost:5432/usersdb?sslmode=disable" up
docker exec -it go_postgres psql -U postgres -d usersdb -c "\dt"
```

---

## 3️ Добавление пользователей
```bash
docker exec -it go_postgres psql -U postgres -d usersdb -c "INSERT INTO users (name, email, balance) VALUES ('Ilya', 'gegedii@example.com', 5000);"
docker exec -it go_postgres psql -U postgres -d usersdb -c "INSERT INTO users (name, email, balance) VALUES ('Denis', 'Petrov@example.com', 8000);"
docker exec -it go_postgres psql -U postgres -d usersdb -c "SELECT * FROM users;"
```

---

## 4️ Запуск Go-кода
```bash
cd cmd/verify
go run .
```
