package user

import (
	"tutorial_bwa_startup/global"
)

type User struct {
	global.BaseModel
	Name           string
	Email          string
	PasswordHash   string
	Occupation     string
	AvatarFileName string
}
