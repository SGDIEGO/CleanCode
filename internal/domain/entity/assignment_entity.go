package entity

type Assignment struct {
	AssignId      int  `json:"id"`
	UserId        int  `json:"userId"`
	GroupId       int  `json:"groupId"`
	Administrator bool `json:"isAdmin"`
}
