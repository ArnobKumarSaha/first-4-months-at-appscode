package schemas

import (
	"fmt"
	"net/http"
)

type Product struct {
	Title string `json:"title"`
	Price  int `json:"price"`
	Type string `json:"type"`
	Id int64 `json:"product_id"`
	OwnerId int64 `json:"owner_id"`
}

type User struct {
	Name string `json:"name"`
	Id int64 `json:"user_id"`
	Contact
}

// Bind on User will run after the unmarshalling is complete, its
// a good time to focus some post-processing after a decoding.
func (u *User) Bind(r *http.Request) error {
	fmt.Println("userPayLoad Bind() method is called.")
	return nil
}

func (u *User) Render(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("userPayLoad Render() method is called.")
	return nil
}


type Contact struct {
	PhoneNumber int64 `json:"phone_number"`
	Address string `json:"address"`
}

type ProductResource struct{}