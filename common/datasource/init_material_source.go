package datasource

import (
	goredis "github.com/go-redis/redis/v8"
	"gobasic/ptc-Game/common/pkg/config"
	"gobasic/ptc-Game/common/pkg/mysql"
	"gobasic/ptc-Game/common/pkg/redis"
	"gorm.io/gorm"
)

var MaterailDataSource *DataMaterialSources

type DataMaterialSources struct {
	DB          *gorm.DB
	RedisClient *goredis.Client
}

func NewMaterialDataSources(config *config.Config) (*DataMaterialSources, error) {
	matedb, err := mysql.New(config.MySQL.Material)
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.New(config.Redis.Default)
	if err != nil {
		return nil, err
	}
	ds := &DataMaterialSources{
		DB:          matedb,
		RedisClient: redisClient,
	}
	MaterailDataSource = ds
	return ds, nil
}

func GetMaterialDataSource() *DataMaterialSources {
	return MaterailDataSource
}
