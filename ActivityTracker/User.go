package main

import "fmt"

type User struct {
	name       string
	email      string
	activities map[string]Activity
}

func RegisterUser(name string, email string, activities map[string]Activity, emailMap map[string]User) *User {
	if exist := emailExist(emailMap, email); exist {
		fmt.Println("Email Already Exist User Registration Failed")
		return nil
	}
	emailMap[email] = User{name: name, email: email, activities: activities}
	user := emailMap[email]
	return &user
}

func emailExist(emailMap map[string]User, email string) bool {
	_, exist := emailMap[email]
	return exist
}

func (u *User) startActivity(activity *Activity) {
	activity.startActivity()
	u.activities[activity.Action] = *activity
}

func (u *User) stopActivity(activity *Activity) {
	activity.stopActivity()

	u.activities[activity.Action] = *activity
}

func (u *User) PrintLog() {
	fmt.Printf("\t--- %s's Activity Log ---\t\n\n", u.name)
	//fmt.Println(u.activities)
	for _, activity := range u.activities {
		fmt.Println(activity.Log())
	}
	fmt.Println()
}
