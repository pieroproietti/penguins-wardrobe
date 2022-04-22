# add group smbguest for samba
groupadd smbguest
useradd -g smbguest -d /dev/null -s /bin/false smbguest
# add group smbuser for samba
groupadd smbuser
# umask 077
mkdir -p /etc/samba/private
touch /etc/samba/private/smbpasswd
# adding group smbuser to user
usermod -aG smbuser $SUDO_USER
# creating sample users
# useradd -d /home/argo -g smbuser -s /bin/false -m argo
# echo argo:evolution | chpasswd
# smbpasswd -a argo
# useradd -d /home/profes -g smbuser -s /bin/false -m profes
# echo profes:evolution | chpasswd
# smbpasswd -a profes
#groupadd workstation
#useradd -d /dev/null -g workstation -s /bin/false samba
#smbpasswd -a -m samba

