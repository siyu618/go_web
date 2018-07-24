package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	"html/template"
	"os"
	"time"
	"crypto/md5"
	"io"
	"strconv"
)

func main() {
	//testOutput()
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
func testOutput() {
	t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err := t.ExecuteTemplate(os.Stdout, "T", template.HTML("<script>alert('you have been pwned')</script>"))


	t, _ = template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	fmt.Println(err)
}
func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		curTime := time.Now().Unix()
		h := md5.New()

		io.WriteString(h, strconv.FormatInt(curTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login4.4.gtpl")
		t.Execute(w, token)
	} else {

		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// valid ok
		} else {
			// invalid
		}

		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}


func sayHelloName(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for k,v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(writer, "Hello astaxie!")
}

func upload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		curTime := time.Now().Unix()
		h := md5.New()

		io.WriteString(h, strconv.FormatInt(curTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32<<20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}