#  Инструкция по запуску Practice 4 

## 1️ Запуск базы данных PostgreSQL
```bash (из корня проекта)
docker-compose up -d
docker ps
```

---

## 2️ Применение миграций
```bash
migrate -path ./internal/db/migrations -database "postgres://postgres:postgres@localhost:5555/usersdb?sslmode=disable" up
docker exec -it go_postgres2 psql -U postgres -d usersdb -c "\dt"
```

---

## 3️ Добавление данных
```bash
docker exec -it go_postgres2 psql -U postgres -d usersdb -c "INSERT INTO categories (name) VALUES ('Electronics');"
docker exec -it go_postgres2 psql -U postgres -d usersdb -c "INSERT INTO categories (name) VALUES ('Books');"
docker exec -it go_postgres2 psql -U postgres -d usersdb -c "INSERT INTO categories (name) VALUES ('Clothes');"

docker exec -it go_postgres2 psql -U postgres -d usersdb -c "INSERT INTO products (name, category_id, price) VALUES ('Smartphone', 1, 150000);"
docker exec -it go_postgres2 psql -U postgres -d usersdb -c "INSERT INTO products (name, category_id, price) VALUES ('Laptop', 1, 350000);"
docker exec -it go_postgres2 psql -U postgres -d usersdb -c "INSERT INTO products (name, category_id, price) VALUES ('Novel', 2, 4000);"
docker exec -it go_postgres2 psql -U postgres -d usersdb -c "INSERT INTO products (name, category_id, price) VALUES ('T-Shirt', 3, 8000);"
```

---

## 4️ Запуск Go-кода
```bash
cd cmd/verify
go run .
```
