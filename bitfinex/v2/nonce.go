package v2

import (
	"strconv"
	"sync/atomic"
	"time"
)

// GetNonce is a naive nonce producer that takes the current Unix nano epoch
// and counts upwards.
// This is a naive approach because the nonce bound to the currently used API
// key and as such needs to be synchronised with other instances using the same
// key in order to avoid race conditions.
func GetNonce() string {
	nonce := uint64(time.Now().Unix()) * 1000000000
	return strconv.FormatUint(atomic.AddUint64(&nonce, 209000000), 10)
}
