package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"area/repository"
	"area/schemas"
	"area/test"
)

func TestActionSave(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewActionRepository(db)
	action := schemas.Action{Name: "Test Action"}

	err = repo.Save(action)
	assert.NoError(t, err)

	var savedAction schemas.Action
	db.First(&savedAction, "name = ?", "Test Action")
	assert.Equal(t, "Test Action", savedAction.Name)
}

func TestActionUpdate(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewActionRepository(db)
	action := schemas.Action{Name: "Test Action"}
	db.Create(&action)

	action.Name = "Updated Action"
	err = repo.Update(action)
	assert.NoError(t, err)

	var updatedAction schemas.Action
	db.First(&updatedAction, action.Id)
	assert.Equal(t, "Updated Action", updatedAction.Name)
}

func TestActionDelete(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewActionRepository(db)
	action := schemas.Action{Name: "Test Action"}
	db.Create(&action)

	err = repo.Delete(action)
	assert.NoError(t, err)

	var deletedAction schemas.Action
	result := db.First(&deletedAction, action.Id)
	assert.Error(t, result.Error)
}

func TestActionFindAll(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewActionRepository(db)
	actions := []schemas.Action{
		{Name: "Action 1"},
		{Name: "Action 2"},
	}
	for _, action := range actions {
		db.Create(&action)
	}

	foundActions, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, foundActions, 2)
}

func TestActionFindByName(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewActionRepository(db)
	action := schemas.Action{Name: "Test Action"}
	db.Create(&action)

	foundActions, err := repo.FindByName("Test Action")
	assert.NoError(t, err)
	assert.Len(t, foundActions, 1)
	assert.Equal(t, "Test Action", foundActions[0].Name)
}

func TestActionFindByServiceId(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewActionRepository(db)
	service := schemas.Service{Name: "Test Service"}
	db.Create(&service)

	action := schemas.Action{Name: "Test Action", ServiceId: service.Id}
	db.Create(&action)

	foundActions, err := repo.FindByServiceId(service.Id)
	assert.NoError(t, err)
	assert.Len(t, foundActions, 1)
	assert.Equal(t, "Test Action", foundActions[0].Name)
}

func TestActionFindByServiceByName(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewActionRepository(db)
	service := schemas.Service{Name: "Test Service"}
	db.Create(&service)

	action := schemas.Action{Name: "Test Action", ServiceId: service.Id}
	db.Create(&action)

	foundActions, err := repo.FindByServiceByName(service.Id, "Test Action")
	assert.NoError(t, err)
	assert.Len(t, foundActions, 1)
	assert.Equal(t, "Test Action", foundActions[0].Name)
}

func TestActionFindById(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewActionRepository(db)
	action := schemas.Action{Name: "Test Action"}
	db.Create(&action)

	foundAction, err := repo.FindById(action.Id)
	assert.NoError(t, err)
	assert.Equal(t, "Test Action", foundAction.Name)
}
