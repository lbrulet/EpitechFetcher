package models

type User struct {
	Login            string       `json:"login"`
	Title            string       `json:"title"`
	InternalEmail    string       `json:"internal_email"`
	LastName         string       `json:"lastname"`
	FirstName        string       `json:"firstname"`
	UserInfo         UserInfo     `json:"userinfo"`
	ReferentUsed     bool         `json:"referent_used"`
	Picture          string       `json:"picture"`
	PictureFun       string       `json:"picture_fun"`
	EmailReferent    string       `json:"email_referent"`
	PassReferent     string       `json:"pass_referent"`
	ScolarYear       string       `json:"scolaryear"`
	Promo            int          `json:"promo"`
	Semester         int          `json:"semester"`
	Location         string       `json:"location"`
	Document         string       `json:"documents"`
	UserDocs         string       `json:"userdocs"`
	Shell            string       `json:"shell"`
	Close            bool         `json:"close"`
	CTime            string       `json:"ctime"`
	MTime            string       `json:"mtime"`
	IdPromo          string       `json:"id_promo"`
	IdHistory        string       `json:"id_history"`
	CourseCode       string       `json:"course_code"`
	SemesterCode     string       `json:"semester_code"`
	SchoolId         string       `json:"school_id"`
	SchoolCode       string       `json:"school_code"`
	SchoolTitle      string       `json:"school_title"`
	OldIdPromo       string       `json:"old_id_promo"`
	OldIdLocation    string       `json:"old_id_location"`
	Rights           RightsInfo   `json:"rights"`
	Invited          bool         `json:"invited"`
	StudentYear      int          `json:"studentyear"`
	Admin            bool         `json:"admin"`
	Editable         bool         `json:"editable"`
	RestrictProfiles bool         `json:"restrictprofiles"`
	Groups           []GroupsInfo `json:"groups"`
	Events           []EventsInfo `json:"events"`
	Credits          int          `json:"credits"`
	GpaInfo          []GpaInfo    `json:"gpa"`
	Spice            SpiceInfo    `json:"spice"`
	NsStat           NsStatInfo   `json:"nsstat"`
}

type Info struct {
	Value  string `json:"value"`
	Adm    bool   `json:"adm"`
	Public bool   `json:"public"`
}

type UserInfo map[string]Info

type RightsInfo map[string]interface{}

type GroupsInfo struct {
	Title string `json:"title"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type EventsInfo struct {
	IdEventFailed    string `json:"id_event_failed"`
	IdUser           string `json:"id_user"`
	Begin            string `json:"begin"`
	IdActiviteFailed string `json:"id_activite_failed"`
}

type GpaInfo struct {
	Gpa   string `json:"gpa"`
	Cycle string `json:"cycle"`
}

type SpiceInfo struct {
	AvailableSpice string `json:"available_spice"`
	ConsumedSpice  int    `json:"consumed_spice"`
}

type NsStatInfo struct {
	Active    float32 `json:"active"`
	Idle      int     `json:"idle"`
	OutActive int     `json:"out_active"`
	OutIdle   int     `json:"out_idle"`
	NsLogNorm int     `json:"nslog_norm"`
}
