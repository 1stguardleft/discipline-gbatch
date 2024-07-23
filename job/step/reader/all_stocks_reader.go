package reader

import (
	"github.com/chararch/gobatch"
)

type allStocksReader struct {
}

func (r *allStocksReader) read(chunkCtx *gobatch.BatchContext) (interface{}, gobatch.BatchError) {
	return "", nil
}
