package entity

type Case struct {
	Id           int      `json:"id" bson:"id"`
	MobileModel  string   `json:"mobile_model" bson:"mobile_model"`
	Status       string   `json:"status" bson:"status"`
	Attachements []string `json:"attachements" bson:"attachements"`
	Color        string   `json:"color" bson:"color"`
	Image        string   `json:"image" bson:"image"`
}
