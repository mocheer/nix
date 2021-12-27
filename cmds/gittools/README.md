
# go-git

# ssh-rsa 生成
1. ssh-keygen -t rsa -C "mocheer@foxmail.com"
2. 添加ssh公钥到github、gitlab、gogs、gitea任一平台
3. 添加ssh私钥进行本地认证

# ssh-Ed25519 
1. ssh-keygen -t ed25519 -f istrongcloud  -C "mocheer@foxmail.com"
2. 添加ssh公钥到github、gitlab、gogs、gitea任一平台
3. 添加ssh私钥进行本地认证

> 需要确保`known_hosts`有添加github、gitlab、gogs、gitea等平台的地址，没有的话，如果是命令行执行的git命令可以根据提示进行添加