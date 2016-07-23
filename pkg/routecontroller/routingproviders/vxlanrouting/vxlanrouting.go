package vxlanrouting

// VXLAN tunnels with layer-2 routing seems to work... tried manually with:

// on each host:
// ip link add vxlan1 type vxlan id 1 dev eth0 dstport 4789
// ip link set vxlan1 address 54:8:64:40:02:01
// ip addr add 10.244.0.0/32 dev vxlan1
// ip link set up vxlan1

// and then on each host, for each peer:
// bridge fdb add to 54:08:64:40:01:01 dst 172.20.27.211 dev vxlan1
// arp -i vxlan1 -s 10.244.100.1 54:08:64:40:01:01
// ip route add 10.244.100.0/32 dev vxlan1
// ip route add 10.244.100.0/24 via 10.244.100.0

