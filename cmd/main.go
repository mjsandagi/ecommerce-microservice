package main
import "log"

func main() {
	cfg := config{
		address: ":8080",
		db:      dbConfig{},
	}
	api := application{
		config: cfg,
	}
	handler := api.mount()
	err := api.run(handler)
	if err != nil {
		log.Fatal("There was an error: ", err) // There's no need to have this followed by an "os.Exit(1)", as log.Fatal is just log.Print followed by an os.Exit(1).
	}
}
