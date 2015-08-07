package crowbar

import (
	"errors"
	"log"
	"strconv"
)

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"username"`
	Email     string `json:"email"`
	Admin     bool   `json:"is_admin"`
	Locked    bool   `json:"locked"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	lastJson  []byte
}

func (o *User) Id() string {
	if o.ID != 0 {
		return strconv.FormatInt(o.ID, 10)
	} else if o.Name != "" {
		return o.Name
	} else {
		log.Panic("User has no ID or name")
		return ""
	}
}

func (o *User) SetId(s string) error {
	if o.ID != 0 || o.Name != "" {
		return errors.New("SetId can only be used on an un-IDed object")
	}
	if id, err := strconv.ParseInt(s, 10, 64); err == nil {
		o.ID = id
	} else {
		o.Name = s
	}
	return nil
}

func (o *User) ApiName() string {
	return "users"
}

func (o *User) setLastJSON(b []byte) {
	o.lastJson = make([]byte, len(b))
	copy(o.lastJson, b)
}

func (o *User) lastJSON() []byte {
	return o.lastJson
}

func Users() (res []*User, err error) {
	res = make([]*User, 0)
	return res, session.list(&res, "users")
}
