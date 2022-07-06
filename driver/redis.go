package driver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"reflect"
	"time"
	"xgo/xgo-throttle/throttle"
)

type ThrottleRedis struct {
	ThrottleDBInterface
	ThrottleDB
	db *redis.Client
}

type ThrottleRedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func (this ThrottleRedis) InitDB(config ThrottleRedisConfig) *redis.Client {
	if this.db == nil {
		//config := this.Config.(ThrottleRedisConfig)
		this.db = redis.NewClient(&redis.Options{
			Addr:     config.Addr,     //"192.168.31.220:6379",
			Password: config.Password, //"", // no password set
			DB:       config.DB,       // 0,  // use default DB
		})
	}

	return this.db
}

//func fff() {
//	var XX ThrottleRedis
//	XX = ThrottleRedis{
//
//	}
//
//
//	fmt.Println(XX)
//}

func (this ThrottleRedis) Client(config DbConfig) {

	if reflect.TypeOf(config).String() == "ThrottleRedisConfig" {
		this.Config = config(config)
		this.InitDB(this.Config.(ThrottleRedisConfig))
	}
}

func (this ThrottleRedis) Get(ctx context.Context, key string) (string, error) {
	//var ctx = context.Background()
	_, err := this.db.Ping(ctx).Result()
	//fmt.Println(pong, err)
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
		return "", errors.New(fmt.Sprintf("连接redis出错，错误信息：%v", err))
	}

	// 结构体转字符串
	rs, err := this.db.Get(ctx, key).Result()
	if err == nil {
		return rs, nil
	} else {
		return "", err
	}
}

func (this ThrottleRedis) Set(ctx context.Context, key string, qpsInfo throttle.QPSData, duration time.Duration) (bool, error) {
	//var ctx = context.Background()
	_, err := this.db.Ping(ctx).Result()
	//fmt.Println(pong, err)
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
		return false, errors.New(fmt.Sprintf("连接redis出错，错误信息：%v", err))
	}

	// 结构体转字符串
	b, _ := json.Marshal(qpsInfo)
	_, err = this.db.Set(ctx, key, string(b), duration).Result()
	if err == nil {
		return true, nil
	} else {
		return false, err
	}
	return false, nil
}
