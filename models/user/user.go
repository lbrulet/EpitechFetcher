package user

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