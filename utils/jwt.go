package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)


	
	
var secretKey string
func init() {
	_ = godotenv.Load()
	secretKey = os.Getenv("SECRET_KEY")
	if secretKey == "" {
		panic("SECRET_KEY environment variable is required")
	}
}

func GenerateToken(email string,userId int64) (string, error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour*5).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
func VerifyToken(token string) (int64,error){
parsedToken,err:=jwt.Parse(token,func(token *jwt.Token)(any,error){
	_,ok:=token.Method.(*jwt.SigningMethodHMAC)
	if !ok{
		return nil,errors.New("unexpected signing method")
	}
return []byte(secretKey), nil
})
if err!=nil{
return 0,errors.New("Couldn't parse token")
}
tokenIsValid:=parsedToken.Valid
if !tokenIsValid{
	return 0,errors.New("Invalid Token")
}
claims,ok:=parsedToken.Claims.(jwt.MapClaims)
if !ok{
		return 0,errors.New("Invalid claims")
	}
	// email:=claims["email"].(string)
	userId:=int64(claims["userId"].(float64))
	return userId,nil
}