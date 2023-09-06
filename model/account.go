package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username string
}

type DynamoAccount struct {
	PK string `dynamodbav:"PK" json:"-"`
	SK string `dynamodbav:"SK" json:"-"`

	Id uint64 `dynamodbav:"accountId" json:"id"`

	CreateAt string `dynamodbav:"createAt" json:"createAt"`
	UpdateAt string `dynamodbav:"updateAt" json:"updateAt"`
}
