package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func main() {

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	user := &User{Id: 1, Name: "mertcakmak2"}
	userId := strconv.Itoa(user.Id)

	err := rdb.Set(ctx, userId, user, 0).Err()
	if err != nil {
		panic(err)
	}

	// res, err := rdb.Del(ctx, userId).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)

	val, err := rdb.Get(ctx, userId).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Cached user:", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

}

type User struct {
	Id   int
	Name string
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
