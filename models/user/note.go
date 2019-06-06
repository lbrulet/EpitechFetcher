package user

func NewNoteIntra() CursusIntra {
	return CursusIntra{}
}

type CursusIntra struct {
	Modules []ModuleInfos `json:"modules"`
	Notes   []NoteInfos   `json:"notes"`
}

type ModuleInfos struct {
	ScolarYear    int    `json:"scolaryear"`
	IdUserHistory string `json:"id_user_history"`
	CodeModule    string `json:"codemodule"`
	CodeInstance  string `json:"codeinstance"`
	Title         string `json:"title"`
	DateIns       string `json:"date_ins"`
	Cycle         string `json:"cycle"`
	Grade         string `json:"grade"`
	Credits       int    `json:"credits"`
	Barrage       int    `json:"barrage"`
}

type NoteInfos struct {
	ScolarYear   int    `json:"scolaryear"`
	CodeModule   string `json:"codemodule"`
	TitleModule  string `json:"titlemodule"`
	CodeInstance string `json:"codeinstance"`
	CodeActi     string `json:"codeacti"`
	Title        string `json:"title"`
	Date         string `json:"date"`
	Correcteur   string `json:"correcteur"`
	FinalNote    float32    `json:"final_note"`
	Comment      string `json:"comment"`
}
