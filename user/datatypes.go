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
	AddressID uint
	Phone     string
}

// The User struct represents the users table with an id, a username and the general information
type User struct {
	ID       uint
	Username string
	Config
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
