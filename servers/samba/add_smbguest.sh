# add group smbguest for samba
groupadd smbguest
useradd -g smbguest -d /dev/null -s /bin/false smbguest
# add group smbuser for samba
groupadd smbuser
useradd -d /home/argo -g smbuser -s /bin/false -m argo
smbpasswd -a argo
useradd -d /home/profes -g smbuser -s /bin/false -m profes
smbpasswd -a profes
groupadd workstation
useradd -d /dev/null -g workstation -s /bin/false samba
smbpasswd -a -m samba

