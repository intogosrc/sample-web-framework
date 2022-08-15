package models

import (
	"github.com/jinzhu/gorm"
)

type TestModel struct {
	Id   int64
	Name string
}

// 实现 GORM 表模型接口
func (*TestModel) TableName() string {
	return "test"
}

func (test *TestModel) Save() error {
	return dbHandle.Save(test).Error
}

func FindAllTest() ([]*TestModel, error) {
	results := make([]*TestModel, 0)
	// 考虑为什么这里需要取地址操作 &
	err := dbHandle.Find(&results).Error
	if gorm.IsRecordNotFoundError(err) {
		return results, nil
	}
	return results, err
}

func FindTest(id int64) (*TestModel, error) {
	result := new(TestModel)
	// 数据库操作语句的输入变量尽可能使用占位符操作
	err := dbHandle.Where("id=?", id).First(result).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return result, err
}
