package cmd

import (
	"cli-app/helpers"
	"encoding/json"
	"fmt"
	"log"

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
				trace(ip)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

}

type ip struct {
	IPAddress string `json:"ip,omitempty"`
	City      string `json:"city,omitempty"`
	Region    string `json:"region,omitempty"`
	Country   string `json:"country,omitempty"`
	Geo       string `json:"loc,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
}

func trace(ipadd string) {
	url := "http://ipinfo.io/" + ipadd + "/geo"
	responseByte := helpers.GetData(url)
	ip := &ip{}
	err := json.Unmarshal(responseByte, &ip)
	if err != nil {
		log.Println("Unable to unmarshall response.")
	}

	fmt.Printf("IP: %s\nCity: %s\nRegion: %s\nCountry: %s\nGeo: %s\nTimezone: %s\n", ip.IPAddress, ip.City, ip.Region, ip.Country, ip.Geo, ip.Timezone)
}
