package openapi_pcf_srv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type DBClient struct {
	redisClient *redis.Client
}

func (d DBClient) DBConnect() (*redis.Client, error) {

	cli := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       6,
		},
	)
	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("error pinging redis: ", err)
		return nil, err
	}

	return cli, err

}

func (d DBClient) PolicyExists(ctx context.Context, smPolicyId string) bool {
	_, err := d.redisClient.Do(ctx, "select", "6").Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	row := d.redisClient.Get(ctx, smPolicyId)
	return !errors.Is(row.Err(), redis.Nil)
}

// return true if created else false
func (d DBClient) PolicyCreate(ctx context.Context, smPolicyId string, smPolicyDes SmPolicyDecision) bool {
	_, err := d.redisClient.Do(ctx, "select", "6").Result()
	if err != nil {
		return false
	}
	row := d.redisClient.Set(ctx, smPolicyId, smPolicyDes, 0)
	return row.Err() == nil
}

// TODO switch to context database
func (d DBClient) ContextCreate(ctx context.Context, smPolicyId string, smPolicyCtx SmPolicyContextData) bool {
	_, err := d.redisClient.Do(ctx, "select", "7").Result()
	if err != nil {
		return false
	}
	row := d.redisClient.Set(ctx, smPolicyId, smPolicyCtx, 0)
	fmt.Println(row.Err())
	return row.Err() == nil
}

func (d DBClient) PolicyDelete(ctx context.Context, smPolicyId string) bool {
	_, err := d.redisClient.Do(ctx, "select", "6").Result()
	if err != nil {
		return false
	}
	res, err := d.redisClient.Del(ctx, smPolicyId).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return res >= 0
}

// TODO switch to context database
func (d DBClient) ContextDelete(ctx context.Context, smPolicyId string) bool {
	_, err := d.redisClient.Do(ctx, "select", "7").Result()
	if err != nil {
		return false
	}
	res, err := d.redisClient.Del(ctx, smPolicyId).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return res >= 0
}

func (d DBClient) PolicyGet(ctx context.Context, smPolicyId string) (SmPolicyDecision, error) {
	_, err := d.redisClient.Do(ctx, "select", "6").Result()
	if err != nil {
		return SmPolicyDecision{}, err
	}
	res, err := d.redisClient.Get(ctx, smPolicyId).Result()
	if err != nil {
		return SmPolicyDecision{}, err
	}
	var smPolDec SmPolicyDecision
	json.Unmarshal([]byte(res), &smPolDec)
	return smPolDec, nil
}

// TODO switch to context data
func (d DBClient) ContextGet(ctx context.Context, smPolicyId string) (SmPolicyContextData, error) {
	_, err := d.redisClient.Do(ctx, "select", "7").Result()
	if err != nil {
		return SmPolicyContextData{}, err
	}
	res, err := d.redisClient.Get(ctx, smPolicyId).Result()
	if err != nil {
		return SmPolicyContextData{}, err
	}
	var smPolCtx SmPolicyContextData
	json.Unmarshal([]byte(res), &smPolCtx)
	return smPolCtx, nil
}

// Pcc rules inserting into redis
func (d DBClient) QosCreate(ctx context.Context, qosId string, qosData QosData) bool {
	_, err := d.redisClient.Do(ctx, "select", "8").Result()
	if err != nil {
		return false
	}
	row := d.redisClient.Set(ctx, qosId, qosData, 0)
	fmt.Println(row.Err())
	return row.Err() == nil
}
