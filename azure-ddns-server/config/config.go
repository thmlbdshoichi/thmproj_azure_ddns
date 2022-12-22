package config

type Config struct {
	AzureDNS AzureDNSConfig `mapstructure:"azure-dns" json:"azure_dns" yaml:"azure-dns"`
}

type AzureDNSConfig struct {
	ClientId       string `mapstructure:"client-id" json:"client_id" yaml:"client-id"`
	ClientSecret   string `mapstructure:"client-secret" json:"client_secret" yaml:"client-secret"`
	Password       string `mapstructure:"password" json:"password" yaml:"password"`
	RecordName     string `mapstructure:"recordname" json:"recordname" yaml:"recordname"`
	ResourceGroup  string `mapstructure:"resource-group" json:"resource_group" yaml:"resource-group"`
	SubscriptionId string `mapstructure:"subscription-id" json:"subscription_id" yaml:"subscription-id"`
	TenantId       string `mapstructure:"tenant-id" json:"tenant_id" yaml:"tenant-id"`
	Username       string `mapstructure:"username" json:"username" yaml:"username"`
	ZoneName       string `mapstructure:"zonename" json:"zonename" yaml:"zonename"`
}
