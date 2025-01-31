# Personal Blogging Platform API

## Описание
API для управления статьями в персональной блог-платформе. Позволяет создавать, редактировать, удалять и получать статьи.

## Структура данных

### Article
| Поле      | Тип данных  | Описание                  |
|-----------|------------|---------------------------|
| id        | UUID       | Уникальный идентификатор |
| title     | String     | Заголовок статьи        |
| body      | Text       | Содержимое статьи       |
| tags      | String     | Тэги                    |
| createdAt | Timestamp  | Дата создания           |
| updatedAt | Timestamp  | Дата последнего обновления |

## API Эндпоинты

### Получение всех статей
`GET /articles`
#### Ответ:
```json
[
  {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "title": "Первая статья",
    "body": "Содержимое первой статьи...",
    "tags": "Тэги",
    "createdAt": "2025-01-01T12:00:00Z",
    "updatedAt": "2025-01-01T12:30:00Z"
  }
]
```

---

### Получение статьи по ID
`GET /article/:id`
#### Ответ:
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "title": "Первая статья",
  "body": "Содержимое первой статьи...",
  "tags": "Тэги",
  "createdAt": "2025-01-01T12:00:00Z",
  "updatedAt": "2025-01-01T12:30:00Z"
}
```

---

### Создание статьи
`POST /article`
#### Тело запроса:
```json
{
  "title": "Новая статья",
  "body": "Текст новой статьи...",
  "tags": "Тэги"
}
```
#### Ответ:
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174001",
  "title": "Новая статья",
  "body": "Текст новой статьи...",
  "tags": "Тэги",
  "createdAt": "2025-01-02T10:00:00Z",
  "updatedAt": "2025-01-02T10:00:00Z"
}
```

---

### Обновление статьи
`PUT /article/:id`
#### Тело запроса:
```json
{
  "title": "Обновленный заголовок",
  "body": "Обновленный текст статьи..."
}
```
#### Ответ:
```json
{
  "id": "1",
  "title": "Обновленный заголовок",
  "body": "Обновленный текст статьи...",
  "tags": "Тэги",
  "createdAt": "2025-01-02T10:00:00Z",
  "updatedAt": "2025-01-02T12:00:00Z"
}
```

---

### Удаление статьи
`DELETE /article/:id`
#### Ответ:
```json
{
  "message": "Article deleted successfully"
}
```