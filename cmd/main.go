package main

import (
	"log"
	"net/http"

	"github.com/AthThobari/simple_music_catalog_go/internal/configs"
	membershipsHandler "github.com/AthThobari/simple_music_catalog_go/internal/handler/memberships"
	trackActivitiesRepo "github.com/AthThobari/simple_music_catalog_go/internal/repository/trackactivities"
	trackHandler "github.com/AthThobari/simple_music_catalog_go/internal/handler/tracks"
	"github.com/AthThobari/simple_music_catalog_go/internal/models/memberships"
	"github.com/AthThobari/simple_music_catalog_go/internal/models/trackactivities"
	membershipsRepo "github.com/AthThobari/simple_music_catalog_go/internal/repository/memberships"
	"github.com/AthThobari/simple_music_catalog_go/internal/repository/spotify"
	membershipsSvc "github.com/AthThobari/simple_music_catalog_go/internal/service/memberships"
	"github.com/AthThobari/simple_music_catalog_go/internal/service/memberships/tracks"
	"github.com/AthThobari/simple_music_catalog_go/pkg/httpclient"
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

	if err := db.AutoMigrate(
		&memberships.User{},
		&trackactivities.TrackActivity{},
	); err != nil {
		log.Fatalf("failed to auto-migrate, err: %+v", err)
	}
	log.Println("AutoMigrate completed successfully")

	r := gin.Default()

	httpclient := httpclient.NewClient(&http.Client{})

	spotifyOutbound := spotify.NewSpotifyOutbound(cfg, *httpclient)

	membershipsRepo := membershipsRepo.NewRepository(db)

	membershipsSvc := membershipsSvc.NewService(cfg, membershipsRepo)
	trackActivitiesRepo := trackActivitiesRepo.NewRepository(db)
	trackSvc := tracks.NewService(spotifyOutbound, trackActivitiesRepo)

	membershipsHandler := membershipsHandler.NewHandler(r, membershipsSvc)
	membershipsHandler.RegisterRoute()

	tracksHandler := trackHandler.NewHandler(r, trackSvc)
	tracksHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
