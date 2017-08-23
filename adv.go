package main

import "golang.org/x/net/ipv6"
import "github.com/skoef/ndp"
import "github.com/pkg/errors"
import "net"
import "time"

type Advertiser struct {
	conn         ipv6.PacketConn
	globalPrefix string // TODO find better datatype
	prefixLength uint8  // 0-128
	ticker       *time.Ticker
	dst          net.Addr
	msg          []byte // via ndp.ICMPRouterAdvertisement.Marshal()
	done         chan bool
}

func (a *Advertiser) tick() {
	// we ignore failures here; they're broadcasts anyway
	a.conn.WriteTo(a.msg, nil, a.dst)
}

func (a *Advertiser) Stop() {
	if a.ticker != nil {
		a.ticker.Stop()
		close(a.done) // kill goroutine
	}
}

func NewAdvertiser(intIF string, d time.Duration) (*Advertiser, error) {
	var adv Advertiser
	var err error

	if d < time.Second*5 {
		return nil, errors.New("Advertiser interval must be at least 5 seconds")
	}

	// TODO resolve netIF to a net.PacketConn

	icmppkt := ndp.ICMPRouterAdvertisement{}
	// TODO set fields in router advertisement
	adv.msg, err = icmppkt.Marshal()
	if err != nil {
		return nil, errors.Wrap(err, "ICMPRouterAdvertisement.Marshal()")
	}

	adv.done = make(chan bool)
	adv.ticker = time.NewTicker(d)
	go func() {
		for {
			select {
			case <-adv.done:
				return
			case <-adv.ticker.C:
				adv.tick()
			}
		}
	}()
	return &adv, nil
}
