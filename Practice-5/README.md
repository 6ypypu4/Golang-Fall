#  Инструкция по запуску Practice 5 

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

## 3️ Добавление данных
```bash
docker exec -it go_postgres psql -U postgres -d usersdb -c "INSERT INTO categories (name) VALUES ('Electronics');"
docker exec -it go_postgres psql -U postgres -d usersdb -c "INSERT INTO categories (name) VALUES ('Books');"
docker exec -it go_postgres psql -U postgres -d usersdb -c "INSERT INTO categories (name) VALUES ('Clothes');"

docker exec -it go_postgres psql -U postgres -d usersdb -c "INSERT INTO products (name, category_id, price) VALUES ('Smartphone', 1, 150000);"
docker exec -it go_postgres psql -U postgres -d usersdb -c "INSERT INTO products (name, category_id, price) VALUES ('Laptop', 1, 350000);"
docker exec -it go_postgres psql -U postgres -d usersdb -c "INSERT INTO products (name, category_id, price) VALUES ('Novel', 2, 4000);"
docker exec -it go_postgres psql -U postgres -d usersdb -c "INSERT INTO products (name, category_id, price) VALUES ('T-Shirt', 3, 8000);"
```

---

## 4️ Запуск Go-кода
```bash

cd cmd/verify
go run .
```

## 5 Тест 
```bash

# 1. Все продукты
curl -X GET "http://localhost:8080/products"

# 2. Фильтр по категории (например, Electronics)
curl -X GET "http://localhost:8080/products?category=Electronics"

# 3. Продукты с ценой >= 1000
curl -X GET "http://localhost:8080/products?min_price=1000"

# 4. Продукты с ценой <= 5000
curl -X GET "http://localhost:8080/products?max_price=5000"

# 5. Сортировка по цене (возрастание)
curl -X GET "http://localhost:8080/products?sort=price_asc"

# 6. Только 5 продуктов, начиная с 0
curl -X GET "http://localhost:8080/products?limit=5"

# 7. Пропустить первые 10 и вернуть следующие 5, отсортированные по убыванию цены
curl -X GET "http://localhost:8080/products?sort=price_desc&offset=10&limit=5"

```