package env

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/sarastee/avito-test-assignment/internal/config"
)

const (
	redisHostEnv        = "REDIS_HOST"
	redisPortEnv        = "REDIS_PORT"
	redisMaxIdleEnv     = "REDIS_MAX_IDLE"
	redisIdleTimeoutEnv = "REDIS_IDLE_TIMEOUT_SECONDS"
	redisTTLMinutes     = "REDIS_TTL_MINUTES"
)

// RedisCfgSearcher searcher for white list redis config.
type RedisCfgSearcher struct{}

// NewRedisCfgSearcher get instance searcher for redis config.
func NewRedisCfgSearcher() *RedisCfgSearcher {
	return &RedisCfgSearcher{}
}

// Get config for Redis connection.
func (s *RedisCfgSearcher) Get() (*config.RedisConfig, error) {
	dbHost := os.Getenv(redisHostEnv)
	if len(dbHost) == 0 {
		return nil, errors.New("redis host not found")
	}

	dbPort := os.Getenv(redisPortEnv)
	if len(dbPort) == 0 {
		return nil, errors.New("redis port not found")
	}

	maxIdle := os.Getenv(redisMaxIdleEnv)
	if len(maxIdle) == 0 {
		return nil, errors.New("redis max idle not found")
	}
	maxIdleInt, err := strconv.Atoi(maxIdle)
	if err != nil {
		return nil, errors.New("redis max idle is not int")
	}

	idleTimeout := os.Getenv(redisIdleTimeoutEnv)
	if len(idleTimeout) == 0 {
		return nil, errors.New("redis idle timeout not found")
	}
	idleTimeoutInt, err := strconv.Atoi(maxIdle)
	if err != nil {
		return nil, errors.New("redis idle timeout is not int")
	}

	ttl := os.Getenv(redisTTLMinutes)
	if len(ttl) == 0 {
		return nil, errors.New("redis time to live not found")
	}

	ttlInt, err := strconv.Atoi(ttl)
	if err != nil {
		return nil, errors.New("redis time to live is not int")
	}

	return &config.RedisConfig{
		Host:        dbHost,
		Port:        dbPort,
		MaxIdle:     maxIdleInt,
		IdleTimeout: time.Duration(idleTimeoutInt) * time.Second,
		TTL:         time.Duration(ttlInt) * time.Minute,
	}, nil
}
