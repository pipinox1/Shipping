package shipping

import (
	"apitest/internal/domain/errors"
	"apitest/internal/domain/model"
	"apitest/internal/domain/orders"
	"apitest/internal/domain/user"
	"context"
	uuid "github.com/satori/go.uuid"
	"time"
)

var (
	availableDistributor  = []string{"oca","andreani","correoargentino"}
)
type Repository interface {
	GetById(ctx context.Context,id string) (*model.Shipment,error)
	Save(ctx context.Context,shipment *model.Shipment) error
}

type Service struct {
	repo Repository
	orderService orders.Service
	userService user.Service
}

func NewService(repository Repository,orderService orders.Service) *Service {
	return &Service{
		repo: repository,
		orderService: orderService,
	}
}


func (s *Service) CreateShipping(ctx context.Context,shipping *model.Shipment) error {
	shipping.DateCreated = time.Now()
	u2:= uuid.NewV4()
	shipping.Id = u2.String()
	err := s.validateShipping(ctx,shipping)
	if err != nil {
		return err
	}
	err = s.orderService.ValidateOrder(ctx,shipping)
	if err != nil {
		return err
	}
	err = s.userService.ValidateUser(ctx,shipping)
	if err != nil {
		return err
	}
	return s.repo.Save(ctx,shipping)
}

func (s *Service) validateShipping(ctx context.Context,shipping *model.Shipment) error{
	//Validate min cap
	if shipping.Cost < 100 {
		return errors.NewBusinessError("invalid cost")
	}

	existDistributor := false
	for _, distributor := range availableDistributor {
		if distributor == shipping.Distributor{
			existDistributor = true
		}
	}
	if !existDistributor {
		return errors.NewBusinessError("invalid distributor")
	}
	return nil

}

/*
	errorChan := make(chan error)
	go func() {
		go func() {
			err := s.validateShipping(ctx,shipping)
			if err != nil {
				errorChan  <- err
			}
		}()
		go func() {
			err := s.orderService.ValidateOrder(ctx,shipping)
			if err != nil {
				errorChan  <-  err
			}
		}()
		go func() {
			err := s.userService.ValidateUser(ctx,shipping)
			if err != nil {
				errorChan  <- err
			}
		}()
		close(errorChan)
	}()


	select {
		case err,open :=  <- errorChan:
			if open {
				return err
			}else {
				break
			}
	}

 */