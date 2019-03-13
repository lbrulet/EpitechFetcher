package EpitechFetcher

import (
	"fmt"
	"github.com/lbrulet/EpitechFetcher/models"
)

func Link(autoLogin, login string) {
	var User models.EpiUser
	User.AutoLogin = autoLogin
	User.Login = login
	if intra, err := FetchIntranetUser(User); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(intra)
	}
	if cursus, err := FetchIntranetNote(User); err != nil {
		fmt.Println(err)
	} else {
		for _, elem := range cursus.Modules {
			fmt.Println(elem)
		}
	}
}
