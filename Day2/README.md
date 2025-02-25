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
