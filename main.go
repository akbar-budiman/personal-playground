package main

var (
	redisAddress = "localhost:6379"
)

func main() {
	RunGoRoutineExample()
	RunAsyncWithCollectionResultExample()

	InitializeLocalRedis(redisAddress)
	SetInitialData()
	RegisterConsumer()
	RegisterRouter()
}
