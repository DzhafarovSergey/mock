# mock

### Запуск проекта

```
go run cmd/main.go 
```

### Авторизиция

Для получения TOKEN необходимо отправить запрос:

```
curl -X POST http://127.0.0.1:8000/sso/oauth/token -H "Content-Type: application/json" -d'{"clientId":"GDgHLESI3dF7av6gwyt2jde3d3", "secretKey":"DR1P3FCMPtKwfzTJoGqrQJaqwrAUxhw4Lv2aZwKbCmbKO5yTb9hMEatKkJ6Mj5h1"}'
```

Результат:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRJZCI6IkdEZ0hMRVNJM2RGN2F2Nmd3eXQyamRlM2QzIiwiZXhwIjoxNzM0MTQxMjgzfQ.0ptvMmoZroYMOdKycVj6Miw1PxntVS6SaUbd-0lcd7c
```

TOKEN доступен в течении 3-х часов