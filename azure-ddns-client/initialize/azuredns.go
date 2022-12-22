package initialize

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-client/global"
)

func SetupAzureDNSConnection() *armdns.RecordSetsClient {
	credentials, err := azidentity.NewClientSecretCredential(global.GB_CONFIG.AzureDNS.TenantId, global.GB_CONFIG.AzureDNS.ClientId, global.GB_CONFIG.AzureDNS.ClientSecret, nil)
	if err != nil {
		log.Fatal("Failed to initialize credentials for Azure DNS")
		return nil
	}

	client, err := armdns.NewRecordSetsClient(global.GB_CONFIG.AzureDNS.SubscriptionId, credentials, nil)
	if err != nil {
		log.Fatal("Failed to initialize connection with Azure DNS")
		return nil
	}

	return client
}
