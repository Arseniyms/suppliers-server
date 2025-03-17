# Suppliers Server

Данный репозиторий содержит серверную часть приложения для управления взаимодействия поставщками компании

## Установка

```bash
    git clone https://github.com/Arseniyms/suppliers-server.git
```

## Запуск

Для запуска необходим **Docker** и **Docker-Compose**

### Установка переменных окружения

1) Скопировать *.env.example* файл и переименовать его в *.env*
2) Заполнить все переменные

```
SMTP_ADDRESS_FROM="адрес, с которого будет отправляться сообщение на почту"
SMTP_PASSWORD_FROM="пароль для почты"
SMTP_HOST="хост SMTP"
SMTP_PORT="пароль от хоста SMTP"

DATABASE_NAME="название базы данных" (может быть любое, прим. suppliers)
COLLECTION_NAME="название коллекции" (может быть любое, прим. main)

JWT_SECRET_KEY="ключ для генерации токенов авторизации" (https://jwtsecret.com/generate)
TEMP_PASSWORD="пароль для входа админа"
```

### Описание Docker-compose

1) mongo - образ MongoDB базы данных для хранения всей информации
2) golang-connectors - миддл слой с логикой получения и обработки информации о поставщиках
3) backup_cron - бекап сервис для создания резерных копий БД каждый день в 5 утра (все копии хранятся в */backup/files*)
4) restore - нужен для восстановления из резервной копии

### Запуск Docker-compose

```bash
    docker-compose up
```

После успешного запуска в консоли должно вывестись:
```
Connected to MongoDB!
Server is running on port 9090
```

## Использование

Для доступа используется порт **9090**

Сделать тестовый запрос GET http://localhost:9090

Ожидается код ответа 200 OK
```http
    Successfully connected
    Now you can make requests
```