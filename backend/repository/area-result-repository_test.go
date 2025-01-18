package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"area/repository"
	"area/schemas"
	"area/test"
)

func TestAreaResult_NewAreaResultRepository(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewAreaResultRepository(db)
	assert.NotNil(t, repo)
}

func TestAreaResult_Save(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewAreaResultRepository(db)
	action := schemas.AreaResult{AreaId: 1, Result: "Test Result"}

	assert.NotPanics(t, func() {
		repo.Save(action)
	})

	var result schemas.AreaResult
	db.First(&result, action.AreaId)
	assert.Equal(t, action.Result, result.Result)
}

// func TestAreaResult_Update(t *testing.T) {
// 	db, err := test.SetupTestDB()
// 	assert.NoError(t, err)

// 	repo := repository.NewAreaResultRepository(db)
// 	action := schemas.AreaResult{AreaId: 1, Result: "Test Result"}
// 	repo.Save(action)

// 	action.Result = "Updated Result"
// 	assert.NotPanics(t, func() {
// 		repo.Update(action)
// 	})

// 	var result schemas.AreaResult
// 	db.First(&result, action.AreaId)
// 	assert.Equal(t, action.Result, result.Result)
// }

// func TestAreaResult_Delete(t *testing.T) {
// 	db, err := test.SetupTestDB()
// 	assert.NoError(t, err)

// 	repo := repository.NewAreaResultRepository(db)
// 	action := schemas.AreaResult{AreaId: 1, Result: "Test Result"}
// 	repo.Save(action)

// 	assert.NotPanics(t, func() {
// 		repo.Delete(action)
// 	})

// 	var result schemas.AreaResult
// 	db.First(&result, action.AreaId)
// 	assert.Equal(t, uint64(0), result.AreaId)
// }

// func TestAreaResult_FindAll(t *testing.T) {
// 	db, err := test.SetupTestDB()
// 	assert.NoError(t, err)

// 	repo := repository.NewAreaResultRepository(db)
// 	action1 := schemas.AreaResult{AreaId: 1, Result: "Test Result 1"}
// 	action2 := schemas.AreaResult{AreaId: 2, Result: "Test Result 2"}
// 	repo.Save(action1)
// 	repo.Save(action2)

// 	results := repo.FindAll()
// 	assert.Len(t, results, 2)
// }

func TestAreaResult_FindByAreaId(t *testing.T) {
	db, err := test.SetupTestDB()
	assert.NoError(t, err)

	repo := repository.NewAreaResultRepository(db)
	action := schemas.AreaResult{AreaId: 1, Result: "Test Result"}
	repo.Save(action)

	results := repo.FindByAreaId(1)
	assert.Len(t, results, 1)
	assert.Equal(t, action.Result, results[0].Result)
}
