## <div align="center"> Introduction (English Coming soon)</div>
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
<br>
<br>
<div align="center"><img width="600" src="https://user-images.githubusercontent.com/26995849/209206603-46f216d6-38cb-49b9-9bb0-18b2aec438ad.png"></div>

</details>


<details>
<summary>Create Service Principal (Azure Active Directory)</summary>

</details>

<details>
<summary>Assign Service Principal as DNS Zone Contributor</summary>

</details>


