package main

func main() {
	cfg := config{
		address: ":8080",
		db:      dbConfig{},
	}
	api := application{
		config: cfg,
	}
}
