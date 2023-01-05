package webmod

import(
	"fmt"
	"net/http"
	"ogredock/contmod"
)

type HelloHandler struct{}
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to OgreDock!")
}

type ContainerListHandler struct{}
func (h *ContainerListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	containers := contmod.GetContainers()

	for _, container := range containers{
		fmt.Fprintf(w, "%s %s %s\n", container.ID[:10], container.Image, container.Names[0])
	}
}

func StartServer() {

	hello := HelloHandler{}
	clist := ContainerListHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/", &hello);
	http.Handle("/list", &clist);

	server.ListenAndServe()
}
