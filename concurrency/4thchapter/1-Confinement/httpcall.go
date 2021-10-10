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



// The ARCHITECTURE
// The main() calls getUsers() , here main is reader & getUsers is writer.
// getUsers() has channel named userStream, which does two things here :
    // 1. It receives data from internal go(), & supplies that data to the receiver end (in main func here)
    // 2. It receives signal from external go function , when that external go() is done. And pass that message to main().
// The external go() uses a waitGroup to enforces all the internal go() to be executed
    // That waitGroup gets a DONE message from all internal go() call. wait until 'numOfUsers' number DONE receives.  which was set by wg.Add()

var getUsers = func(numOfUsers int) <-chan PostHttpCall {
	// struct-typed channel
	// a channel of type PostHttpCall of size numOfUsers
	userStream := make(chan PostHttpCall, numOfUsers)

	// Anonymous go function is called, which is using one directional channel  === confinement
	go func() {
		wg := &sync.WaitGroup{}
		defer close(userStream)
		for i := 1; i <= numOfUsers; i++ {
			wg.Add(1)

			go func(id int, in_wg *sync.WaitGroup) { // go function inside go function
				defer in_wg.Done()
				// In each internal go call, just do the Http request & make the response struct
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
	naiveApproach()
	efficientApproach()
}

// The efficientApproach takes 103 ms
func efficientApproach()  {
	begin := time.Now()
	for pst := range getUsers(100) {
		if pst.err != nil {
			log.Println("error while getting pst data: ", pst.err)
			continue
		}
		//log.Println(pst.Post)
	}
	duration := time.Since(begin)
	fmt.Println(duration.Milliseconds())
}

// where The naiveApproach takes 6830  ms
func naiveApproach()  {
	begin := time.Now()
	for i:=1 ; i<=100; i+=1 {
		pst, err := getPostById(i)
		if err != nil{
			fmt.Println(pst, err)
		}
	}
	duration := time.Since(begin)
	fmt.Println(duration.Milliseconds())
}