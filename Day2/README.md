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
