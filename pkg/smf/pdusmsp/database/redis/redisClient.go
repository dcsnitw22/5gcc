// TODO gracefull closing client

// TODO create redisManager and start
package redisClient

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"k8s.io/klog"
	"go.opentelemetry.io/otel/codes"
	"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/tracing"
)

type Database int

func NewRedisDbManager() *RedisInfo {
	ctx := context.Background()

	//Redis Client initialization
	client, err := NewRedisClient(ctx, 0)
	if err != nil {
		klog.Errorf("unable to connect to database:ERROR:%s", err.Error())
	}
	klog.Info("Connected to redis Database")
	return &RedisInfo{
		Ctx:    ctx,
		Client: client,
	}
}

const (
	SessionDb Database = 0
	UserDb    Database = 1
)

type RedisInfo struct {
	Ctx    context.Context
	Client *redis.Client
}

func NewRedisClient(ctx context.Context, redisDb Database) (*redis.Client, error) {
	/*
		Creates new client
		parameters : [context, Database]
		returns: [redisClient,error]
	*/
	klog.Info("establishing connection with redis server")
	// _, span := tracing.Tracer.Start(context.Background(), "NewRedisClient")
	// defer span.End()
	db := database2Int(redisDb)

	//TODO get from config
	client := redis.NewClient(&redis.Options{
		Addr:     "172.16.27.15:6379",
		Password: "redisWipro",
		DB:       db,
	})

	//TODO put it in start function
	_, err := client.Ping(ctx).Result()
	if err != nil {
		klog.Error("Unable to establish connection")
		return nil, err
	}
	// span.SetStatus(codes.Ok, "Connection with redis established")
	return client, nil
}

func (rc *RedisInfo) ChangeRedisDatabase(redisDb Database) {
	/*
		Change database client by creating new Database client
		Note: Unable to find a method to change db without creating new client
		parameters : [Database]
		returns: None
	*/
	_, span := tracing.Tracer.Start(context.Background(), "ChangeRedisDatabase")
	defer span.End()
	var err error
	// db := database2Int(redisDb)
	rc.Client, err = NewRedisClient(rc.Ctx, redisDb)
	// rc.Client.Options().DB = db
	if err != nil {
		klog.Error(err.Error())
	}
	span.SetStatus(codes.Ok, "Database changed")
	klog.Infof("Database changed to : %+v| database requested:%+v ", int2Database(rc.Client.Options().DB), redisDb)
}

// TODO return DATABSE,error
func int2Database(dbnum int) Database {
	/*
		Convert int values to Database type values
		parameters : [databaseNumber]
		returns: [Database]
	*/
	// _, span := tracing.Tracer.Start(context.Background(), "int2Database")
	// defer span.End()
	if dbnum == 0 {
		return SessionDb
	}
	if dbnum == 1 {
		return UserDb
	}
	// span.SetStatus(codes.Ok, "int2Database")
	return SessionDb
}

// TODO return INT,error
func database2Int(redisDb Database) int {

	/*
		Convert database type value into int
		parameters : [Database]
		returns: [integer]
	*/
	// _, span := tracing.Tracer.Start(context.Background(), "database2Int")
	// defer span.End()
	if redisDb == SessionDb {
		return 0
	}
	if redisDb == UserDb {
		return 1
	}
	// span.SetStatus(codes.Ok, "database2Int")
	return 0
}

func (rc *RedisInfo) Create(key string, value string, database Database) (string, error) {
	/*
		create a key value pair in given client database
		parameters : [key,value,Database]
		returns: [outputString,error]
	*/
	_, span := tracing.Tracer.Start(context.Background(), "Create")
	defer span.End()
	ctx := context.Background()
	rc.ChangeRedisDatabase(database)
	data, err := rc.Client.Set(ctx, key, value, 0).Result()
	if err != nil {
		return "", errors.New("failed to create data in Redis: " + err.Error())
	}
	span.SetStatus(codes.Ok, "Create")
	return data, err
}

// Retrieve data in the specified Redis database
func (rc *RedisInfo) Read(key string, database Database) (string, error) {
	/*
		get  value corresponding to given key in given client database
		parameters : [key,Database]
		returns: [outputString,error]
	*/
	_, span := tracing.Tracer.Start(context.Background(), "Read")
	defer span.End()
	ctx := context.Background()
	rc.ChangeRedisDatabase(database)
	data, err := rc.Client.Get(ctx, key).Result()
	if err != nil {
		return "", errors.New("failed to get data in Redis: " + err.Error())
	}
	span.SetStatus(codes.Ok, "Read")
	return data, err
}

// Release data in the specified Redis database
func (rc *RedisInfo) Delete(key string, database Database) (int64, error) {
	/*
		Delete a key value pair in given client database
		parameters : [key,Database]
		returns: [outputString,error]
	*/
	_, span := tracing.Tracer.Start(context.Background(), "Delete")
	defer span.End()
	ctx := context.Background()
	rc.ChangeRedisDatabase(database)
	data, err := rc.Client.Del(ctx, key).Result()
	if err != nil {
		return -1, errors.New("failed to release data in Redis: " + err.Error())
	}
	span.SetStatus(codes.Ok, "Delete")
	return data, err
}

// Update updates data in the specified Redis database
func (rc *RedisInfo) Update(key, value string, database Database) (string, error) {
	/*
		U[datae a key value pair in given client database
		parameters : [key,Database]
		returns: [outputString,error]
	*/
	_, span := tracing.Tracer.Start(context.Background(), "Update")
	defer span.End()
	ctx := context.Background()
	rc.ChangeRedisDatabase(database)
	data, err := rc.Client.Set(ctx, key, value, 0).Result()
	if err != nil {
		return "", errors.New("failed to update data in Redis: " + err.Error())
	}
	span.SetStatus(codes.Ok, "Update")
	return data, nil
}

// // MarshalJSON marshals an interface to JSON
// func MarshalJSON(v interface{}) (string, error) {
// 	data, err := json.Marshal(v)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(data), nil
// }

// // UnmarshalJSON unmarshals JSON data to an interface
// func UnmarshalJSON(data string, v interface{}) error {
// 	return json.Unmarshal([]byte(data), v)
// }
