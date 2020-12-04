package table

import (
	"context"
	"database/sql"
	"github.com/kallydev/privacy/database"
	"github.com/kallydev/privacy/ent"
	"github.com/kallydev/privacy/ent/jdmodel"
)

var (
	_ database.Database = &JDDatabase{}
	_ database.Model    = &JDModel{}
)

type JDDatabase struct {
	Client *ent.Client
}

func (db *JDDatabase) QueryByQQNumber(ctx context.Context, qqNumber int64) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *JDDatabase) QueryByEmail(ctx context.Context, email string) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *JDDatabase) QueryByIDNumber(ctx context.Context, idNumber string) ([]database.Model, error) {
	return []database.Model{}, nil
}

func (db *JDDatabase) QueryByPhoneNumber(ctx context.Context, phoneNumber int64) ([]database.Model, error) {
	models, err := db.Client.JDModel.
		Query().
		Where(jdmodel.PhoneNumberEQ(phoneNumber)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return entModelsToJDModels(models), nil
}

type JDModel struct {
	Name        sql.NullString
	PhoneNumber sql.NullInt64
	Address     sql.NullString
}

func (model *JDModel) GetName() (name string, valid bool) {
	return model.Name.String, model.Name.Valid
}

func (model *JDModel) GetNickname() (nickname string, valid bool) {
	return "", false
}

func (model *JDModel) GetPassword() (password string, valid bool) {
	return "", false
}

func (model *JDModel) GetEmail() (email string, valid bool) {
	return "", false
}

func (model *JDModel) GetQQNumber() (qqNumber int64, valid bool) {
	return 0, false
}

func (model *JDModel) GetIDNumber() (idNumber string, valid bool) {
	return "", false
}

func (model *JDModel) GetPhoneNumber() (phoneNumber int64, valid bool) {
	return model.PhoneNumber.Int64, model.PhoneNumber.Valid
}

func (model *JDModel) GetAddress() (address string, valid bool) {
	return model.Address.String, model.Address.Valid
}

func entModelsToJDModels(endModels []*ent.JDModel) []database.Model {
	models := make([]database.Model, len(endModels))
	for i, model := range endModels {
		models[i] = &JDModel{
			Name: sql.NullString{
				String: model.Name,
				Valid:  model.Name != "",
			},
			PhoneNumber: sql.NullInt64{
				Int64: model.PhoneNumber,
				Valid: model.PhoneNumber != 0,
			},

		}
	}
	return models
}
