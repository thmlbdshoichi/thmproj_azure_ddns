package ddns

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-client/global"
)

func UpdateRecord(recordName string, recordType armdns.RecordType, IP string, ttl int64, client *armdns.RecordSetsClient) error {
	record := armdns.RecordSet{
		Properties: &armdns.RecordSetProperties{
			ARecords: []*armdns.ARecord{
				{
					IPv4Address: to.Ptr(IP),
				}},
			TTL: to.Ptr[int64](ttl),
			Metadata: map[string]*string{
				"updatedBy": to.Ptr("HAT-Server | DNS Updater"),
				"updatedAt": to.Ptr(time.Now().String()),
			},
		},
	}

	recordOptions := &armdns.RecordSetsClientCreateOrUpdateOptions{
		IfMatch:     nil,
		IfNoneMatch: nil,
	}

	_, err := client.CreateOrUpdate(
		context.Background(),
		global.GB_CONFIG.AzureDNS.ResourceGroup,
		global.GB_CONFIG.AzureDNS.ZoneName,
		recordName, recordType, record, recordOptions)

	if err != nil {
		return err
	}

	return nil
}

func GetPublicIP() (string, error) {
	req, err := http.Get("https://ifconfig.me")
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
