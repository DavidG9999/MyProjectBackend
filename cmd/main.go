package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/DavidG9999/MyProject/base"
	"github.com/DavidG9999/MyProject/implementation"
	"github.com/DavidG9999/MyProject/transport"
	httptransport "github.com/DavidG9999/MyProject/transport/http"
	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {

	db, err := base.ConnectToDb()
	if err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
	defer db.Close()
	productRepo := implementation.NewProductRepo(db)
	unitRepo := implementation.NewUnitRepo(db)
	svc := implementation.NewService(productRepo, unitRepo)
	endpoints := transport.MakeEndpoints(svc)
	h := httptransport.NewService(endpoints, []kithttp.ServerOption{})
	fmt.Println("Listening on 7071...")
	if err := http.ListenAndServe(":7071", h); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
