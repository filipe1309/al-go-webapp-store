# DevOnTheRun Notes

> notes taken during the course

<!-- https://gitignore.io -->

https://www.toptal.com/developers/gitignore/api/go

go run main.go

psql -U postgres
\l
\c go_store
\dt
SELECT \* FROM products;

https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
https://github.com/cosmtrek/air
https://medium.com/easyread/today-i-learned-golang-live-reload-for-development-using-docker-compose-air-ecc688ee076

godoc.org
https://pkg.go.dev/

docker-compose exec store go get github.com/lib/pq
go mod vendor

products := []Product{
{Name: "T-shirt", Description: "Blue, very beatiful", Price: 39, Quantity: 5},
{"Shoes", "Confortable", 89, 3},
{"Headphode", "Very good", 59, 2},
{"New Prodcut", "Very nice", 99, 1},
}
