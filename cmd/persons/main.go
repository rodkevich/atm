package main

import (
	"fmt"
	"github.com/rodkevich/atm/persons"
	"github.com/rodkevich/profile/repository"
	"github.com/rodkevich/profile/repository/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	s := grpc.NewServer()
	srv := &persons.GRPCServer{}
	persons.RegisterPersonsServer(s, srv)
	l, err := net.Listen("tcp", ":9090")

	err = initDataBase(err)

	fmt.Println("Server is running")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func initDataBase(err error) error {
	_ = os.Setenv("DATABASE_URL", "postgresql://postgres:postgres@0.0.0.0:5432/postgres")

	var (
		rep repository.Repository
		//err error
	)

	rep, err = postgres.NewRepository() //init postgres rep db

	if err != nil {
		panic(err)
	}

	//create tables if not presented
	err = rep.Up()

	//personsList list
	var personsList []*repository.PersonModel

	//some test personsList
	var Peter = &repository.PersonModel{
		ProjectID:   "S0me_pr0j3ct_1d",
		ProjectEnv:  "S0me_pr0j3ct_3nv",
		CompanyID:   "S0me_c0mp4ny_1d",
		TeamID:      "S0me_t34m_1d",
		GroupID:     "S0me_gr0up_1d",
		FirstName:   "Peter",
		LastName:    "Bar",
		Email:       "Peter@bar.foo",
		Phone:       "08123456789",
		Description: "Some Description 1",
	}

	var John = &repository.PersonModel{
		ProjectID:   "S0me_pr0j3ct_1d",
		ProjectEnv:  "S0me_pr0j3ct_3nv",
		CompanyID:   "S0me_c0mp4ny_1d",
		TeamID:      "S0me_t34m_1d",
		GroupID:     "S0me_gr0up_2d",
		FirstName:   "John",
		LastName:    "Foo",
		Email:       "John@foo.bar",
		Phone:       "9887654210",
		Description: "Some Description 2",
	}

	//append to list
	personsList = append(personsList, Peter)
	personsList = append(personsList, John)

	//create users
	for _, person := range personsList {
		id, _ := rep.Create(person)
		log.Println(id)
	}

	// search for users
	users, err := rep.Find()
	if err != nil {
		log.Println(err)
	}
	log.Println(users)

	if err != nil {
		panic(err)
	}
	return err

}
