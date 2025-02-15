package repositori

import (
	"COURSE/pkg/models"
	"context"
)

func (repo *PGRepo) GetAllUser() (data []models.Users, err error) {
	row, err := repo.pool.Query(context.Background(), `SELECT id, login, data_of_register, email
	FROM public.users;`)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var u models.Users
		err := row.Scan(
			&u.Id,
			&u.Login,
			&u.DataOfRegister,
			&u.Email,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, u)
	}

	return data, err
}
