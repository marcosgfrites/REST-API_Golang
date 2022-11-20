package common

const (
	//	user register constants

	RegisterError       = "An error occurred while obtaining the user data: "
	RequiredEmailError  = "The email is a required field"
	PasswordLengthError = "The minimum password length is 8"
	UserAlreadyExists   = "Already exists an user with the specified email"
	UserSaveError       = "An error occurred while trying to save a new user in the database: "
	UserPersistError    = "Persist new user data in the database has been failed"
)

const (
	// database constants

	DatabaseConnectionFailed     = "Database connection failed"
	DatabaseConnectionLost       = "Database connection has been lost"
	DatabaseConnectionSuccessful = "The connection to the database is successful"
)
