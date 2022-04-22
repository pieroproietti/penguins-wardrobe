# create /etc/samba/private/smbpasswd
# umask 077
mkdir -p /etc/samba/private
touch /etc/samba/private/smbpasswd

# add groups for samba
groupadd smbguest
groupadd smbuser
groupadd workstation

# add user smdguest for samba
useradd -g smbguest -d /dev/null -s /bin/false smbguest

# adding user $SUDO_USER to group smbuser
usermod -aG smbuser $SUDO_USER

# creating sample users
# useradd -d /home/argo -g smbuser -s /bin/false -m argo
# echo argo:evolution | chpasswd
# echo argo:evolution | smbpasswd -a argo
# useradd -d /home/profes -g smbuser -s /bin/false -m profes
# echo profes:evolution | chpasswd
# echo argo:evolution |smbpasswd -a profes

# adding workstation
# useradd -d /dev/null -g workstation -s /bin/false samba
# smbpasswd -a -m samba

