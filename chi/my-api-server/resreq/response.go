package resreq

import (
	"fmt"
	"github.com/Arnobkumarsaha/myserver/schemas"
	"github.com/go-chi/render"
	"net/http"
)

type ProductResponse struct {
	*schemas.Product
	User *schemas.User `json:"user,omitempty"`
}

func (rd *ProductResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	w.Write([]byte("ProductResponse Render() method is called.") )
	return nil
}


func NewProductResponse(product *schemas.Product) *ProductResponse {
	resp := &ProductResponse{Product: product}
	if resp.User == nil {
		if user, _ := dbGetUser(resp.OwnerId); user != nil {
			resp.User = user
		}
	}
	return resp
}

func NewProductListResponse(products []*schemas.Product) []render.Renderer {
	list := []render.Renderer{}
	for _, article := range products {
		list = append(list, NewProductResponse(article))
	}
	fmt.Println("In NewArticleListResponse done.")
	return list
}


/*
ABOUT RENDERING (call stack):
render.RenderList & render.Render calls renderer() & Respond().
renderer() calls Render() (not render.render(), but the Render() method of the structure.)

Respond() calls DefaultResponder(), then DefaultResponder() call JSON().
and JSON() actually writes, what We see in the responses.
*/