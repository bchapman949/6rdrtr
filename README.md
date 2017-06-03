# 6rdrtr
Set up 6RD on your router. Run `6rdrtr`. Suddenly IPv6.

This utility is meant to be a set-and-forget tool that looks at your interfaces and just Does The Right Thing.

## Reference Links
1. [radvd](https://github.com/reubenhwk/radvd), the de-facto Linux app for stateless configuration of clients
2. [The golang.org/x/net/icmp package](https://godoc.org/golang.org/x/net/icmp)
3. [skoef's ndp package](https://github.com/skoef/ndp), which defines useful ICMPv6 datatypes
4. [skoef's dhcpv6 package](https://github.com/skoef/dhcpv6)

## Core Functionality
* `6rdrtr` should require two arguments: your external ipv6 interface (eg `6rd` or `sit0`), and your internal interface (eg `eth0` or `lan0`).
* It should be able to automatically add an address to your internal interface that is within the prefix of your 6rd interface.
* It should be able to monitor your external interface for changes, to detect if the IPv6 address has changed, and:
  * deprecate the old route immediately
  * update the address of your internal interface to match
  * immediately begin advertising a new prefix and route

## Additional Functionality
* `6rdrtr` should be able to detect if your Linux kernel is configured incorrectly by reading `/proc`
* It should also be able to automatically correct some issues given an appropriate argument (ala `fsck -p`)

## Current Non-goals
* `6rdrtr` won't configure IP routing on your router. Your initscripts/systemd will be better equipped to handle dependency graphs and such.
* It won't configure your 6rd interface for you, though this is not impossible down the road.
* It won't actually route packets. This is up to your kernel.

## Why not `radvd`?
`radvd` operates on a straightforward principle: the prefix you configure is the prefix it advertises. It has a couple of hacks for meeting some narrow usecases, like classic 6to4 addressing, but by-and-large it is focused on statically assigned prefixes. It is very awkward and frustrating to try to contort `radvd` into advertising a prefix that changes every once in a while, such as a 6rd prefix you might receive from your DSL provider.

In addition, while `radvd` will monitor an interface for changes (well... kind of) it will not actually create a routable address on your internal interface for your internal clients to use. So short of more contorting and flailing, a single utility that handles both (related) tasks seems simpler.
