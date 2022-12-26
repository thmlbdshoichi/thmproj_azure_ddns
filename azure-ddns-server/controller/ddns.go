package controller

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/gin-gonic/gin"
	"github.com/thmlbdshoichi/thmproj_azure_ddns/azure-ddns-server/global"
)

type DDNSController interface {
	DNSUpdate(ctx *gin.Context)
}

type ddnsController struct {
}

func NewDDNSController() DDNSController {
	return &ddnsController{}
}

func (c *ddnsController) DNSUpdate(ctx *gin.Context) {

	// Authentication Basic Auth
	if len(ctx.Request.Header["Authorization"]) == 0 {
		ctx.String(http.StatusUnauthorized, "badauth")
		return
	}
	base64key := strings.Split(ctx.Request.Header["Authorization"][0], " ")
	if len(base64key) < 2 {
		ctx.String(http.StatusUnauthorized, "badauth")
		return
	}
	ok, err := c.BasicAuth(base64key[1])
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if !ok {
		ctx.String(http.StatusUnauthorized, "badauth")
		return
	}

	//----------------------------------------------------------------------------------------------------
	// RECORD UPDATE INFORMATION
	hostname := ctx.Query("hostname")
	if hostname == "" {
		ctx.String(http.StatusBadRequest, "nohost")
		return
	}
	record_names := strings.Split(hostname, ".")
	if len(record_names) < 3 {
		ctx.String(http.StatusBadRequest, "nohost")
		return
	}
	record_name := record_names[0]
	if record_name == "" {
		ctx.String(http.StatusBadRequest, "nohost")
		return
	}
	recordType := armdns.RecordTypeA
	myip := ctx.Query("myip")
	var ttl int64 = 3600

	client := global.GB_AZDNS

	currentIP, err := c.GetIPFromDNSRecord(record_name, recordType, client)
	if err != nil {
		//แสดงว่าอาจจะยังไม่มี DNS Record
		err := c.UpdateRecord(record_name, recordType, myip, ttl, client)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "911")
			return
		}
	}
	if currentIP == myip {
		ctx.String(http.StatusOK, "nochg %s", myip)
		return
	}

	err = c.UpdateRecord(record_name, recordType, myip, ttl, client)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "911")
		return
	}
	ctx.String(http.StatusOK, "good %s", myip)
}

// Helper function
func (c *ddnsController) BasicAuth(base64key string) (bool, error) {

	decode_key, err := base64.StdEncoding.DecodeString(base64key)
	if err != nil {
		return false, err
	}

	credentials := strings.Split(string(decode_key), ":")
	username := credentials[0]
	password := credentials[1]

	//global.GVA_LOG.Info(fmt.Sprintf("username: %q, password: %q", username, password))

	if (username == global.GB_CONFIG.AzureDNS.Username) && (password == global.GB_CONFIG.AzureDNS.Password) {
		return true, nil
	}

	return false, nil

}

func (c *ddnsController) GetIPFromDNSRecord(recordName string, recordType armdns.RecordType, client *armdns.RecordSetsClient) (string, error) {
	res, err := client.Get(
		context.Background(),
		global.GB_CONFIG.AzureDNS.ResourceGroup,
		global.GB_CONFIG.AzureDNS.ZoneName,
		recordName,
		recordType, nil)

	if err != nil {
		return "0.0.0.0", err
	}

	if len(res.RecordSet.Properties.ARecords) == 0 {
		return "0.0.0.0", nil
	}

	// Return first IPv4 Address in the Record
	currentIP := *res.RecordSet.Properties.ARecords[0].IPv4Address
	return currentIP, nil
}

func (c *ddnsController) UpdateRecord(recordName string, recordType armdns.RecordType, IP string, ttl int64, client *armdns.RecordSetsClient) error {
	record := armdns.RecordSet{
		Properties: &armdns.RecordSetProperties{
			ARecords: []*armdns.ARecord{
				{
					IPv4Address: to.Ptr(IP),
				}},
			TTL: to.Ptr[int64](ttl),
			Metadata: map[string]*string{
				"updatedBy": to.Ptr("DNS Updater Server (GO)"),
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
