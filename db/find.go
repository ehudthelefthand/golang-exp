package main

func getUserByID(id uint) (*User, error) {
	user := User{}
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
