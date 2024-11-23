package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/david-gurley/gopen"
)

const (
	PSM               = "psm-ent.p6o.net"
	PSM_USERNAME      = "admin"
	PSM_PASSWORD      = "Pensando0$"
	QUARANTINE_POLICY = "secops-responses"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("got a requst")
		if r.Method != "POST" {
			http.Error(w, "method is not supported", http.StatusNotFound)
		}
		var quarantinePayload QuarantinePayload
		err := json.NewDecoder(r.Body).Decode(&quarantinePayload)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		c := gopen.NewPSMClient(&gopen.PSMClientConfig{
			Username: PSM_USERNAME,
			Password: PSM_PASSWORD,
			Hostname: PSM,
		})
		err = c.Login()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("submitting quarantine action for: %s\n", quarantinePayload.IP)
		err = gopen.QuarantineWorkload(c, QUARANTINE_POLICY, quarantinePayload.IP)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	})
	fmt.Printf("started server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type QuarantinePayload struct {
	IP string `json:"ip"`
}
