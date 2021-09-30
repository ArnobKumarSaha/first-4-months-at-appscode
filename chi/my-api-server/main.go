package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

var products []*Product
var users []*User

func main()  {
	port := "8080"
	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//r.Get("/pp", check)

	r.Post("/signin", SignIn)
	r.Get("/welcome", Welcome)
	r.Post("/logout", LogOut)

	r.Mount("/products", productResource{}.Routes())

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


func init()  {
	adduser()
	addProduct()
}
func adduser()  {
	user1 := &User{
		Name:    "Arnob kumar saha",
		Id:      10,
		Contact: Contact{
			PhoneNumber: 123123123,
			Address:     "uttara",
		},
	}
	user2 := &User{
		Name:    "Tasdidur rahman",
		Id:      15,
		Contact: Contact{
			PhoneNumber: 4556132465,
			Address:     "banani",
		},
	}
	user3 := &User{
		Name:    "Rakibul hossain",
		Id:      12,
		Contact: Contact{
			PhoneNumber: 54132133,
			Address:     "dhanmondi",
		},
	}
	users = append(users, user1, user2, user3)
}
func addProduct()  {
	pr1 := &Product{
		Title:   "samsung j7",
		Price:   123,
		Type:    "phone",
		Id:      2,
		OwnerId: 12,
	}
	pr2 := &Product{
		Title:   "asus 3453",
		Price:   6200,
		Type:    "laptop",
		Id:      3,
		OwnerId: 10,
	}
	pr3 := &Product{
		Title:   "redme note4",
		Price:   10023,
		Type:    "phone",
		Id:      6,
		OwnerId: 15,
	}
	products = append(products, pr1, pr2, pr3)
}