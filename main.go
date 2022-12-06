package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

func add_user(AdSoyad string, Telefon string, Eposta string) bool {
	db, err := sql.Open("mysql", "root:1Mhszxisq4r@tcp(127.0.0.1:3306)/deneme1")
	if err != nil {
		panic(err)
	}
	add, err := db.Query("INSERT INTO  kullanici (AdSoyad,Telefon,Eposta)VALUES (?,?,?)", (AdSoyad), (Telefon), (Eposta))
	if err != nil {
		panic(err)
	}
	fmt.Println(add)
	defer db.Close()
	return true
}

func check_user(AdSoyad string, Telefon string, Eposta string) bool {
	db, err := sql.Open("mysql", "root:1Mhszxisq4r@tcp(127.0.0.1:3306)/deneme1")
	if err != nil {
		panic(err)
	}
	var exists bool
	var query string
	query = fmt.Sprintf("SELECT EXISTS(SELECT AdSoyad FROM kullanici WHERE AdSoyad='%s' AND Telefon='%s'AND Eposta='%s')", (AdSoyad), (Telefon), (Eposta))
	row := db.QueryRow(query).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	fmt.Println(row)
	defer db.Close()
	return exists
}
func login(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("templates/login.html"))
	tmplt.Execute(w, nil)
}
func sign_up(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("templates/sign_up.html"))
	tmplt.Execute(w, nil)
}
func signup_user(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var AdSoyad = r.Form["fname"]
	var Telefon = r.Form["phone"]
	var Eposta = r.Form["mail"]
	fmt.Println(AdSoyad, Telefon, Eposta)
	if add_user(AdSoyad[0], Telefon[0], Eposta[0]) {
		var tmplt = template.Must(template.ParseFiles("templates/index.html"))
		tmplt.Execute(w, nil)
	} else {
		var tmplt = template.Must(template.ParseFiles("templates/error.html"))
		tmplt.Execute(w, nil)
	}
}
func login_user(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var AdSoyad = r.Form["fname"]
	var Telefon = r.Form["phone"]
	var Eposta = r.Form["mail"]
	fmt.Println(AdSoyad, Telefon, Eposta)
	if check_user(AdSoyad[0], Telefon[0], Eposta[0]) {
		var tmplt = template.Must(template.ParseFiles("templates/index.html"))
		tmplt.Execute(w, nil)
	} else {
		var tmplt = template.Must(template.ParseFiles("templates/error.html"))
		tmplt.Execute(w, nil)
	}
}

func main() {

	http.HandleFunc("/", login)
	http.HandleFunc("/sign_up", sign_up)
	http.HandleFunc("/login_user", login_user)
	http.HandleFunc("/signup_user", signup_user)
	http.ListenAndServe(":8000", nil)
}
