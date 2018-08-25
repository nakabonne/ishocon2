#!/bin/bash
set -e
cd /home/ishocon/webapp
git pull
sudo cp etc/nginx.conf /etc/nginx/nginx.conf
cd /home/ishocon/webapp/go
#make TODO: ビルド
sudo /usr/sbin/nginx -t
sudo service nginx reload
#sudo service isubata.golang restart TODO: リスタート
