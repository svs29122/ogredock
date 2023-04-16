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

		containers := contmod.GetContainers()
		selectedContainers := []string{}
		for _, container := range containers {
			selected := r.FormValue(container.ID)
			if selected != "" {
				selectedContainers = append(selectedContainers, container.ID)
			}
		}

		//fmt.Printf("%s %s %s %s\n", cmds, name, net, img);

		if cmds[0] != "" {
			for i := 0; i < len(selectedContainers); i++ {
				err := contmod.StartContainer(selectedContainers[i])
				if err != nil {
					fmt.Printf("Start Container failed")
				}
			}
			if name != "" && net != "" && img != "" {
				contId, err := contmod.CreateContainer(name, net, img, "");
				err = contmod.StartContainer(contId)
				if err != nil {
					fmt.Printf("Start Containerfailed)")
				}
			}
		} else if cmds[1] != "" {
			for i := 0; i < len(selectedContainers); i++ {
				err := contmod.StopContainer(selectedContainers[i])
				if err != nil {
					fmt.Printf("Stop Container failed")
				}
			}
		} else if cmds[2] != "" {
			_, err := contmod.CreateContainer(name, net, img, "");
			if err != nil {
				fmt.Printf("Create Container failed")
			}
		} else if cmds[3] != "" {
			for i := 0; i < len(selectedContainers); i++ {
				err := contmod.DestroyContainer(selectedContainers[i])
				if err != nil {
					fmt.Printf("Destroy Container failed")
				}
			}
		}
	}

	containers := contmod.GetContainers()

	t, _ := template.ParseFiles("./webmod/templates/ctable.html")
	t.Execute(w, containers)
}


func networkManagement (w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		r.ParseForm();

		name := r.FormValue("Name")
		pif := r.FormValue("ParentIF")
		drv := r.FormValue("Driver")
		ipr := r.FormValue("IPRange")
		sn := r.FormValue("Subnet")
		gw := r.FormValue("Gateway")
		cmds := []string{r.FormValue("create_button"),
							r.FormValue("destroy_button"),
							r.FormValue("inspect_button")}

		networks := contmod.GetNetworks()
		selectedNetworks := []string{}
		for _, network := range networks {
			selected := r.FormValue(network.ID)
			if selected != "" {
				selectedNetworks = append(selectedNetworks, network.ID)
			}
		}

		//fmt.Printf("%s\n %s %s %s\n %s %s %s\n", cmds, name, pif, drv, ipr, sn, gw)

		if cmds[0] != "" {
			contmod.CreateNetwork(name, pif, drv, ipr, sn, gw)
		} else if cmds[1] != "" {
			for i := 0; i < len(selectedNetworks); i++ {
				err := contmod.DestroyNetwork(selectedNetworks[i])
				if err != nil {
					fmt.Printf("Destroy Network Failed\n")
				}
			}
		} else if cmds[2] != "" {
			for i := 0; i < len(selectedNetworks); i++ {
				//err := contmod.DestroyNetwork(selectedNetworks[i])
				//if err != nil {
					fmt.Printf("Inspect Network Failed\n")
				//}
			}
		}
	}

	networks := contmod.GetNetworks()

	t, _ := template.ParseFiles("./webmod/templates/networks.html")
	t.Execute(w, networks)
}

func bulkGeneration (w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("./webmod/templates/generate.html")
	t.Execute(w, "Bulk Generation")

}

func StartServer() {

	fmt.Printf("starting server...\n")

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", home);
	http.HandleFunc("/cmv", containerManagement);
	http.HandleFunc("/nmv", networkManagement);
	http.HandleFunc("/bgv", bulkGeneration);

	server.ListenAndServe()
}
