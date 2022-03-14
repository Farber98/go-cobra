package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  "Trace the IP",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("IP needed to trace.")
		} else {
			for i, ip := range args {
				fmt.Printf(" ===================== IP INFORMATION #%d ====================== \n", i+1)
				showData(ip)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

}

/*  ip: "1.1.1.1"
hostname: "one.one.one.one"
anycast: true
city: "Miami"
region: "Florida"
country: "US"
loc: "25.7867,-80.1800"
org: "AS13335 Cloudflare, Inc."
postal: "33132"
timezone: "America/New_York"
asn: Object
asn: "AS13335"
name: "Cloudflare, Inc."
domain: "cloudflare.com"
route: "1.1.1.0/24"
type: "hosting"
company: Object
name: "APNIC and Cloudflare DNS Resolver project"
domain: "cloudflare.com"
type: "hosting"
privacy: Object
vpn: false
proxy: false
tor: false
relay: false
hosting: true
service: ""
abuse: Object
address: "PO Box 3646, South Brisbane, QLD 4101, Australia"
country: "AU"
email: "helpdesk@apnic.net"
name: "APNIC RESEARCH"
network: "1.1.1.0/24"
phone: "+61-7-3858-3188"
domains: Object
ip: "1.1.1.1"
total: 26968
domains: Array
0: "axutongxue.com"
1: "nfmovies.com"
2: "xmoon.club"
3: "christacoetzee.com"
4: "authrock.com" */
type IP struct {
	IPAddress string `json:"ip,omitempty"`
	City      string `json:"city,omitempty"`
	Region    string `json:"region,omitempty"`
	Country   string `json:"country,omitempty"`
	Geo       string `json:"loc,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
}

func showData(ipadd string) {
	url := "http://ipinfo.io/" + ipadd + "/geo"
	responseByte := getData(url)
	ip := &IP{}
	err := json.Unmarshal(responseByte, &ip)
	if err != nil {
		log.Println("Unable to unmarshall response.")
	}

	fmt.Printf("IP: %s\nCity: %s\nRegion: %s\nCountry: %s\nGeo: %s\nTimezone: %s\n", ip.IPAddress, ip.City, ip.Region, ip.Country, ip.Geo, ip.Timezone)
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get response.")
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error getting the response.")
	}
	return responseByte
}
