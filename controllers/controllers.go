package controllers

import (
	"backend/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var db = database.CreateDatabase()

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	log.Println(r)
	rows, err := db.Query("SELECT * FROM user;")
	if err != nil {
		log.Println("Database query failed cause by:", err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.CompanyId, &user.Name, &user.HashedPassword, &user.Phone, &user.NotifyType, &user.LineNotifyToken)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("HTTP request decode failed cause by :", err)
		return
	}
	r.Body.Close()
	stmt, _ := db.Prepare("INSERT user SET companyid=?, name=?, hashedPassword=?, phone=?, notifyType=?, lineNotifyToken=?")
	if _, err := stmt.Exec(user.CompanyId, user.Name, user.HashedPassword, user.Phone, user.NotifyType, user.LineNotifyToken); err != nil {
		log.Println("HTTP request decode failed cause by :", err)
		return
	}
	stmt.Close()
	responsewithJSON(w, http.StatusCreated, map[string]string{"message": "success"})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	param := mux.Vars(r)
	rows := db.QueryRow("SELECT * FROM user where id=?", param["id"])
	if err := rows.Scan(&user.Id, &user.CompanyId, &user.Name, &user.HashedPassword, &user.Phone, &user.NotifyType, &user.LineNotifyToken); err != nil {
		log.Println("Database Query failed cause by :", err)
		responsewithJSON(w, http.StatusNotFound, map[string]string{"404 page not found": "Can't find user"})
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])
	if err := db.QueryRow("SELECT * FROM user where id=?", param["id"]).Scan(&user); err != nil {
		responsewithJSON(w, http.StatusNotFound, map[string]string{"404 page not found": "Can't find user"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("HTTP request decode failed cause by :", err)
		return
	}
	r.Body.Close()
	stmt, _ := db.Prepare("UPDATE user SET companyid=?, name=?, hashedPassword=?, phone=?, notifyType=?, lineNotifyToken=? where id=?")
	stmt.Exec(user.CompanyId, user.Name, user.HashedPassword, user.Phone, user.NotifyType, user.LineNotifyToken, id)
	stmt.Close()
	responsewithJSON(w, http.StatusOK, map[string]string{"message": "success"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])
	if err := db.QueryRow("SELECT * FROM user where id=?", param["id"]).Scan(&user); err != nil {
		responsewithJSON(w, http.StatusNotFound, map[string]string{"404 page not found": "Can't find user"})
		return
	}
	stmt, _ := db.Prepare("DELETE from user where id=?")
	stmt.Exec(id)
	stmt.Close()
	responsewithJSON(w, http.StatusOK, map[string]string{"message": "success"})
}

func responsewithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type User struct {
	Id              int
	CompanyId       int
	Name            string
	HashedPassword  string
	Phone           string
	NotifyType      string
	LineNotifyToken string
}
