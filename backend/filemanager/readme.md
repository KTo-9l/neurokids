# Filemanager

## О сервисе
Здесь находятся сервисы для работы с файлами в MongoDB GridFS

## Сервисы внутри
### getAllFiles
Возвращает массив всех файлов в виде JSON
### getFileById
Возвращает файл по его <code>id</code> на загрузку
### getFileByPath
Возвращает файл по его <code>path</code> на загрузку<br>
Путь должен быть записан в следующем формате: <code>some/path</code>, разделитель - <code>/</code>
### getFileInfoById
Возвращает информацию о файле <code>id</code> в формате JSON

### updateFile
Заменяет файл по его <code>id</code> на первый вложенный в тело запроса
### updateFile
Заменяет путь файла по его <code>id</code>

### uploadFiles
Загружает файлы в MongoDB GridFS из formdata. Принимает:<br>
- <code>files</code> - файлы;
- <code>path</code> - "путь", который будет иметься у загруженных файлов. Итоговый путь каждого файла при загрузке будет следующим - <code>path + filename</code>

### upsertFileByPath
Заменяет по пути или загружает первый вложенный в formdata файл в MongoDB GridFS. Принимает те же поля, что и <code>uploadFiles</code>
