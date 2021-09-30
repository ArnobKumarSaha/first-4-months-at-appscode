package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
}

type omit *struct{}
type PublicUser struct {
	*User
	Password omit `json:"password,omitempty"`
}

type BlogPost struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type Analytics struct {
	Visitors  int `json:"visitors"`
	PageViews int `json:"page_views"`
}

func main()  {
	// REMOVING A FIELD EXAMPLE :
	user := User{
		Email:    "arnob@gmail.com",
		Password: "1234",
		Name:     "Arnob",
	}

	b1, _ := json.Marshal(PublicUser{
		User: &user,
	})
	fmt.Println(string(b1))

	
	
	//	COMBINING MULTIPLE STRUCT OBJECTS
	blogObj := BlogPost{
		URL:   "attilaolah@gmail.com",
		Title: "Attila's Blog",
	}
	anaObj := Analytics{
		Visitors:  7,
		PageViews: 45,
	}
	b2, _ := json.Marshal(struct{
		*BlogPost
		*Analytics
	}{&blogObj, &anaObj})
	fmt.Println(string(b2))



	// SPLITTING OBJECT
	var newBlogObj BlogPost
	var newAnaObj Analytics
	_ = json.Unmarshal(b2, &struct {
		*BlogPost
		*Analytics
	}{&newBlogObj, &newAnaObj})

	fmt.Println(newBlogObj, " ++ ", newAnaObj)
}
