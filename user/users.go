package user

type Users struct {
	Users map[string]*User
}

func NewUsers() *Users {
	return &Users{
		Users: map[string]*User{},
	}
}

func (u *Users) AddUsers(user *User) string {
	u.Users[user.Name] = user
	return "Adding user successfully"
}
