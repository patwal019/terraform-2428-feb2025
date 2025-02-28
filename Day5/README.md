# Day 5

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
