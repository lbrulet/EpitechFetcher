package user

type Info struct {
	Value  string `json:"value"`
	Adm    bool   `json:"adm"`
	Public bool   `json:"public"`
}

type UserInfo map[string]Info
