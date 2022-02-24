package model

import "time"

type Member struct {
	MemberId       string    `bson:"memberId" json:"memberId"`
	MemberName     string    `bson:"memberName" json:"memberName"`
	MemberPassword string    `bson:"memberPassword" json:"memberPassword"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
}
