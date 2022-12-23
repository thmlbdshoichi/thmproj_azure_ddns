import base64
from typing import Union
from azure.mgmt.dns import DnsManagementClient
from azure.identity import ClientSecretCredential

class AzureDDNS:
    def __init__(self, client_id: str, client_secret: str, resource_group: str, subscription_id: str, tenant_id: str, zonename: str, username: str, password: str):
        self.client_id = client_id
        self.client_secret = client_secret
        self.resource_group = resource_group
        self.subscription_id = subscription_id
        self.tenant_id = tenant_id
        self.zonename = zonename
        self.username = username
        self.password = password
        self.dns_client = self.SetupAzureDNSConnection()
    
    def SetupAzureDNSConnection(self) -> DnsManagementClient:
        credentials = ClientSecretCredential(self.tenant_id, self.client_id, self.client_secret)
        dns_client = DnsManagementClient(credentials, self.subscription_id)
        return dns_client
    
    def BasicAuth(self, base64key: str) -> bool:
        decode_key = base64.b64decode(base64key).decode('utf-8')
        credentials = decode_key.split(":")
        username = credentials[0]
        password = credentials[1]
        
        if (username == self.username) and (password == self.password):
            return True
        else:
            return False
    
    def GetIPFromDNSRecord(self, record_name) -> Union[str,bool]:
        try:
            record_set = self.dns_client.record_sets.get(self.resource_group, self.zonename, record_name, "A")
            ip_address = record_set.a_records[0].ipv4_address
            return ip_address, True
        except:
            print("Exception, an Error occur during DNS Query")
            return "0.0.0.0", False

    def UpdateRecord(self, record_name: str, ipv4: str, ttl: int):
        record_set = self.dns_client.record_sets.create_or_update(self.resource_group, 
                                                                self.zonename,record_name, 
                                                                "A",
                                                                {
                                                                "ttl": ttl, 
                                                                "arecords": [{"ipv4_address": ipv4}],
                                                                "metadata": {
                                                                    "ipv4": ipv4,
                                                                    "updatedBy": "Azure DDNS Server (Python)",
                                                                },})
