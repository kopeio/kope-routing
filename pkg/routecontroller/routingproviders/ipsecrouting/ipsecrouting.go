package vxlanrouting

on each node:

apt-get install --yes  ipsec-tools

// To dump
// setkey -D && setkey -DP

// To reset
// setkey -F && setkey -FP



node 10.244.1.1/24 172.20.6.74

ME=172.20.6.74
REMOTE=172.20.6.75
ENCAP=esp-udp
# Would be great to use esp-udp

setkey -F && setkey -FP
setkey -c <<EOF
add ${ME} ${REMOTE} ah 24500 -A hmac-md5 "1234567890123456";
add ${ME} ${REMOTE} ${ENCAP} 24501 -E 3des-cbc "123456789012123456789012";
add ${REMOTE} ${ME} ah 24502 -A hmac-md5 "1234567890123456";
add ${REMOTE} ${ME} ${ENCAP}  24503 -E 3des-cbc "123456789012123456789012";

spdadd ${ME} ${REMOTE}[4500] udp -P out prio 100 none;
spdadd ${ME} ${REMOTE}[4500] udp -P in prio 100 none;
spdadd ${REMOTE} ${ME}[4500] udp -P in prio 100 none;
spdadd ${REMOTE} ${ME}[4500] udp -P out prio 100 none;

spdadd ${ME} ${REMOTE} any -P out ipsec esp/transport//require ah/transport//require;
spdadd ${REMOTE} ${ME} any -P in ipsec esp/transport//require ah/transport//require;

EOF
setkey -D
setkey -DP


node2 10.244.2.1/24 172.20.6.75

REMOTE=172.20.6.74
ME=172.20.6.75
ENCAP=esp-udp

setkey -F && setkey -FP
setkey -c <<EOF
add ${REMOTE} ${ME} ah 24500 -A hmac-md5 "1234567890123456";
add ${REMOTE} ${ME} ${ENCAP} 24501 -E 3des-cbc "123456789012123456789012";
add ${ME} ${REMOTE} ah 24502 -A hmac-md5 "1234567890123456";
add ${ME} ${REMOTE} ${ENCAP} 24503 -E 3des-cbc "123456789012123456789012";

spdadd ${ME} ${REMOTE}[4500] udp -P out prio 100 none;
spdadd ${ME} ${REMOTE}[4500] udp -P in prio 100 none;
spdadd ${REMOTE} ${ME}[4500] udp -P in prio 100 none;
spdadd ${REMOTE} ${ME}[4500] udp -P out prio 100 none;

spdadd ${ME} ${REMOTE} any -P out ipsec esp/transport//require  ah/transport//require;
spdadd ${REMOTE} ${ME} any -P in ipsec esp/transport//require  ah/transport//require;

EOF
setkey -D
setkey -DP

==========================

# NODE1

ip tunnel add tonode2 mode gre remote ${NODE2} local ${NODE1} ttl 255
ip link set tonode2 up
ip route add ${CIDR2} dev tonode2


# NODE2

ip tunnel add tonode1 mode gre remote ${NODE1} local ${NODE2} ttl 255
ip link set tonode1 up
ip route add ${CIDR1} dev tonode1





// Tunnel mode (ARP problems)

master 10.244.1.1/24 172.20.6.74

setkey -F && setkey -FP

echo 'add 172.20.6.74 172.20.6.75 esp 34501 -m tunnel -E 3des-cbc "123456789012123456789012";' | setkey -c
echo "spdadd 10.244.1.1/24 10.244.2.1/24 any -P out ipsec esp/tunnel/172.20.6.74-172.20.6.75/require;" | setkey -c

echo 'add 172.20.6.75 172.20.6.74 esp 34502 -m tunnel -E 3des-cbc "123456789012123456789012";' | setkey -c
echo "spdadd 10.244.2.1/24 10.244.1.1/24 any -P in ipsec esp/tunnel/172.20.6.75-172.20.6.74/require;" | setkey -c

ip route add 10.244.2.0/24 via 172.20.6.75

node 10.244.2.1/24 172.20.6.75

setkey -F && setkey -FP

echo 'add 172.20.6.74 172.20.6.75 esp 34501 -m tunnel -E 3des-cbc "123456789012123456789012";' | setkey -c
echo "spdadd 10.244.1.1/24 10.244.2.1/24 any -P in ipsec esp/tunnel/172.20.6.74-172.20.6.75/require;" | setkey -c

echo 'add 172.20.6.75 172.20.6.74 esp 34502 -m tunnel -E 3des-cbc "123456789012123456789012";' | setkey -c
echo "spdadd 10.244.2.1/24 10.244.1.1/24 any -P out ipsec esp/tunnel/172.20.6.75-172.20.6.74/require;" | setkey -c

ip route add 10.244.1.0/24 via 172.20.6.74





// Tunnel mode with fake IP

#node1 10.244.1.1/24 172.20.6.74

setkey -F && setkey -FP

echo 'add 172.20.6.74 172.20.6.75 esp 34501 -m tunnel -E 3des-cbc "123456789012123456789012";' | setkey -c
echo "spdadd 10.244.1.0/24 10.244.2.0/24 any -P out ipsec esp/tunnel/172.20.6.74-172.20.6.75/require;" | setkey -c

echo 'add 172.20.6.75 172.20.6.74 esp 34502 -m tunnel -E 3des-cbc "123456789012123456789012";' | setkey -c
echo "spdadd 10.244.2.0/24 10.244.1.0/24 any -P in ipsec esp/tunnel/172.20.6.75-172.20.6.74/require;" | setkey -c

ip route add 10.244.2.0/24 dev eth0

#node2 10.244.2.1/24 172.20.6.75

setkey -F && setkey -FP

echo 'add 172.20.6.74 172.20.6.75 esp 34501 -m tunnel -E 3des-cbc "123456789012123456789012";' | setkey -c
echo "spdadd 10.244.1.0/24 10.244.2.0/24 any -P in ipsec esp/tunnel/172.20.6.74-172.20.6.75/require;" | setkey -c

echo 'add 172.20.6.75 172.20.6.74 esp 34502 -m tunnel -E 3des-cbc "123456789012123456789012";' | setkey -c
echo "spdadd 10.244.2.0/24 10.244.1.0/24 any -P out ipsec esp/tunnel/172.20.6.75-172.20.6.74/require;" | setkey -c

ip route add 10.244.1.0/24 dev eth0
