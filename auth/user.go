package auth

type User struct {
	login     string
	key       string
	buildings []string
}

func NewUser(login string, key string) *User {
	return &User{
		login: login,
		key:   key,
	}
}

func (u *User) Login() string {
	return u.login
}

func (u *User) Key() string {
	return u.key
}

func (u *User) AddBuilding(bld string) {
	u.buildings = append(u.buildings, bld)
}

func (u *User) Buildings() *[]string {
	return &u.buildings
}
