package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/usmonzodasomon/time-tracker/internal/model"
	"strings"
	"time"
)

type UserRepoI interface {
	GetAllUsers(filter model.UserFilter) ([]model.User, error)
	GetUserTimeSpent(userID int, startPeriod, endPeriod time.Time) ([]model.TaskTimeSpent, error)

	GetUser(id int) (model.User, error)
	CreateUser(user model.User) (int, error)
	UpdateUser(user model.User) error
	DeleteUser(id int) error
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetAllUsers(filter model.UserFilter) ([]model.User, error) {
	q := `SELECT id, passport_serie, passport_number, name, surname, patronymic, address FROM users WHERE is_deleted = false`

	var conditions []string
	var args []interface{}
	argId := 1

	if filter.ID != nil {
		conditions = append(conditions, fmt.Sprintf("id = $%d", argId))
		args = append(args, *filter.ID)
		argId++
	}
	if filter.PassportSerie != nil {
		conditions = append(conditions, fmt.Sprintf("passport_serie = $%d", argId))
		args = append(args, *filter.PassportSerie)
		argId++
	}
	if filter.PassportNumber != nil {
		conditions = append(conditions, fmt.Sprintf("passport_number = $%d", argId))
		args = append(args, *filter.PassportNumber)
		argId++
	}
	if filter.Name != nil {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", argId))
		args = append(args, "%"+*filter.Name+"%")
		argId++
	}
	if filter.Surname != nil {
		conditions = append(conditions, fmt.Sprintf("surname ILIKE $%d", argId))
		args = append(args, "%"+*filter.Surname+"%")
		argId++
	}
	if filter.Patronymic != nil {
		conditions = append(conditions, fmt.Sprintf("patronymic ILIKE $%d", argId))
		args = append(args, "%"+*filter.Patronymic+"%")
		argId++
	}
	if filter.Address != nil {
		conditions = append(conditions, fmt.Sprintf("address ILIKE $%d", argId))
		args = append(args, "%"+*filter.Address+"%")
		argId++
	}

	if len(conditions) > 0 {
		q += " AND " + strings.Join(conditions, " AND ")
	}

	q += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argId, argId+1)
	args = append(args, filter.PerPage, (filter.Page-1)*filter.PerPage)

	var users []model.User
	if err := r.db.Select(&users, q, args...); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) GetUser(id int) (model.User, error) {
	q := `SELECT id, passport_serie, passport_number, name, surname, patronymic, address FROM users 
    	WHERE id = $1 AND is_deleted = false`
	user := model.User{}
	if err := r.db.Get(&user, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, model.ErrUserNotFound
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepo) GetUserTimeSpent(userID int, startPeriod, endPeriod time.Time) ([]model.TaskTimeSpent, error) {
	q := `
        SELECT 
            t.id AS task_id, 
            SUM(EXTRACT(EPOCH FROM (COALESCE(te.end_time, CURRENT_TIMESTAMP) - te.start_time))) / 60 AS total_minutes
        FROM 
            users u
        JOIN 
            tasks t ON u.id = t.user_id
        LEFT JOIN 
            time_entries te ON t.id = te.task_id
        WHERE 
            u.id = $1
            AND te.start_time >= $2
            AND (te.end_time <= $3 OR te.end_time IS NULL)
            AND u.is_deleted = false
        GROUP BY 
            t.id
        ORDER BY 
            total_minutes DESC;
    `
	var tasks []model.TaskTimeSpent
	if err := r.db.Select(&tasks, q, userID, startPeriod, endPeriod); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *UserRepo) CreateUser(user model.User) (int, error) {
	q := `INSERT INTO users
    (passport_serie, passport_number, name, surname, patronymic, address)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	if err := r.db.QueryRowx(q, user.PassportSerie, user.PassportNumber, user.Name, user.Surname, user.Patronymic, user.Address).
		Scan(&user.ID); err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UserRepo) UpdateUser(user model.User) error {
	q := `UPDATE users SET passport_serie = $1, passport_number = $2, name = $3, surname = $4, patronymic = $5, address = $6 WHERE id = $7`
	res, err := r.db.Exec(q, user.PassportSerie, user.PassportNumber, user.Name, user.Surname, user.Patronymic, user.Address, user.ID)
	if err != nil {
		return err
	}
	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return model.ErrUserNotFound
	}
	return nil
}

func (r *UserRepo) DeleteUser(taskID int) error {
	q := `UPDATE users SET is_deleted = true WHERE id = $1`
	res, err := r.db.Exec(q, taskID)
	if err != nil {
		return err
	}
	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return model.ErrUserNotFound
	}
	return nil
}
