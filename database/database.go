package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ckenkub/finalexam/types"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	createTb := `CREATE TABLE IF NOT EXISTS customer (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
		);`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table customer ", err)
	}
}

//Conn get database
func Conn() *sql.DB {
	return db
}

//GetCustomers all
func GetCustomers() ([]*types.Customer, error) {

	items := []*types.Customer{}
	stmt, err := Conn().Prepare("SELECT id, name, email, status FROM customer")

	rows, err := stmt.Query()
	for rows.Next() {

		var id int
		var name, email, status string
		err = rows.Scan(&id, &name, &email, &status)
		if err != nil {
			return nil, fmt.Errorf("can't scan customer items: %w", err)
		}
		item := types.Customer{id, name, email, status}
		items = append(items, &item)
	}

	return items, nil
}

//GetCustomerByID int
func GetCustomerByID(id string) (types.Customer, error) {

	stmt, err := Conn().Prepare("SELECT id, name, email, status FROM customer WHERE id =$1")

	if err != nil {
		return types.Customer{}, fmt.Errorf("can't prepare select statement: %w", err)
	}

	row := stmt.QueryRow(id)

	var idResponse int
	var name, email, status string
	err = row.Scan(&idResponse, &name, &email, &status)
	if err != nil {
		return types.Customer{}, fmt.Errorf("can't scan customer item: %w", err)
	}

	item := types.Customer{idResponse, name, email, status}

	return item, nil
}

//CreateCustomer customer.Customer
func CreateCustomer(customer types.Customer) (types.Customer, error) {
	stmt, err := Conn().Prepare("INSERT INTO customer (name, email, status) VALUES($1, $2, $3) RETURNING id")
	log.Println("start statement")
	if err != nil {
		return customer, fmt.Errorf("can't prepare insert statement: %w", err)
	}

	log.Println("start insert")
	row := stmt.QueryRow(customer.Name, customer.Email, customer.Status)

	log.Println("start scan")
	err = row.Scan(&customer.ID)
	if err != nil {
		return customer, fmt.Errorf("can't scan create customer: %w", err)
	}

	return customer, nil
}

//UpdateCustomer customer.Customer
func UpdateCustomer(customer types.Customer, id string) error {
	stmt, err := Conn().Prepare("UPDATE customer SET name=$1, email=$2, status=$3 where id=$4")
	if err != nil {
		return fmt.Errorf("can't prepare update statement: %w", err)
	}

	if _, err := stmt.Exec(customer.Name, customer.Email, customer.Status, id); err != nil {
		return fmt.Errorf("can't execute update: %w", err)
	}

	return nil
}

func DeleteCustomerById(id string) error {
	stmt, err := Conn().Prepare("DELETE FROM customer WHERE id =$1")
	if err != nil {
		return fmt.Errorf("can't prepare delete statement: %w", err)
	}

	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("can't execute delete: %w", err)
	}

	return nil
}
