#!/bin/bash
COSTUME=$(hostname)
IPADDRESS=$(hostname -i)

# /etc/hostname
echo ${COSTUME} > /etc/hostname

# /etc/hosts
cat << 'EOF' > /etc/hosts
127.0.0.1 localhost localhost.localdomain
IPADDRESS COSTUME COSTUME.localdomain pvelocalhost
# The following lines are desirable for IPv6 capable hosts
:: 1     ip6 - localhost ip6 - loopback
fe00:: 0 ip6 - localnet
ff00:: 0 ip6 - mcastprefix
ff02:: 1 ip6 - allnodes
ff02:: 2 ip6 - allrouters
ff02:: 3 ip6 - allhosts
EOF

# replace COSTUME
sed -i 's@COSTUME@'"${COSTUME}"'@g' /etc/hosts
sed -i 's@IPADDRESS@'"${IPADDRESS}"'@g' /etc/hosts
