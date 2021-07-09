package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	Ctx          = context.Background()
	DbConnection *redis.Client
	hashKey      = "geometry"
)

func InitializeLocalRedis(address string) {
	fmt.Println("Initializing local redis")

	DbConnection = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	if err := DbConnection.Ping(Ctx).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Local redis initialized")
}

func SetInitialData() {
	fmt.Println("Setting initial data")

	firstObj, _ := json.Marshal(Rect{GeometryObject: GeometryObject{Id: 1}, Width: 2, Length: 3})
	result := DbConnection.HSet(Ctx, hashKey, "1", firstObj)
	if result.Err() != nil {
		panic(result.Err())
	}

	secondObj, _ := json.Marshal(Circle{GeometryObject: GeometryObject{Id: 2}, Radius: float64(10)})
	DbConnection.HSet(Ctx, hashKey, "2", secondObj)

	fmt.Println("Initial data have been set")
}

func SetData(id int, obj string) {
	result := DbConnection.HSet(Ctx, hashKey, strconv.Itoa(id), obj)
	if result.Err() != nil {
		panic(result.Err())
	}
}

func GetData(id int) *Rect {
	val, err := DbConnection.HGet(Ctx, hashKey, strconv.Itoa(id)).Result()
	if err == nil {
		var geometryObj Rect
		errJson := json.Unmarshal([]byte(val), &geometryObj)
		if errJson != nil {
			panic(errJson)
		}
		return &geometryObj
	} else {
		panic(err)
	}
}
