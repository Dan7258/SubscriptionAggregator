# SubscriptionAggregator
REST-сервис для агрегации данных об онлайн-подписках пользователей

## Используемый стек
- Gin
- Swagger
- GORM
- PostgreSQL
- Docker Compose

## Эндпоинты
- `GET /`: Получить список всех подписок.
- `POST /subscriptions`: Создать новую подписку.
- `GET /subscriptions/{id}`: Получить подписку по ID.
- `PATCH /subscriptions/{id}`: Обновить подписку по ID.
- `DELETE /subscriptions/{id}`: Удалить подписку по ID.
- `POST /subscriptions/filters`: Получить подписки по фильтрам.
- `GET /swagger/index.html` : Открыть интерфейс Swagger UI для просмотра и тестирования API.

## Установка и запуск
1. Клонируйте репозиторий и перейдите в директорию проекта.
2. Выполните команду для запуска проекта с помощью Docker Compose:
   ```bash
   docker-compose up --build
   ```
3. Сервис будет доступен на `http://localhost:8080`.