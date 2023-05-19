### <div align="center"> Introduction (English Coming soon)</div>
<div align="left">
Repository นี้สร้างขึ้นเพื่อเป็นข้อมูลเกี่ยวกับการทำ DDNS โดยใช้ Azure DNS Service โดยปกติแล้ว Azure DNS ไม่ได้มีฟังก์ชั่นรองรับการทำ Dynamic DNS แต่ทาง Microsft Azure ได้ออกตัว Azure SDK ที่ทำให้เราสามารถแก้ไข Service ต่างๆของ Azure ได้ เช่น
<br>
1. Azure SDK for GO | https://github.com/Azure/azure-sdk-for-go
<br>
2. Azure SDK for Python | https://github.com/Azure/azure-sdk-for-python
<br>
หากเรารู้หลักการทำงานของพวก DDNS Service และรู้วิธีที่ Router สื่อสารกับ API ของ DDNS เราก็สามารถสร้าง DDNS Server ขึ้นมาเองได้ :) 
</div>

## <div align="center"> Azure DNS Zone </div>
<div align="center"> การจัดการทุกอย่างเกี่ยวกับตัว Azure จะทำผ่าน Azure Cloud Portal </div>
<br>
<details>
<summary>Create Azure DNS Zone</summary>
<br>
1. สร้าง DNS Zone ไปที่เมนู "Home -> Create a resource -> Create a resource -> DNS Zone" แล้วกดปุ่ม Create
<br>
<br>
<div align="center"><img width="300" src="https://user-images.githubusercontent.com/26995849/209205707-1e4ed6fc-3c7b-4763-b25a-8aa8e2f561f7.png"></div>
<br>
2. กรอกรายละเอียด เลือก Subscription, Resource group (ถ้าไม่มีให้กดสร้างใหม่) 
  ตรงช่อง Name คือชื่อ DNS Zone ของเรา เช่น thmddns.net | hostname ประกอบด้วย {recordname}.{zonename} เช่น test.thmddns.net เมื่อเสร็จแล้วกด "Review + Create"
<div align="center"><img width="600" src="https://user-images.githubusercontent.com/26995849/209206603-46f216d6-38cb-49b9-9bb0-18b2aec438ad.png"></div>

3. เราจะได้ข้อมูลบางอย่างในหน้า DNS Zone (Overview) เพื่อนำไปใส่ในไฟล์ config.yaml (ตัวอย่าง configexample.yaml) เช่น<br>
  3.1. [subscription-id] Subscription ID<br>
  3.2. [resource-group] ชื่อ Resource group<br>
  3.3. [zonename] มาจาก DNS Zone ที่เรากำหนดไว้ ดูตรงซ้ายบน<br>
  
![Screenshot 2023-05-19 112315](https://github.com/thmlbdshoichi/thmproj_azure_ddns/assets/26995849/32ae2f8f-91bd-4732-8b3c-758bba0be8df)
<br><br>
</details>


<details>
<summary>Create Service Principal (Azure Active Directory)</summary>
ลิงค์เพิ่มเติม: https://learn.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal
<br><br>
1. ไปที่ Azure Active Directory -> App registrations -> New Registration เพื่อสร้าง Service Principal ขึ้นมา
<br><br>
2. เมื่อสร้างเสร็จแล้วให้ไปที่ DNS Zone ของเราเพื่อกำหนดสิทธิ์ให้ SP ที่สร้างขึ้นมาเป็น "DNS Zone Contributor"
<br><br>
3. เราจะได้ข้อมูลบางอย่างในหน้า Service Principal (Overview) มาเพื่อนำไปใส่ในไฟล์ config.yaml (ตัวอย่าง configexample.yaml) เช่น<br>
  3.1. [tenant-id] Directory (tenant) ID<br>
  3.2. [client-id] Application Client ID<br>
  3.3. [client-secret] Client Secrets ซึ่งสามารถ Generate ได้จากเมนู Certificates & Secret<br>

![Screenshot 2023-05-19 113556](https://github.com/thmlbdshoichi/thmproj_azure_ddns/assets/26995849/e11697e2-c314-4e77-a8b7-aa01d97dcafc)

<br><br>
</details>

<details>
<summary>Trying to Run Server or Client to Update DNS</summary>
<br><br>
1. ไฟล์ YAML Config ของ Server จะเหลือในส่วนของ Username, Password เราสามารถตั้งเองได้ เป็น BasicAuth ที่ Router ใช้ในการส่งข้อมูลมาหาตัว Server<br><br>
2. เปิดตัว Server<br>
  2.1. azure-ddns-server (go run main.go)<br>
  2.2. azure-ddns-server-python (pip install ให้เรียบร้อย และ python main.py หรือ uvicorn main:app --reload)<br><br>
3. ทดสอบการอัพเดท DNS Record ใน ตั้งค่า DDNS Configuration ใน Router เช่น (3BB, AIS) หรือลองใช้โฟล์เดอร์ azure-ddns-client (ต้องใส่ข้อมูลใน config.yaml ก่อน)
</details>

#### <div align="center"> ขออภัยในความไม่ละเอียดยังไม่มีเวลาเขียน และ ตกแต่งไม่เป็น TT</div>
