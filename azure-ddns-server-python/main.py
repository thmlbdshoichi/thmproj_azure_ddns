import yaml
import azdnshelper
import uvicorn

from typing import Union
from fastapi import FastAPI, Header
from fastapi.responses import PlainTextResponse

with open('config.yaml', 'r') as file:
    config = yaml.safe_load(file)
    
client_id = config["azure-dns"]["client-id"]
client_secret = config["azure-dns"]["client-secret"]
resource_group = config["azure-dns"]["resource-group"]
subscription_id = config["azure-dns"]["subscription-id"]
tenant_id = config["azure-dns"]["tenant-id"]
zonename = config["azure-dns"]["zonename"]
username = config["azure-dns"]["username"]
password = config["azure-dns"]["password"]

dns_client = azdnshelper.AzureDDNS(client_id, client_secret, resource_group, subscription_id, tenant_id, zonename, username, password)

app = FastAPI()

@app.get("/nic/update", response_class=PlainTextResponse)
async def DNSUpdater(hostname: str, myip: str, Authorization: list[str] | None = Header(default=None)):
    # Authentication by BasicAuth | BASE64
    # IF DON'T WANT BASIC AUTH JUST COMMENT IT
    if Authorization is None:
        return "emptyauth"
    base64key = Authorization[0]
    base64key = base64key.split(" ")
    if len(base64key) < 2:
        return "invalidauth"
    ok = dns_client.BasicAuth(base64key[1])
    if not ok:
        return "badauth"
    ##################################################
    ttl = 3600
    record_names = hostname.split(".")
    if len(record_names) < 3:
        return "invalidhostname"
    record_name = record_names[0]
    
    current_ip, ok = dns_client.GetIPFromDNSRecord(record_name)
    if not ok:
        dns_client.UpdateRecord(record_name, myip, ttl)
        return f"good {myip}"
    
    if myip == current_ip:
        return f"nochg {myip}"
    
    dns_client.UpdateRecord(record_name, myip, ttl)
    
    return f"good {myip}"

if __name__ == "__main__":
    uvicorn.run(app, host=config["fast-api"]["IP"], port=config["fast-api"]["PORT"])