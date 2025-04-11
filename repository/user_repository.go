package repository

import (
	"github.com/cody/go-demo1/common"
	"github.com/cody/go-demo1/config"
	"github.com/cody/go-demo1/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: config.GetDB(),
	}
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *UserRepository) CreateUser(user *model.User) error {
	result := r.db.Create(user)
	return result.Error
}

// GetUsersByCondition 新增根据条件查询用户列表的方法
func (r *UserRepository) GetUsersByCondition(condition any) ([]model.User, error) {
	var users []model.User
	result := r.db.Where(condition).Find(&users)
	return users, result.Error
}

func (r *UserRepository) DeleteUser(id any) error {
	result := r.db.Delete(&model.User{}, id)
	return result.Error
}

func (r *UserRepository) GetUsersByPage(param *common.PageParam, condition any) common.PageResult[model.User] {
	result := common.Paginate[model.User](r.db, param, condition)
	return result
}
