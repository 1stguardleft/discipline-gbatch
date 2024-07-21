package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/chararch/gobatch"
	"github.com/chararch/gobatch/util"
	_ "github.com/go-sql-driver/mysql"

	"github.com/1stguardleft/discipline-gbatch/job/step/handler"
	"github.com/1stguardleft/discipline-gbatch/pkg/setting"
)

func init() {
	setting.Setup()
}

// simple task
func mytask() {
	fmt.Println("mytask executed")
}

func main() {
	var db *sql.DB
	var err error

	dbtype, dbstring := setting.MakeDBString()
	fmt.Println(dbstring)
	db, err = sql.Open(dbtype, dbstring)
	if err != nil {
		panic(err)
	}
	gobatch.SetDB(db)

	//build steps
	step1 := gobatch.NewStep("mytask").Handler(handler.PrintHello).Build()
	//step2 := gobatch.NewStep("my_step").Handler(&myReader{}, &myProcessor{}, &myWriter{}).Build()
	//step2 := gobatch.NewStep("my_step").Reader(&myReader{}).Processor(&myProcessor{}).Writer(&myWriter{}).ChunkSize(10).Build()

	//build job
	job := gobatch.NewJob("myjob").Step(step1).Build()

	gobatch.Register(job)

	params, _ := util.JsonString(map[string]interface{}{
		"rand": time.Now().Nanosecond(),
	})

	gobatch.Start(context.Background(), job.Name(), params)
}
