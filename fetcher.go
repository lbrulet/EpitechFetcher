package EpitechFetcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lbrulet/EpitechFetcher/models/user"

	"github.com/lbrulet/EpitechFetcher/models"
)

//FetchIntranetUser going to fetch the profil of the user
func FetchIntranetUser(userInfo models.EpiUser) (user.UserIntra, error) {
	client := &http.Client{}
	if req, err := http.NewRequest("GET", "https://intra.epitech.eu/"+userInfo.AutoLogin+"/user?format=json", nil); err != nil {
		return user.UserIntra{}, err
	} else {
		response, err := client.Do(req)
		if err != nil {
			return user.UserIntra{}, err
		} else {
			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return user.UserIntra{}, err
			} else {
				usr := user.NewUserIntra()
				if err := json.Unmarshal(data, &usr); err != nil {
					return user.UserIntra{}, err
				} else {
					if err := response.Body.Close(); err != nil {
						return user.UserIntra{}, err
					} else {
						return usr, nil
					}
				}
			}
		}
	}
}

//FetchIntranetNote going to fetch note of the user
func FetchIntranetNote(userInfo models.EpiUser) (user.CursusIntra, error) {
	client := &http.Client{}
	if req, err := http.NewRequest("GET", "https://intra.epitech.eu/"+userInfo.AutoLogin+"/user/"+userInfo.Login+"/notes/?format=json", nil); err != nil {
		return user.CursusIntra{}, err
	} else {
		response, err := client.Do(req)
		if err != nil {
			return user.CursusIntra{}, err
		} else {
			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return user.CursusIntra{}, err
			} else {
				cursus := user.NewNoteIntra()
				if err := json.Unmarshal(data, &cursus); err != nil {
					return user.CursusIntra{}, err
				} else {
					if err := response.Body.Close(); err != nil {
						return user.CursusIntra{}, err
					} else {
						return cursus, nil
					}
				}
			}
		}
	}
}
