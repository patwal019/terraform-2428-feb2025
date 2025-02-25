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
