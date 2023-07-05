# Тестовое задание TAGES (golang)

http сервис, который хранит в памяти информацию о пользователях в виде:
```go
type User struct {
	ID    int
	Token string
	Name  string
	Age   int
}
```
Сервис может обрабатывать два вида запросов:
1. `GET /user/{id}` - отдает пользователю структуру User перекодированную в json.
2. `POST /user/{id}` - обновляет информацию о пользователе раскодируя json внутрь объекта User, но клиент не может изменять `ID` и `Token`

## Примеры использования
```shell
curl http://localhost:8080/user/1
```
>`{"id":1,"token":"token1","name":"John Doe","age":25}`

```shell
curl -X POST -H "Content-Type: application/json" -d "{\"id\":1,\"token\":\"token1\",\"name\":\"New Name\",\"age\":26}" http://localhost:8080/user/1
```
>`User info updated succesfully`

```shell
curl http://localhost:8080/user/1
```
>`{"id":1,"token":"token1","name":"New Name","age":26}`

```shell
curl http://localhost:8080/user/100
```
>`404 page not found`

```shell
curl -X POST -H "Content-Type: application/json" -d "{\"id\":2,\"token\":\"token1\",\"name\":\"New Name\",\"age\":26}" http://localhost:8080/user/1
```
>`Cannot change ID or Token`

```shell
curl -X POST -H "Content-Type: application/json" -d "{\"id\":1,\"token\":\"token2\",\"name\":\"New Name\",\"age\":26}" http://localhost:8080/user/1
```
>`Cannot change ID or Token`

```shell
curl -X POST -H "Content-Type: application/json" -d "{\"invalid\":json}" http://localhost:8080/user/1
```
>`Invalid JSON`
