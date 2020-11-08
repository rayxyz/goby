package db

import (
	comdb "goby/pkg/db"
	"goby/pkg/dict"
	"goby/services/helloworld/model"
	"strings"

	"goby/pkg/conf"

	_ "github.com/go-sql-driver/mysql" //
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

var (
	basic = conf.GetBasic()
	svc   = conf.GetService("helloworld")
)

// DataSource :
type DataSource struct {
	db *sqlx.DB
}

// GetDataSource : Get datasource
func GetDataSource() *DataSource {
	ds := &DataSource{
		db: comdb.GetDB(&comdb.Config{
			Driver: svc.DSN.DBDriver,
			DB:     svc.DSN.DB,
		}),
	}
	ds.db.Mapper = reflectx.NewMapperFunc(basic.MapperTag, func(str string) string {
		return str
	})
	return ds
}

// const saveServiceListSQL = `insert into service(` + "`key`" + `, name, scheme, host, port,
// 	path_prefix, register_time, heartbeat_interval, max_heartbeat_try)
// 	values (:key, :name, :schem, :host, :port, :path_prefix, :register_time, :heartbeat_interval, :max_heartbeat_try);`

// // SaveServiceList :
// func (ds *DataSource) SaveServiceList(serviceList []*pb.Service) (bool, error) {
// 	tx := ds.db.MustBegin()

// 	for _, v := range serviceList {
// 		_, err := tx.NamedExec(saveServiceListSQL, v)
// 		if err != nil {
// 			tx.Rollback()
// 			return false, err
// 		}
// 	}

// 	if err := tx.Commit(); err != nil {
// 		tx.Rollback()
// 		return false, err
// 	}

// 	return true, nil
// }

const saveAdviceSQL = `insert into advice(content) values (:content);`

// SaveAdvice :
func (ds *DataSource) SaveAdvice(advice *model.Advice) (int64, error) {
	tx := ds.db.MustBegin()

	result, err := tx.NamedExec(saveAdviceSQL, advice)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	advice.ID, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return advice.ID, nil
}

func genGetAdviceListSQL(queryObj *model.AdviceListQueryObj) (string, []interface{}) {
	querySQL := `select 
		t.id, t.content, t.create_time
	from
		advice t
	where
		1 = 1`
	var args []interface{}

	if !strings.EqualFold(queryObj.StartTime, dict.Blank) &&
		!strings.EqualFold(queryObj.EndTime, dict.Blank) {
		querySQL += ` and t.create_time between cast(? as datetime) and cast(? as datetime)`
		args = append(args, queryObj.StartTime, queryObj.EndTime)
	}

	querySQL += ` order by t.create_time desc`

	return querySQL, args
}

// GetAdviceList :
func (ds *DataSource) GetAdviceList(queryObj *model.AdviceListQueryObj) ([]*model.Advice, error) {
	querySQL, args := genGetAdviceListSQL(queryObj)
	if queryObj.Start >= 0 && queryObj.Size > 0 {
		querySQL += ` limit ?, ?`
		args = append(args, queryObj.Start, queryObj.Size)
	}

	var adviceList []*model.Advice
	if err := ds.db.Select(&adviceList, querySQL, args...); err != nil {
		return nil, err
	}

	return adviceList, nil
}
