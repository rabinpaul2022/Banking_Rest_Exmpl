package app

import (
	"Banking_Rest_Exmpl/domain"
	"Banking_Rest_Exmpl/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	/*http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World Rabin.....")
	})*/

	//mux := http.NewServeMux()

	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//mux.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", ch.getallcustomers).Methods(http.MethodGet)
	//mux.HandleFunc("/newcustomer", createCustomer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8333", router))
}
