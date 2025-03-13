# Используем официальный образ MongoDB
FROM mongo:latest

# Устанавливаем необходимые пакеты
RUN apt-get update && apt-get install -y cron && rm -rf /var/lib/apt/lists/*

# Создаем директорию для скриптов и бэкапов
RUN mkdir -p /files

# Копируем скрипт для бэкапа
COPY backup.sh /backup.sh

# Даем права на выполнение скрипта
RUN chmod +x /backup.sh

# Настраиваем cron
COPY crontab /etc/cron.d/mongo-backup-cron

# Обеспечиваем, что cron файл имеет правильные права доступа
RUN chmod 0644 /etc/cron.d/mongo-backup-cron

# Применяем cron файл
RUN crontab /etc/cron.d/mongo-backup-cron

# Запускаем cron и MongoDB
CMD ["cron", "-f"]