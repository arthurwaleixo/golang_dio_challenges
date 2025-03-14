package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Address struct {
	Street string `json:"street,omitempty"`
	City   string `json:"city,omitempty"`
	State  string `json:"state,omitempty"`
	Zip    string `json:"zip,omitempty"`
}

type Customer struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Email   string   `json:"email,omitempty"`
	Phone   string   `json:"phone,omitempty"`
	Address *Address `json:"address,omitempty"`
}

var customer []Customer

// GetAllCustomers
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(customer)
}

// GetCustomerByID
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range customer {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Customer{})
}

// CreateCustomer
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var cust Customer
	_ = json.NewDecoder(r.Body).Decode(&cust)
	cust.ID = params["id"]
	customer = append(customer, cust)
	json.NewEncoder(w).Encode(customer)
}

// UpdateCustomer
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range customer {
		if item.ID == params["id"] {
			customer = append(customer[:index], customer[index+1:]...)
			var cust Customer
			_ = json.NewDecoder(r.Body).Decode(&cust)
			cust.ID = params["id"]
			customer = append(customer, cust)
			json.NewEncoder(w).Encode(cust)
			return
		}
	}
	json.NewEncoder(w).Encode(customer)
}

// DeleteCustomer
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range customer {
		if item.ID == params["id"] {
			customer = append(customer[:index], customer[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(customer)
}

func main() {
	router := mux.NewRouter()
	customer = append(customer, Customer{ID: "1", Name: "John Doe", Email: "john.doe@example.com", Phone: "1234567890", Address: &Address{Street: "123 Main St", City: "Anytown", State: "CA", Zip: "12345"}})
	customer = append(customer, Customer{ID: "2", Name: "Jane Doe", Email: "jane.doe@example.com", Phone: "0987654321", Address: &Address{Street: "456 Oak St", City: "Anytown", State: "CA", Zip: "12345"}})
	customer = append(customer, Customer{ID: "3", Name: "Bob Smith", Email: "bob.smith@example.com", Phone: "5555555555", Address: &Address{Street: "789 Pine St", City: "Anytown", State: "CA", Zip: "12345"}})
	router.HandleFunc("/customers", GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", GetCustomerByID).Methods("GET")
	router.HandleFunc("/customers/{id}", CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", DeleteCustomer).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
