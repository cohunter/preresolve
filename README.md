# Module to use Cloudflare DNS-over-HTTPS

## DNS resolution in the default HTTP/S client without depending on the underlying system

For a cross-compiled Go program, DNS lookups on Android were failing due to no /etc/resolv.conf

Simple fix, just make it work: Replace the default http(s) transport with one that does name resolution itself by querying Cloudflare's DoH at 1.1.1.1

## Usage:

Just import it in the program where it is needed:

    import _ github.com/cohunter/preresolve

That's it. Names passed to the transport used by the default HTTP/S client will be replaced by an IP address from a DNS-over-HTTPS lookup. Will panic() on error.

This worked for my purposes, but I do not warrant that it is suitable for any particular use case.

License: CC0 or 0BSD