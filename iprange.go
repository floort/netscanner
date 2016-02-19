package netscanner

import (
	"log"
	"net"
	"sync"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func MultiRangeCIDR(cidrranges []string) <-chan net.IP {
	var wg sync.WaitGroup
	addrs := make(chan net.IP)
	for _, cidrrange := range cidrranges {
		wg.Add(1)
		go func(addrs chan net.IP, cidrrange string) {
			ip, ipnet, err := net.ParseCIDR(cidrrange)
			if err != nil {
				log.Println(err)
				wg.Done()
				return
			}
			for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
				addrs <- ip
			}
			wg.Done()
		}(addrs, cidrrange)
	}
	go func() {
		wg.Wait()
		close(addrs)
	}()
	return addrs
}
