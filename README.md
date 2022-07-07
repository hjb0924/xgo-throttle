# xgo-throttle
基于golang + redis 的 高频访问限制与节流，限定用户在一段时间内的访问次数 QPS

---

## 快速使用
```
// 初始化实例，配置redis
qps := throttle.GetInstance(driver.DbTypeRedis, driver.ThrottleRedisConfig{
		Addr:     "192.168.31.220:6379", // your redis ip:port
		Password: "", // no password set
		DB:       0,  // use default DB
	})

rs := qps.CheckQPS("http://asdasd", 1, 2)
if rs == true {
    // TODO 正常逻辑代码
}else{
    panic("请勿频繁请求！")
}

```

---
## Future
+ 增加内存驱动 driver/memory
+ 增加更加详细的超频错误信息
+ 其它调优
+ 其它驱动
---
## Thinks
+ Redis驱动： `github.com/go-redis/redis`