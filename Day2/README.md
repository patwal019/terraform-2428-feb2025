# Day 2

## Lab - Building a custom Rocky linux ansible node container image
```
cd ~/terraform-2428-feb2025
git pull
cd Day1/CustomDockerAnsibleNodeImages/rockylinux-ansible
cp ~/.ssh/id_ed25519.pub authorized_keys
docker build -t tektutor/rocky-ansible-node:1.0 .
rm authorized_keys
docker images
```

Expected output
![image](https://github.com/user-attachments/assets/d49df2b4-f2bb-4d96-bae4-02fdbdcc9934)
![image](https://github.com/user-attachments/assets/86fa4f59-2ae3-4704-97b9-e95ad59397a4)


## Lab - Let's create couple of rocky linux ansible node containers using the above custom docker image
```
cd ~
docker run -d --name rocky1 --hostname rocky1 -p 2003:22 -p 8003:80 tektutor/rocky-ansible-node:1.0
docker run -d --name rocky2 --hostname rocky2 -p 2004:22 -p 8004:80 tektutor/rocky-ansible-node:1.0
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/0a769789-5ff5-4f4d-94b5-3e2d5e078c0b)

## Lab - Try to connect to rocky1 and rocky2 ansible node container via SSH
```
ssh -p 2003 root@localhost
exit

ssh -p 2004 root@localhost
exit
```

Expected output
![image](https://github.com/user-attachments/assets/162b811a-aab4-4696-8e0c-85645de10f6d)

## Lab - Let's update inventory to add the rocky linux ansible node containers
```
cd ~/terraform-2428-feb2025
git pull
cd Day2/ansible
cat hosts
ansible all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/c4c8270a-fc33-4970-9a2f-90f3738e9440)


## Lab - Let's run the refactored playbook that installs nginx on rocky and ubuntu based ansible nodes
```
cd ~/terraform-2428-feb2025
git pull
cd Day2/ansible
ansible-playbook install-nginx-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/b2fa4278-fac6-4317-899f-08e395f3998b)
![image](https://github.com/user-attachments/assets/9e30a343-72a9-44a9-b1b2-7c103e8b6520)
![image](https://github.com/user-attachments/assets/38e748af-754f-415f-9b04-69352d755373)

## Info - Ansible roles
<pre>
- is a way we can resue code in Ansible
- it follows certain directory structure and coding practices
- roles can't be executed directly ( they are like Windows DLL )
- roles can be invoked in playbook
- when we execute playbook,it can invoke one or more roles
- generally one ansible role, will install one software
- to develop ansible role, we will be using a tool called ansible-galaxy
- there is also a web portal from where we can download and use opensource roles written by community
</pre>

## Lab - Creating a custom Ansible role
<pre>
cd ~/terraform-2428-feb2025
git pull
cd Day2/ansible/roles
ansible-galaxy init nginx
tree nginx
</pre>

Expected output
![image](https://github.com/user-attachments/assets/476a7b87-f76b-46d9-8684-7317c0edb87f)

In the above screenshot,
<pre>
vars - captures user-defined variables
defaults - also captures user-defined variables which are normally read-only variables
files - the src files referred by the ansible file modules will be maintained here
handlers - this folder maintains code related to managing services, these tasks will be invoked based on some notifications
meta - has documentation that explains about the roles, the platforms(OS) supported by the roles, author, license under which the role is published
README.md - documentation
tasks - this folder will have all the code we normally write in an ansible playbook
templates - this folder has files referred by the template module
tests - has test inventory and test playbook that demonstrates on how one can invoke custom ansible role
</pre>

## Lab - Invoking custom role from Ansible Playbook
```
cd ~/terraform-2428-feb2025
git pull
cd Day2/ansible/roles
ansible-playbook playbook.yml

curl http://localhost:8001
curl http://localhost:8002
curl http://localhost:8003
curl http://localhost:8004
```

Expected output
![image](https://github.com/user-attachments/assets/c2c07f1c-b7a7-44b0-b540-1696b0398352)
![image](https://github.com/user-attachments/assets/1f3518f5-0ff0-412c-907a-46932949cc0a)
![image](https://github.com/user-attachments/assets/bf06652b-6845-45c3-809e-69538203afbc)
![image](https://github.com/user-attachments/assets/a3ffc956-04b7-47d0-bfbd-8b1ebb42d2c2)
![image](https://github.com/user-attachments/assets/6187c9de-584e-4e1d-bea8-5aba08e3b089)
![image](https://github.com/user-attachments/assets/306f8ed5-f5ac-4f98-8b54-d448d26b283f)
![image](https://github.com/user-attachments/assets/5d0fb005-8dba-48d4-ad10-592579ef46c4)

## Info - Ansible error handling
Refer https://docs.ansible.com/ansible/latest/playbook_guide/playbooks_error_handling.html

## Info - Ansible variable precedence
https://docs.ansible.com/ansible/latest/playbook_guide/playbooks_variables.html

## Lab - Ansible vault
<pre>
- ansible vault helps us protect sensitive data like login credentials, public/private keys, etc by encrypting with a password
- the password protected ansible vault file data can be securely accessed from playbooks
- Hashicorp vault is an alternate to Ansible vault
</pre>

Creating a vault protected file, when prompts for password type any password
```
ansible-vault create mysql-login-credentials.yml
cat mysql-login-credentials.yml
ansible-vault view mysql-login-credentials.yml
ansible-vault edit mysql-login-credentials.yml
ansible-vault decrypt mysql-login-credentials.yml
ansible-vault encrypt mysql-login-credentials.yml
```

Expected output
![image](https://github.com/user-attachments/assets/a0b70c3a-84c3-440d-9063-e7e7be5ea49d)
![image](https://github.com/user-attachments/assets/7692a7b5-f676-4c40-b3c0-ce6724723b5a)
![image](https://github.com/user-attachments/assets/555b6949-3eba-457c-8840-4baa92d425c7)
![image](https://github.com/user-attachments/assets/52a62bca-d116-484a-b303-d123dff9ce9e)
![image](https://github.com/user-attachments/assets/2ceb1285-7e65-4032-be0d-9b2976f2b720)


Accessing vault protected file from a playbook
```
cd ~/terraform-2428-feb2025
git pull
cd Day2/ansible/vault
cat mysql-login-credentials.yml
ansible-playbook playbook --ask-vault-pass
cat mysql-login-credentials.yml
```

Expected output
![image](https://github.com/user-attachments/assets/5655bd9d-f50a-4825-ae4b-b9093c083305)
![image](https://github.com/user-attachments/assets/56190462-a523-43b7-b114-aae0606d68c6)

## Info - Ansible Automation Platform
<pre>
- was called Ansible Tower earlier
- comes in 2 flavours
  - AWX ( open source )
    - this is developed on top of Ansible core opensource
    - this supports web interface
  - Red Hat Ansible Automation Platform ( enterprise - commercial product )
    - this is developed on top of opensource AWX
    - this supports web interface
</pre>

Info - Installing Ansible Tower (AWX)
https://medium.com/@jegan_50867/installing-ansible-tower-awx-e46d5231357d

## Lab - Launching AWX - Ansible Tower opensource variant
```
minikube start
kubectl get nodes
minikube service awx-demo-service --url -n ansible-awx
```

To retrieve AWX password ( save the password in a file to avoid typing this lengthy command each time )
```
kubectl get secret awx-demo-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```

Expected output
![image](https://github.com/user-attachments/assets/148c0187-52a5-40b8-af99-2beaa8391c36)

AWX Tower Dashboard can be accessed at the below url
<pre>
http://192.168.49.2:31225  
</pre>

Expected output
![image](https://github.com/user-attachments/assets/d4f5e833-2a2a-40e8-aee6-242004934de0)
![image](https://github.com/user-attachments/assets/3ab099d8-58b7-4523-a634-35437e289b94)

## Lab - Creating credential to store the private key file
Navigate to AWX Dashboard
![image](https://github.com/user-attachments/assets/3a203648-2d3f-43d1-8597-c3e7d043237b)
Click "Resource --> Credentials"
![image](https://github.com/user-attachments/assets/2b7246c9-4bf7-45d1-94aa-497d5d3c21fe)
Click "Add"
![image](https://github.com/user-attachments/assets/a8631719-e941-4715-80b5-5ed3b7d1ce8d)
![image](https://github.com/user-attachments/assets/cfdca274-081d-4d75-b134-cc873541bcac)
To print the private key ( type this command in your terminal )
```
cat ~/.ssh/id_ed25519
```
![image](https://github.com/user-attachments/assets/cfd87474-634e-46b6-bdfe-c805ad2c3cde)
![image](https://github.com/user-attachments/assets/b8b2646b-ae5a-4189-9692-97d890476896)
![image](https://github.com/user-attachments/assets/186e739f-a9b8-4260-9059-e623c40b4dff)
Click "Save"
![image](https://github.com/user-attachments/assets/a8127662-542c-403e-8917-38abc4631227)

## Lab - Let's create a project in Ansible Tower
Navigate to AWX Dashboard
![image](https://github.com/user-attachments/assets/3a203648-2d3f-43d1-8597-c3e7d043237b)

Click "Resource --> Projects"
![image](https://github.com/user-attachments/assets/e8d12363-73b5-4ad0-ab83-dfa51a15f437)
Click "Add"
![image](https://github.com/user-attachments/assets/4a3fd653-53af-4a76-a113-fa7a3a1ac53d)
![image](https://github.com/user-attachments/assets/111ab649-4f27-4669-81ac-51f81451c234)
Select "Git"
![image](https://github.com/user-attachments/assets/b89b0667-e1aa-43c2-a6ff-d9f6b6a63296)
![image](https://github.com/user-attachments/assets/05db3e4b-4cd2-43e0-8406-67cf779ee202)
Click "Save"
![image](https://github.com/user-attachments/assets/b7c949cb-02e2-42be-85b2-1f9d43a22504)
![image](https://github.com/user-attachments/assets/6489c52b-b7c6-4010-88eb-7f3595d6f28c)
![image](https://github.com/user-attachments/assets/c607344b-bfe8-40db-9b04-6fa92721d6e6)
![image](https://github.com/user-attachments/assets/27ad7838-c423-45cf-996b-278e91df0d05)
![image](https://github.com/user-attachments/assets/6745431e-9059-4a3c-9604-066e85cf744b)

## Lab - Creating an Inventory in Ansible Tower
Navigate to AWX Dashboard
![image](https://github.com/user-attachments/assets/3a203648-2d3f-43d1-8597-c3e7d043237b)

Click "Resource --> Inventories"
![image](https://github.com/user-attachments/assets/68a674aa-4a19-4b36-9d7c-eef082f47417)
Click "Add"
![image](https://github.com/user-attachments/assets/52012c2b-d80c-4e5f-a7ba-001c1351059b)
Select "Add Inventory"
![image](https://github.com/user-attachments/assets/524e97d0-ef79-46ca-b75d-e04e4a63bdbe)
![image](https://github.com/user-attachments/assets/1e39edb9-7d08-430a-9adc-8aa858dd3fc2)
Click "Save"
![image](https://github.com/user-attachments/assets/32bccb46-3fd5-4266-bf17-ad79c37189b2)

Click "Hosts" Tab within the Inventory we created
![image](https://github.com/user-attachments/assets/f8100566-4c99-46ce-b74e-9bad538d10b5)
Click "Add"
![image](https://github.com/user-attachments/assets/50b825a1-669c-4f96-98a5-5e8fd30ba8c1)
![image](https://github.com/user-attachments/assets/c62255ac-d0ad-4c13-b624-5d63259a4c40)
Click "Save"
![image](https://github.com/user-attachments/assets/e1679c16-1a45-46de-a5b9-438f4ff70319)

Let's add Ubuntu2 Hosts
Click "Add"
![image](https://github.com/user-attachments/assets/2111a1ba-f991-4c86-98a6-a8d195d71642)
Click "Save"
![image](https://github.com/user-attachments/assets/98510b92-6893-4e6b-a082-9ef1b5ebeca2)


Let's add Rocky1 Host
Click "Add"
![image](https://github.com/user-attachments/assets/679a2346-02e5-4f00-9ed5-e458a43ff298)
![image](https://github.com/user-attachments/assets/d36a2683-bec3-42ec-bc9c-34aa36358763)
Click "Save"
![image](https://github.com/user-attachments/assets/d4695dda-3dd8-4334-8bcb-d2f0f6418650)

Let's add Rocky2 host
![image](https://github.com/user-attachments/assets/4a568e5c-ca4a-46d8-a74a-584cedd0dfb4)
Click "Add"
![image](https://github.com/user-attachments/assets/e7c92359-4625-422a-b101-6ade0f715a4f)
Click "Save"
![image](https://github.com/user-attachments/assets/8ae01dbc-69a5-4522-bf6d-2c1be2a60f2d)
![image](https://github.com/user-attachments/assets/5d1a7df8-3fd7-4ad1-bd18-1df5d120236c)


## Lab - Let's create Template to run an Ansible Playbook in Ansible Tower
Navigate to AWX Dashboard
![image](https://github.com/user-attachments/assets/3a203648-2d3f-43d1-8597-c3e7d043237b)

Click "Resource --> Templates"
![image](https://github.com/user-attachments/assets/df29b26f-9808-4a44-ae4c-2d1f1e387173)
Click "Add"
![image](https://github.com/user-attachments/assets/a02636a5-53c7-41f3-adac-f7252220b359)
Select "Add Job template"
![image](https://github.com/user-attachments/assets/82cd3440-0c2a-4a37-bd5f-30113ce069fe)
Search and Select Inventory "Docker Inventory"
![image](https://github.com/user-attachments/assets/7702e175-2648-411b-b264-be4db7246e0c)
Search and Select Project "TekTutor Training Repositoyr"
![image](https://github.com/user-attachments/assets/78d83e59-4fc5-4f55-a0c5-cbefee30f594)
![image](https://github.com/user-attachments/assets/f3e94196-dc8d-4cea-87e5-16bac0499dbe)
![image](https://github.com/user-attachments/assets/dcbba9fb-dd0d-41ac-b7b8-987456c717b5)
Under Playbook, select "Day2/ansible/after-refactoring-playbook/install-nginx-playbook.yml
![image](https://github.com/user-attachments/assets/dcfd607c-11af-49ac-b763-8bbc65cb5f85)

Under Credentials, 
![image](https://github.com/user-attachments/assets/bdec8a60-635f-4471-8dd6-23e4efd56caa)
Select "PrivateKey"
![image](https://github.com/user-attachments/assets/78c4a4ea-01d6-48da-81be-de4ae04c2ad4)
![image](https://github.com/user-attachments/assets/ced637b4-2ea7-489b-b8ca-878589d5f212)
Click "Save"
![image](https://github.com/user-attachments/assets/d5885161-080c-4263-82eb-bc3fb70c44ab)
![image](https://github.com/user-attachments/assets/2fbf4525-4e93-462f-9ece-bc0d137219da)
![image](https://github.com/user-attachments/assets/c28174aa-3eaa-4db0-b558-055f2f989acb)
![image](https://github.com/user-attachments/assets/8fd2a3b2-d379-4f60-94ae-177eb3955966)
![image](https://github.com/user-attachments/assets/639ff489-e040-471d-9517-29e0275ca045)
![image](https://github.com/user-attachments/assets/e391978a-82dd-4481-b4ba-7dbd78328f25)

## Info - Golang Overview
<pre>
- is developed Google in C programming language
- it is compiled language
- it is strongly typed
- when golang was developed, the inventors wanted to keep it simple and support only relevant unambiguous features
- golang supports many built-in data types
  - int, int8, int16,int32, int64
  - byte,
  - float32, float64
  - string
  - struct ( user-defined data type )
  - there is no class
- golang loops
  - only for loop is supported
- if else, switch case are supported
- performance
  - it is faster than any interpreted language
  - golang compiles faster than most compiled languages
- golang is command-line, but you have opensource framework to develop desktop, mobile applications
- golang supports web development
- one could develop REST API and/or microservices in golang
</pre>

## Lab - Writing your first golang program

Create a file named hello.go with the below content
<pre>
package main

import "fmt"

//Entry point function - default function that would be invoked first
//This function must be defined under package main only
func main() {
  fmt.Println( "Hello Golang !" )
}  
</pre>

You may run the program as shown below
```
go run ./hello.go
```

Expected output
![image](https://github.com/user-attachments/assets/af1f5b7f-00c0-4d47-a52e-61a71284639f)
