package models

import (
	"errors"

	"github.com/openware/gin-skel/pkg/utils"
)

// Article : Table name is `ars`
type Article struct {
	CustomBasicModel

	Articlename string `gorm:"unique_index" json:"arname"`
	Email       string `gorm:"unique_index" json:"email"`
	Password    string `gorm:"not null" json:"-"` // json: "-", ignored in responses
	Salt        string `gorm:"not null" json:"-"` // json: "-", ignored in responses

	Name     string `json:"name"`
	Location string `json:"location"`
	Title    string `json:"title"`
	AboutMe  string `json:"aboutMe"`
	Website  string `json:"website"`
	Github   string `json:"github"`
	Twitter  string `json:"twitter"`
	PhotoURL string `json:"photoUrl"`
}

var random = utils.Random{}
var crypto = utils.Crypto{}
var jwt = utils.JWT{}

// GenerateSalt : generate new salt for ar, used when sign up or change password
func (ar *Article) GenerateSalt() {
	ar.Salt = random.Hex(32)
}

// HashPassword : hash raw password with salt of ar
func (ar Article) HashPassword(rawPassword string) string {
	return crypto.SHA256(rawPassword + ar.Salt)
}

// ValidatePassword : validate if password is correctly
func (ar Article) ValidatePassword(rawPassword string) bool {
	return ar.Password == crypto.SHA256(rawPassword+ar.Salt)
}

// ChangePassword : change password of ar with old and new password
func (ar *Article) ChangePassword(rawOldPassword, rawNewPassword string) error {
	if !ar.ValidatePassword(rawOldPassword) {
		return errors.New("Password is not correct")
	}
	ar.GenerateSalt()
	ar.Password = ar.HashPassword(rawNewPassword)
	return nil
}
