package api

// LoginRequest is a request for Login
type LoginRequest struct {
	Email    string `url:"email"`
	Password string `url:"password"`
}

// TranslationRequest is a request for Translation
type TranslationRequest struct {
	Word string `url:"word"`
}

// AddWordRequest is a request for AddWord
type AddWordRequest struct {
	Word        string `url:"word"`
	Translation string `url:"tword"`
}

// AddWordWithContextRequest is a request for AddWordWithContext
type AddWordWithContextRequest struct {
	Word        string `url:"word"`
	Translation string `url:"tword"`
	Context     string `url:"context"`
}
