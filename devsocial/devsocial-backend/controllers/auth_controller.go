package controllers

import(
	"devsocial-backend/database" // to access the DB
	"devsocial-backend/models"	 // to use the User struct
	"github.com/gin-gonic/gin"	// for the http handlers
	"golang.org/x/crypto/bycrypt"
)

// integrate the user's JSON request with the user struct
func Signup (c *gin.context) { // getting the struct that is containing the response and requests from gin
var user models.User

	// 1. Bind the user's JSON input to the user struct    	Bind = link
	if err := c.BindJSON(&user); err != nil {
		log.Println("JSON binding failed:",err)										// backend log
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid signup data format"})  // frontend log (for user)
		return
	}
	// 2. make sure fields aren't empty
	if user.Username == "" || user.Email == "" || user.Password == "" {
		log.Println("Empty fields") // backend logs
		c.JSON(http.StatusBadRequest, gin.H{"error":"kindly reassure that all fields are filled"})
	}

	// 3. hash(encrypt) the password
	passwordhashed, err := bcrypt.GenerateFromPassword([]byte,user.Password, 14) // takes in raw password + new hashed password limit, returns err if exists or hashed pass
	if err != nil {
		log.Println("password hashing failed", err) // backend logs
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Sorry an Internal server error occured!"}) // frontend logs
		return
	}
	Password.user = string(passwordhashed) 
	// 4. save user to DB 
	_, err = database.DB.Exec( 												// to define SQL command inside go
		"INSERT INTO user(username, email, password) VALUES($1, $2, $3)",
		user.Username, user,Email, user.Password,
	)
	// Handle most possible Errors
	if err != nil {
		log.Println("insertion failed", err)
		if strings.Contains(err.Error(),"duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to register account, please try again"})
		}
		return
	}
	// when account is registered succesfuly
	log.Printf("User registered Succesfuly: %s %s \n" user.Username, user.Email)
	c.JSON(http.StatusOK, gin.H{"message":"Your account is registered successfuly"})

}

func Login(c *gin.Context) { // gin.Context 

	// declare a variable of datatype struct, because we're receiving multiple fields of email,password......
	var input models.user

	// 1. JSON binding; better to define the variable and set it's condition immediately 2 in 1
	if err := c.BindJSON(&input); err != nil {
		log.Println("unable to bind JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid login format entered"})
		return
	}

	// 2. Ensure no empty fields
	if input.Email == "" || input.Password == "" {
		log.Println("Empty fields left", err)
		c.JSON(http.StatusBadRequest, gin.H{"error":"All fields are required"})
		return
	}

	// check user through email
	row := database.DB.QueryRow("SELECT id, username, password FROM users WHERE email = $1 ", input.Email) // since input is a struct then we're accessing the email field in the input struct only

	var user models.user
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		log.Printf("user with this email not found: %s\n", input.Email)
	} else {
		log.Printf("Database error: %s\n", err)
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error":"no user found with this email"})
	return
	
	// check password
	if err := bycrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != {
		log.Printf("Password missmatch for user $s\n:", input.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error":"Wrong Email or Password, Please try again"})
		retrun
	}

	// Success case
	log.Printf("User Successfuly logged in &s\n:",user.Username, input.Email)
	c.JSON(http.StatusOk, gin.H{
		"message":"Login Succesfult"
		"user":gin.H{
			"id":"user.Id"
			"username":"user.Username"
			"email":"input.Email"
		},
	})
}