package user

import "fmt"

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

func (a Address) String() string {
	s := fmt.Sprintf("Address: %d\n", a.ID)
	s += fmt.Sprintf("Street:\t%s %d %s\n", a.Street, a.Houseno, a.Additional)
	s += fmt.Sprintf("City:\t%s %s\n", a.Postcode, a.City)
	s += fmt.Sprintf("Country:\t%s\n", a.Country)
	return s
}

// The Config struct represents a users general information
type Config struct {
	Admin     bool
	Email     string
	Password  string
	Salt      string
	Image     string
	Address   Address
	AddressID uint
	Phone     string
}

// The User struct represents the users table with an id, a username and the general information
type User struct {
	ID       uint
	Username string
	Config
}

func (u User) String() string {
	s := fmt.Sprintf("User: %d - %s\n", u.ID, u.Username)
	return s
}

func (u *User) setConfig(cfg *Config) {
	u.Admin = cfg.Admin
	u.Email = cfg.Email
	u.Password = cfg.Password
	u.Salt = cfg.Salt
	u.Image = cfg.Image
	u.Address = cfg.Address
	u.AddressID = cfg.AddressID
	u.Phone = cfg.Phone
}

// The Customer struct represents the customers table, extending the User by a tier level and a company name
type Customer struct {
	User    User
	UserID  uint
	Tier    int
	Company string
}
