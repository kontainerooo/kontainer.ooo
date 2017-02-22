package user

// The Address structs is used to represent a real world address
type Address struct {
	ID         uint
	Postcode   string
	City       string
	Country    string
	Street     string
	Houseno    int
	Additional string
}

// The Config struct represents a users general information
type Config struct {
	Admin     bool
	Email     string
	Password  string
	Salt      string
	Image     string
	Address   Address
	AddressID int
	Phone     string
}

// The User struct represents the users table with an id, a username and the general information
type User struct {
	ID       uint
	Username string
	Config
}

// The Customer struct represents the customers table, extending the User by a tier level and a company name
type Customer struct {
	User    User
	UserID  int
	Tier    int
	Company string
}
