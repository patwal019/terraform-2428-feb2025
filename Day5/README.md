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

## Info - Terraform Provider Development best practices and recommended naming conventions
<pre>
- Provider name must be terraform-provider-nameoftheprovider, must be all lower case
- Resource names must start with nameoftheprovider-resource i.e docker_container, docker is the provider name while the resource managed is container
- resource and data source name must be all lowercase separated by underscore, and recommened to restrict to 2 or 3 words at the max
- 

</pre>
