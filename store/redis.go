package store

import (
	"fmt"
	"sync"

	"github.com/Dataman-Cloud/omega-napp/config"
	log "github.com/Sirupsen/logrus"
	redis "github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func Open() redis.Conn {
	if pool != nil {
		return pool.Get()
	}

	mutex := &sync.Mutex{}
	mutex.Lock()
	InitCache()
	defer mutex.Unlock()

	return pool.Get()
}

func initConn() (redis.Conn, error) {
	conf := config.Pairs()
	addr := fmt.Sprintf("%s:%d", conf.Cache.Host, conf.Cache.Port)
	log.Debugf("init redis conn addr: %s", addr)
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return c, err
}

func InitCache() {
	conf := config.Pairs()
	pool = redis.NewPool(initConn, conf.Cache.PoolSize)
	conn := Open()
	defer conn.Close()
	pong, err := conn.Do("ping")
	if err != nil {
		log.Error("got err", err)
		log.Fatal("can't connect cache server: ", conf.Cache)
	}
	log.Debug("reach cache server ", pong)
	log.Debug("initialized cache: ", conf.Cache)
}
