package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tamada/morningglory/common"
)

type KeyPhrase struct {
	KeyPhrase string `json:"token"`
}

func (key *KeyPhrase) Get() string {
	return common.Md5sum(key.KeyPhrase)
}

func readBody(context *gin.Context) *KeyPhrase {
	var keyphrase = KeyPhrase{}
	context.Bind(&keyphrase)
	return &keyphrase
}

func RegisterUser(context *gin.Context) error {
	var userName = context.Param("userName")
	if userName == "" {
		return fmt.Errorf("user name not found")
	}
	var keyphrase = readBody(context)
	return common.RegisterUser(userName, keyphrase.Get())
}

func UpdateKeyPhrase(context *gin.Context) error {
	if err := Authenticate(context); err != nil {
		return err
	}
	var userName = context.Param("userName")
	if userName == "" {
		return fmt.Errorf("user name not found")
	}
	var keyphrase = readBody(context)
	return common.UpdateKeyPhrase(userName, keyphrase.Get())
}

func DeleteUser(context *gin.Context) error {
	if err := Authenticate(context); err != nil {
		return err
	}
	common.DeleteUser(context.Param("userName"))
	return nil
}

func FindUser(context *gin.Context) (*common.User, error) {
	var userName = context.Param("userName")
	if userName == "" {
		return nil, fmt.Errorf("user name not found")
	}
	var user = common.FindUser(userName)
	if user == nil {
		return nil, fmt.Errorf("%s: user not found", userName)
	}
	return user, nil
}

/*
Authenticate performs authentication from given http request.
*/
func Authenticate(context *gin.Context) error {
	var userName = context.Param("userName")
	if userName == "" {
		return fmt.Errorf("user name not found")
	}
	var keyPhrase = context.GetHeader("X-USER-TOKEN")
	if keyPhrase == "" {
		return fmt.Errorf("%s: key phrase not found", userName)
	}
	var md5Hash = common.Md5sum(keyPhrase)
	if err := common.Authenticate(userName, md5Hash); err != nil {
		return err
	}
	return nil
}
