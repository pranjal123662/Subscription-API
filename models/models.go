package model

//PayloadForAPI
type DataPaylod struct {
	Email string `json:"email,omitempty" bson:"email,omitempty"`
}

type EncodeData struct {
	Code    string `json:"code,omitempty" bson:"code,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}
