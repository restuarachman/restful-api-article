package repository

import (
	"tech-test/internal/article/entity"
	"time"
)

var (
	articleEntity = entity.Article{
		ID:         1,
		Author_ID:  1,
		Title:      "title",
		Body:       "body",
		Created_at: time.Now(),
	}
)

// func TestStore(t *testing.T) {
// 	mockedDB, mockObj, err := sqlmock.New()
// 	db, err := gorm.Open(mysql.Dialector{
// 		&mysql.Config{
// 			Conn:                      mockedDB,
// 			SkipInitializeWithVersion: true,
// 		},
// 	}, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	defer mockedDB.Close()

// 	mockObj.ExpectBegin()
// 	mockObj.ExpectExec(regexp.QuoteMeta("INSERT INTO articles (author_id, title, body, created_at) VALUES (?,?,?,?)")).WithArgs(1, "title", "body", time.Now()).WillReturnResult(sqlmock.NewResult(1, 1))
// 	mockObj.ExpectCommit()

// 	articleRepo := NewArticleRepo(db)
// 	err = articleRepo.Store(articleEntity)
// assert.NoError(t, err)
// 	assert.NotEmpty(t, res)

// }
