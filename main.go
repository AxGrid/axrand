package main

import (
	"axrand/internal"
	"context"
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/ironstar-io/chizerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

/*
 _____  __  __  _____  _____  ___  _____
/  _  \/  \/  \/   __\/  _  \/___\|  _  \
|  _  |>-    -<|  |_ ||  _  <|   ||  |  |
\__|__/\__/\__/\_____/\__|\_/\___/|_____/
zed (17.09.2024)
*/

var (
	host           = ":8000"
	release        = false
	logFormat      = "2006-01-02 15:04:05"
	reSeedInterval = 10000
	bufferSize     = 1000
	workerCount    = 10
)

func init() {
	flag.BoolVar(&release, "release", false, "Release mode")
	flag.StringVar(&host, "host", ":8000", "Bind host")
	flag.IntVar(&reSeedInterval, "reseed-interval", 10000, "ReSeed interval")
	flag.IntVar(&bufferSize, "buffer-size", 1000, "Buffer size")
	flag.IntVar(&workerCount, "worker-count", 10, "Worker count")
	flag.Parse()
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if release {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: logFormat}).Level(zerolog.InfoLevel)
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: logFormat}).Level(zerolog.DebugLevel)
	}
	r := chi.NewRouter()
	r.Use(chizerolog.LoggerMiddleware(&log.Logger))
	ctx := context.Background()
	service, err := internal.NewRandomGenerationService(ctx, workerCount, bufferSize, reSeedInterval)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to create random generation service")
	}

	r.Get("/api/uint64", getHandler(internal.RequestTypeUint64, service))
	r.Get("/api/int64", getHandler(internal.RequestTypeInt64, service))
	r.Get("/api/float64", getHandler(internal.RequestTypeFloat64, service))
	log.Info().Str("host", host).Msg("start random server")
	if err := http.ListenAndServe(host, r); err != nil {
		log.Fatal().Err(err).Msg("fail to start server")
	}
}

func getHandler(t internal.RequestTypes, service *internal.RandomGenerationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := service.GetRandomResultJson(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(res)
	}
}
