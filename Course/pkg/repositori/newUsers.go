package repositori

import (
	"COURSE/pkg/models"
	"context"
)

func (repo *PGRepo) NewUsers(NewUsers models.Users) error {
	_, err := repo.pool.Exec(context.Background(), `INSERT INTO public.users(
	id, login, email)
	VALUES ($1, $2, $3);`, NewUsers.Id, NewUsers.Login, NewUsers.Email)
	if err != nil {
		return err
	}
	return err
}
