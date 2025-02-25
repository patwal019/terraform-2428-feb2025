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


## Lab - Let's run the refactored playbook which install nginx on rocky and ubuntu based ansible nodes
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
