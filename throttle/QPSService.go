package throttle

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type QPSData struct {
	LastTime int64 //`json:"last_time"`
	Num      int64
	//IsOut int
}

func (this *Throrrle) CheckQPS(key string, qps int, mill int64) bool {

	var ctx = context.Background()

	cacheKey := "qps_group:" + key

	// 缓存有效时间（秒）默认 1800
	cacheMaxTime := 1800
	if mill/1000 > 1800 {
		cacheMaxTime = int(mill / 1000)
	}

	// 以毫秒为单位
	now := time.Now().UnixMilli() //microtime(true);

	info, _ := this.getCache(ctx, cacheKey)

	if len(info) == 0 {
		qpsInfo := QPSData{now, 0}
		qpsInfo.Num++
		rs, err := this.setCache(ctx, cacheKey, qpsInfo, time.Duration(cacheMaxTime)*time.Second)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("add qps cache-key:"+cacheKey, rs, "expire:", time.Duration(cacheMaxTime)*time.Second)
		}
	} else {

		b := []byte(info)
		qpsInfo := QPSData{}
		err := json.Unmarshal(b, &qpsInfo)
		if err == nil {
			// 是否超出限定时间
			if now-qpsInfo.LastTime > mill {
				qpsInfo := QPSData{now, 0}
				qpsInfo.Num++
				this.setCache(ctx, cacheKey, qpsInfo, time.Duration(cacheMaxTime)*time.Second)
			} else {
				qpsInfo.Num++
				//isOut := qpsInfo.IsOut;

				if qpsInfo.Num > int64(qps) {
					//qpsInfo.IsOut = 1;
					this.setCache(ctx, cacheKey, qpsInfo, time.Duration(cacheMaxTime)*time.Second)
					//if(isOut == 0){
					//	// TODO something
					//}

					return false
				} else {
					this.setCache(ctx, cacheKey, qpsInfo, time.Duration(cacheMaxTime)*time.Second)
				}

				return true
			}

		} else {
			fmt.Println(err.Error())
		}

	}
	return true
}

func (this *Throrrle) setCache(ctx context.Context, key string, qpsInfo QPSData, duration time.Duration) (bool, error) {
	////var ctx = context.Background()
	//_, err := rdb.Ping(ctx).Result()
	////fmt.Println(pong, err)
	//if err != nil {
	//	fmt.Printf("连接redis出错，错误信息：%v", err)
	//	return false, errors.New(fmt.Sprintf("连接redis出错，错误信息：%v", err))
	//}
	//
	//// 结构体转字符串
	//b, _ := json.Marshal(qpsInfo)
	//_, err = rdb.Set(ctx, key, string(b), duration).Result()
	//if err == nil {
	//	return true, nil
	//} else {
	//	return false, err
	//}

	return this.Db.Set(ctx, key, qpsInfo, duration)
}

func (this *Throrrle) getCache(ctx context.Context, key string) (string, error) {
	////var ctx = context.Background()
	//_, err := driver.rdb.Ping(ctx).Result()
	////fmt.Println(pong, err)
	//if err != nil {
	//	fmt.Printf("连接redis出错，错误信息：%v", err)
	//	return "", errors.New(fmt.Sprintf("连接redis出错，错误信息：%v", err))
	//}
	//
	//// 结构体转字符串
	//rs, err := rdb.Get(ctx, key).Result()
	//if err == nil {
	//	return rs, nil
	//} else {
	//	return "", err
	//}

	return this.Db.Get(ctx, key)
}
