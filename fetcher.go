package EpitechFetcher

import (
	"encoding/json"
	"fmt"
	"github.com/lbrulet/EpitechFetcher/models/user"
	"io/ioutil"
	"net/http"

	"github.com/lbrulet/EpitechFetcher/models"
)

func FetchIntranetUser(userInfo models.EpiUser) (user.UserIntra, error) {
	fmt.Println(userInfo)
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
				usr := NewUserIntra()
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

func FetchIntranetNote(userInfo models.EpiUser) (user.CursusIntra, error) {
	fmt.Println(userInfo)
	client := &http.Client{}
	if req, err := http.NewRequest("GET", "https://intra.epitech.eu/"+userInfo.AutoLogin+"/user/"+ userInfo.Login+"/notes/?format=json", nil); err != nil {
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
				cursus := NewNoteIntra()
				if err := json.Unmarshal(data, &cursus); err != nil {
					return user.CursusIntra{}, err
				} else {
					fmt.Println(string(data))
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
