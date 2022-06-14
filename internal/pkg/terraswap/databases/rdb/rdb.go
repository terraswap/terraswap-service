package rdb

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/terraswap/terraswap-service/configs"
	"github.com/terraswap/terraswap-service/internal/pkg/terraswap"
)

type TerraswapRdb interface {
	GetZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error)
}

type rdb struct {
	*sql.DB
}

func New(c configs.RdbConfig) *rdb {
	if c.Host == "" || c.Port == "" || c.Username == "" || c.Password == "" || c.Database == "" {
		panic(errors.New("must provide RDB options properly"))
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.Username, c.Password, c.Database,
	)

	var rdb = &rdb{}
	var err error

	rdb.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	rdb.prepareStmt()

	return rdb
}

func (r *rdb) Close() {
	r.DB.Close()
}

var (
	stmtZeroPool *sql.Stmt
)

const (
	sqlZeroPool = `
	SELECT
		pair
	FROM (
		SELECT
			DISTINCT ON(pair) pair,
			token_0_reserve * token_1_reserve as pool
		FROM 
			pair_day_data
		ORDER BY
			pair, "timestamp" DESC
		) t1
	WHERE 
		t1.pool = 0;`
)

func (r *rdb) prepareStmt() {
	var err error
	stmtZeroPool, err = r.Prepare(sqlZeroPool)
	if err != nil {
		panic(err)
	}
}

func (r *rdb) GetZeroPoolPairs(pairs []terraswap.Pair) (map[string]bool, error) {
	zeroPoolPairs := make(map[string]bool)

	rows, err := stmtZeroPool.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pair := ""
		err = rows.Scan(&pair)
		if err != nil {
			return nil, err
		}

		zeroPoolPairs[pair] = true
	}

	return zeroPoolPairs, nil
}
