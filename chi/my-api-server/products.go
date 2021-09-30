package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type productResource struct{}

func isEqual(a interface{},b interface{}) bool{
	var A, B []byte

	switch u := a.(type) {
	case int:
		A = []byte(strconv.Itoa(u))
	case string:
		A = []byte(u)
	case int64:
		A = []byte(strconv.FormatInt(int64(u),10))
	}

	switch u := b.(type) {
	case int:
		B = []byte(strconv.Itoa(u))
	case string:
		B = []byte(u)
	case int64:
		B = []byte(strconv.FormatInt(int64(u),10))
	}

	for len(A) != len(B){
		return false
	}
	for i,j := range A{
		if B[i] != j{
			return false
		}
	}
	return true
}

func (rs productResource) Routes() chi.Router {
	r := chi.NewRouter()

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

func (rs productResource) ParseProductFromRequestBody(w http.ResponseWriter, r *http.Request)  Product{
	var newProduct Product
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Decoding is not successful in ParseProduct function.")
	}
	return newProduct
}

func (rs productResource) ProductCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "prod_id", chi.URLParam(r, "prod_id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


/*
ABOUT RENDERING (call stack):
render.RenderList & render.Render calls renderer() & Respond().
renderer() calls Render() (not render.render(), but the Render() method of the structure.)

Respond() calls DefaultResponder(), then DefaultResponder() call JSON().
and JSON() actually writes, what We see in the responses.
*/


// **************************************    CONTROLLERS    **************************************



func (rs productResource) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Products (GetAllProducts)= ", products)
	tmp := NewProductListResponse(products)
	fmt.Println("tmp = ", tmp)
	if err := render.RenderList(w, r, tmp); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (rs productResource) GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.Context().Value("prod_id").(string)

	for _ , p := range products{
		if isEqual(id, p.Id){
			tmp := NewProductResponse(p)
			fmt.Println("Found ! ", p, tmp)
			rdObj := render.Renderer(tmp)

			if err := render.Render(w, r, rdObj); err != nil {
				render.Render(w, r, ErrRender(err))
				return
			}
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("No products found with id = " + id))
}

func (rs productResource) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newProduct := rs.ParseProductFromRequestBody(w,r)

	fmt.Println(newProduct)

	products = append(products, &newProduct)
	fmt.Println("After appending : ", products)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("added !! "))
}

func (rs productResource) UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	updatedProduct := rs.ParseProductFromRequestBody(w,r)

	id := r.Context().Value("prod_id").(string)
	fmt.Println("id in UpdateProduct() = " , id)

	if !isEqual(updatedProduct.Id, id){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id can not be changed when updating.."))
		return
	}

	for idx , p := range products{
		if isEqual(id, p.Id){
			products[idx] = &updatedProduct
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Product updated. \n"))
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("No products found with id = " + id))
}

func (rs productResource) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.Context().Value("prod_id").(string)
	fmt.Println("id in DeleteProduct() = " , id)

	for idx , p := range products{
		if isEqual(id, p.Id){
			// Deleting this indexed product
			products = append(products[:idx], products[idx+1:]...)
			fmt.Println("After Deleting : ", products)
			return
		}
	}
	fmt.Println("No product found with product_id = ", id)
	return
}

