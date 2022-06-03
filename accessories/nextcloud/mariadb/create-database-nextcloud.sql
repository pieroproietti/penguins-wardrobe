CREATE USER 'nextcloud'@'localhost' IDENTIFIED BY 't3rr4s1n1';
CREATE DATABASE nextcloud;
GRANT ALL PRIVILEGES ON nextcloud.* TO 'nextcloud'@'localhost';
FLUSH PRIVILEGES;
QUIT
