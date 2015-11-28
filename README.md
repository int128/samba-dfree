Disk space correction tool for Samba
====================================

Samba service may not report correct disk space if the published path consists from several disks.
`samba-dfree` returns total and available space calculated from all directories in the current directory.

For example, if the current directory has following directories,

* ./mountpoint1 -> /dev/sdb
* ./mountpoint2 -> /dev/sdc
* ./mountpoint3 -> /dev/sdd

`samba-dfree` will return total and available kilo-bytes of above 3 disks.

How to use
----------

Install `samba-dfree` into `/etc/samba`.

Configure samba.

```ini
# smb.conf
[global]
dfree command = /etc/samba/samba-dfree
dfree cache time = 60
```

Restart samba service.

Contributions
-------------

This is an open source software licensed under the Apache License Version 2.0. Feel free to open issues or pull requests.

