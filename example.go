package main

import (
	"xgo/xgo-throttle/driver"
	"xgo/xgo-throttle/throttle"
)

func test() {
	qps := throttle.GetInstance(driver.DbTypeRedis, driver.ThrottleRedisConfig{
		Addr:     "192.168.31.220:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	qps.CheckQPS("http://asdasd", 1, 2)
}
