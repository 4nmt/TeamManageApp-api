package models

import (
	"errors"
	"time"

	"github.com/4nmt/TeamManageApp-api/db"
	"github.com/4nmt/TeamManageApp-api/forms"
)

//UserProject ...
type UserProject struct {
	ID        int   `db:"id, primarykey, autoincrement" json:"id"`
	UserID    int   `db:"user_id" json:"user_id"`
	ProjectID int   `db:"project_id" json:"project_id"`
	UpdatedAt int64 `db:"updated_at" json:"updated_at"`
	CreatedAt int64 `db:"created_at" json:"created_at"`
}

//UserProjectModel ...
type UserProjectModel struct{}

func (m UserProjectModel) Assign(form forms.UserProjectForm) (userProjectID int, err error) {
	getDb := db.GetDB()

	checkUser, err := getDb.SelectInt("SELECT count(id) FROM public.users WHERE id=$1 LIMIT 1", form.UserID)
	if err != nil {
		return -1, err
	}

	if checkUser <= 0 {
		return -1, errors.New("user doesn't exists")
	}

	checkProject, err := getDb.SelectInt("SELECT count(id) FROM public.projects WHERE id=$1 LIMIT 1", form.ProjectID)
	if err != nil {
		return -1, err
	}
	if checkProject <= 0 {
		return -1, errors.New("project doesn't exists")
	}

	err = getDb.QueryRow("INSERT INTO public.user_project(user_id, project_id, updated_at, created_at) VALUES($1, $2, $3, $4) RETURNING id", form.UserID, form.ProjectID, time.Now().Unix(), time.Now().Unix()).Scan(&userProjectID)

	if err != nil {
		return -1, errors.New("Not registered")
	}

	return userProjectID, nil
}

//One ...
func (m UserProjectModel) One(userProjectID int) (userProject UserProject, err error) {
	err = db.GetDB().SelectOne(&userProject, "SELECT id FROM public.user_project WHERE id=$1 LIMIT 1", userProjectID)
	return userProject, err
}

func (m UserProjectModel) Remove(userProjectID int) (err error) {
	getDb := db.GetDB()
	_, err = m.One(userProjectID)

	if err != nil {
		return errors.New("user_project not found")
	}

	_, err = getDb.Exec("DELETE FROM public.user_project WHERE id=$1", userProjectID)

	return err
}
