package main

func createUser(user *User) error {
	return db.Create(user).Error
}
