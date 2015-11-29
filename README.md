Disk space calculator for Samba
===============================

Samba daemon does not report correct disk space if the exported path consists from several disks.

For example, it assumes `/home` is exported on Samba and consists from 3 disks as follows.
Samba daemon should report free space of 210GB but actually reports 10GB.

Mount point | Disk     | Free space
------------|----------|-----------
/           | /dev/sda | 10GB
/home/fiz   | /dev/sdb | 100GB
/home/buz   | /dev/sdc | 100GB

`samba-dfree` is an executable binary and called by samba daemon.
It shows disk space calculated from all directories in the exported path.

How to configure
----------------

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
