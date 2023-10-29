package isp

type User struct {
	id    int
	name  string
	phone int
	email string
}

func (u *User) GetSender() *User {
	return &User{
		id:    1,
		name:  "Thiago",
		phone: 21999999999,
		email: "thiagothalves@gmail.com",
	}
}

func (u *User) GetReceiver() *User {
	return &User{
		id:    2,
		name:  "Marcos",
		phone: 21888888888,
		email: "marcos@gmail.com",
	}
}
