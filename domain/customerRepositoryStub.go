package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Rabin", City: "Kolkata", Zipcode: "700123", DateofBirth: "04011983", Status: "Active"},
		{Id: "1002", Name: "Sangita", City: "Kolkata", Zipcode: "700123", DateofBirth: "06011986", Status: "Active"},
		{Id: "1003", Name: "Ayushman", City: "Kolkata", Zipcode: "700123", DateofBirth: "30092015", Status: "Active"},
	}

	return CustomerRepositoryStub{customers}
}
