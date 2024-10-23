package main

import (
	"log"

	"github.com/AthThobari/simple_music_catalog_go/internal/configs"
	"github.com/AthThobari/simple_music_catalog_go/internal/models/memberships"
	membershipsRepo "github.com/AthThobari/simple_music_catalog_go/internal/repository/memberships"
	membershipsSvc "github.com/AthThobari/simple_music_catalog_go/internal/service/memberships"
	membershipsHandler "github.com/AthThobari/simple_music_catalog_go/internal/handler/memberships"
	"github.com/AthThobari/simple_music_catalog_go/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database, err: %+v", err)
	}
	log.Println("Connecting to database: ", cfg.Database.DataSourceName)

	if err := db.AutoMigrate(&memberships.User{}); err != nil {
		log.Fatalf("failed to auto-migrate, err: %+v", err)
	}
	log.Println("AutoMigrate completed successfully")
	
	r:= gin.Default()

	membershipsRepo := membershipsRepo.NewRepository(db)

    membershipsSvc := membershipsSvc.NewService(cfg, membershipsRepo)	

	membershipsHandler := membershipsHandler.NewHandler(r, membershipsSvc)

	membershipsHandler.RegisterRoute()
	

	r.Run(cfg.Service.Port)
}
