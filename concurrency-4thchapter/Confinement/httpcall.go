package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Post struct {
	UserId     int `json:"userId"`
	Id		int `json:"id"`
	Title    string `json:"title"`
	Body string `json:"body"`
}

type PostHttpCall struct{
	Post
	err error
}

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


// struct-typed channel
var getUsers = func(numOfUsers int) <-chan PostHttpCall {
	userStream := make(chan PostHttpCall, numOfUsers)

	// Anonymous go function is called, which is using one directional channel  === confinement
	go func() {
		wg := &sync.WaitGroup{}
		defer close(userStream)
		for i := 1; i <= numOfUsers; i++ {
			wg.Add(1)
			go func(id int, group *sync.WaitGroup) {
				defer group.Done()
				user, err := getPostById(id)
				userStream <- PostHttpCall{
					Post: user,
					err:  err,
				}
			}(i, wg)
		}
		wg.Wait()
	}()
	return userStream
}
func main() {
	begin := time.Now()
	for pst := range getUsers(100) {
		if pst.err != nil {
			log.Println("error while getting pst data: ", pst.err)
			continue
		}
		//log.Println(pst.Post)
	}
	duration := time.Since(begin)
	fmt.Println(duration)

	// Nanoseconds as int64
	fmt.Println(duration.Milliseconds())
}