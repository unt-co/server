package db

import (
	"github.com/go-redis/redis"
)

var client *redis.Client

//InitConnection initializes the connection with redis.
//Panics if there's inndur with the databse.
func InitConnection() (*redis.Client, error) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use DB 1
	})

	_, err := client.Ping().Result()
	return client, err

	/*
		val, err := client.Get("key").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", val)
	*/

	/*
		err := client.Set("key", "value", 0).Err()
		if err != nil {
			panic(err)
		}*/
}
