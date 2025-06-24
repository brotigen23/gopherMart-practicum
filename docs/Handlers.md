
# Хэндлеры
## Содержание
* [Регистрация пользователя](#регистрация-пользователя)
* [Аутентификация пользователя](#аутентификация-пользователя)
* [Загрузка пользователем номера заказа](#загрузка-пользователем-номера-заказа-для-расчёта)
* [Получение списка заказов пользователя](#получение-списка-загруженных-пользователем-номеров-заказов-статусов-их-обработки-и-информации-о-начислениях)
* [Получение текущего баланса](#получение-текущего-баланса-счёта-баллов-лояльности-пользователя)
* [Запрос на списание баллов](#запрос-на-списание-баллов-с-накопительного-счёта-в-счёт-оплаты-нового-заказа)
* [Информация о выводе средств](#получение-информации-о-выводе-средств-с-накопительного-счёта-пользователем)




## Регистрация пользователя
### Структура запроса
```http
POST /api/user/register HTTP/1.1
Content-Type: application/json

{
    "login": "string",
    "password": "string"
} 
```
### Ответы
```http
HTTP/1.1 200 OK - успешная регистрация и аутентификация
```
```http
HTTP/1.1 400 BadRequest - неверный формат запроса
```
```http
HTTP/1.1 409 Conflict  - логин уже занят
```
```http
HTTP/1.1 500 Internal Server Error - внутренняя ошибка сервера
```

## Аутентификация пользователя
### Структура запроса
```http
POST /api/user/login HTTP/1.1
Content-Type: application/json

{
    "login": "string",
    "password": "string"
} 
```
### Ответы
```http
HTTP/1.1 200 OK - успешная аутентификация
```
```http
HTTP/1.1 400 BadRequest - неверный формат запроса
```
```http
HTTP/1.1 401 Unauthorized  - неверный логин/пароль
```
```http
HTTP/1.1 500 Internal Server Error - внутренняя ошибка сервера
```

## Загрузка пользователем номера заказа для расчёта
```http
POST /api/user/orders HTTP/1.1
Content-Type: text/plain

12345678903 
```
### Ответы
```http
HTTP/1.1 200 OK - заказ уже был загружен
```
```http
HTTP/1.1 202 Accept - заказ принят на обработку
```
```http
HTTP/1.1 400 BadRequest - неверный формат запроса
```
```http
HTTP/1.1 401 Unauthorized  - пользователь не авторизован
```
```http
HTTP/1.1 409 Conflict  - номер заказа загружен другим пользователем
```
```http
HTTP/1.1 422 Unprocessable Entity   - неверный формат номера заказа
```
```http
HTTP/1.1 500 Internal Server Error - внутренняя ошибка сервера
```

## Получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
```http
GET /api/user/orders HTTP/1.1
Content-Length: 0
```
### Ответы
```http
HTTP/1.1 200 OK
Content-Type: application/json

[
    {
        "number": "9278923470",
        "status": "PROCESSED",
        "accrual": 500,
        "uploaded_at": "2020-12-10T15:15:45+03:00"
    }
]
```
```http
HTTP/1.1 204 No Content - нет данных
```
```http
HTTP/1.1 401 Unauthorized  - пользователь не авторизован
```
```http
HTTP/1.1 500 Internal Server Error - внутренняя ошибка сервера
```


## Получение текущего баланса счёта баллов лояльности пользователя
```http
GET /api/user/balance HTTP/1.1
Content-Length: 0
```
### Ответы
```http
HTTP/1.1 200 OK
Content-Type: application/json

{
    "current": 500.5,
    "withdrawn": 42
}
```
```http
HTTP/1.1 400 BadRequest - неверный формат запроса
```
```http
HTTP/1.1 401 Unauthorized  - пользователь не авторизован
```
```http
HTTP/1.1 500 Internal Server Error - внутренняя ошибка сервера
```

## Запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
```http
POST /api/user/balance/withdraw HTTP/1.1
Content-Type: application/json

{
    "order": "2377225624",
    "sum": 751
} 
```
### Ответы
```http
HTTP/1.1 200 OK - успешная аутентификация
```
```http
HTTP/1.1 400 BadRequest - неверный формат запроса
```
```http
HTTP/1.1 401 Unauthorized  - пользователь не авторизован
```
```http
HTTP/1.1 402 Payment Required - недостаточно средств
```
```http
HTTP/1.1 422 Unprocessable Entity - неверный номер заказа
```
```http
HTTP/1.1 500 Internal Server Error - внутренняя ошибка сервера
```

## Получение информации о выводе средств с накопительного счёта пользователем
```http
GET /api/user/withdrawals HTTP/1.1
Content-Length: 0
```
### Ответы
```http
HTTP/1.1 200 OK 
Content-Type: application/json

[
    {
        "order": "2377225624",
        "sum": 500,
        "processed_at": "2020-12-09T16:09:57+03:00"
    }
]
```
```http
HTTP/1.1 204 No Content - нет ни одного списания
```
```http
HTTP/1.1 401 Unauthorized  - пользователь не авторизован
```
```http
HTTP/1.1 500 Internal Server Error - неверный формат запроса
```


## Взаимодействие с системой расчёта начислений баллов лояльности
```http
GET /api/orders/{number} HTTP/1.1
Content-Length: 0
```
### Ответы
```http
HTTP/1.1 200 OK 
Content-Type: application/json

{
    "order": "<number>",
    "status": "PROCESSED",
    "accrual": 500
}
```
```http
HTTP/1.1 204 No Content - заказ не зарегистрирован в системе расчёта
```
```http
HTTP/1.1 429 Too Many Requests   - превышено количество запросов к сервису
```
```http
HTTP/1.1 500 Internal Server Error - неверный формат запроса
```
