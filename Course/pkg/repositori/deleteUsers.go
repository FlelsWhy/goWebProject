package repositori

import "context"

func (repo *PGRepo) DeleteUsers(id int) (err error) {
	_, err = repo.pool.Exec(context.Background(), ` DELETE FROM users
	WHERE id = $1;`, id)
	if err != nil {
		return err
	}
	return err
}
