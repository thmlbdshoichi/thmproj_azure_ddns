package config

type Config struct {
	AzureDNS AzureDNSConfig `mapstructure:"azure-dns" json:"azure_dns" yaml:"azure-dns"`
}

type AzureDNSConfig struct {
	ClientId       string `mapstructure:"client-id" json:"client_id" yaml:"client-id"`
	ClientSecret   string `mapstructure:"client-secret" json:"client_secret" yaml:"client-secret"`
	RecordName     string `mapstructure:"record-name" json:"record_name" yaml:"record-name"`
	ResourceGroup  string `mapstructure:"resource-group" json:"resource_group" yaml:"resource-group"`
	SubscriptionId string `mapstructure:"subscription-id" json:"subscription_id" yaml:"subscription-id"`
	TenantId       string `mapstructure:"tenant-id" json:"tenant_id" yaml:"tenant-id"`
	TTL            int64  `mapstructure:"ttl" json:"ttl" yaml:"ttl"`
	ZoneName       string `mapstructure:"zonename" json:"zonename" yaml:"zonename"`
}
