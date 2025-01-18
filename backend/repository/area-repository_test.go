package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"area/repository"
	"area/schemas"
	"area/test"
)

func TestSaveArea(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewAreaRepository(db)
	area := schemas.Area{
		ActionRefreshRate: 10,
	}

	areaID, err := repo.SaveArea(area)
	assert.NoError(t, err)
	assert.NotZero(t, areaID)
}

func TestSave(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewAreaRepository(db)
	area := schemas.Area{
		ActionRefreshRate: 10,
	}

	err = repo.Save(area)
	assert.NoError(t, err)
}

// func TestUpdate(t *testing.T) {
// 	db, err := test.SetupTestDB()
// 	assert.NoError(t, err)

// 	repo := repository.NewAreaRepository(db)
// 	area := schemas.Area{
// 		ActionRefreshRate: 10,
// 	}
// 	err = repo.Save(area)
// 	assert.NoError(t, err)

// 	area.ActionRefreshRate = 20
// 	err = repo.Update(area)
// 	assert.NoError(t, err)

// 	var updatedArea schemas.Area
// 	db.First(&updatedArea, area.Id)
// 	assert.Equal(t, area.ActionRefreshRate, updatedArea.ActionRefreshRate)
// }

// func TestDelete(t *testing.T) {
// 	db, err := test.SetupTestDB()
// 	assert.NoError(t, err)

// 	repo := repository.NewAreaRepository(db)
// 	area := schemas.Area{
// 		ActionRefreshRate: 10,
// 	}
// 	err = repo.Save(area)
// 	assert.NoError(t, err)

// 	err = repo.Delete(area)
// 	assert.NoError(t, err)

// 	var deletedArea schemas.Area
// 	result := db.First(&deletedArea, area.Id)
// 	assert.Error(t, result.Error)
// }

// func TestFindAll(t *testing.T) {
// 	db, err := test.SetupTestDB()
// 	assert.NoError(t, err)

// 	repo := repository.NewAreaRepository(db)
// 	area1 := schemas.Area{
// 		ActionRefreshRate: 10,
// 	}
// 	area2 := schemas.Area{
// 		ActionRefreshRate: 20,
// 	}
// 	err = repo.Save(area1)
// 	assert.NoError(t, err)
// 	err = repo.Save(area2)
// 	assert.NoError(t, err)

// 	areas, err := repo.FindAll()
// 	assert.NoError(t, err)
// 	assert.Len(t, areas, 2)
// }

func TestFindByUserId(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewAreaRepository(db)
	user := schemas.User{
		Username: "Test User",
	}
	db.Create(&user)

	area := schemas.Area{
		UserId:            user.Id,
		ActionRefreshRate: 10,
	}
	err = repo.Save(area)
	assert.NoError(t, err)

	areas, err := repo.FindByUserId(user.Id)
	assert.NoError(t, err)
	assert.Len(t, areas, 1)
	assert.Equal(t, user.Id, areas[0].UserId)
}

func TestAreaFindById(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewAreaRepository(db)
	area := schemas.Area{
		ActionRefreshRate: 10,
	}
	err = repo.Save(area)
	assert.NoError(t, err)

	foundArea, err := repo.FindById(area.Id)
	assert.NoError(t, err)
	assert.Equal(t, area.Id, foundArea.UserId)
}
