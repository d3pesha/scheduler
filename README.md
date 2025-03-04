## Установка

1. Клонируйте репозиторий и установите зависимости:
   ```bash
   git clone https://github.com/d3pesha/scheduler.git
   cd scheduler
   go mod tidy

2. Запустите тесты:
    ```bash
   go test ./...
   
3. Запустите приложение:
    ```bash
   go run cmd/main.go
    
### Сервис запустится на порту 8080

## Доступные эндпоинты API

- `POST /jobs` - Создает новую задачу.
- `GET /jobs` - Получает список всех задач.
- `GET /jobs/{id}` - Получает информацию о задаче по её ID.
- `POST /jobs/{id}/run` - Принудительно запускает задачу.
- `DELETE /jobs/{id}` - Отменяет задачу.
