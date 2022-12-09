package entity

import "github.com/zebresel-com/mongodm"

type Order struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`
	Timestamp            uint        `json:"timestamp" bson:"timestamp"`
	Status               string      `json:"status" bson:"status"`
	Cases                interface{} `json:"orders" bson:"orders" model:"Case" relation:"1n" autosave:"true"`
}
