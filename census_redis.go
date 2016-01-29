package census

import (
	"github.com/garyburd/redigo/redis"
	"os"
	"fmt"
)

var ErrNoRedisURL = fmt.Errorf("census: no redis url specified.  Use SetRedisURL or REDIS_URL")
var ErrRedisCachingNotEnabled = fmt.Errorf("census: redis caching is not enabled")

func (c *Census) SetRedisURL(url string) {
	//if (c.conn != nil) {
	//	panic("Cannot change the url with an active redis connection")
	//}
	c.redisURL = url
}

func (c *Census) RedisConnect() error {
	if c.cacheType != CensusCacheRedis {
	}
	var url string
	if c.redisURL != "" {
		url = c.redisURL
	} else if os.Getenv("REDIS_URL") != "" {
		url = os.Getenv("REDIS_URL")
	} else {
		return ErrNoRedisURL
	}

	var err error
	c.conn, err = redis.DialURL(url)
	
	if err != nil {
		return err
	}
	return nil
}
