package repository

import (
	"testing"
)

func TestBlogsRepo_Create(t *testing.T) {
	// t.Parallel()

	// db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	// require.NoError(t, err)
	// defer db.Close()

	// sqlxDB := sqlx.NewDb(db, "sqlmock")
	// defer sqlxDB.Close()

	// commRepo := NewActorsRepository(sqlxDB)

	// t.Run("Create", func(t *testing.T) {
	// 	newsUID := uuid.New()
	// 	title := "title"

	// 	rows := sqlmock.NewRows([]string{"id", "title"}).AddRow(newsUID, title)

	// 	blog := &models.Blog{
	// 		ID:        newsUID,
	// 		Title:     title,
	// 		CreatedAt: time.Now(),
	// 	}

	// 	mock.ExpectQuery("").WithArgs(blog.ID, blog.Title).WillReturnRows(rows)

	// 	createdBlog, err := commRepo.Create(context.Background(), blog)

	// 	require.NoError(t, err)
	// 	require.NotNil(t, createdBlog)
	// 	// require.Equal(t, createdBlog, blog)
	// })

	// t.Run("Create ERR", func(t *testing.T) {
	// 	newsUID := uuid.New()
	// 	title := "title"
	// 	createErr := errors.New("Create blog error")

	// 	blog := &models.Blog{
	// 		ID:    newsUID,
	// 		Title: title,
	// 	}

	// 	mock.ExpectQuery("").WithArgs(blog.ID, blog.Title).WillReturnError(createErr)

	// 	createdToDo, err := commRepo.Create(context.Background(), blog)

	// 	require.Nil(t, createdToDo)
	// 	require.NotNil(t, err)
	// })
}

func TestBlogsRepo_Update(t *testing.T) {

}

func TestBlogsRepo_Delete(t *testing.T) {

}
