// Copyright 2015 Eleme Inc. All rights reserved.

// exampleAlert.go is a script which will alert msg out by banshee
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/eleme/banshee/models"
)

type msg struct {
	Project *models.Project `json:"project"`
	Metric  *models.Metric  `json:"metric"`
	User    *models.User    `json:"user"`
}

func main() {
	str := os.Args[1]
	res := msg{}
	json.Unmarshal([]byte(str), &res)

	if res.User.EnableEmail {
		fmt.Printf("Alert metric - project:%s metric:%s value:%5f average:%5f score:%5f timestamp:%d",
			res.Project.Name, res.Metric.Name, res.Metric.Value, res.Metric.Average, res.Metric.Score,
			res.Metric.Stamp)
		fmt.Printf("User info - name:%s email:%s", res.User.Name, res.User.Email)
	}

	if res.User.EnablePhone {
		fmt.Printf("Alert metric - project:%s metric:%s value:%5f average:%5f score:%5f timestamp:%d",
			res.Project.Name, res.Metric.Name, res.Metric.Value, res.Metric.Average, res.Metric.Score,
			res.Metric.Stamp)
		fmt.Printf("User info - name:%s email:%s", res.User.Name, res.User.Phone)
	}

}
