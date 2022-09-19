package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mutahirlatif/go-clean-architecture/auth"
	"github.com/mutahirlatif/go-clean-architecture/task"

	authhttp "github.com/mutahirlatif/go-clean-architecture/auth/delivery/http"
	authpostgres "github.com/mutahirlatif/go-clean-architecture/auth/repository/postgres"
	authusecase "github.com/mutahirlatif/go-clean-architecture/auth/usecase"

	thttp "github.com/mutahirlatif/go-clean-architecture/task/delivery/http"
	tpostgres "github.com/mutahirlatif/go-clean-architecture/task/repository/postgres"
	tusecase "github.com/mutahirlatif/go-clean-architecture/task/usecase"
)

type App struct {
	httpServer *http.Server

	// bookmarkUC bookmark.UseCase
	authUC auth.UseCase
	taskUC task.UseCase
}

func NewApp() *App {
	db := initPostGresDB()

	// userRepo := authmongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	// bookmarkRepo := bmmongo.NewBookmarkRepository(db, viper.GetString("mongo.bookmark_collection"))
	// taskRepo := tmongo.NewTaskRepository(db, viper.GetString("mongo.task_collection"))
	userRepo := authpostgres.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	taskRepo := tpostgres.NewTaskRepository(db, viper.GetString("mongo.user_collection"))

	return &App{
		// bookmarkUC: bmusecase.NewBookmarkUseCase(bookmarkRepo),
		authUC: authusecase.NewAuthUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl"),
		),
		taskUC: tusecase.NewTaskUseCase(taskRepo),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	// SignUp/SignIn endpoints
	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	// API endpoints
	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware)
	// bmhttp.RegisterHTTPEndpoints(api, a.bookmarkUC)
	thttp.RegisterHTTPEndpoints(api, a.taskUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

// func initDB() *mongo.Database {

// 	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
// 	if err != nil {
// 		log.Fatalf("Error occured while establishing connection to mongoDB")
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(context.Background(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return client.Database(viper.GetString("mongo.name"))
// }

func initPostGresDB() *gorm.DB {
	dbURL := viper.GetString("postgres.uri")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", "postgres")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", "postgres")
	}

	return db
}
