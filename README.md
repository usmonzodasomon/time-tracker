
# Time Tracker

Это приложение является REST API для трекинга времени пользователей на выполнение задач. Включает функции управления пользователями, задачами и учета времени.

## Установка

Для начала работы с проектом убедитесь, что у вас установлен Go и PostgreSQL.

## Настройка

Создайте файл `.env` в корневой директории проекта и добавьте следующие переменные окружения:

```env
GO_ENV=local
PORT=8080
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_PORT=5432
POSTGRES_DATABASE=database

EXTERNAL_API_URL=https://api.passportdata.com
```

## Запуск

Для запуска приложения используйте следующую команду:

```sh
make run
```

Эта команда выполнит `go run cmd/time-tracker/main.go`, запустив сервер на порту, указанном в `.env` файле (по умолчанию 8080).

## Миграции

Для применения миграций базы данных используйте следующие команды:

### Применить миграции:

```sh
make migrate_up
```
`//NOTE: Замените конфигурационные данные БД в Makefile на свои`
### Откатить миграции:

```sh
make migrate_down
```

## API Роуты

### Swagger

Документация API доступна по пути:

```
/swagger/index.html
```

### Ping

Проверка состояния сервера:

```
GET /api/ping
```

Ответ:
```json
{
  "message": "pong"
}
```

## Описание Таблиц

### Таблица `users`

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    passport_serie INTEGER NOT NULL,
    passport_number INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    address TEXT NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    UNIQUE (passport_serie, passport_number)
);
```

### Таблица `tasks`

```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users,
    name VARCHAR(255) NOT NULL,
    description TEXT
);
```

### Таблица `time_entries`

```sql
CREATE TABLE time_entries (
    id SERIAL PRIMARY KEY,
    task_id INTEGER NOT NULL REFERENCES tasks,
    start_time TIMESTAMP DEFAULT now(),
    end_time TIMESTAMP
);
```
