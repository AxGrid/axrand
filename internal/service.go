package internal

import (
	"context"
	"github.com/rs/zerolog/log"
)

/*
 _____  __  __  _____  _____  ___  _____
/  _  \/  \/  \/   __\/  _  \/___\|  _  \
|  _  |>-    -<|  |_ ||  _  <|   ||  |  |
\__|__/\__/\__/\_____/\__|\_/\___/|_____/
zed (17.09.2024)
*/

type RandomGenerationService struct {
	requestChan chan *RandomRequest
}

func NewRandomGenerationService(ctx context.Context, workerCount, bufferSize, reSeedInterval int) (*RandomGenerationService, error) {
	res := &RandomGenerationService{
		requestChan: make(chan *RandomRequest, bufferSize),
	}
	for i := 0; i < workerCount; i++ {
		_, err := NewRandomWorker(ctx, i, res.requestChan, reSeedInterval)
		if err != nil {
			return nil, err
		}
	}
	log.Debug().Msgf("start %d workers", workerCount)
	return res, nil
}

func (s *RandomGenerationService) C() chan *RandomRequest {
	return s.requestChan
}
