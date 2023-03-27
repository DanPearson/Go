package main

import (
   "net/http"
   "net/http/httptest"
   "testing"
)

func TestGetUser(t *testing.T) {
   // Create a test server to mock the API endpoint
   server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
       // Set the response headers and body
       rw.Header().Set("Content-Type", "application/json")
       rw.WriteHeader(http.StatusOK)
       rw.Write([]byte(`{"name":"John Doe","username":"jdoe","email":"jdoe@example.com"}`))
   }))
   defer server.Close()

   // Make a GET request to the test server
   response, err := http.Get(server.URL)
   if err != nil {
       t.Fatal(err)
   }
   defer response.Body.Close()

   // Check the response status code
   if response.StatusCode != http.StatusOK {
       t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.StatusCode)
   }

   // Test the decoding of the response body
   var user User
   err = json.NewDecoder(response.Body).Decode(&user)
   if err != nil {
       t.Fatal(err)
   }

   // Check the decoded values
   if user.Name != "John Doe" {
       t.Errorf("Expected name to be %q, but got %q", "John Doe", user.Name)
   }
   if user.Username != "jdoe" {
       t.Errorf("Expected username to be %q, but got %q", "jdoe", user.Username)
   }
   if user.Email != "jdoe@example.com" {
       t.Errorf("Expected email to be %q, but got %q", "jdoe@example.com", user.Email)
   }
}
