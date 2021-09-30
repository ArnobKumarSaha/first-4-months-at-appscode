package main

import (
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)


type ProductRequest struct {
	*Product
	User *User `json:"user,omitempty"`
}

func (a *ProductRequest) Bind(r *http.Request) error {
	fmt.Println("ArticleRequest Bind() method is called.")
	if a.Product == nil {
		return errors.New("missing required Article fields.")
	}

	a.Product.Title = strings.ToLower(a.Product.Title) // as an example, we down-case
	return nil
}


type ProductResponse struct {
	*Product
	User *User `json:"user,omitempty"`
}

func (rd *ProductResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	w.Write([]byte("ProductResponse Render() method is called.") )
	return nil
}


func NewProductResponse(product *Product) *ProductResponse {
	resp := &ProductResponse{Product: product}
	if resp.User == nil {
		if user, _ := dbGetUser(resp.OwnerId); user != nil {
			resp.User = user
		}
	}
	return resp
}

func NewProductListResponse(products []*Product) []render.Renderer {
	list := []render.Renderer{}
	for _, article := range products {
		list = append(list, NewProductResponse(article))
	}
	fmt.Println("In NewArticleListResponse done.")
	return list
}

func dbGetUser(id int64) (*User, error) {
	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found.")
}