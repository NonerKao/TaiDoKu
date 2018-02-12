package tdk

import (
	"fmt"
	"os"
)

type IP_STATE uint8

const (
	IP_EXEC    IP_STATE = 0
	IP_SYNCING IP_STATE = 1
	IP_HALT    IP_STATE = 15
)

type Sync struct {
	IP     []*IP
	Number int
	Ping   chan bool
}

type TDKProgram struct {
	Tile []*Tile
	IP   []*IP
	Sync [16]Sync
}

func (ip *IP) Run(ctrl chan<- *IP, ack <-chan bool) {
	for {
		for ip.State == IP_EXEC {
			op := ip.OPLookup()
			ip.Execute(op)
			ip.Refine()
		}

		ctrl <- ip
		<-ack

		switch ip.State {
		case IP_SYNCING:
			<-ip.Ch
		case IP_HALT:
			return
		}
	}
}

func (tdk *TDKProgram) Run() {

	ctrl := make(chan *IP)
	ack := make(chan bool)
	for _, ip := range tdk.IP {
		ip.State = IP_EXEC
		go ip.Run(ctrl, ack)
	}

	var ip *IP
	for len(tdk.IP) != 0 {
		ip = <-ctrl
		switch ip.State {
		case IP_SYNCING:
			if ip.Partner < 2 {
				panic("Invalid sync: not a group")
			}
			if len(tdk.Sync[ip.Token].IP) == 0 {
				tdk.Sync[ip.Token].Number = ip.Partner
				tdk.Sync[ip.Token].IP = append(tdk.Sync[ip.Token].IP, ip)
				go func(s *Sync) {
					for i := 1; i < s.Number; i++ {
						<-s.Ping
					}

					tdk.Sync[ip.Token].Number = 0
					tdk.Sync[ip.Token].IP = nil
					ip.State = IP_EXEC
					ip.Partner = 0
					ip.Token = -1

					for _, ip := range s.IP {
						ip.Ch <- true
					}

				}(&tdk.Sync[ip.Token])
			} else if tdk.Sync[ip.Token].Number == ip.Partner {
				tdk.Sync[ip.Token].IP = append(tdk.Sync[ip.Token].IP, ip)
				tdk.Sync[ip.Token].Ping <- true
			} else {
				panic("Invalid sync: different criteria")
			}

		case IP_HALT:
			var i int
			var p *IP
			for i, p = range tdk.IP {
				if p == ip {
					break
				}
			}

			if p != ip {
				panic("ip not found")
			}

			copy(tdk.IP[i:], tdk.IP[i+1:])
			tdk.IP[len(tdk.IP)-1] = nil
			tdk.IP = tdk.IP[:len(tdk.IP)-1]

			ip.Tile.IP = nil
			ip.Tile = nil

			fmt.Println(len(tdk.IP))
		}
		ack <- true
	}
}

func InitProgram(fn string) *TDKProgram {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var nt uint64

	fmt.Fscan(f, &nt)
	tdk := TDKProgram{
		Tile: make([]*Tile, nt),
		IP:   make([]*IP, 0),
	}

	for k := 0; k < int(nt); k++ {
		tdk.Tile[k] = NewRand(nil, nil, nil)

		var hasip int
		fmt.Fscan(f, &hasip)
		if hasip == 1 {
			ip := IP{
				Tile:    nil,
				x:       0,
				y:       0,
				Partner: 0,
				Token:   -1,
				Ch:      make(chan bool),
			}
			fmt.Fscanf(f, "%X %X", &ip.x, &ip.y)
			tdk.IP = append(tdk.IP, &ip)
			ip.Tile = tdk.Tile[k]
			tdk.Tile[k].IP = &ip
		}

		for i := 0; i < 16; i += 1 {
			for j := 0; j < 16; j += 1 {
				fmt.Fscanf(f, "%X", &tdk.Tile[k].a[i][j])
			}
			var c rune
			fmt.Fscanf(f, "\n", &c)
		}
	}

	for i := 0; i < 16; i += 1 {
		tdk.Sync[i].Ping = make(chan bool)
	}

	return &tdk
}

func InitRandProgram() *TDKProgram {
	tdk := TDKProgram{
		Tile: make([]*Tile, 1),
		IP:   make([]*IP, 1),
	}
	ip := IP{
		Tile:    NewRand(nil, nil, nil),
		x:       0,
		y:       0,
		Partner: 0,
		Token:   -1,
		Ch:      make(chan bool),
	}

	tdk.IP[0] = &ip
	tdk.Tile[0] = tdk.IP[0].Tile
	tdk.Tile[0].IP = &ip

	for i := 0; i < 16; i += 1 {
		tdk.Sync[i].Ping = make(chan bool)
	}
	return &tdk
}
