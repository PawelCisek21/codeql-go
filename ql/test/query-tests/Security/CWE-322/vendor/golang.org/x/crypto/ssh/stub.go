// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for golang.org/x/crypto/ssh, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: golang.org/x/crypto/ssh (exports: ClientConfig,AuthMethod,HostKeyCallback,PublicKey; functions: InsecureIgnoreHostKey)

// Package ssh is a stub of golang.org/x/crypto/ssh, generated by depstubber.
package ssh

import (
	io "io"
	net "net"
	time "time"
)

type AuthMethod interface{}

type BannerCallback func(string) error

type ClientConfig struct {
	Config            Config
	User              string
	Auth              []AuthMethod
	HostKeyCallback   HostKeyCallback
	BannerCallback    BannerCallback
	ClientVersion     string
	HostKeyAlgorithms []string
	Timeout           time.Duration
}

func (_ *ClientConfig) SetDefaults() {}

type Config struct {
	Rand           io.Reader
	RekeyThreshold uint64
	KeyExchanges   []string
	Ciphers        []string
	MACs           []string
}

func (_ *Config) SetDefaults() {}

type HostKeyCallback func(string, net.Addr, PublicKey) error

func InsecureIgnoreHostKey() HostKeyCallback {
	return nil
}

type PublicKey interface {
	Marshal() []byte
	Type() string
	Verify(_ []byte, _ *Signature) error
}

type Signature struct {
	Format string
	Blob   []byte
	Rest   []byte
}
