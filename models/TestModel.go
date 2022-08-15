package models

import (
	"github.com/jinzhu/gorm"
)

type TestModel struct {
	Id   int64
	Name string
}

// implements GORM table Interface
func (*TestModel) TableName() string {
	return "test"
}

func (test *TestModel) Save() error {
	return dbHandle.Save(test).Error
}

func FindAllTest() ([]*TestModel, error) {
	results := make([]*TestModel, 0)
	err := dbHandle.Find(&results).Error
	if gorm.IsRecordNotFoundError(err) {
		return results, nil
	}
	return results, err
}

func FindTest(id int64) (*TestModel, error) {
	result := new(TestModel)
	err := dbHandle.Where("id=?", id).First(result).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return result, err
}
