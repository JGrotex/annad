package redisstorage

import (
	"sync"

	"github.com/garyburd/redigo/redis"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeRedisStorage represents the object type of the redis storage
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeRedisStorage spec.ObjectType = "redis-storage"
)

// Config represents the configuration used to create a new redis storage
// object.
type Config struct {
	Log  spec.Log
	Pool *redis.Pool
}

func DefaultConfigWithConn(redisConn redis.Conn) Config {
	newPoolConfig := DefaultRedisPoolConfig()
	newMockDialConfig := defaultMockDialConfig()
	newMockDialConfig.RedisConn = redisConn
	newPoolConfig.Dial = newMockDial(newMockDialConfig)
	newPool := NewRedisPool(newPoolConfig)

	newStorageConfig := DefaultConfig()
	newStorageConfig.Pool = newPool

	return newStorageConfig
}

// DefaultConfig provides a default configuration to create a new redis storage
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Log:  log.NewLog(log.DefaultConfig()),
		Pool: NewRedisPool(DefaultRedisPoolConfig()),
	}

	return newConfig
}

// NewRedisStorage creates a new configured redis storage object.
func NewRedisStorage(config Config) spec.Storage {
	newStorage := &storage{
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Config: config,
		Type:   ObjectTypeRedisStorage,
	}

	newStorage.Log.Register(newStorage.GetType())

	return newStorage
}

type storage struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *storage) Get(key string) (string, error) {
	conn := s.Pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", maskAny(err)
	}

	return value, nil
}

func (s *storage) GetElementsByScore(key string, score float64, maxElements int) ([]string, error) {
	conn := s.Pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("ZREVRANGEBYSCORE", key, score, score, "LIMIT", 0, maxElements))
	if err != nil {
		return nil, maskAny(err)
	}

	newList := []string{}
	for _, v := range values {
		newList = append(newList, v.(string))
	}

	return newList, nil
}

func (s *storage) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	conn := s.Pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("ZREVRANGE", key, 0, maxElements-1, "WITHSCORES"))
	if err != nil {
		return nil, maskAny(err)
	}

	newList := []string{}
	for _, v := range values {
		newList = append(newList, v.(string))
	}

	return newList, nil
}

func (s *storage) Set(key, value string) error {
	conn := s.Pool.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("SET", key, value))
	if err != nil {
		return maskAny(err)
	}

	if reply != "OK" {
		return maskAnyf(queryExecutionFailedError, "SET not executed correctly")
	}

	return nil
}

func (s *storage) SetElementByScore(key, element string, score float64) error {
	conn := s.Pool.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("ZADD", key, score, element))
	if err != nil {
		return maskAny(err)
	}

	if reply != 1 {
		return maskAnyf(queryExecutionFailedError, "ZADD not executed correctly")
	}

	return nil
}
