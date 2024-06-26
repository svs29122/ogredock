package webmod

import(
	"fmt"
	"strconv"

	"net/http"
	"html/template"
	"ogredock/contmod"
)

func index (w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("./webmod/templates/index.html")
	t.Execute(w, "Welcome to")
}

func multiWindowView (w http.ResponseWriter, r *http.Request){

	t, _ := template.ParseFiles("./webmod/templates/multiView.html")
	t.Execute(w, nil)

}

func containerManagement (w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		r.ParseForm();

		name := r.FormValue("Name")
		net := r.FormValue("Net")
		img := r.FormValue("Img")
		cmds := []string{r.FormValue("start_button"),
								r.FormValue("stop_button"),
								r.FormValue("create_button"),
								r.FormValue("destroy_button")}

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
			if name != "" && pif != "" && drv != "" && ipr != "" && sn != ""  && gw != "" {
				contmod.CreateNetwork(name, pif, drv, ipr, sn, gw)
			}
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

	//TODO: consider sorting this slice
	networks := contmod.GetNetworks()

	t, _ := template.ParseFiles("./webmod/templates/networks.html")
	t.Execute(w, networks)
}

func containerGeneration (w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		r.ParseForm();

		bname := r.FormValue("BName")
		net := r.FormValue("Net")
		img := r.FormValue("Img")
		rlow := r.FormValue("RLow")
		rhigh := r.FormValue("RHigh")
		cmds := []string{ r.FormValue("generate_button"),
							   r.FormValue("terminate_button")}

		//fmt.Printf("%s %s %s %s\n", cmds, bname, net, img);

		if cmds[0] != "" {
			if bname != "" && net != "" && img != "" && rlow != "" && rhigh != "" {
				//get the ranges
				rangeLow, errLow := strconv.Atoi(rlow)
				if errLow != nil {
					panic(errLow)
				}
				rangeHigh, errHigh := strconv.Atoi(rhigh)
				if errHigh != nil {
					panic(errHigh)
				}

				//start the containers
				for i := rangeLow; i <= rangeHigh; i++ {
					name := bname + "-" + strconv.Itoa(i)
					contId, err := contmod.CreateContainer(name, net, img, "");
					err = contmod.StartContainer(contId)
					if err != nil {
						fmt.Println("Start Container failed)")
					}
				}
			} else {
					fmt.Println("Generate Container failed")
			}
		} else if cmds[1] != "" {
			if bname != "" && rlow != "" && rhigh != "" {
				//get the ranges
				rangeLow, errLow := strconv.Atoi(rlow)
				if errLow != nil {
					panic(errLow)
				}
				rangeHigh, errHigh := strconv.Atoi(rhigh)
				if errHigh != nil {
					panic(errHigh)
				}

				destroyList := make(map[string]bool)
				for i := rangeLow; i <= rangeHigh; i++ {
					name := "/" + bname + "-" + strconv.Itoa(i)
					destroyList[name] = true
				}

				//terminate the containers
				containers := contmod.GetContainers()
				for _, container := range containers {
					if destroyList[container.Names[0]] {
						errd := contmod.DestroyContainer(container.ID)
						if errd != nil {
							panic(errd)
						}
					}
				}
			} else {
					fmt.Println("Terminate Container failed")
			}
		}
	}

	containers := contmod.GetContainers()

	t, _ := template.ParseFiles("./webmod/templates/generate.html")
	t.Execute(w, containers)
}

func StartServer() {

	fmt.Printf("starting server...\n")

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	fs := http.FileServer(http.Dir("webmod/public"))
	http.Handle("/webmod/public/", http.StripPrefix("/webmod/public/", fs))

	http.HandleFunc("/", index);
	http.HandleFunc("/cmv", containerManagement);
	http.HandleFunc("/nmv", networkManagement);
	http.HandleFunc("/cgv", containerGeneration);
	http.HandleFunc("/mwv", multiWindowView);

	server.ListenAndServe()
}
