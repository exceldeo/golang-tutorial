package user

type UserInteface interface {
	Follow(following *User) string
	UploadPost() string
}
