#!/bin/bash
set -e
cd /home/ishocon/webapp
git pull
sudo cp etc/nginx.conf /etc/nginx/nginx.conf
cd /home/ishocon/webapp/go
make build
sudo /usr/sbin/nginx -t
sudo service nginx reload
#sudo service isubata.golang restart TODO: リスタート
