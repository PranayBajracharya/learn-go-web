package repositories

import (
	"database/sql"
	"log"

	"learn-go/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user models.User) int64 {
	query := `
	INSERT INTO users (name, email)
	VALUES (?, ?)`

	result, err := r.db.Exec(query, user.Name, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func (r *UserRepository) List() []models.User {
	var users []models.User
	query := `
	SELECT id, name, email, created_at
	FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func (r *UserRepository) Get(id int64) *models.User {
	var user models.User
	query := `
	SELECT id, name, email, created_at
	FROM users
	WHERE id = ?`

	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Fatal(err)
	}

	return &user
}

func (r *UserRepository) Update(id int64, user models.User) {
	query := `
	UPDATE users
	SET name = ?, email = ?
	WHERE id = ?`

	_, err := r.db.Exec(query, user.Name, user.Email, id)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *UserRepository) Delete(id int64) {
	query := `
	DELETE FROM users
	WHERE id = ?`

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
}
