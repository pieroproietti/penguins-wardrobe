sudo apt update
sudo apt -y install mariadb-server mariadb-client
sudo mysql_secure_installation 
sudo mysql -u root -p
CREATE USER 'nextcloud'@'localhost' IDENTIFIED BY 't3rr4s1n1';
CREATE DATABASE nextcloud;
GRANT ALL PRIVILEGES ON nextcloud.* TO 'nextcloud'@'localhost';
FLUSH PRIVILEGES;
QUIT

sudo apt -y install php php-{cli,xml,zip,curl,gd,cgi,mysql,mbstring}
sudo apt -y install apache2 libapache2-mod-php

sudo nano /etc/php/*/apache2/php.ini
date.timezone = Europe/Roma
memory_limit = 512M
upload_max_filesize = 500M
post_max_size = 500M
max_execution_time = 300

sudo systemctl restart apache2

sudo apt -y install wget curl unzip
wget https://download.nextcloud.com/server/releases/latest.zip
unzip latest.zip
rm -f latest.zip

sudo mv nextcloud /var/www/html/
sudo chown -R www-data:www-data /var/www/html/nextcloud
sudo sudo chmod -R 755 /var/www/html/nextcloud

sudo a2dissite 000-default.conf
sudo rm /var/www/html/index.html
sudo systemctl restart apache2


