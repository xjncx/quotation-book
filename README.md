# Quotation Book API

REST API сервис для управления коллекцией цитат, разработанный на Go с использованием PostgreSQL.

## Описание

Микросервис для работы с цитатами, включая:
- Добавление новых цитат с проверкой на дубликаты
- Получение списка цитат с возможностью фильтрации по автору
- Получение случайной цитаты
- Удаление цитат по ID
- Валидация входных данных и обработка ошибок

## Технический стек

- **Backend**: Go 1.24
- **База данных**: PostgreSQL 15
- **Контейнеризация**: Docker & Docker Compose
- **Архитектура**: Clean Architecture с разделением на слои (handler, service, repository)

## Быстрый запуск

```bash
git clone https://github.com/xjncx/quotation-book.git
cd quotation-book
docker compose up --build
```

API будет доступен по адресу: `http://localhost:8080`

## Тестирование функционала

### 1. Добавление цитаты

```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Marcus Aurelius", "quote":"You have power over your mind - not outside events."}'
```

### 2. Получение всех цитат

```bash
curl http://localhost:8080/quotes
```

### 3. Фильтрация по автору

```bash
curl "http://localhost:8080/quotes?author=Marcus%20Aurelius"
```

### 4. Случайная цитата

```bash
curl http://localhost:8080/quotes/random
```

### 5. Проверка дубликатов

```bash
# Повторное добавление той же цитаты вернет ошибку 409
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Marcus Aurelius", "quote":"You have power over your mind - not outside events."}'
```

Ответ:
```json
{"error":"duplicate quote"}
```

### 6. Удаление цитаты

```bash
# Сначала получите ID из списка цитат, затем удалите
curl -X DELETE http://localhost:8080/quotes/1
```

## API Спецификация

| Метод | Эндпоинт | Описание | Параметры |
|-------|----------|----------|-----------|
| POST | `/quotes` | Добавить новую цитату | `{"author": "string", "quote": "string"}` |
| GET | `/quotes` | Получить все цитаты | `?author=string` (опционально) |
| GET | `/quotes/random` | Получить случайную цитату | - |
| DELETE | `/quotes/{id}` | Удалить цитату по ID | `id` в URL |

## Структура проекта

```
.
├── cmd/
│   └── main.go           # Точка входа приложения
├── internal/
│   ├── handler/          # HTTP обработчики
│   ├── service/          # Бизнес-логика
│   ├── repository/       # Работа с БД
│   └── models/           # Структуры данных
├── docker-compose.yml    # Конфигурация контейнеров
├── Dockerfile           # Образ приложения
└── README.md
```

## Примечания

- Приложение автоматически создает необходимые таблицы при первом запуске
- Индексы настроены для оптимизации поиска по автору и случайной выборки
- Все эндпоинты возвращают JSON с соответствующими HTTP статус-кодами
- Логирование настроено для отслеживания запросов и ошибок
