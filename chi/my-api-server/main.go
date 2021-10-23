package main

import (
	"github.com/Arnobkumarsaha/myserver/auth"
	"github.com/Arnobkumarsaha/myserver/controllers"
	"github.com/Arnobkumarsaha/myserver/schemas"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)



func main()  {
	schemas.AddRequiredData()
	port := "8080"
	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//r.Get("/pp", check)

	r.Post("/signin", auth.SignIn)
	r.Get("/welcome", auth.Welcome)
	r.Post("/logout", auth.LogOut)

	/*
		ControllerProductResource should be initialized if it has a pointer indside.
		But as AuthProductResource has no fields, we hadn't to do that.
		For details :: https://play.golang.org/p/x6q-2HUfTH
	*/
	r.Mount("/products", (&controllers.ControllerProductResource{}).Routes())

	// (&controllers.ControllerProductResource{}).AuthMiddleware()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(http.ListenAndServe(":8080", r))

}

/*
func check(w http.ResponseWriter, r *http.Request){
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Dhaka&appid=e035ca5c00b6f72b3e2447c49dd92c57")


	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resByte, errr := io.ReadAll(resp.Body)
	if errr != nil {
		fmt.Println("i am fucked", errr)
	}
	fmt.Println("....",string(resByte))

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
*/
