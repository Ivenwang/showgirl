package redis

import (
	//	"github.com/astaxie/beego"
	//  _ "github.com/astaxie/beego/cache/redis"

	//	"errors"
	//	"strconv"
	//	"strings"
	//	"time"

	"github.com/garyburd/redigo/redis"
)

//FUNCTION LIMIT_API_CALL(ip):
//current = GET(ip)
//IF current != NULL AND current > 10 THEN
//ERROR “too many requests per second”
//ELSE
//value = INCR(ip)
//IF value == 1 THEN
//EXPIRE(ip,1)
//END
//PERFORM_API_CALL()
//END

func LimitCall(svrname string, key string, threshold, seconds int) (bool, error) {
	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()
	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	current, err := redis.Int64(RedisClient.Do("GET", finalKey))
	if err != nil && err != redis.ErrNil {
		return false, err
	}
	if err == redis.ErrNil {
		val := int32(1)

		_, err := RedisClient.Do("SET", finalKey, val, "EX", seconds)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	if current >= (int64)(threshold) {
		return false, nil
	}
	_, err = RedisClient.Do("INCR", finalKey)
	if err != nil {
		return false, err
	}
	return true, nil
}
