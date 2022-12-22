package main

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-client/ddns"
	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-client/global"
	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-client/initialize"
)

func init() {
	yaml, err := initialize.ReadYamlFile("config.yaml")
	if err != nil {
		log.Fatal("Error loading config.yaml file")
	}
	global.GB_CONFIG = yaml
	global.GB_AZDNS = initialize.SetupAzureDNSConnection()
}

func main() {
	record_name := global.GB_CONFIG.AzureDNS.RecordName
	zone_name := global.GB_CONFIG.AzureDNS.ZoneName
	record_type := armdns.RecordTypeA
	ttl := global.GB_CONFIG.AzureDNS.TTL
	ip_address, err := ddns.GetPublicIP()
	if err != nil {
		log.Fatal("Error getting PubicIP", err.Error())
	}

	err = ddns.UpdateRecord(record_name, record_type, ip_address, ttl, global.GB_AZDNS)
	if err != nil {
		log.Fatal("Error updating azure dns record", err.Error())
	}

	fmt.Printf("Successfully updated IP Address: %s to Hostname: %s.%s", ip_address, record_name, zone_name)
	fmt.Printf("\nPress any key to close this window.")
	fmt.Scanln()
}
