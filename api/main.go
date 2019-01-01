package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/askft/wloggr/api/route"
	"github.com/askft/wloggr/api/store"
	"github.com/askft/wloggr/api/util"
)

func main() {
	usage := fmt.Sprintf("usage: %s -port=[port]\n", os.Args[0])
	if len(os.Args) > 2 {
		fmt.Println(usage)
		return
	}
	portFlag := flag.String("port", util.Config.DefaultPort, "API server port")
	flag.Parse()
	if len(*portFlag) == 0 {
		fmt.Println(usage)
		return
	}
	port := fmt.Sprintf(":%s", *portFlag)

	err := store.SetupDB(util.Config.DBConnectionString)
	if err != nil {
		fmt.Println("Could not set up database.")
		panic(err)
	}

	r := route.SetupRouter()

	fmt.Printf("Starting API server on http://127.0.0.1%s/.\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
