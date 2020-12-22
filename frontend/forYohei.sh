#!/bin/sh

sudo apt update && sudo apt upgrade -y
sudo apt install docker.io -y
sudo groupadd docker
sudo gpasswd -a $USER docker

sudo systemctl enable docker
sudo systemctl restart docker

echo "再起動してね！"
echo "何もしなければ　1分後に　再起動するよ"
echo "shutdown -c"
echo "で，取り消し可能"
