package main

import (
     "fmt"
     "log"
     "encoding/json"
     "net/http"

     "golang-challenge/dataservice"
)


func main() {

        dump := new(dataservice.Dump) 

        if err := dump.Init("/Users/sundar/go/src/interview/golang-challenge/dataservice/small.gob"); err != nil{
		fmt.Println("Init Failed")	
        }
        
	//rest end point
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
                name := r.URL.Query()["name"][0]

	        //Basic Validation
	        if name == "" || name == "null" || name == "NULL" {
		     http.Error(w, "Missing Param", http.StatusInternalServerError)
    		     return 
	        }
		w.Header().Add("Content-Type", "application/json")
		result, err := dump.Get(name) 
                if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	        w.Header().Add("Content-Type", "application/json")
                json.NewEncoder(w).Encode(result)
	        w.WriteHeader(http.StatusOK)
	})

    	//webserver
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

