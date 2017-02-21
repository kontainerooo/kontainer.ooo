package user

// The Address structs is used to represent a real world address
type Address struct {
	postcode   string
	city       string
	country    string
	street     string
	houseno    int
	additional string
}

// The Config struct represents a users general information
type Config struct {
	admin    bool
	email    string
	password string
	salt     string
	image    string
	address  Address
	phone    string
}

// The User struct represents the users table with an id, a username and the general information
type User struct {
	ID       uint64
	username string
	Config
}

// The Customer struct represents the customers table, extending the User by a tier level and a company name
type Customer struct {
	User    `pg:",override"`
	tier    int
	company string
}
