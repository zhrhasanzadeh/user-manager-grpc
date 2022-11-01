package postgres

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"golang.org/x/net/context"
	"grpcprj/internal/model"
	"time"
)

type PostgresRepository struct {
	db *pgx.Conn
}

func NewPostgresRepository(db *pgx.Conn) *PostgresRepository {
	return &PostgresRepository{db}
}

func (p *PostgresRepository) GetUsers() string {
	rows, err := p.db.Query(context.Background(), `select * from "User"`)
	if err != nil {
		fmt.Println(err)
	}
	var allUsers []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Username, &user.Email,
			&user.MobileNo, &user.Password, &user.Birthdate,
			&user.RegisterDate, &user.LastLoginTime)
		if err != nil {
			println(err)
		}
		allUsers = append(allUsers, user)
	}
	rows.Close()
	j, _ := json.Marshal(allUsers)
	return string(j)
}

func (p *PostgresRepository) SignUp(user model.User) error {
	rows, err := p.db.Query(context.Background(),
		`insert into "User" (username, email, mobileno, pass, birthdate, registerdate, lastlogintime) VALUES
                                                ($1,$2,$3,$4,$5,$6,$7)`,
		user.Username, user.Email, user.MobileNo, user.Password, user.Birthdate, time.Now(), nil)
	rows.Close()
	return err
}

func (p *PostgresRepository) Login(UserName string, Password string) (bool, error) {
	rows, err := p.db.Query(context.Background(), `select * from "User" where username = $1 and pass = $2;`,
		UserName, Password)
	rows.Close()
	if err != nil {
		return false, err
	}
	rows, err = p.db.Query(context.Background(),
		`update "User" set lastlogintime = $1 where username = $2 and pass = $3`,
		time.Now(), UserName, Password)
	rows.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *PostgresRepository) Delete(username string) (bool, error) {
	rows, err := p.db.Query(context.Background(), `delete from "User" where username = $1`, username)
	rows.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}
