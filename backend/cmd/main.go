package main

import (
	"github.com/DomEscobar/erp-dev-bench/internal/config"
	"github.com/DomEscobar/erp-dev-bench/internal/db"
	"github.com/DomEscobar/erp-dev-bench/internal/server"
)

func main() {
	db.Init()
	defer db.Close()

	cfg := config.Load()
	server.Start(cfg)
}
