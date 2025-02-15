package repositori

import (
	"COURSE/pkg/models"
	"context"
)

func (repo *PGRepo) GetUsersFromID(id int) (data models.Users, err error) {
	err = repo.pool.QueryRow(context.Background(), `SELECT id, login, data_of_register, email
	FROM public.users WHERE id = $1`, id).Scan(&data.Id, &data.Login, &data.DataOfRegister, &data.Email)
	if err != nil {
		return data, err
	}
	return data, err
}
