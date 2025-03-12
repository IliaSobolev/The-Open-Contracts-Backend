package main

import (
	"TOC/internal/api"
	codeblockRepo "TOC/internal/codeblock/repository/mongo"
	codeblockUsecase "TOC/internal/codeblock/usecase"
	"TOC/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strconv"
)

func main() {
	//logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	pretty, err := strconv.ParseBool(os.Getenv("LOG_PRETTY"))
	if err != nil {
		pretty = false
	}
	if pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	switch os.Getenv("LOG_LEVEL") {
	case "DISABLED":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	case "PANIC":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "FATAL":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	//DB
	log.Debug().Msg("DB Connection")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal().Err(err).Msg("DB Connection")
	}
	db := client.Database(utils.GetEnv("MONGO_DB", "toc"))

	//PING
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal().Err(err).Msg("DB ping")
	}

	//Repositories
	codeblockRepo := codeblockRepo.NewCodeblockRepo(db)

	//Usecases
	codeblockUC := codeblockUsecase.NewCodeblockUsecase(codeblockRepo)

	//rest api
	r := gin.New()
	api := api.NewAPI(r, codeblockUC)
	err = api.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("API Start")
	}
}
