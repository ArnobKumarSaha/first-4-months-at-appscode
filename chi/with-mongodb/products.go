package main

import (
	"context"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

type productsResource struct{}

func (rs productsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.GetAllProducts)    // GET /posts - Read a list of posts.
	//r.Post("/", rs.Create) // POST /posts - Create a new post.


	r.Route("/{id}", func(r chi.Router) {
		r.Use(PostCtx)
		r.Get("/", rs.GetSingleProduct)       // GET /posts/{id} - Read a single post by :id.
		//r.Put("/", rs.Update)    // PUT /posts/{id} - Update a single post by :id.
		//r.Delete("/", rs.Delete) // DELETE /posts/{id} - Delete a single post by :id.
	})

	return r
}

func (rs productsResource) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs productsResource) GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
