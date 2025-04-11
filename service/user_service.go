package service

import (
	"github.com/acmestack/gorm-plus/gplus"
	"github.com/cody/go-demo1/common"
	"github.com/cody/go-demo1/model"
	"github.com/cody/go-demo1/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (s *UserService) GetAllUsers(param *common.PageParam, condition any) common.PageResult[model.User] {
	return s.userRepository.GetUsersByPage(param, condition)
}

func (s *UserService) GetAllUsersByGPlus(param *common.PageParam, condition any) common.PageResult[model.User] {
	p := gplus.NewPage[model.User](param.Page, param.PageSize)
	query, _ := gplus.NewQuery[model.User]()
	//query.Select(condition)
	p, _ = gplus.SelectPageGeneric(p, query)
	return common.PageResult[model.User]{
		Total: p.Total,
		Rows:  p.Records,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepository.CreateUser(user)
}

// GetUsersByCondition 新增根据条件查询用户列表的服务方法
func (s *UserService) GetUsersByCondition(condition interface{}) ([]model.User, error) {
	return s.userRepository.GetUsersByCondition(condition)
}

func (s *UserService) DeleteUserById(id any) {
	err := s.userRepository.DeleteUser(id)
	if err != nil {
		return
	}
}
