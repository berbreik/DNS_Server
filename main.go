package DNS

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
)

var dnsRecords map[string]DNSRecord

func server(input dns.Msg) {

}

func loadDnsRecords(Record string) error {
	return nil
}

func main() {
	if err := loadDnsRecords("Record.json"); err != nil {
		log.Println("Failed to fetch record")
	}

	addr := net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 8053,
		Zone: "",
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("failed to start dns server")
	}

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Failed to close the connection")
		}
	}(conn)

}
