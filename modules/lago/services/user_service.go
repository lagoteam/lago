package services

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"

	"goravel/modules/lago/models"
)

type UserService struct {
	BaseService
	Dao orm.Query
}

func NewUserService() UserService {
	return UserService{
		BaseService: NewBaseService(),
		Dao:         facades.Orm().Query().Model(&models.User{}),
	}
}

func (s *UserService) List(page int, limit int, search string) ([]models.User, int64, error) {
	var items []models.User
	var total int64
	err := s.Dao.Paginate(page, limit, &items, &total)
	return items, total, err
}

func (s *UserService) GetUser(id uint) (models.User, error) {
	var user models.User
	err := s.Dao.FindOrFail(&user, id)
	return user, err
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	// 判断密码是否需要加密
	if len(user.Password) > 0 && facades.Hash().NeedsRehash(user.Password) {
		user.Password, _ = facades.Hash().Make(user.Password)
	}
	err := s.Dao.Create(&user)
	return user, err
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	err := s.Dao.Save(&user)
	return user, err
}

func (s *UserService) DeleteUser(ids []uint) (*orm.Result, error) {
	//ids := []uint{1, 2, 3}
	result, err := s.Dao.Delete(ids)
	return result, err
}
