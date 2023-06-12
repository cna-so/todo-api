package models

// import (
// 	"net/http"
//
// 	"github.com/cna-so/todo-api/initializers"
// 	"github.com/cna-so/todo-api/models/db"
// )

// func (u *db.UserModels) CreateUserAccount() (status int, err error) {
// initializers.Db.Create(u)
// _, err = initializers.Db.Exec(`INSERT INTO users
// (user_uid ,first_name ,last_name , email ,password, role)
// VALUES ($1,$2,$3,$4,$5,$6)`,
// 	uuid.New(), u.FirstName, u.LastName, u.Email, u.Password, "U") // role is defualt to user we have S as superuser and A as admin
//  initializers.Db.Create(u)
// if err != nil {
// 	return http.StatusForbidden, err
// }
// return http.StatusOK, nil
// }

// func (u *UserModels) GetUserAccount() (UserModels, error) {
// row := initializers.Db.QueryRow("SELECT user_uid ,first_name ,last_name , email ,password, role FROM users where email=$1", u.Email)
// if row.Err() != nil {
// 	return UserModels{}, row.Err()
// }
// var user UserModels
// row.Scan(&user.UserUID, &user.FirstName, &user.LastName, &user.Email, &user.Password, user.Role)
// return user, nil
// }
