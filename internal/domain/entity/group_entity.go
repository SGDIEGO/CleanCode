package entity

type Group struct {
	GroupId   int    `json:"id"`
	GroupName string `json:"name"`
	Admin     string `json:"admin"`
	Members   int    `json:"members"`
	Created   string `json:"created"`
}
