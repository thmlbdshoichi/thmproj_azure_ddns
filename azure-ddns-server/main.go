package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-server/controller"
	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-server/global"
	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-server/initialize"
)

const (
	IP   string = "0.0.0.0"
	Port string = "9999"
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
	// Controller
	var ddnscontroller controller.DDNSController = controller.NewDDNSController()

	r := gin.Default()
	r.GET("/nic/update", ddnscontroller.DNSUpdate)
	r.Run(IP + ":" + Port)
}
