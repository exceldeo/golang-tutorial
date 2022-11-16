package user

type User struct {
	Name     string
	Follower map[string]*User
	Post     bool
}

func NewUser(name string) UserInteface {
	return &User{
		Name: name,
		// Follower: map[string]*User{},
		Follower: make(map[string]*User),
		Post:     false,
	}
}

func (u *User) Follow(following *User) string {
	if _, ok := u.Follower[following.Name]; ok {
		return "you already following"
	}
	u.Follower[following.Name] = following
	return "following successfully"
}

func (u *User) UploadPost() string {
	if u.Post {
		return "you already post"
	}
	u.Post = true
	return "post successfully"
}
