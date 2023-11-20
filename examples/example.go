package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"rorm"
)

const redisApi = "IP:PORT"
const redisPwd = "REDIS_PASSWORD"

var ctx = context.Background()
var Client *redis.Client

// AuthCode Your Data Struct
// Assuming that you put record to redis when user use mail to login
type AuthCode struct {
	Mail       string `rorm:"key;name:mail;"` // set mail to be the ID
	Code       string `rorm:"name:code;"`     // name code
	VisitCount int    // not tag, default name: visit_count
	Test       string `rorm:"-"` // ignore this field
}

func main() {
	// link to your redis server
	Client = redis.NewClient(&redis.Options{
		Addr:     redisApi,
		Password: redisPwd,
		DB:       0,
	})

	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return
	}

	// init rorm DB
	rdb := rorm.Open(Client)

	// add new rorm model
	rdb.AppendModel(&AuthCode{})

	// set record
	record := AuthCode{Mail: "xxx@gmail.com", Code: "172837", VisitCount: 0}
	res := rdb.Set(&record)
	if res {
		fmt.Printf("record ID %s has been successfully settled\n", record.Mail)
	}

	// get record from redis client
	target := AuthCode{Mail: "xxx@gmail.com"}
	res = rdb.Get(&target)
	if !res {
		fmt.Printf("can not find record Authcode where phone = %s\n", target.Mail)
	} else {
		fmt.Printf("the login code of Mail %s is %s", target.Mail, target.Code)
	}

	// change record's field value
	rdb.IncrBy(&record, "VisitCount", 1) // change record
	rdb.Get(&record)
	fmt.Println(record)
	// VisitCount : 1
	rdb.IncrBy(&record, "visit_count", 1) // change record
	rdb.Get(&record)
	fmt.Println(record)
	// VisitCount : 2
}
