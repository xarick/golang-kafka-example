# golang-kafka-example

- SQL code to add a table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    fio VARCHAR(255) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

- Docker Compose ni ishga tushirish (docker-compose.yml mavjud)
```bash
docker-compose up -d
```
