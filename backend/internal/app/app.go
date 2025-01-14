package app

import (
	"backend/config"

	"backend/internal/controllers/http/middleware"

	// handlers
	handleritem "backend/internal/controllers/http/v1/item"
	handleruser "backend/internal/controllers/http/v1/user"
	handleruserorder "backend/internal/controllers/http/v1/user_order"

	// entities
	entityitem "backend/internal/entity/item"
	entityuser "backend/internal/entity/user"
	entityuserorder "backend/internal/entity/user_order"

	// services
	serviceitem "backend/internal/service/item"
	serviceuser "backend/internal/service/user"
	serviceuserorder "backend/internal/service/user_order"

	// storages
	repoitem "backend/internal/storage/mysql/item"
	repouser "backend/internal/storage/mysql/user"
	repouserorder "backend/internal/storage/mysql/user_order"

	"backend/pkg/db"
	"backend/pkg/httpserver"
	"backend/pkg/logger"
	"backend/pkg/validator"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(cfg config.Config) {
	l := logger.NewZap("")
	validator := validator.Validator{}

	// Connect to storage
	dbConn, err := db.NewPostgres(cfg.Postgres, l)
	if err != nil {
		l.Fatal("Unable to make postgres connection: ", err)
	}

	// Auto migrations
	if err = dbConn.DB.AutoMigrate(
		&entityitem.Item{},
		&entityuser.User{},
		&entityuserorder.UserOrder{},
		&entityuserorder.UserOrderItem{},
	); err != nil {
		l.Fatal("Auto migration failed: ", err)
	}

	// Storages
	itemRepo := repoitem.NewRepo(dbConn)
	userRepo := repouser.NewRepo(dbConn)
	userOrderRepo := repouserorder.NewRepo(dbConn)

	// Services
	itemSvc := serviceitem.NewService(itemRepo, l)
	userSvc := serviceuser.NewService(userRepo, userOrderRepo, l, cfg.JWT)
	userOrderSvc := serviceuserorder.NewService(userOrderRepo, l)

	// HTTP server
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	g.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	// Handlers
	middle := middleware.NewMiddleware(cfg.JWT)

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	handleritem.InitHandler(
		g, l,
		itemSvc,
		validator,
		middle,
	)

	handleruser.InitHandler(
		g, l,
		userSvc,
		validator, middle,
	)

	handleruserorder.InitHandler(
		g, l,
		userOrderSvc,
		validator, middle,
	)

	server := httpserver.New(g, httpserver.Port(cfg.Server.Port))

	// Waiting signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Fatal("Signal interrupt error: ", s.String())
	case err := <-server.Notify():
		l.Fatal("Server notify err", err)
	}

	// Shutdown server
	err = server.Shutdown()
	if err != nil {
		l.Info("Server shutdown err: ", err)
	}
}
