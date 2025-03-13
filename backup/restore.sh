#!/bin/bash

# Параметры
BACKUP_FILE=$1

if [ -z "$BACKUP_FILE" ]; then
  echo "Нужно передать название бекапа: $0 <backup_file>"
  exit 1
fi

# Восстановление из бэкапа
mongorestore --uri mongodb://mongo:27017 --gzip --archive=$BACKUP_FILE

echo "Database restored from $BACKUP_FILE"