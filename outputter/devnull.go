package outputter

import config "github.com/coccyx/gogen/internal"

type devnull struct{}

func (foo devnull) Send(item *config.OutQueueItem) error {
	return nil
}
