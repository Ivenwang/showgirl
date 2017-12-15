package redis

import (
	"fmt"

	"github.com/astaxie/beego"
	//_ "github.com/astaxie/beego/cache/redis"

	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	GRedisPool   *redis.Pool
	RedisHost    string
	RedisTimeout time.Duration
	RedisLogTime int
)

var GRedisPoolMap = make(map[string]*redis.Pool)

var G_RedisExpire = int64(5184000)

func redisDialer(instanceIndex int, redisHost string, RedisTimeout time.Duration, redisAuth string) func() (redis.Conn, error) {
	return func() (redis.Conn, error) {
		c, err := redis.DialTimeout("tcp", redisHost, RedisTimeout, RedisTimeout, RedisTimeout)
		if err != nil {
			beego.Warn("redis dial failed, instanceIndex = ", instanceIndex, " RedisHost = ", redisHost, " err =", err.Error())
			return nil, err
		}
		if redisAuth != "" {
			if _, err := c.Do("AUTH", redisAuth); err != nil {
				beego.Warn("redis auth failed, err =", err.Error())
				c.Close()
				return nil, err
			}
		}
		beego.Debug("redis dial sucess, instanceIndex = ", instanceIndex, " RedisHost = ", redisHost)
		return c, nil
	}
}

func init() {
	beego.Debug("init redis")

	RedisTimeout = time.Duration(beego.AppConfig.DefaultInt("redis::timeout", 2)) * time.Second

	G_RedisExpire, _ = beego.AppConfig.Int64("redis::Expire")
	if G_RedisExpire <= 0 {
		G_RedisExpire = 5184000
	}

	RedisLogTime = beego.AppConfig.DefaultInt("redis::redislogtime", 3)
	// 建立默认redis连接池
	RedisHost = beego.AppConfig.String("redis::host")
	RedisAuth := beego.AppConfig.String("redis::auth")
	GRedisPool = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis::maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis::maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial:        redisDialer(0, RedisHost, RedisTimeout, RedisAuth),
	}
	beego.Debug("redis pool func init, RedisHost =", RedisHost)

	// 建立扩展redis连接池
	expRedisCount := beego.AppConfig.DefaultInt("redis::expRedisCount", 1)
	for i := 1; i <= expRedisCount; i++ {
		ExpRedisHost := beego.AppConfig.String(fmt.Sprintf("redis::host_%d", i))
		ExpRedisAuth := beego.AppConfig.String(fmt.Sprintf("redis::auth_%d", i))
		RedisPool := &redis.Pool{
			// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
			MaxIdle:     beego.AppConfig.DefaultInt(fmt.Sprintf("redis::maxidle_%d", i), 1),
			MaxActive:   beego.AppConfig.DefaultInt(fmt.Sprintf("redis::maxactive_%d", i), 10),
			IdleTimeout: 180 * time.Second,
			Dial:        redisDialer(i, ExpRedisHost, RedisTimeout, ExpRedisAuth),
		}
		beego.Debug("expRedis pool func init, instanceIndex = ", i, " RedisHost =", ExpRedisHost)

		//读取配置
		//设置特殊处理redis
		strRedisModule := beego.AppConfig.String(fmt.Sprintf("redis::redisModule_%d", i))
		arryRedisModule := strings.Split(strRedisModule, ",")

		for _, value := range arryRedisModule {
			GRedisPoolMap[value] = RedisPool
		}
	}

	beego.Debug("expRedis pool map=", GRedisPoolMap)

	return
}

func startTime() int64 {
	return time.Now().UnixNano()
}
func getTime(start int64) int {
	return int(time.Now().UnixNano()-start) / 1000000
}

func Set(svrname string, key string, val interface{}) error {
	start := startTime()
	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("SET", finalKey, val)

	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:SET")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}
	return nil
}

func SetAndExpire(svrname string, key string, val interface{}, seconds uint32) error {
	start := startTime()
	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("SET", finalKey, val, "EX", seconds)

	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:SETEX")
	}

	if err != nil {
		return err
	}
	return nil
}

func Del(svrname string, key string) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("DEL", finalKey)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:DEL")
	}
	if err != nil {
		return err
	}
	return nil
}

func IsExist(svrname string, key string) bool {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	beego.Debug("redis IsExist, finalKey =", finalKey)

	value, err := redis.Bool(RedisClient.Do("EXISTS", finalKey))
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:EXISTS")
	}
	if err != nil {
		beego.Warn("redis IsExist, err =", err.Error())
		return false
	}

	return value
}

//@return seconds int32 剩余的秒数(-2:key不存在，-1:key存在没有设置有效期，其他自然数：剩余秒数)
func TTL(svrname string, key string) int32 {
	start := startTime()

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
	beego.Debug("redis TTL, finalKey= ", finalKey)

	seconds, err := redis.Int(RedisClient.Do("TTL", finalKey))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:TTL")
	}
	if err != nil {
		beego.Warn("redis_TTL_err, err=", err.Error())
		return -2
	}
	if seconds < -2 {
		return -2
	}

	return int32(seconds)
}

func Expire(svrname string, key string, seconds uint32) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	beego.Debug("redis Expire, finalKey=", finalKey, " seconds=", seconds)

	value, err := redis.Bool(RedisClient.Do("EXPIRE", finalKey, seconds))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:EXPIRE")
	}
	if err != nil {
		beego.Warn("redis Expire, err=", err.Error())
		return err
	}
	if value != true {
		return errors.New("EXPIRE failed")
	}

	return nil
}

func GetInt64(svrname string, key string) (int64, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	value, err := redis.Int64(RedisClient.Do("GET", finalKey))
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:GET")
	}
	if err != nil && err != redis.ErrNil {
		return 0, err
	}

	return value, nil
}

func GetInt32(svrname string, key string) (int32, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	value, err := redis.Int(RedisClient.Do("GET", finalKey))
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:GET")
	}
	if err != nil && err != redis.ErrNil {
		return 0, err
	}

	return int32(value), nil
}

func GetString(svrname string, key string) (string, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key

	beego.Debug("redis GetString, finalKey =", finalKey)

	value, err := redis.String(RedisClient.Do("GET", finalKey))
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:GET")
	}
	if err != nil && err != redis.ErrNil {
		return "", err
	}

	return value, nil
}

func Incr(svrname string, key string) (int64, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	MaxNumber, err := redis.Int64(RedisClient.Do("INCR", finalKey))
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:INCR")
	}
	if err != nil {
		return 0, err
	}
	return MaxNumber, nil
}

func IncrBy(svrname string, key string, add int64) (int64, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	MaxNumber, err := redis.Int64(RedisClient.Do("INCRBY", finalKey, add))
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:INCRBY")
	}
	if err != nil {
		return 0, err
	}
	return MaxNumber, nil
}

func GetMulti(svrname string, keylist []string) ([]string, error) {
	start := startTime()

	if len(keylist) == 0 {
		return nil, errors.New("empty keylist")
	}

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	var args []interface{}
	for _, k := range keylist {
		finalKey := svrname + "_" + k

		beego.Debug("redis GetMulti, finalKey =", finalKey)

		args = append(args, finalKey)

		//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)
	}
	values, err := redis.Strings(RedisClient.Do("MGET", args...))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key: keysize " + strconv.Itoa(len(keylist)) + ", take time:" + strconv.Itoa(time) + "ms, method:MGET")
	}
	if err != nil {
		return nil, err
	}

	//resultList := []string{}

	//for _, v := range values {
	//	resultList = append(resultList, v)
	//}

	return values, nil
}

func GetMultiByte(svrname string, keylist []string) ([][]byte, error) {
	start := startTime()

	if len(keylist) == 0 {
		return nil, errors.New("empty keylist")
	}

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	var args []interface{}
	for _, k := range keylist {
		finalKey := svrname + "_" + k
		args = append(args, finalKey)

		//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)
	}
	values, err := redis.ByteSlices(RedisClient.Do("MGET", args...))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key: keysize " + strconv.Itoa(len(keylist)) + ", take time:" + strconv.Itoa(time) + "ms, method:SET")
	}
	if err != nil {
		return nil, err
	}

	//resultList := []string{}

	//for _, v := range values {
	//	resultList = append(resultList, v)
	//}

	return values, nil
}

//redis sorted set相关操作
func ZRangeNoScore(svrname string, key string, start int, end int) ([]string, error) {
	startTime := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	values, err := redis.Strings(RedisClient.Do("ZRANGE", finalKey, start, end))
	if time := getTime(startTime); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZRANGE")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return nil, err
	}

	return values, nil
}

//获取无序数组中的元素
func SScanByKey(svrname string, key string, startCursor int32, count int) ([]string, int32, error) {
	startTime := startTime()
	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()
	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	values, err := redis.Values(RedisClient.Do("SSCAN", finalKey, startCursor, "count", count))
	if err != nil {
		return nil, 0, err
	}
	newCursor, _ := redis.Int(values[0], nil)
	regId, _ := redis.Strings(values[1], nil)

	if time := getTime(startTime); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:SSCAN")
	}

	return regId, int32(newCursor), nil
}

//redis sorted set相关操作
//获取一段range中的用户，按score从大到小排序
func ZRevRangeNoScore(svrname string, key string, start int, end int) ([]string, error) {
	startTime := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	values, err := redis.Strings(RedisClient.Do("ZREVRANGE", finalKey, start, end))
	if time := getTime(startTime); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZREVRANGE")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return nil, err
	}

	return values, nil
}

//redis sorted set相关操作
//获取一段range中的用户（从大到小排序）
func ZRevRangeWithScore(svrname string, key string, start int, end int) ([]string, error) {
	startTime := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	values, err := redis.Strings(RedisClient.Do("ZREVRANGE", finalKey, start, end, "WITHSCORES"))
	if time := getTime(startTime); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZREVRANGE")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return nil, err
	}

	return values, nil
}

//redis sorted set相关操作
//获取一段range中的用户(从小到大排序)
func ZRangeWithScore(svrname string, key string, start int, end int) ([]string, error) {
	startTime := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	values, err := redis.Strings(RedisClient.Do("ZRANGE", finalKey, start, end, "WITHSCORES"))
	if time := getTime(startTime); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZRANGE")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return nil, err
	}

	return values, nil
}

//获取一段range中的用户(int)
func ZRevRangeNoScoreByInt(svrname string, key string, start int, end int) ([]int, error) {
	startTime := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	values, err := redis.Ints(RedisClient.Do("ZREVRANGE", finalKey, start, end))
	if time := getTime(startTime); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZREVRANGE")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return nil, err
	}

	return values, nil
}

//获取一段range中的用户(int64)
func ZRevRangeNoScoreByInt64(svrname string, key string, start int, end int) ([]int64, error) {
	startTime := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	values, err := redis.Strings(RedisClient.Do("ZREVRANGE", finalKey, start, end))
	if time := getTime(startTime); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZREVRANGE")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return nil, err
	}

	var ReturnValues []int64

	for idx := range values {
		tmpValue, StrErr := strconv.ParseInt(values[idx], 10, 64)
		if StrErr != nil {
			continue
		}
		ReturnValues = append(ReturnValues, tmpValue)
	}

	return ReturnValues, nil
}

func ZAddScoreMember(svrname string, key string, score int, member int64) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("ZADD", finalKey, score, member)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZADD")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}

	return nil
}

func ZAddScoreString(svrname string, key string, score int, member string) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("ZADD", finalKey, score, member)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZADD")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}

	return nil
}

//增加无序集合元素
func SAddString(svrname string, key string, member string) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("SADD", finalKey, member)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:SADD")
	}

	if err != nil {
		return err
	}

	return nil
}

func ZAddInt64ScoreString(svrname string, key string, score int64, member string) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("ZADD", finalKey, score, member)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZADD")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}

	return nil
}

func ZRemScoreMember(svrname string, key string, member int64) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("ZREM", finalKey, member)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZREM")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}

	return nil
}

func ZRemScoreString(svrname string, key string, member string) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("ZREM", finalKey, member)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZREM")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}

	return nil
}

func SRemString(svrname string, key string, member string) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("SREM", finalKey, member)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:SREM")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}

	return nil
}

func ZCountAddMember(svrname string, key string) (int, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	num, err := redis.Int(RedisClient.Do("ZCARD", finalKey))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZCARD")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return 0, err
	}

	return num, nil
}

func ZRemMemberList(svrname string, key string, MemberList [][]byte) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	var args []interface{}

	finalKey := svrname + "_" + key

	args = append(args, finalKey)

	for _, Member := range MemberList {
		args = append(args, Member)
	}

	beego.Debug("ZRemMemberList, before del user len = %d, key = %s", len(args), finalKey)

	values, err := redis.Int(RedisClient.Do("ZREM", args...))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZREM")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}

	beego.Debug("ZRemMemberList, del user len = %d, key = %s", values, finalKey)

	return nil
}

func ZIncrBy(svrname string, key string, member string, add int64) (int64, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	MaxNumber, err := redis.Int64(RedisClient.Do("ZINCRBY", finalKey, add, member))
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZINCRBY")
	}
	if err != nil {
		return 0, err
	}
	return MaxNumber, nil
}

//redis hash相关操作
func HashSet(svrname string, field string, key string, value string) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + field

	beego.Debug("redis HashSet, finalKey = %s, key = %s, value = %s", finalKey, key, value)

	_, err := redis.Int(RedisClient.Do("HSET", finalKey, key, value))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:HSET")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}

	return nil
}

func HashDel(svrname string, field string, key string) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + field

	beego.Debug("redis HashDel, finalKey = %s, key = %s", finalKey, key)

	_, err := redis.Int(RedisClient.Do("HDEL", finalKey, key))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:HDEL")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}

	return nil
}

func HashGet(svrname string, field string, key string) (string, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + field

	beego.Debug("redis HashGet, finalKey = %s, key = %s", finalKey, key)

	value, err := redis.String(RedisClient.Do("HGET", finalKey, key))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:HGET")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil && err != redis.ErrNil {
		return "", err
	}

	return value, nil
}

//获取某个member的score值
func ZScoreByMember(svrname string, key string, member string) (string, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		RedisClient = redisPool.Get()

	} else {
		RedisClient = GRedisPool.Get()
	}
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	value, err := redis.String(RedisClient.Do("ZSCORE", finalKey, member))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZSCORE")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil && err != redis.ErrNil {
		return "", err
	}

	return value, nil
}

// Zrevrank 反向排名，rank = -1 表示不在zset中 -2表示报错
func ZRevRank(svrname string, key string, member string) (int64, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		RedisClient = redisPool.Get()

	} else {
		RedisClient = GRedisPool.Get()
	}
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	reply, err := RedisClient.Do("ZREVRANK", finalKey, member)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:ZREVRANK")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil && err != redis.ErrNil {
		return -2, err
	}

	if reply == nil {
		return -1, err
	}

	value, err := redis.Int64(reply, err)

	return value, nil
}

//插入队尾
func ListRPush(svrname string, key string, member string) (int, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		RedisClient = redisPool.Get()

	} else {
		RedisClient = GRedisPool.Get()
	}
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	value, err := redis.Int(RedisClient.Do("RPUSH", finalKey, member))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:RPUSH")
	}
	if err != nil && err != redis.ErrNil {
		return 0, err
	}

	return value, nil
}

//阻塞移除队头
func ListBlockLPop(svrname string, key string, timeoutSec int) (*string, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		RedisClient = redisPool.Get()
	} else {
		RedisClient = GRedisPool.Get()
	}
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	value, err := redis.Strings(RedisClient.Do("BLPOP", finalKey, timeoutSec))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:BLPOP")
	}
	if err != nil && err != redis.ErrNil {
		return nil, err
	}
	if len(value) == 0 {
		return nil, nil
	}

	//redis> BLPOP job command request 0       # job 列表为空，被跳过，紧接着 command 列表的第一个元素被弹出。
	//1) "command"                             # 弹出元素所属的列表
	//2) "update system..."                    # 弹出元素所属的值
	if len(value) == 1 {
		return nil, errors.New(fmt.Sprintf("BLPOP error, value:%v", value))
	}

	return &value[1], nil
}

//移除队头
func ListLPop(svrname string, key string) (*string, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		RedisClient = redisPool.Get()
	} else {
		RedisClient = GRedisPool.Get()
	}
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	//redis> LPOP course  # 移除头元素
	//"algorithm001"
	value, err := redis.String(RedisClient.Do("LPOP", finalKey))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:LPOP")
	}
	if err != nil && err != redis.ErrNil {
		return nil, err
	}
	if len(value) == 0 {
		return nil, nil
	}

	return &value, nil
}

func SIsMember(svrname string, key string, member string) (bool, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		RedisClient = redisPool.Get()

	} else {
		RedisClient = GRedisPool.Get()
	}
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	value, err := redis.Bool(RedisClient.Do("SISMEMBER", finalKey, member))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:SISMEMBER")
	}
	if err != nil {
		return false, err
	}

	return value, nil
}

func SAdd(svrname string, key string, members []string) (int, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		RedisClient = redisPool.Get()

	} else {
		RedisClient = GRedisPool.Get()
	}
	defer RedisClient.Close()

	if len(members) <= 0 {
		return 0, errors.New("member has no element")
	}

	finalKey := svrname + "_" + key
	cmd := []interface{}{}
	cmd = append(cmd, finalKey)
	for _, v := range members {
		cmd = append(cmd, v)
	}

	value, err := redis.Int(RedisClient.Do("SADD", cmd...))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:SADD")
	}
	if err != nil {
		return 0, err
	}

	return value, nil
}

func SRem(svrname string, key string, member string) (int, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		RedisClient = redisPool.Get()

	} else {
		RedisClient = GRedisPool.Get()
	}
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	value, err := redis.Int(RedisClient.Do("SREM", finalKey, member))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:SREM")
	}
	if err != nil {
		return 0, err
	}

	return value, nil
}

//获取匹配的所有的key
func Keys(svrname string, pattern string) ([]string, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + pattern
	values, err := redis.Strings(RedisClient.Do("KEYS", finalKey))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:KEYS")
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}

//一次删除多个key的缓存
func DelMulti(svrname string, keylist []string) error {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	var args []interface{}
	for _, k := range keylist {
		args = append(args, k)
	}

	// finalKey := svrname + "_" + key
	_, err := RedisClient.Do("DEL", args...)
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key: keysize " + strconv.Itoa(len(keylist)) + ", take time:" + strconv.Itoa(time) + "ms, method:DEL")
	}

	if err != nil {
		return err
	}
	return nil
}

// 获取list的长度
func ListLen(svrname string, key string) (int, error) {
	start := startTime()

	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池

	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	value, err := redis.Int(RedisClient.Do("LLEN", finalKey))
	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:LLEN")
	}
	if err != nil {
		return 0, err
	}

	return value, nil
}

func SetByExpire(svrname string, key string, expiretimes int64, val interface{}) error {
	start := startTime()
	var RedisClient redis.Conn
	if redisPool, found := GRedisPoolMap[svrname]; found {
		// 从特殊池里获取连接
		RedisClient = redisPool.Get()

	} else {
		// 从池里获取连接
		RedisClient = GRedisPool.Get()
	}

	// 用完后将连接放回连接池
	defer RedisClient.Close()

	finalKey := svrname + "_" + key
	_, err := RedisClient.Do("SET", finalKey, val, "EX", expiretimes)

	if time := getTime(start); time > RedisLogTime {
		beego.Info("redis key:" + finalKey + ", take time:" + strconv.Itoa(time) + "ms, method:SET")
	}
	//RedisClient.Do("EXPIRE", finalKey, G_RedisExpire)

	if err != nil {
		return err
	}
	return nil
}
