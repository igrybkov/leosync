package response

type apiResponse struct {
	ErrorMsg string `json:"error_msg"`
}

// UserDataResponse is an entity which represents UserDataResponse
type UserDataResponse struct {
	Username string `json:"nickname"`
	ID       int    `json:"user_id"`
	Key      string `json:"autologin_key"`
}

// LoginResponse is an entity which represents LoginResponse
type LoginResponse struct {
	apiResponse
	User UserDataResponse `json:"user"`
}

// Word is a word representation
type Word struct {
	apiResponse
	ID         uint   `json:"word_id"`
	Value      string `json:"word_value"`
	Transcript string `json:"transcription"`
	//	Created     time.Time       `json:"created_at"`
	//	LastUpdated time.Time       `json:"last_updated_at"`
	Translations []UserTranslate `json:"translate"`
	TranslateID  int             `json:"translate_id"`
	SoundURL     string          `json:"sound_url"`
	PictureURL   string          `json:"pic_url"`
}

// UserTranslate is an entity which represents UserTranslate
type UserTranslate struct {
	Value      string `json:"value"`
	PictureURL string `json:"pic_url"`
}

// UserDict is an entity which represents UserDict
type UserDict struct {
	Name  string `json:"name"`
	Count uint   `json:"count"`
	Words []Word `json:"words"`
}
