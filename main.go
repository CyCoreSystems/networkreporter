package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/CyCoreSystems/netdiscover/discover"
)

const defaultBindAddr = ":9995"

var bindAddr string

func init() {
	bindAddr = defaultBindAddr
	if os.Getenv("ADDR") != "" {
		bindAddr = os.Getenv("ADDR")
	}
}

func main() {

	var discoverer discover.Discoverer
	switch os.Getenv("CLOUD") {
	case "aws":
		discoverer = discover.NewAWSDiscoverer()
	case "azure":
		discoverer = discover.NewAzureDiscoverer()
	case "do":
		discoverer = discover.NewDigitalOceanDiscoverer()
	case "gcp":
		discoverer = discover.NewGCPDiscoverer()
	case "":
		discoverer = discover.NewDiscoverer()
	default:
		log.Fatal("unsupported CLOUD; leave empty for best-effort")
	}

	http.HandleFunc("/netinfo", func(w http.ResponseWriter, r *http.Request) {
		priv4, err := discoverer.PrivateIPv4()
		if err != nil {
			log.Println("failed to resolve private IPv4:", err)
		}
		pub4, err := discoverer.PublicIPv4()
		if err != nil {
			log.Println("failed to resolve public IPv4:", err)
		}
		pub6, err := discoverer.PublicIPv6()
		if err != nil {
			log.Println("failed to resolve public IPv6:", err)
		}
		hostname, err := discoverer.Hostname()
		if err != nil {
			log.Println("failed to resolve public hostname:", err)
		}

		err = json.NewEncoder(w).Encode(struct {
			PrivateIPv4 string `json:"privateIPv4"`
			PublicIPv4  string `json:"publicIPv4"`
			PublicIPv6  string `json:"publicIPv6"`
			Hostname    string `json:"hostname"`
		}{
			PrivateIPv4: priv4.String(),
			PublicIPv4:  pub4.String(),
			PublicIPv6:  pub6.String(),
			Hostname:    hostname,
		})
		if err != nil {
			log.Println("failed to write response:", err)
		}
	})

	log.Fatal(http.ListenAndServe(bindAddr, nil))
}
