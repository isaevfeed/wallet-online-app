# Клиент обмена валютой

Небольшой эхо-сервер, написанный ради интереса. Каждый подключаемый клиент получает свой кашелёк с 10000-ми некой валюты. Любой пользователь может перевести себе 100 единиц валюты в реальном времени.

Написано в учебных целях

### Запуск

- Запустить сервер: `go run ./cmd/server/main.go`
- Запустить клиент (сколько угодно): `go run ./cmd/client/main.go`