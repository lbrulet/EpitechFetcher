package EpitechFetcher

import (
	"fmt"
	"github.com/lbrulet/EpitechFetcher/models"
)

func Link(autoLogin, login string) {
	var User models.EpiUser
	User.AutoLogin = autoLogin
	User.Login = login
	err := FetchIntranet(User)
	fmt.Println(err)
}
