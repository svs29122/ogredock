package webmod

import(
	"fmt"
	"net/http"
	"html/template"
	"ogredock/contmod"
)

func home (w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("./webmod/templates/index.html")
	t.Execute(w, "Welcome to OgreDock")

	//writes text to the Response Writer directly as opposed to using a template
	//fmt.Fprintf(w, "Welcome to OgreDock!")
}

func containerManagement (w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		r.ParseForm();

		name := r.FormValue("Name")
		net := r.FormValue("Net")
		img := r.FormValue("Img")
		cmds := []string{r.FormValue("start_button"), r.FormValue("stop_button"),r.FormValue("create_button"),r.FormValue("destroy_button")}

		if cmds[0] != "" {
			contId, err := contmod.CreateContainer(name, net, img, "");
			if err != nil {
				fmt.Printf("Create Container failed")
			} else {
				err = contmod.RunContainer(contId)
				if err != nil {
					fmt.Printf("Run Container failed")
				} else {
						fmt.Printf("%s %s %s %s\n", cmds[0], name, net, img);
				}
			}
		} else if cmds[1] != "" {
			fmt.Printf("%s %s %s %s\n", cmds[0], name, net, img);
		} else if cmds[2] != "" {
			_, err := contmod.CreateContainer(name, net, img, "");
			if err != nil {
				fmt.Printf("Create Container failed")
			} else {
				fmt.Printf("%s %s %s %s\n", cmds[0], name, net, img);
			}
		} else if cmds[3] != "" {
			fmt.Printf("%s %s %s %s\n", cmds[0], name, net, img);
		}
	}

	containers := contmod.GetContainers()

	t, _ := template.ParseFiles("./webmod/templates/ctable.html")
	t.Execute(w, containers)
}

func StartServer() {

	fmt.Printf("starting server...\n")

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", home);
	http.HandleFunc("/cmv", containerManagement);

	server.ListenAndServe()
}
