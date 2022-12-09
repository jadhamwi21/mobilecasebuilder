package entity

import "github.com/zebresel-com/mongodm"

type User struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`
	FirstName            string      `json:"first_name" bson:"first_name"`
	LastName             string      `json:"last_name" bson:"last_name"`
	Email                string      `json:"email" bson:"email"`
	Password             string      `json:"password" bson:"password"`
	Role                 string      `json:"role" bson:"role"`
	Orders               interface{} `json:"orders" bson:"orders" model:"Order" relation:"1n" autosave:"true"`
}
