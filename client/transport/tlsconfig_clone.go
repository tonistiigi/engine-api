// +build !go1.7

package transport

import "crypto/tls"

// TLSConfigClone returns a shallow copy of tls.Config. This function is
// provided for compatibility with go1.7 that doesn't include this method in stdlib.
func TLSConfigClone(c *tls.Config) *tls.Config {
	return c.Clone()
}
