package EpitechFetcher

import (
	"github.com/lbrulet/EpitechFetcher/models"
	"github.com/lbrulet/EpitechFetcher/models/user"
)

//Link going to fetch all information about the user
func Link(autoLogin, login string) (*user.UserIntra, *user.CursusIntra, error) {
	var User models.EpiUser
	User.AutoLogin = autoLogin
	User.Login = login
	if intra, err := FetchIntranetUser(User); err != nil {
		return nil, nil, err
	} else {
		if cursus, err := FetchIntranetNote(User); err != nil {
			return nil, nil, err
		} else {
			return &intra, &cursus, nil
		}
	}
}
