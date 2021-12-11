package interop

import (
	"testing"
	"time"

	ccs "github.com/Hyperledger-TWGC/ccs-gm/sm3"
	pku "github.com/Hyperledger-TWGC/pku-gm/gmssl"
	tj "github.com/Hyperledger-TWGC/tjfoc-gm/sm3"
)

var base_format = "2006-01-02 15:04:05"

func TestSM3(t *testing.T) {
	// generate a random string as data
	time := time.Now()
	str_time := time.Format(base_format)
	msg := []byte(str_time)
	// generate key from tj
	tj_digest := tj.Sm3Sum([]byte(str_time))
	ccs_digest := ccs.SumSM3(msg)
	sm3hash := pku.Ne