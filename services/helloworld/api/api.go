package api

import (
	"goby/services/helloworld/model"

	"golang.org/x/net/context"

	"goby/services/helloworld/db"

	"goby/pkg/conf"
	"goby/pkg/redis"

	log "github.com/sirupsen/logrus"
)

var (
	svcConf = conf.GetService("helloworld")
	ds      = db.GetDataSource()

	redisConf = conf.GetRedis()
	// rc: redis client
	rc = redis.NewClient(&redis.ConnectConf{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       (svcConf.Extra["redis_db"]).(int),
	})
)

// SaveAdvice :
func SaveAdvice(ctx context.Context, advice *model.Advice) (int64, error) {
	id, err := ds.SaveAdvice(advice)
	if err != nil {
		log.Error("SaveAdvice => ", err)
		return 0, err
	}
	return id, nil
}

// GetAdviceList :
func GetAdviceList(ctx context.Context, queryObj *model.AdviceListQueryObj) ([]*model.Advice, error) {
	data, err := ds.GetAdviceList(queryObj)
	if err != nil {
		log.Error("GetAdviceList => ", err)
		return nil, err
	}
	return data, nil
}
