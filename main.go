package main

var (
	redisAddress = "localhost:6379"
)

func main() {
	InitializeLocalRedis(redisAddress)
	SetInitialData()
	RegisterRouter()
}
