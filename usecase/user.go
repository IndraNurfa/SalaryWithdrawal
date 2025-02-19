package usecase

import (
	"context"
	"errors"
	"log"
	"self-payrol/model"
	"self-payrol/request"

	"gorm.io/gorm"
)

type userUsecase struct {
	userRepository model.UserRepository
	positionRepo   model.PositionRepository
	companyRepo    model.CompanyRepository
}

func NewUserUsecase(user model.UserRepository, post model.PositionRepository, company model.CompanyRepository) model.UserUsecase {
	return &userUsecase{userRepository: user, positionRepo: post, companyRepo: company}
}

func (p *userUsecase) WithdrawSalary(ctx context.Context, req *request.WithdrawRequest) error {
	user, err := p.userRepository.FindByID(ctx, req.ID)
	if err != nil {
		log.Println("Error getting user:", err)
		return err
	}

	if user == nil { // Ensure user is not nil
		log.Println("User not found")
		return errors.New("user not found")
	}

	if user.SecretID != req.SecretID {
		log.Println("secret ID is not valid")
		return errors.New("secret id not valid")
	}

	notes := user.Name + " withdraw salary "

	pos, err := p.positionRepo.FindByID(ctx, user.PositionID)
	if err != nil {
		log.Println("Error getting posistion:", err)
		return err
	}

	if pos == nil {
		return errors.New("user position is not set")
	}

	err = p.companyRepo.DebitBalance(ctx, pos.Salary, notes)
	if err != nil {
		log.Println("Not debitting balance")
		return err
	}

	return nil
}

func (p *userUsecase) GetByID(ctx context.Context, id int) (*model.User, error) {
	user, err := p.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *userUsecase) FetchUser(ctx context.Context, limit, offset int) ([]*model.User, error) {

	users, err := p.userRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (p *userUsecase) DestroyUser(ctx context.Context, id int) error {
	err := p.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *userUsecase) EditUser(ctx context.Context, id int, req *request.UserRequest) (*model.User, error) {
	_, err := p.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user, err := p.userRepository.UpdateByID(ctx, id, &model.User{
		SecretID:   req.SecretID,
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Address:    req.Address,
		PositionID: req.PositionID,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *userUsecase) StoreUser(ctx context.Context, req *request.UserRequest) (*model.User, error) {
	newUser := &model.User{
		SecretID:   req.SecretID,
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Address:    req.Address,
		PositionID: req.PositionID,
	}

	_, err := p.positionRepo.FindByID(ctx, req.PositionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("position id not valid ")
		}

		return nil, err
	}

	user, err := p.userRepository.Create(ctx, newUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}
