package main

import (
   "encoding/json"
   "fmt"
   "net/http"
   "log"
)

type User struct {
   Name     string `json:"name"`
   Username string `json:"username"`
   Email    string `json:"email"`
}

func main() {
   // Make a GET request to the API endpoint
   response, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
   if err != nil {
       log.Fatal(err)
   }
   defer response.Body.Close()

   // Decode the response into a User struct
   var user User
   err = json.NewDecoder(response.Body).Decode(&user)
   if err != nil {
       log.Fatal(err)
   }

   // Print the user's name, username, and email
   fmt.Println("Name:", user.Name)
   fmt.Println("Username:", user.Username)
   fmt.Println("Email:", user.Email)
}
