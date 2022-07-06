package driver

import (
	"context"
	"time"
	"xgo/xgo-throttle/throttle"
)

type ThrottleDBInterface interface {
	Client(config DbConfig)
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, qpsInfo throttle.QPSData, duration time.Duration) (bool, error)
}

type DbType string
type DbConfig = interface{}

const (
	DbTypeRedis DbType = "redis"
)

type ThrottleDB struct {
	Config DbConfig
}

func (this ThrottleDB) SetDbConfig(config DbConfig) {
	this.Config = config
}

//type ThrottleRedisX struct {
//	ThrottleDB
//}

//
//func (this *ThrottleDB) GetInstance() *ThrottleDB {
//	db := new(ThrottleRedis)
//	return db.GetInstance()
//}
//
//func (this *ThrottleDB) TestGet() {
//	//db := new(ThrottleRedis)
//	//return db.GetInstance()
//
//	rds := ThrottleRedis{}
//	this.Get(&rds, context.Background(), "")
//}
//
//func (this *ThrottleDB) Get(db ThrottleDBInterface, ctx context.Context, key string) {
//	db.Get(ctx, key)
//}
//func (this *ThrottleDB) Set(ctx context.Context, key string, data string, duration time.Duration) {
//
//}
