package srp

type User struct {
	id    int
	name  string
	phone int
}

func (u *User) GetSender() *User {
	return &User{
		id:    1,
		name:  "Thiago",
		phone: 21999999999,
	}
}

func (u *User) GetReceiver() *User {
	return &User{
		id:    2,
		name:  "Marcos",
		phone: 21888888888,
	}
}
