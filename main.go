package main

import(
	"fmt"
	"os"
	// "github.com/mediocregopher/radix.v2/redis"
	"github.com/gomodule/redigo/redis"
	"time"
	"encoding/json"

)

type person struct{
	RandID   int64    `json:"id"`
	Username string   `json:"username"`
	Name     string   `json:"name"`
	Role     string   `json:"role"`
}

func checkErr(e error){
	if e != nil{
		panic(e)
	}
}

func redisConn(personObject person){
	//connection Section
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkErr(err)
	checkConn, err := redis.String(conn.Do("PING"))
	checkErr(err)

	fmt.Printf("PING Response = %s\n", checkConn)
//end connection Section

//Convert JSON Section
personAsJSON, err := json.Marshal(personObject)
		checkErr(err)
		 fmt.Println("My Json: ", string(personAsJSON))
//End Convert JSON Section

prefix:="user:"

//HMSET method
// _, err = conn.Do("HMSET",
//  redis.Args{"person_Object"}.AddFlat(personObject)...)

	//prefix+personAsJSON.username
	_, err = conn.Do("SET", prefix+personObject.Username, personAsJSON)
	checkErr(err)
	value, err := redis.String(conn.Do("GET", prefix+personObject.Username))
	if err == redis.ErrNil {
		fmt.Println("User does not exist")
	} else if err != nil {
	checkErr(err)
	}

personFromRedis := person{}
	err = json.Unmarshal([]byte(value), &personFromRedis)
	fmt.Printf("Person from Redis: %+v\n", personFromRedis)


//HMGET method get values
// value, err := redis.Values(conn.Do(“HGETALL”, key))
// err = redis.ScanStruct(value, &object)

defer conn.Close()
}

func main()  {
	 personObject := &person{
		 time.Now().UnixNano(),
		 "WolfusFlow",
		"Michael Salamakha",
		"programmer", 
		}
		fmt.Println(personObject)


	 redisConn(*personObject)
}
