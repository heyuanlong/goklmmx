package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	klog "goklmmx/lib/log"
)

func nextAccountIdKey() string {
	return "nextAccountId"
}
func accountKey(accountId int) string {
	return fmt.Sprintf( "id_%d",accountId)
}

func GetNextAccountId() int {
	rc := RedisClient.Get()
	defer rc.Close()
	v, err := redis.Int(rc.Do("INCR", nextAccountIdKey()))

	if err !=nil {
		klog.Klog.Println(err)
		return 0
	}
	return v
}

func SetAuth(accountId int ,value string) error  {
	rc := RedisClient.Get()
	defer rc.Close()

	if _, err := rc.Do("SET", accountKey(accountId), value) ; err !=nil {
		klog.Klog.Println(err)
		return err
	}
	if  _, err := redis.Int(rc.Do("EXPIRE", accountKey(accountId), 60)); err !=nil {
		klog.Klog.Println(err)
		return err
	}
	return nil
}