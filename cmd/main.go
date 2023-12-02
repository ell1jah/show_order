package main

import (
	"fmt"
	"net/http"

	"github.com/ell1jah/db_cp/pkg/middleware"
	"github.com/ell1jah/show_order/internal/config"
	"github.com/ell1jah/show_order/internal/delivery/rest"
	"github.com/ell1jah/show_order/internal/delivery/stan"
	"github.com/ell1jah/show_order/internal/logic"
	inmemory "github.com/ell1jah/show_order/internal/repository/immemory"
	pgRepo "github.com/ell1jah/show_order/internal/repository/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())
	sugaredLogger := logger.Sugar()

	config, err := config.Parse()
	if err != nil {
		sugaredLogger.Fatalf("error parse config: %s", err)
	}
	sugaredLogger.Info("config was loaded")

	params := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
		config.DB.User,
		config.DB.DBName,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(params), &gorm.Config{})
	if err != nil {
		sugaredLogger.Fatalf("error connect to db: %s", err)
	}
	sugaredLogger.Info("db was connected")

	orderRepo := pgRepo.NewRepository(db)
	orderCache := inmemory.NewCache()

	orderLogic := logic.NewLogic(orderCache, orderRepo)

	natsServ := stan.NewStanManager(orderLogic, logger, config)
	err = natsServ.Run()
	if err != nil {
		sugaredLogger.Fatalf("error run nats: %s", err)
	}
	sugaredLogger.Info("nats was runned")
	defer natsServ.Stop()

	orderHandler := rest.NewOrderRest(orderLogic, *logger)

	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, config.App.StaticDir)
	})
	r.HandleFunc("/", http.FileServer(http.Dir(config.App.StaticDir)).ServeHTTP).
		Methods("GET")
	r.HandleFunc("/orders/{ORDER_UID:[0-9]+}", orderHandler.GetOrder).Methods("GET")

	mux := middleware.AccessLog(sugaredLogger, r)
	mux = middleware.Panic(sugaredLogger, mux)

	sugaredLogger.Infow("starting server",
		"port", config.App.Port,
	)

	sugaredLogger.Errorln(http.ListenAndServe(config.App.Port, mux))

	err = logger.Sync()
	if err != nil {
		fmt.Println(err)
	}
}
