#!/bin/bash

# Параметры
BACKUP_DIR=/backup/files
TIMESTAMP=$(date +"%F:%H:%M")
BACKUP_FILE="$BACKUP_DIR/suppliers_backup_$TIMESTAMP.gz"

# Создание директории для бэкапов, если она не существует
mkdir -p $BACKUP_DIR

# Выполнение бэкапа
mongodump --host mongo --port 27017 --gzip --archive=$BACKUP_FILE

echo "Backup created at $BACKUP_FILE"

find /files -type f -name '*.gz' -mtime +7 -exec rm {} ;