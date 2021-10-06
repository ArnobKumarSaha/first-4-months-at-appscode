package main

import (
	"fmt"
	"log"
	"time"
)
/*
type Post struct {
	UserId     int `json:"userId"`
	Id		int `json:"id"`
	Title    string `json:"title"`
	Body string `json:"body"`
}
*/
func main() {
	begin := time.Now()

	for id:=1; id<=100; id+=1{
		_, err := getPostById(id)
		if err != nil {
			log.Println("error while getting user data: ", err)
		}
		//log.Println(pst)
	}
	duration := time.Since(begin)
	fmt.Println(duration)

	// Nanoseconds as int64
	fmt.Println(duration.Milliseconds())
}
/*
func getPostById(id int) (Post, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%v", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Post{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}
	str, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}
	var pst Post
	pst = Post{}
	if err := json.Unmarshal(str, &pst); err != nil {
		fmt.Println("err = " , err)
		return Post{}, err
	}
	return pst, nil
}
*/