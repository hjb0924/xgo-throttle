package throttle

import "xgo/xgo-throttle/driver"

type Throrrle struct {
	Db driver.ThrottleDBInterface
}

func GetInstance(dbType driver.DbType, config driver.DbConfig) *Throrrle {
	if dbType == driver.DbTypeRedis {
		redis := driver.ThrottleRedis{}
		redis.Client(config)
		return &Throrrle{
			Db: redis,
		}
	}
	return nil
}
