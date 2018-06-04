使用公司提供的私有GitLab仓库的同时，使用GitHub的共有仓库。这两个账户的邮箱不同，需要使用两个ssh-key
* 生成公司用的GitLab的SSH-Key
```
$ ssh-keygen -t rsa -C "1email@company.com" -f ~/.ssh/id-rsa
```
* 生成GitHub的SSH-key
```
$ ssh-keygen -t rsa -C "2email@github.com” -f ~/.ssh/github-rsa
```
此时，.ssh目录下会有四个文件，id_rsa.pub,id_rsa,github_rsa,github_rsa.pub
* 为电脑添加两个私钥
```
$ ssh-add ~/.ssh/id-rsa 
$ ssh-add ~/.ssh/github-rsa

# 如果执行ssh-add时提示”Could not open a connection to your authentication agent”，可以现执行命令:    
$ ssh-agent bash
# 然后再运行ssh-add命令。

# 可以通过 ssh-add -l 来确私钥列表
$ ssh-add -l
# 可以通过 ssh-add -D 来清空私钥列表
$ ssh-add -D
```
* 修改本地git的配置文件
```
# 若.ssh目录下无config文件，那么创建
touch config

# 添加以下内容
# gitlab
Host gitlab.com
    HostName gitlab.com
    PreferredAuthentications publickey
    IdentityFile ~/.ssh/id_rsa
# github
Host github.com
    HostName github.com
    PreferredAuthentications publickey
    IdentityFile ~/.ssh/github_rsa
```
* 验证是否可用
```
$ ssh -T git@github.com
Hi githublwp! You've successfully authenticated, but GitHub does not provide shell access.

$ ssh -T git@git.inspur.com
Welcome to GitLab, 刘伟鹏!

```
