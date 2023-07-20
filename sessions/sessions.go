package sessions

import (
	"../util"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"
)

const SESSION_BYTES = 64
package sessions

import (
	"../util"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"
)

const (
	SESSION_BYTES         = 64
	insertSessionQuery    = "INSERT INTO sessions (id, user_id, created_at) VALUES ($1, $2, $3);"
)

type Session struct {
	Id        []byte
	UserId    int64
	CreatedAt time.Time
}

func New(uid int64) (*Session, error) {
	bytes := make([]byte, SESSION_BYTES)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return &Session{
		Id:        encodeBase64(bytes),
		UserId:    uid,
		CreatedAt: time.Now().UTC(), 
	}, nil
}

func encodeBase64(in []byte) []byte {
	out_bytes := base64.StdEncoding.EncodedLen(len(in))
	out := make([]byte, out_bytes)
	base64.StdEncoding.Encode(out, in)
	return out
}

func (s *Session) Save(db util.DB) error {
	_, err := db.Exec(insertSessionQuery, s.Id, s.UserId, s.CreatedAt)
	if err != nil {
		return fmt.Errorf("error while saving session to database: %v", err)
	}
	return nil
}

func (s *Session) String() string {
	return fmt.Sprintf("Session{id:'%s', user_id:%v, created_at:%v}", s.Id, s.UserId, s.CreatedAt)
}

type Session struct {
	Id        []byte
	UserId    int64
	CreatedAt time.Time
}

func New(uid int64) (s *Session, err error) {
	bytes := make([]byte, SESSION_BYTES)
	n, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	if n != len(bytes) {
		return nil, errors.New("Not enough random bytes generated")
	}
	return &Session{
		Id:     encodeBase64(bytes),
		UserId: uid}, nil
}
func encodeBase64(in []byte) []byte {
	out_bytes := base64.StdEncoding.EncodedLen(SESSION_BYTES)
	out := make([]byte, out_bytes)
	base64.StdEncoding.Encode(out, in)
	return out
}
func (s *Session) Save(db util.DB) error {
	_, err := db.Exec("INSERT INTO sessions (id, user_id) VALUES ($1, $2);", s.Id, s.UserId)
	return err
}

func (s *Session) String() string {
	return fmt.Sprintf("Session{id:'%s', user_id:%v, created_at:%v}", s.Id, s.UserId, s.CreatedAt)
}
