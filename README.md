### Network CLI tool

based on https://www.youtube.com/watch?v=i2p0Snwk4gc

```
#example below

$ ./cli 
NAME:
   Website Lookup CLI - Let's you query IPs, CNAMEs, MX records and Name Servers!

USAGE:
   cli [COMMAND] --host [HOSTNAME]

VERSION:
   0.1

COMMANDS:
   ns       Looks Up the NameServers for a Particular Host
   ip       Looks Up the IP addresses for a Particular Host
   cn       Looks Up the CNAME for a Particular Host
   mx       Looks Up the MX records for a Particular Host
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

$ ./cli ns --host google.com
ns4.google.com.
ns3.google.com.
ns2.google.com.
ns1.google.com.

$ ./cli ip --host google.com
172.217.28.238
2800:3f0:4001:801::200e

$ ./cli cn --host google.com
google.com.

$ ./cli mx --host google.com
aspmx.l.google.com. 10
alt1.aspmx.l.google.com. 20
alt2.aspmx.l.google.com. 30
alt3.aspmx.l.google.com. 40
alt4.aspmx.l.google.com. 50

#if you do not specifc the host, receive error
$ ./cli mx
2020/11/07 12:11:34 lookup : no such host
```