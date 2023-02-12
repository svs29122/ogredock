package webmod

import(
	"fmt"
	"net/http"
	"html/template"
	"ogredock/contmod"
)

func hello (w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("./webmod/templates/tmpl.html")
	t.Execute(w, "Welcome to OgreDock")
	//fmt.Fprintf(w, "Welcome to OgreDock!")
}

func listContainers (w http.ResponseWriter, r *http.Request){
	containers := contmod.GetContainers()

	t, _ := template.ParseFiles("./webmod/templates/ctable.html")

	t.Execute(w, containers)
}

func StartServer() {

	fmt.Printf("starting server...")

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", hello);
	http.HandleFunc("/list", listContainers);

	server.ListenAndServe()
}
