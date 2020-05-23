package main

func updateUser(user *User) error {
	// update user
	// set username = ?, password = ?
	// where id = ?
	return db.Model(user).Update(user).Error
}
