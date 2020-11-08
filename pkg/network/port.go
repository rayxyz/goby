package network

import (
	"errors"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// ObtainAvailablePort :
func ObtainAvailablePort(min, max int) (int, error) {
	for i := 0; i < 50; i++ {
		time.Sleep(time.Millisecond * 50)
		p := rand.Intn(max-min) + min
		l, err := net.Listen("tcp", ":"+strconv.Itoa(p))
		if err != nil {
			continue
		}
		l.Close()
		return p, nil
	}

	return 0, errors.New("cannot obtain free network port")
}
