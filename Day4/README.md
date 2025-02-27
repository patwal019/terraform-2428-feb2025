# Day 4

## Info - Terraform Overview
<pre>
- is a Infrastructure as a Code (IOC) tool
- in other a provisioner tool
- cloud newtral
- alternates to tranform
  - AWS Cloud formation
- terraform supports state management of resources 
- is developed in Golang
- comes in 2 flavours
  - Terraform core - opensource and CLI
  - Terraform Enterprise - Web Interface ( Commercial product from HashiCorp )
- Terraform uses a DSL called HCL(HashiCorp Configuration Language ) for writing automation scripts
- terraform depends on providers to provision resources
- For examples
  - aws provider helps terraform provision resources in AWS cloud
  - azure provider helps terraform provision resources in Azure cloud
  - docker provider that helps terraform provision resources like image, container, network, etc
  - local provider that helps terraform provision files in the local machine
- Terraform supports a provider plugin framework to develop custom terraform providers
- Terraform providers are developed in Golang as the Terraform Provider Plugin Framework is natively written in Golang
</pre>

## Lab - Checking the terraform version
```
terraform version
```

Expected output
![image](https://github.com/user-attachments/assets/a96b8bb3-88df-4b57-9549-03cec717d9b5)

## Lab - Writing your first Terraform manifest file (automation script)
```
cd ~/terraform-2428-feb2025
git pull
cd Day4/terraform/pull-docker-image
cat main.tf
```

Downloading the terraform provider plugin mentioned in the main.tf
```
cd ~/terraform-2428-feb2025
git pull
cd Day4/terraform/pull-docker-image
terraform init
```

Expected output
![image](https://github.com/user-attachments/assets/fee26516-4183-4eaa-9782-22d0e58317d9)
![image](https://github.com/user-attachments/assets/f44b2072-3f2d-472a-a526-4882f8feba58)
![image](https://github.com/user-attachments/assets/11a0000e-58ed-4078-aeb8-27b122ea5931)


We can check the plan before we execute the terraform manifest scripts
```
terraform plan
```

Expected output
![image](https://github.com/user-attachments/assets/5a6e87c8-6346-442a-94d5-3574c379b1fa)


To run the terraform manifest scripts
```
terraform apply
docker images
```

Expected output
![image](https://github.com/user-attachments/assets/101fdd28-19e4-4239-b6a8-4ed9c0862a0a)
![image](https://github.com/user-attachments/assets/59d19833-f109-4aa0-bd31-839598899c89)
![image](https://github.com/user-attachments/assets/dac4f04a-3a21-4322-b08e-f139192a7518)


You can check the current state of the resouces provisioned by terraform
```
terraform show
```
Expected
![image](https://github.com/user-attachments/assets/bb1b143e-c74c-42e5-b57e-1bb63db928e1)

Once you are done with this exercises, it is a good practice to dispose the resources provisioned by terraform
```
terraform destroy
```

Expected output
![image](https://github.com/user-attachments/assets/6bf939b4-035a-4480-832a-8fd71f21e89e)
![image](https://github.com/user-attachments/assets/ed403027-128e-447b-961d-4e073a19a568)
