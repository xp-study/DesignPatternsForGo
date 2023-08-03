package main

import (
	"fmt"
	"time"
	"xp-study/DesignPatternsForGo/BuildPattern/dbpool"
)

func main() {
	dbPool, err := dbpool.Builder().DSN("localhost:3306").MaxOpenConn(50).MaxConnLifeTime(0 * time.Second).Build()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbPool)
}
