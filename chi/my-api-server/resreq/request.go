package resreq

import (
	"errors"
	"fmt"
	"github.com/Arnobkumarsaha/myserver/schemas"
	"net/http"
	"strings"
)

type ProductRequest struct {
	*schemas.Product
	User *schemas.User `json:"user,omitempty"`
}

func (a *ProductRequest) Bind(r *http.Request) error {
	fmt.Println("ArticleRequest Bind() method is called.")
	if a.Product == nil {
		return errors.New("missing required Article fields.")
	}

	a.Product.Title = strings.ToLower(a.Product.Title) // as an example, we down-case
	return nil
}

func dbGetUser(id int64) (*schemas.User, error) {
	for _, u := range schemas.Users {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found.")
}