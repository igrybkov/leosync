package api

type apiResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type UserDataResponse struct {
	Username string `json:"nickname"`
	Id       int    `json:"user_id"`
	Key      string `json:"autologin_key"`
}

type LoginResponse struct {
	apiResponse
	User UserDataResponse `json:"user"`
}

type Word struct {
	apiResponse
	Id         uint   `json:"word_id"`
	Value      string `json:"word_value"`
	Transcript string `json:"transcription"`
	//	Created     time.Time       `json:"created_at"`
	//	LastUpdated time.Time       `json:"last_updated_at"`
	Translations []UserTranslate `json:"translate"`
	SoundUrl     string          `json:"sound_url"`
	PictureUrl   string          `json:"pic_url"`
}

type UserTranslate struct {
	Value      string `json:"value"`
	PictureUrl string `json:"pic_url"`
}

type Userdict struct {
	Name  string `json:"name"`
	Count uint   `json:"count"`
	Words []Word `json:"words"`
}
