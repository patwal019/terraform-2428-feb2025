# Day 5

## Info - Terraform Module vs Terraform Provider
<pre>
- Terraform Provider is developed by a Terraform Provider developer using golang programming language
- Terraform Module is written by DevOps Engineers who use the Terraform provider
- Terraform Module is written in HCL(Hashicorp Configuration Language) will file extensions *.tf
- Terraform Module is written in declarative code
- Terraform Modules are reusable code that can be invoked from Terraform Root modules
- Terraform Modules can be reused across many Terraform Projects
- Terraform Modules depends on one or more Terraform Providers
- Terraform Projects is nothing but the top level folder that has the below
  - Terraform Providers
  - Terraform Root Module
    - Terrafoorm child Module(s) optionally
</pre>

## Lab - Printing Resources and Datasources supported by any Terraform Provider 
```
cd ~/terraform-2428-feb2025
git pull
cd Day4/custom-terraform-provider/test-custom-terraform-docker-provider

terraform providers schema -json | jq
```

Expected output
![image](https://github.com/user-attachments/assets/6b5e7265-7c64-42eb-a686-3dd4bc44f47f)
![image](https://github.com/user-attachments/assets/5a00f5e0-b18a-4626-8ada-e328f3a042f4)

## Info - Terraform Provider Development best practices and recommended naming conventions
<pre>
- Provider name must be terraform-provider-nameoftheprovider, must be all lower case
- Resource names must start with nameoftheprovider-resource i.e docker_container, docker is the provider name while the resource managed is container
- resource and data source name must be all lowercase separated by underscore, and recommened to restrict to 2 or 3 words at the max
- 

</pre>
