package users

import (
	"../util"
	"errors"
	"github.com/jameskeane/bcrypt"
)

type User struct {
	Id     int64
	Email  string
	PwHash []byte
	saved  bool
}

var ErrUserDoesNotExist = errors.New("User does not exist")

func scanUser(row *sql.Row) (u *User, err error) {
	u = &User{}
	err = row.Scan(&u.Id, &u.Email, &u.PwHash)
	if err != nil {
		return nil, err
	}
	u.saved = true
	return u, nil
}

func GetByEmail(db util.DB, email string) (u *User, err error) {
	row := db.QueryRow("SELECT id, email, pw_hash FROM users WHERE email = $1;", email)
	return scanUser(row)
}

func GetById(db util.DB, id int64) (u *User, err error) {
	row := db.QueryRow("SELECT id, email, pw_hash FROM users WHERE id = $1;", id)
	return scanUser(row)
}

func (u *User) String() string {
	return fmt.Sprintf("User{id: %v, email:%v, pw_hash:%v, saved: %v}", u.Id, u.Email, u.PwHash, u.saved)
}

func New(email, password string) (*User, error) {
	if !isValidEmail(email) {
		return nil, errors.New("Invalid email format")
	}
	if !isValidPassword(password) {
		return nil, errors.New("Password must be at least 8 characters long and include at least one number")
	}

	u := &User{Email: email}
	hash, err := bcrypt.HashBytes([]byte(password))
	if err != nil {
		return nil, err
	}
	u.PwHash = hash
	return u, nil
}

func (u *User) Verify(password string) bool {
	return bcrypt.MatchBytes([]byte(password), u.PwHash)
}

func (u *User) saveNew(db util.DB) error {
	_, err := db.Exec("INSERT INTO users (email, pw_hash) VALUES ($1, $2);", u.Email, u.PwHash)
	if err != nil {
		return err
	}
	u2, err := GetByEmail(db, u.Email)
	if err != nil {
		return err
	}
	u.saved = true
	u.Id = u2.Id
	return nil
}

func (u *User) update(db util.DB) error {
	_, err := db.Exec("UPDATE users SET email=$1, pw_hash=$2 WHERE id=$3", u.Email, u.PwHash, u.Id)
	return err
}

func (u *User) Save(db util.DB) error {
	if u.saved {
		return u.update(db)
	}
	return u.saveNew(db)
}

func (u *User) Delete(db util.DB) error {
    _, err := db.Exec("DELETE FROM users WHERE id=$1", u.Id)
    return err
}

func (u *User) UpdatePassword(db util.DB, newPassword string) error {
    hash, err := bcrypt.HashBytes([]byte(newPassword))
    if err != nil {
        return err
    }
    u.PwHash = hash
    return u.update(db)
}

func GetAll(db util.DB) ([]*User, error) {
    rows, err := db.Query("SELECT id, email, pw_hash FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    users := []*User{}
    for rows.Next() {
        user, err := scanUser(rows)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    
    return users, nil
}

func isValidEmail(email string) bool {
	// Basic regex to check email format
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	// Example: Require at least 8 characters, including at least one number
	re := regexp.MustCompile(`^(?=.*[0-9]).{8,}$`)
	return re.MatchString(password)
}
