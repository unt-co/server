package main

import (
	"fmt"
	"net/http"
	"regexp"

	"./db"
	"github.com/go-redis/redis"
)

type env struct {
	rediscl *redis.Client
}

func main() {

	redisDB, err := db.InitConnection()
	if err != nil {
		panic(err)
	}
	env := &env{redisDB}

	http.Handle("/", handler(env))
	http.ListenAndServe(":8081", nil)

	fmt.Println("Hello world!")
}

func handler(env *env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		shortLink := r.URL.Path[1:]
		if len(shortLink) == 0 {
			//redirect to home page
			w.Write([]byte("Homepage!"))
			return
		}
		var location string
		ok := false //If true => Result comes from DB

		match, _ := regexp.MatchString("^([\\w\\d\\S]){2,}$", shortLink)
		if !match {
			fmt.Println("Wrong shortlink!")
			http.Redirect(w, r, "/", 301)
		}

		//Redis check
		//		Redis blacklist check

		rs := env.rediscl.HMGet(shortLink, "location", "banned")
		if rs.Err() == nil {

			result, err := rs.Result()
			if err == nil {
				if result[1] == nil {
					//Not banned
					var locationVal string
					locationVal, ok = result[0].(string)
					if ok == true {
						location = locationVal
						//recuperato da DB!
					}
				} else {
					//Banned link
					fmt.Println("Limited Link hit - " + result[1].(string))
					fmt.Println(result[1])
					return
				}

			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(rs.Err())
		}


		if location == "" {
			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				}}

			resp, err := client.Get("https://t.co/" + shortLink)
			location = resp.Header.Get("Location")
			if err != nil {
				fmt.Println(err)
			}

			if rs.Err() == nil {
				amap := make(map[string]interface{})
				amap["location"] = location
				//amap["banned"] = "0"
				is := env.rediscl.HMSet(shortLink, amap)
				if is.Err() == nil {
				} else {
					fmt.Println(err)
				}
			}
		}
		//}
		http.Redirect(w, r, location, 301)
		message := " - Link obtained from "
		if ok {
			message += "database"
		} else {
			message += "request"
		}
		fmt.Println("Redirecting " + r.RequestURI + " to " + location + message)
		return
	})
}
