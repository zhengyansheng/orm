package orm

import "github.com/go-redis/redis"

var (
	Rds *redis.Client
)

func InitRedis(addr, password string, db int) error {
	Rds = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	return nil
}
