package app

import (
	"Banking_Rest_Exmpl/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type customer struct {
	Name    string `json:"Full_Name" xml:"name"`
	City    string `json:"City_Name" xml:"city"`
	Zipcode string `json:"Zip_code" xml:"zip"`
}

/*
	func greet(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World Rabin.....")
	}
*/

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getallcustomers(w http.ResponseWriter, r *http.Request) {

	/*customers := []customer{
		{Name: "Rabin", City: "Kolkata", Zipcode: "700123"},
		{Name: "Sangita", City: "Kolkata", Zipcode: "700123"},
		{Name: "Ayushman", City: "Kolkata", Zipcode: "700123"},
	}*/

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

/*
func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post request received")
}
*/
