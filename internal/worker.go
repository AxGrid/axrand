package internal

import (
	"context"
	cryptorand "crypto/rand"
	"encoding/binary"
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/seehuhn/mt19937"
	"math/rand"
	"time"
)

/*
 _____  __  __  _____  _____  ___  _____
/  _  \/  \/  \/   __\/  _  \/___\|  _  \
|  _  |>-    -<|  |_ ||  _  <|   ||  |  |
\__|__/\__/\__/\_____/\__|\_/\___/|_____/
zed (17.09.2024)
*/

type RandomWorker struct {
	rng            *rand.Rand
	requestChan    chan *RandomRequest
	ctx            context.Context
	reSeedCount    int
	reSeedInterval int
	logger         zerolog.Logger
}

func NewRandomWorker(ctx context.Context, id int, requestChan chan *RandomRequest, reSeedInterval int) (*RandomWorker, error) {
	res := &RandomWorker{
		requestChan:    requestChan,
		rng:            rand.New(mt19937.New()),
		ctx:            ctx,
		reSeedInterval: reSeedInterval,
		logger:         log.With().Int("worker_id", id).Logger(),
	}
	if err := res.ReSeed(); err != nil {
		return nil, err
	}
	go res.run()
	return res, nil
}

func (w *RandomWorker) ReSeed() error {
	secureInt, err := CryptoInt64LE()
	if err != nil {
		return err
	}
	w.rng.Seed(time.Now().UnixNano() + secureInt)
	w.reSeedCount += w.reSeedInterval
	w.logger.Debug().Msg("generated new seed")
	return nil
}

func (w *RandomWorker) GetRandomResult(req *RandomRequest) (*RandomResponse, error) {
	if w.reSeedCount <= 0 {
		if err := w.ReSeed(); err != nil {
			return nil, err
		}
	}
	w.reSeedCount--
	switch req.RequestType {
	case RequestTypeInt64:
		return &RandomResponse{Out: w.rng.Int63()}, nil
	case RequestTypeUint64:
		return &RandomResponse{Out: w.rng.Uint64()}, nil
	case RequestTypeFloat64:
		return &RandomResponse{Out: w.rng.Float64()}, nil
	default:
		return nil, errors.New("not implemented request type")
	}
}

func (w *RandomWorker) run() {

	for {
		select {
		case req := <-w.requestChan:
			resp, err := w.GetRandomResult(req)
			if err != nil {
				req.Return <- &RandomResponse{Err: err}
			} else {
				req.Return <- resp
			}
		case <-w.ctx.Done():
			return
		}
	}
}

func CryptoInt64LE() (int64, error) {
	bytes8 := make([]byte, 8)
	if n, err := cryptorand.Read(bytes8); err != nil || n == 0 {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(bytes8)), nil
}
