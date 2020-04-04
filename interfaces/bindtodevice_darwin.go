// +build darwin solaris

package interfaces

import (
	"fmt"
	"net"

	"golang.org/x/sys/unix"
)

// BindToInterface emulates linux's SO_BINDTODEVICE option for a socket by using
// IP_RECVIF.
func BindToInterface(fd int, ifname string) error {
	iface, err := net.InterfaceByName(ifname)
	if err != nil {
		return err
	}
	if err := unix.SetsockoptInt(fd, unix.IPPROTO_IP, unix.IP_RECVIF, iface.Index); err != nil {
		return fmt.Errorf("Could not bind to an interface (receive path): %w", err)
	}

	socktype, err := unix.GetsockoptInt(fd, unix.SOL_SOCKET, unix.SO_TYPE)
	if err != nil {
		return fmt.Errorf("Could not decide socket address type: %w", err)
	}
	switch socktype {
	case unix.IPPROTO_IP:
		err = unix.SetsockoptInt(fd, unix.IPPROTO_IP, unix.IP_BOUND_IF, iface.Index)
	case unix.IPPROTO_IPV6:
		err = unix.SetsockoptInt(fd, unix.IPPROTO_IPV6, unix.IPV6_BOUND_IF, iface.Index)
	default:
		err = fmt.Errorf("Unsupported address family %d", socktype)
	}

	return err
}
