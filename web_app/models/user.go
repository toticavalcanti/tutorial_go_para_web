package models

func RegisterUser(username, password string) error{
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	return client.Set("user: " + username, hash, 0).Err()
}