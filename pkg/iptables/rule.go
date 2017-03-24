// Package iptables is a wrapper around the iptables binary
package iptables

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
)

var (
	// ErrTargetUnknown occurs when the target is unknown
	ErrTargetUnknown = errors.New("Unknown target")

	// ErrWrongChain occurs when a rule cannot be added to a certain
	// chain (i.e redirect in anything other than prerouting)
	ErrWrongChain = errors.New("Rule cannot be added to this chain")

	// ErrNoPorts occurs when src and dest ports are required but not provided
	ErrNoPorts = errors.New("Ports were not provided")

	// ErrNoDestination occurs when a destination is needed but not provided
	ErrNoDestination = errors.New("No destination provided")

	// ErrNoDestinationPort occurs when a destination port is needed but not provided
	ErrNoDestinationPort = errors.New("No destination port provided")

	// ErrWrongProtocol occurs when a protocol is provided that does not fit the target
	ErrWrongProtocol = errors.New("Wrong protocol")

	// ErrNoInterfaces occurs when interfaces are required but not provided
	ErrNoInterfaces = errors.New("No interfaces provided")

	// ErrIPWrongFormat occurs when the supplied IP address or range is malformed
	ErrIPWrongFormat = errors.New("Malformed IP Address (range)")
)

func isIP(ip string) bool {
	r := regexp.MustCompile("((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.|$)){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\/(0|3[0-2]|[1-2][0-9]|[0-9]))?")
	return r.MatchString(ip)
}

// ToString returns the string representation of a rule
func (r *Rule) ToString() (string, error) {
	switch r.Target {
	case "REDIRECT":
		if r.Chain != "PREROUTING" {
			return "", ErrWrongChain
		}
		if r.SourcePort == 0 || r.DestinationPort == 0 {
			return "", ErrNoPorts
		}
		if r.Destination == "" {
			return "", ErrNoDestination
		}
		if !isIP(r.Destination) {
			return "", ErrIPWrongFormat
		}
		if !(r.Protocol == "tcp" || r.Protocol == "udp") {
			return "", ErrWrongProtocol
		}

		str := fmt.Sprintf("-t nat -A PREROUTING --dst %s -p %s --dport %d -j REDIRECT --to-port %d", r.Destination, r.Protocol, r.SourcePort, r.DestinationPort)

		return str, nil

	case "ACCEPT", "DROP":
		if !(r.Chain == "INPUT" || r.Chain == "OUTPUT" || r.Chain == "REDIRECT") {
			return "", ErrWrongChain
		}

		if r.Destination == "" {
			return "", ErrNoDestination
		}
		if !isIP(r.Destination) {
			return "", ErrIPWrongFormat
		}
		str := fmt.Sprintf("-A %s -j %s --dst %s", r.Chain, r.Target, r.Destination)

		if r.Protocol != "" {
			str = fmt.Sprintf("%s -p %s", str, r.Protocol)
		}

		if r.DestinationPort != 0 {
			str = fmt.Sprintf("%s --dport %d", str, r.DestinationPort)
		}

		if r.In != "" {
			str = fmt.Sprintf("%s -i %s", str, r.In)
		}

		if r.Out != "" {
			str = fmt.Sprintf("%s -o %s", str, r.Out)
		}

		if r.State != "" {
			str = fmt.Sprintf("%s -m state --state %s", str, r.State)
		}

		return str, nil
	default:
		return "", ErrTargetUnknown
	}
}

// GetHash returns the hash value of a rule
func (r *Rule) GetHash() string {
	str, err := r.ToString()
	if err != nil {
		return ""
	}
	data := []byte(str)
	return fmt.Sprintf("%x", sha256.Sum256(data))
}
