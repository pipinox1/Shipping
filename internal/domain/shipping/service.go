package shipping

import (
	"apitest/internal/domain/model"
	"apitest/internal/domain/orders"
	"apitest/internal/domain/user"
	"apitest/internal/tools/errors"
	"context"
	uuid "github.com/satori/go.uuid"
	"time"
)

var (
	availableDistributor = []string{"oca", "andreani", "correoargentino"}
)

type Service interface {
	CreateShipping(ctx context.Context, shipping *model.Shipment) error
	GetShipping(ctx context.Context, id string) (*model.Shipment, error)
}

type Repository interface {
	GetById(ctx context.Context, id string) (*model.Shipment, error)
	Save(ctx context.Context, shipment *model.Shipment) error
}

type service struct {
	repo         Repository
	orderService orders.Service
	userService  user.Service
}

func NewService(repository Repository, orderService orders.Service, userService user.Service) Service {
	return &service{
		repo:         repository,
		orderService: orderService,
		userService:  userService,
	}
}

func (s *service) GetShipping(ctx context.Context, id string) (*model.Shipment, error) {
	shipment, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errCast, ok := err.(*errors.RestClientError); ok {
			if errCast.StatusCode == 404 {
				return nil, errors.NotFound
			}
		}
		return nil, err
	}
	return shipment, nil
}

func (s *service) CreateShipping(ctx context.Context, shipping *model.Shipment) error {
	shipping.DateCreated = time.Now()
	u2 := uuid.NewV4()
	shipping.Id = u2.String()
	err := s.validateShipping(ctx, shipping)
	if err != nil {
		return err
	}
	err = s.orderService.ValidateOrder(ctx, shipping)
	if err != nil {
		return err
	}
	err = s.userService.ValidateUser(ctx, shipping)
	if err != nil {
		return err
	}
	return s.repo.Save(ctx, shipping)
}

func (s *service) validateShipping(ctx context.Context, shipping *model.Shipment) error {
	//Validate min cap
	if shipping.Cost < 100 {
		return errors.NewBusinessError("invalid cost")
	}

	existDistributor := false
	for _, distributor := range availableDistributor {
		if distributor == shipping.Distributor {
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
