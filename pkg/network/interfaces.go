package network

import (
	"net"
)

// LS is like the ls command in unix which returns the list of files
// in the current directory. In this case Ls return the network interfaces
// that can be accessed.
func LS() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, i := range interfaces {
		names = append(names, i.Name)
	}
	return names, nil
}

// Interface returns the name of the actual interface given a paticular name.
func Interface(name string) (*net.Interface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range interfaces {
		if name == i.Name {
			return &i, nil
		}
	}
	return nil, errInterfaceNotFound

}
