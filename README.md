# Практическая работа #3
## Реализация простого HTTP-сервера на стандартной библиотеке net/http. Обработка запросов GET/POST
## Саттаров Булат Рамилевич ЭФМО-01-25

---


## Реализованные маршруты
- `GET    /health` – состояние сервера
- `GET    /tasks` — получить список всех задач
- `GET    /tasks/{id}` — получить задачу по id
- `POST   /tasks` — создать новую задачу (поля: title, description)
### Дополнительные маршруты
- `PATCH  /tasks/{id}` — обновить задачу по id 
- `DELETE /tasks/{id}` — удалить задачу по id 

#### Примеры запросов с использованием curl/postman находятся в файле [requests.md](docs/requests.md).

## Реализована дополнительная часть

- ### Middleware для поддержки CORS (разрешает кросс-доменные запросы).

![cors](/docs/screenshots/CORS.png)

- ### Тесты [handlers_test.go](internal/api/handlers_test.go)

![test](/docs/screenshots/test.png)

- ### Порт через переменную окружения PORT 

![port](/docs/screenshots/port.png)

- ### Валидация входящих данных для создания и обновления задач пример показан в [requests.md](docs/requests.md).

- ### [Makefile](Makefile) с задачами run/build/test 

## Запуск проекта и примеры запросов
Для запуска проекта используйте команду:
```
make run
```
Для сборки бинарного файла используйте:
```
make build
```
Для запуска тестов выполните:
```
make test
```

