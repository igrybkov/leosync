package api

type LoginRequest struct {
	Email    string `url:"email"`
	Password string `url:"password"`
}

type TranslationRequest struct {
	Word string `url:"word"`
}

type AddWordRequest struct {
	Word        string `url:"word"`
	Translation string `url:"tword"`
}

type AddWordWithContextRequest struct {
	Word        string `url:"word"`
	Translation string `url:"tword"`
	Context string `url:"context"`
}
