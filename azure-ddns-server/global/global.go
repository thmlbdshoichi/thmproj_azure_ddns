package global

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-server/config"
)

var (
	GB_AZDNS  *armdns.RecordSetsClient
	GB_CONFIG *config.Config
)
