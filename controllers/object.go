package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	//beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}
type User struct {
	FirstName string `form:"fname"`
	LastName  string `form:"lname"`
	Phone     string `form:"phone"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	DoB       string `form:"dob"`
}

func (u *User) Valid(v *validation.Validation) {
	if strings.Contains(u.FirstName, "admin") {
		// Set error messages of Name by SetError and HasErrors will return true
		v.SetError("FName", "Can't contain 'admin' in Name")
	}
	if strings.Contains(u.LastName, "admin") {
		// Set error messages of Name by SetError and HasErrors will return true
		v.SetError("LName", "Can't contain 'admin' in Name")
	}
	rep := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	ph := rep.MatchString(u.Phone)

	rem := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	em := rem.MatchString(u.Email)

	if !ph {
		v.SetError("Phone", "Format supported 8801XXXXXXX")
	}

	if !em {
		v.SetError("Email", "Format supported any@any.any")
	}

}
func (o *ObjectController) Post() {
	o.TplName = "index.tpl"
	flash := beego.NewFlash()
	ob := User{}
	fname := o.GetString("fname")
	lname := o.GetString("lname")
	phone := o.GetString("phone")
	email := o.GetString("email")
	pass := o.GetString("password")
	dob := o.GetString("dob")
	if err := o.ParseForm(&ob); err != nil {
		//handle error
		fmt.Println("Error")
	}
	hash, _ := HashPassword(ob.Password)
	valid := validation.Validation{}
	b, err := valid.Valid(&ob)
	if err != nil {
		// handle error
		fmt.Println("Input Invalid")
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			if err.Key == "FName" {
				o.Data["Fname"] = "Can't contain 'admin' in Name"
				flash.Error("Can't contain 'admin' in Name")
				flash.Store(&o.Controller)
				return
			}
			if err.Key == "LName" {
				o.Data["Lname"] = "Can't contain 'admin' in Name"
				flash.Error("Can't contain 'admin' in Name")
				flash.Store(&o.Controller)
				return

			}
			if err.Key == "Email" {
				o.Data["Email"] = "Email invalid! (any@any.any)"
				flash.Notice("Email invalid! (any@any.any)")
				flash.Store(&o.Controller)
				return

			}
			if err.Key == "Phone" {
				o.Data["Phone"] = "Phone number invalid! (8801XXXXXXX)"
				flash.Warning("Phone number invalid! (8801XXXXXXX)")
				flash.Store(&o.Controller)
				return

			}

		}
	} else {
		o.Data["Success"] = "Succesfully data sent to API"
		flash.Success("Succesfully data sent to API")
		flash.Store(&o.Controller)
		fmt.Println(fname, lname, phone, email, pass, dob, hash)
		postBody, _ := json.Marshal(map[string]string{
			"FirstName": fname,
			"LastName":  lname,
			"Password":  pass,
			"Email":     email,
			"Phone":     phone,
			"DoB":       dob,
		})
		//fmt.Println(postBody)
		responseBody := bytes.NewBuffer(postBody)
		//Leverage Go's HTTP Post function to make request
		resp, err := http.Post("http://localhost:8080/v1/object", "application/json", responseBody)
		//Handle Error
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)

		fmt.Println("All OKAY", sb)
	}
	// var ob models.Object
	// json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	// objectemail := models.AddOne(ob)
	// o.Data["json"] = map[string]string{"Email": ob.Email}
	// o.ServeJSON()
	//Objects[object.Email] = &object
	fmt.Println("ALL GOOD")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
