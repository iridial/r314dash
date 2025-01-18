//go:generate npm run build

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"r314dash/config"
	"r314dash/handler"

	"github.com/gorilla/mux"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		panic(err)
	}

	fmt.Println("Serving at port:", config.Main.Settings.Port)
	fmt.Println("Serving at path:", config.Main.Settings.Path)
	fmt.Println("Static  at path:", config.Main.Settings.StaticPath)
	fs := http.FileServer(http.Dir("static"))
	router := mux.NewRouter()
	router.PathPrefix(config.Main.Settings.StaticPath).Handler(http.StripPrefix(config.Main.Settings.StaticPath, fs))
	router.HandleFunc(config.Main.Settings.Path, handler.HomeHandler).Methods("GET")
	if config.Main.Settings.Stats.Enabled {
		router.HandleFunc(config.Main.Settings.StatsPath, handler.StatsHandler).Methods("GET")
		fmt.Println("Stats   at path:", config.Main.Settings.StatsPath)
	} else {
		fmt.Println("Stats   disabled")
	}

	http.ListenAndServe(":"+strconv.Itoa(config.Main.Settings.Port), router)
}
