package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Arnobkumarsaha/myserver/auth"
	"github.com/Arnobkumarsaha/myserver/schemas"
	"github.com/go-chi/chi/v5"
	"net/http"
)

/*
If I set alias like , type ControllerProductResource auth.AuthProductResource
then, ControllerProductResource will not find the methods (for example AuthMiddleWare() ) of AuthProductResource.
It will only get the fields .
*/


type ControllerProductResource struct {
	*auth.AuthProductResource
}

//type ControllerProductResource auth.AuthProductResource

func (rs *ControllerProductResource) Routes() chi.Router {
	r := chi.NewRouter()

	/*oo := ControllerProductResource{
		rs.AuthProductResource,
	}*/

	r.Use(rs.AuthMiddleware)
	r.Get("/", rs.GetAllProducts)    // GET /todos - read a list of todos
	r.Post("/create", rs.CreateProduct) // POST /todos - create a new todo and persist it

	/*
		r.Put("/", rs.Delete)
	*/
	r.Route("/{prod_id}", func(r chi.Router) {
		r.Use(rs.ProductCtx)
		r.Get("/", rs.GetSingleProduct)
		r.Put("/", rs.UpdateProduct)
		r.Delete("/", rs.DeleteProduct)
	})

	return r
}

func (rs *ControllerProductResource) ParseProductFromRequestBody(w http.ResponseWriter, r *http.Request) schemas.Product {
	var newProduct schemas.Product
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP errorhandler
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Decoding is not successful in ParseProduct function.")
	}
	return newProduct
}

func (rs *ControllerProductResource) ProductCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "prod_id", chi.URLParam(r, "prod_id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
