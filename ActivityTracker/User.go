package main

type User struct {
	name       string
	email      string
	activities []*Activity
}

func RegisterUser(name string, email string, activities []*Activity, users []User) (*User, bool) {
	if exist := emailExist(users, email); exist {
		return nil, false
	}
	users = append(users, User{name: name, email: email, activities: activities})
	return &User{name: name, email: email, activities: activities}, true
}

func emailExist(users []User, email string) bool {
	for _, user := range users {
		if user.email == email {
			return false
		}
	}
	return true
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) startActivity(activity Activity) {
	activity.startActivity()
}

func (u *User) stopActivity(activity Activity) {
	activity.stopActivity()
	u.activities = append(u.activities, &activity)
}
