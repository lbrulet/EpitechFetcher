package EpitechFetcher

import (
	"encoding/json"
	"fmt"
	"github.com/lbrulet/EpitechFetcher/models"
	"io/ioutil"
	"net/http"
)

func FetchIntranet(user models.EpiUser) error {
	fmt.Println(user)
	client := &http.Client{}
	if req, err := http.NewRequest("GET", "https://intra.epitech.eu/"+ user.AutoLogin +"/user?format=json", nil); err != nil {
		return err
	} else {
		response, err := client.Do(req)
		if err != nil {
			return err
		} else {
			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			} else {
				var user models.User
				if err := json.Unmarshal(data, &user); err != nil {
					return err
				} else {
					rankingsJson, _ := json.Marshal(user)
					err = ioutil.WriteFile("output.json", rankingsJson, 0644)
					fmt.Printf("%s", string(rankingsJson))
					if err := response.Body.Close(); err != nil {
						return err
					} else {
						return nil
					}
				}
			}
		}
	}
}
