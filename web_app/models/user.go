package models

import (
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(username, password string) error {
	hash, err := client.Get("user: " + username).Bytes()
	if err == redis.Nil{
		return
	} else if err != nil {
		
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		
		return
	}

}

func RegisterUser(username, password string) error{
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	return client.Set("user: " + username, hash, 0).Err()
}