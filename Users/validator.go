package Users

import (
	"fmt"
	"regexp"
	"time"
)

func DateCheck(bod, check time.Time) bool {
	return check.After(bod)
}

func (input RegisterInput) RegisterValiator() string {

	var pass string

	if !ValidateEmail(input.Email) {
		pass = "Email address is invalid"
	} 
	if len(input.Password) < 8 {
		pass =  "Password should more than 8"
	}

	if input.Password != input.PasswordConfirmation{
		pass =  "Confirmation password must same as your password "
	}

	BOD, _ := time.Parse(time.RFC3339, input.BOD)

	CurrentDate, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	fmt.Println(input.BOD,CurrentDate)


	if !DateCheck(BOD, CurrentDate){
		
		pass = "Your birthday cannot more than today"
	}
	
	return pass
}

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}