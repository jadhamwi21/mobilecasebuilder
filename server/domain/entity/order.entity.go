package entity

type Order struct {
	Id        int    `json:"id" bson:"id"`
	Timestamp uint   `json:"timestamp" bson:"timestamp"`
	Status    string `json:"status" bson:"status"`
	Cases     []int  `json:"cases" bson:"cases"`
}
