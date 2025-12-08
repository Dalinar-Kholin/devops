package dbs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

var rdb *redis.Client

func TakeRedisDB() *redis.Client {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
	}
	return rdb
}

func QueryDb[T any](ctx context.Context, m interface{}) (*T, error) {
	marshal, err := bson.Marshal(m)
	key := string(marshal)
	fmt.Printf("key := %v\n", key)
	if err != nil {
		return nil, err
	}

	data, err := TakeRedisDB().Get(ctx, key).Bytes()
	var r T
	if !errors.Is(err, redis.Nil) {
		fmt.Printf("cashed\n")
		if err := json.Unmarshal(data, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	if err := GetDataBase("inz", VoteCollection).FindOne(ctx, m).Decode(&r); err == nil {
		b, err := json.Marshal(r)
		if err != nil {
			panic(err)
		}
		TakeRedisDB().Set(ctx, key, b, time.Second*5)
		return &r, nil
	}
	return nil, nil
}
