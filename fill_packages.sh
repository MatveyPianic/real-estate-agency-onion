#!/bin/bash

# 1. Находим все пустые .go файлы
empty_files=$(find . -name "*.go" -type f -empty)

# Если пустых файлов нет, то и делать нечего
if [ -z "$empty_files" ]; then
    echo "👍 Пустых файлов .go не найдено!"
    exit 0
fi

# 2. Выводим список пользователю
echo "🔍 Найдены следующие пустые файлы:"
echo "$empty_files"
echo "---------------------------------"

# 3. Запрашиваем подтверждение
read -p "✍️ Заполнить их автоматически объявлениями пакетов? (y/n): " confirm

if [[ "$confirm" == "y" || "$confirm" == "Y" ]]; then
    # 4. Если пользователь согласен — запускаем цикл
    echo "$empty_files" | while read -r f; do
        # Берем имя папки, в которой лежит файл
        dir=$(basename "$(dirname "$f")")
        # Записываем "package имя_папки" в файл
        echo "package $dir" > "$f"
        echo "✅ Заполнен: $f -> package $dir"
    done
    echo "🎉 Всё готово!"
else
    echo "❌ Действие отменено."
fi