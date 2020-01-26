package model

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	gorm.Model `json:"-"`
	Email      string `gorm:"type:varchar(64);unique_index;not null"`
	Name       string `gorm:"type:varchar(64);unique;not null"`
	Password   string `gorm:"type:varchar(64);not null"`
}

func (user *User) GetUser() error {
	return dao.Where(user).First(user).Error
}

func (user *User) GetUserNameById(id uint) string {
	dao.Model(&User{}).Where("id = ?", id).Select("name").First(user)
	return user.Name
}

func (user *User) CreateUser() error {
	return dao.Create(user).Error
}

func (user *User) Verification(password string) bool {
	if user.Password == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}

}

func (user *User) CreateToken() (string, error) {
	now := time.Now()
	hs := jwt.NewHS256([]byte(user.Password))
	type customPayload struct {
		jwt.Payload
		Id uint `json:"id,omitempty"`
	}
	pl := customPayload{
		Payload: jwt.Payload{
			ExpirationTime: jwt.NumericDate(now.Add(time.Hour * 24 * 7)),
			IssuedAt:       jwt.NumericDate(now)},
		Id: user.ID,
	}
	token, err := jwt.Sign(pl, hs)
	//fmt.Println(string(token))
	return string(token), err
}

func (user *User) CheckToken(token string) bool {

	sep1 := strings.IndexByte(token, '.')
	if sep1 == 0 {
		return false
	}
	sep2 := strings.IndexByte(token[sep1+1:], '.')
	if sep2 == 0 {
		return false
	}

	js, err := base64.RawURLEncoding.DecodeString(token[sep1+1 : sep1+1+sep2])
	if err != nil {
		return false
	}
	type customPayload struct {
		jwt.Payload
		Id uint `json:"id,omitempty"`
	}
	pl := new(customPayload)
	if err := json.Unmarshal(js, pl); err != nil {
		return false
	}
	user.ID = pl.Id

	if err = user.GetUser(); err != nil {
		return false
	}
	now := time.Now()
	hs := jwt.NewHS256([]byte(user.Password))
	_, err = jwt.Verify([]byte(token), hs, pl)
	if err != nil {
		return false
	}
	return pl.ExpirationTime.After(now)
}
