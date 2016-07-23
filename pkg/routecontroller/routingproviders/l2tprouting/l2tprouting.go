package l2tprouting

// L2TP routing seems to have the problem (for our use case) that we must allocate a unique
// UDP port for each tunnel...

//import (
//	"fmt"
//	"github.com/golang/glog"
//	"github.com/kopeio/route-controller/pkg/routecontroller"
//	"github.com/kopeio/route-controller/pkg/routecontroller/routingproviders"
//	"k8s.io/kubernetes/pkg/api"
//	"net"
//	"strings"
//	"encoding/binary"
//	"strconv"
//	"os/exec"
//)
//
//type L2TPRoutingProvider struct {
//}
//
//var _ routingproviders.RoutingProvider = &L2TPRoutingProvider{}
//
//func NewL2TPRoutingProvider() (*L2TPRoutingProvider, error) {
//	p := &L2TPRoutingProvider{}
//	return p, nil
//}
//
//func (p *L2TPRoutingProvider) EnsureCIDRs(me *api.Node, allNodes []api.Node) error {
//	// TODO: modprobe l2tp_eth
//
//	// To delete:  ip l2tp del tunnel tunnel_id 2886994109
//
//	localIP := routecontroller.FindInternalIPAddress(me)
//	if localIP == "" {
//		glog.Infof("self-node does not yet have internalIP; delaying configuration")
//		return nil
//	}
//
//	//links, err := routecontroller.QueryIPLinks()
//	//if err != nil {
//	//	return err
//	//}
//	//
//	//routes, err := routecontroller.QueryIPRoutes()
//	//if err != nil {
//	//	return err
//	//}
//
//	cidrMap := routecontroller.BuildCIDRMap(me, allNodes)
//
//	for _ /*remoteCIDRString*/, remoteIP := range cidrMap {
//		//_, remoteCIDR, err := net.ParseCIDR(remoteCIDRString)
//		//if err != nil {
//		//	return fmt.Errorf("error parsing PodCidr %q: %v", remoteCIDRString, err)
//		//}
//
//		//name := "gre-" + strings.Replace(remoteCIDR.IP.String(), ".", "-", -1)
//
//		// The tunnel id is the IP address of the other side
//		// (we could also make it the CIDR of the other side...)
//		localTunnelID := toUint32(remoteIP)
//		peerTunnelID := toUint32(localIP)
//		localSessionID := localTunnelID
//		peerSessionID := peerTunnelID
//		localPort := 5000
//		remotePort := 5000
//
//
//
//		//err = links.EnsureL2TPTunnel(localTunnelID, peerTunnelID, encap, localIP, remoteIP, localPort, remotePort)
//		//if err != nil {
//		//	return err
//		//}
//		{
//			conf := make(map[string]string)
//			conf["tunnel_id"] = strconv.FormatUint(uint64(localTunnelID), 10)
//			conf["peer_tunnel_id"] = strconv.FormatUint(uint64(peerTunnelID), 10)
//			conf["udp_sport"] = strconv.Itoa(localPort)
//			conf["udp_dport"] = strconv.Itoa(remotePort)
//			conf["encap"] = "udp"
//			conf["local"] = localIP
//			conf["remote"] = remoteIP
//
//			argv := []string{"ip", "l2tp", "add", "tunnel" }
//			for k, v := range conf {
//				argv = append(argv, k, v)
//			}
//			humanArgv := strings.Join(argv, " ")
//
//			glog.V(2).Infof("Running %q", humanArgv)
//			cmd := exec.Command(argv[0], argv[1:]...)
//
//			out, err := cmd.CombinedOutput()
//			if err != nil {
//				return fmt.Errorf("error running %q: %v: %q", humanArgv, err, string(out))
//			}
//		}
//
//		//err = links.EnsureL2TPSession(localTunnelID, sessionID, peerSessionID)
//		//if err != nil {
//		//	return err
//		//}
//		{
//			conf := make(map[string]string)
//			conf["tunnel_id"] = strconv.FormatUint(uint64(localTunnelID), 10)
//			conf["session_id"] = strconv.FormatUint(uint64(localSessionID), 10)
//			conf["peer_session_id"] = strconv.FormatUint(uint64(peerSessionID), 10)
//
//			argv := []string{"ip", "l2tp", "add", "session" }
//			for k, v := range conf {
//				argv = append(argv, k, v)
//			}
//			humanArgv := strings.Join(argv, " ")
//
//			glog.V(2).Infof("Running %q", humanArgv)
//			cmd := exec.Command(argv[0], argv[1:]...)
//
//			out, err := cmd.CombinedOutput()
//			if err != nil {
//				return fmt.Errorf("error running %q: %v: %q", humanArgv, err, string(out))
//			}
//		}
//
//		//err = routes.EnsureRouteToDevice(remoteCIDR.String(), name)
//		//if err != nil {
//		//	return err
//		//}
//
//		// TODO: Delete any extra tunnels
//	}
//
//	return nil
//}
//
//func toUint32(ipString string) uint32 {
//	ip := net.ParseIP(ipString)
//	if ip == nil {
//		glog.Fatalf("IP not valid: %v", ipString)
//	}
//	ip4 := ip.To4()
//	if ip4 != nil {
//		n := binary.BigEndian.Uint32(ip4)
//		return n
//	}
//	glog.Fatalf("IPv6 not supported: %v", ip)
//	return 0
//}