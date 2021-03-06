<strong>bash code<strong>
```bash
#!/usr/bin/env bash

echo "This is nsswitch.conf :"
cat nsswitch.conf
echo
echo

echo "Do not show whiteline :"
grep -v '^\s*$' nsswitch.conf
echo
echo

echo "Do not show # :"
grep -v '#' nsswitch.conf
echo
echo
echo "Do not show whiteline and # :"
grep -v '#\|^\s*$' nsswitch.conf
echo
echo
read -p "Press enter to end..." 

```

nsswitch.conf 內容

<pre><code>#
# /etc/nsswitch.conf
#
# An example Name Service Switch config file. This file should be
# sorted with the most-used services at the beginning.
#
# The entry '[NOTFOUND=return]' means that the search for an
# entry should stop if the search in the previous entry turned
# up nothing. Note that if the search failed due to some other reason
# (like no NIS server responding) then the search continues with the
# next entry.
#
# Legal entries are:
#
#       compat                  Use compatibility setup
#       nisplus                 Use NIS+ (NIS version 3)
#       nis                     Use NIS (NIS version 2), also called YP
#       dns                     Use DNS (Domain Name Service)
#       files                   Use the local files
#       [NOTFOUND=return]       Stop searching if not found so far
#
# For more information, please read the nsswitch.conf.5 manual page.
#

passwd: compat
group:  compat
shadow: compat

hosts:  files dns
networks:       files dns

services:       files
protocols:      files
rpc:    files
ethers: files
netmasks:       files
netgroup:       files
publickey:      files

bootparams:     files
automount:      files nis
aliases:        files

</code></pre>

說明

<pre><code>grep -v '^\s*$' nsswitch.conf   過濾空白列
grep -v '#\|^\s*$' nsswitch.conf   過濾#
grep -v "key1\|key2\|key3"   過濾多個條件
grep -v '#\|^\s*$' nsswitch.conf 過濾#和空白列

</code></pre>

測試結果

![image](https://github.com/HongScarlet/homework/blob/master/bash/img/06shownsswitch.png)


***
