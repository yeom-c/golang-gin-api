package repository

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/yeom-c/golang-gin-api/model"
)

type AccountDynamoRepository struct {
	db *dynamodb.Client
}

func GetAccountPk(accountId uint64) string {
	return "ACCOUNT#" + strconv.FormatUint(accountId, 10)
}

func GetAccountKey(accountId uint64) map[string]types.AttributeValue {
	pk := GetAccountPk(accountId)

	return map[string]types.AttributeValue{
		"PK": &types.AttributeValueMemberS{Value: pk},
		"SK": &types.AttributeValueMemberS{Value: "METADATA#ACCOUNT"},
	}
}

func NewAccountDynamoRepository(db *dynamodb.Client) *AccountDynamoRepository {
	return &AccountDynamoRepository{db}
}

func (r *AccountDynamoRepository) FindById(id int) (*model.DynamoAccount, error) {
	tableName := "Dev-Souls-game-user"
	key := GetAccountKey(uint64(id))

	output, err := r.db.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: &tableName,
		Key:       key,
	})
	if err != nil {
		return nil, err
	}

	if output.Item == nil {
		return nil, err
	} else {
		var item model.DynamoAccount
		err = attributevalue.UnmarshalMap(output.Item, &item)
		if err != nil {
			return nil, err
		}
		return &item, err
	}
}
