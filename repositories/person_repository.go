package repositories

import (
	"health-history/config"
	"health-history/models"
)

type PersonRepository struct{}

func (r *PersonRepository) CreatePerson(person *models.Person) error {
	result := config.DB.Create(person)
	return result.Error
}

func (r *PersonRepository) UpdatePerson(person *models.Person) error {
	result := config.DB.Save(person)
	return result.Error
}

func (r *PersonRepository) DeletePerson(id int) error {
	result := config.DB.Delete(&models.Person{}, id)
	return result.Error
}

func (r *PersonRepository) GetPersonById(id int) (*models.Person, error) {
	var person models.Person
	result := config.DB.First(&person, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &person, nil
}

func (r *PersonRepository) GetAllPersons() ([]models.Person, error) {
	var persons []models.Person
	result := config.DB.Find(&persons)
	return persons, result.Error
}
