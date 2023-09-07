package sm

import (
	"context"

	"github.com/go-redis/redis/v8"
	"k8s.io/klog"
)

func NewRedisClient(ctx context.Context, db int) (*redis.Client, error) {
	klog.Info("establishing connection with redis server")
	client := redis.NewClient(&redis.Options{
		Addr:     "10.250.108.128:6379",
		Password: "redisWipro",
		DB:       db,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		klog.Error("Unable to establish connection")
		return nil, err
	}
	return client, nil
}
