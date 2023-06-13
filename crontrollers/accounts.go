package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mau005/KryManagerUpdate/db"
	"github.com/Mau005/KryManagerUpdate/models"
	"github.com/golang-jwt/jwt/v5"
)

// Funcion que me permite ingresar el usuario
func Login(w http.ResponseWriter, r *http.Request) {
	accounts := &models.Accounts{}

	err := json.NewDecoder(r.Body).Decode(&accounts)

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOne(accounts.Email, accounts.Password) //Revisamos y chekeamos en la bd
	json.NewEncoder(w).Encode(resp)
}

// BD se encuentra en sha1 por ende tenemos que revisarla como tal
func PreparaingSha(password string) string {
	preparaing := sha1.New()
	preparaing.Write([]byte(password))
	return hex.EncodeToString(preparaing.Sum(nil))

}

func FindOne(email, password string) map[string]interface{} {
	user := &models.Accounts{}
	if err := db.DB.Where("email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	if PreparaingSha(password) != user.Password { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Claims{
		Name:             user.Name,
		Email:            user.Email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = tk
	return resp
}
