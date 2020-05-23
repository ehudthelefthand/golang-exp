package main

func deleteUser(id uint) error {
	return db.Where("id = ?", id).Delete(&User{}).Error
}

// func deleteUser(user *User) error {
// 	return db.Delete(&User{}).Error
// }
