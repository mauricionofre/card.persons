package service

import (
	"github.com/mauricionofre/person-api/internal/model"
	"github.com/mauricionofre/person-api/internal/repository"
	"github.com/mauricionofre/person-api/pkg/rabbitmq"
)

type PersonService struct {
	repo     *repository.PersonRepository
	rabbitMQ *rabbitmq.RabbitMQ
}

func NewPersonService(repo *repository.PersonRepository, rabbitMQ *rabbitmq.RabbitMQ) *PersonService {
	return &PersonService{repo: repo, rabbitMQ: rabbitMQ}
}

func (s *PersonService) Create(p *model.Person) error {
	err := s.repo.Create(p)
	if err != nil {
		return err
	}
	return s.rabbitMQ.PublishEvent("person.created", p)
}

func (s *PersonService) GetByID(id int) (*model.Person, error) {
	return s.repo.GetByID(id)
}

func (s *PersonService) Update(p *model.Person) error {
	err := s.repo.Update(p)
	if err != nil {
		return err
	}
	return s.rabbitMQ.PublishEvent("person.updated", p)
}

func (s *PersonService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return s.rabbitMQ.PublishEvent("person.deleted", id)
}
