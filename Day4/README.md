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

## Lab - Provision a docker container using Terraform
```
cd ~/terraform-2428-feb2025
git pull
cd Day4/terraform/provision-docker-container
cat providers.tf
cat main.tf
terraform init
terraform apply --auto-approve
docker images
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/006f4feb-54c1-41e3-9650-971c805f456f)
![image](https://github.com/user-attachments/assets/8fcec5ce-6754-45d2-82e3-e6ab344a6697)
![image](https://github.com/user-attachments/assets/76b6f670-34ab-433e-ad00-49ab2430c073)
![image](https://github.com/user-attachments/assets/5e9f8bb1-f816-43a0-b359-af2d8d90df40)
![image](https://github.com/user-attachments/assets/1655e04c-36bd-48ea-a58a-06f1e1acfb6b)
![image](https://github.com/user-attachments/assets/1e4595e7-621b-43fe-a227-a6257e625391)
![image](https://github.com/user-attachments/assets/316354e9-66bd-4dbe-9c10-8760f9823e8d)
![image](https://github.com/user-attachments/assets/831830b0-8b6e-419e-a2be-b8dc0f8bb7d9)
![image](https://github.com/user-attachments/assets/d9e4cd11-eacb-40c1-99f2-5ca536c4eee0)
![image](https://github.com/user-attachments/assets/8b49af6c-3c2b-4fac-91b7-9480901511e4)
![image](https://github.com/user-attachments/assets/2795ce5e-bed4-4403-9f0c-742a65124964)

## Lab - Invoking ansible playbook from terraform 
```
cd ~/terraform-2428-feb2025
git pull
cd Day4/terraform/invoking-ansible-playbook-from-terraform
terraform init
terraform plan
terraform apply --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/b750112a-8964-4a53-b1da-056884e6f7dc)
![image](https://github.com/user-attachments/assets/49188ed8-351e-44b6-9ecd-1a4b37603d17)
![image](https://github.com/user-attachments/assets/02f7630a-1471-458f-a508-93abaf518ea9)
![image](https://github.com/user-attachments/assets/1db05416-7aa3-484b-a495-353dfc6a75fc)
![image](https://github.com/user-attachments/assets/88b1b73d-d411-4a28-b629-10bc3146daf3)
![image](https://github.com/user-attachments/assets/301c5294-7774-4bbb-8ddb-12adbdb609d8)
![image](https://github.com/user-attachments/assets/83ae0e58-0363-4e96-8d04-4bd2e11ea85b)
![image](https://github.com/user-attachments/assets/b9dc63ef-e424-4c87-905a-859659448c77)

Once you are done with this exercise, you may dispose all the resources provisioned by terraform
```
terraform destroy --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/d8ab9783-c89d-4817-97c8-39d5e3ebc803)
![image](https://github.com/user-attachments/assets/50f27285-495c-4582-89bc-c4522855107f)
![image](https://github.com/user-attachments/assets/6f8929b3-a0d3-4df6-817a-b4bc920e2be1)
![image](https://github.com/user-attachments/assets/36bf19be-1aee-4137-9ba2-cd01089164e7)

## Lab - Understanding Terraform remote-exec provisioner block
```
cd ~/terraform-2428-feb2025
git pull
cd Day4/terraform/remote-exec
cat vars.tf
cat main.tf
terraform init
terraform plan
terraform apply --auto-approve
terraform destroy --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/ac72fc80-cf79-4d45-8ed6-1eb33d4236c0)
![image](https://github.com/user-attachments/assets/8f966a83-17f9-4f53-a24f-5389a2de79b1)
![image](https://github.com/user-attachments/assets/ce3a3573-a745-4457-bda0-bd53d34d21e3)
![image](https://github.com/user-attachments/assets/3f4c0983-48a6-40e2-9564-fb511b66cf27)
![image](https://github.com/user-attachments/assets/70ed7c7b-d71d-471f-8d68-192a406e6cf5)
![image](https://github.com/user-attachments/assets/3cb51b07-0cfa-43e0-a6fa-577e810b464e)
![image](https://github.com/user-attachments/assets/68aca640-00cd-43e4-a279-4d58d67b68cb)
![image](https://github.com/user-attachments/assets/a54d2fad-72b6-4616-a7ff-81df574c0d40)
![image](https://github.com/user-attachments/assets/c08d816d-d981-4b59-99e3-79b5c3317e34)
![image](https://github.com/user-attachments/assets/59ba84de-016e-4f87-83cb-6baf41e0cf0c)
