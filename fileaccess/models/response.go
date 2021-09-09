package models

// Response container object for holding data and error
type Response struct {
	UserCopy     *UserApplication
	ErrorMessage error
}
