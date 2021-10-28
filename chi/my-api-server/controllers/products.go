package controllers

import (
	"fmt"
	"github.com/Arnobkumarsaha/myserver/errorhandler"
	"github.com/Arnobkumarsaha/myserver/resreq"
	"github.com/Arnobkumarsaha/myserver/schemas"
	"github.com/go-chi/render"
	"net/http"
)

/*
resreq.NewProductResponse has a Render() method,
also errorhandler.ErrRender() gives a errResponse, which has a Render() method

and we need these to use render.Render() & render.RenderList() functions.
 */

func (rs *ControllerProductResource) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Products (GetAllProducts)= ", schemas.Products)
	tmp := resreq.NewProductListResponse(schemas.Products)
	fmt.Println("tmp = ", tmp)
	if err := render.RenderList(w, r, tmp); err != nil {
		render.Render(w, r, errorhandler.ErrRender(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (rs *ControllerProductResource) GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.Context().Value("prod_id").(string)

	for _ , p := range schemas.Products {
		if isEqual(id, p.Id){
			tmp := resreq.NewProductResponse(p)
			fmt.Println("Found ! ", p, tmp)
			rdObj := render.Renderer(tmp)

			if err := render.Render(w, r, rdObj); err != nil {
				render.Render(w, r, errorhandler.ErrRender(err))
				return
			}
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("No products found with id = " + id))
}

func (rs *ControllerProductResource) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newProduct := rs.ParseProductFromRequestBody(w,r)

	fmt.Println(newProduct)

	schemas.Products = append(schemas.Products, &newProduct)
	fmt.Println("After appending : ", schemas.Products)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("added !! "))
}

func (rs *ControllerProductResource) UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	updatedProduct := rs.ParseProductFromRequestBody(w,r)

	id := r.Context().Value("prod_id").(string)
	fmt.Println("id in UpdateProduct() = " , id)

	if !isEqual(updatedProduct.Id, id){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Id can not be changed when updating.."))
		return
	}

	for idx , p := range schemas.Products {
		if isEqual(id, p.Id){
			schemas.Products[idx] = &updatedProduct
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Product updated. \n"))
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("No products found with id = " + id))
}

func (rs *ControllerProductResource) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.Context().Value("prod_id").(string)
	fmt.Println("id in DeleteProduct() = " , id)

	for idx , p := range schemas.Products {
		if isEqual(id, p.Id){
			// Deleting this indexed product
			schemas.Products = append(schemas.Products[:idx], schemas.Products[idx+1:]...)
			fmt.Println("After Deleting : ", schemas.Products)
			return
		}
	}
	fmt.Println("No product found with product_id = ", id)
	return
}

