seekerd
=======

`seekerd` uses the Linux `inotify` functionality to see when SudoSH has 
written log output, then it calls rsync to move it to another host.

This needs to run in the background for Spyglass to work correctly. Make
sure it starts using your system's init system of choice (systemd, upstart, 
sysvinit)

## Use

Simply `seekerd /path/to/seekerd.conf`

`seekerd.conf` is a simple JSON-based config file:
```
{
  "LogLocation": "user@10.10.10.10:audit/.",
  "RsyncOpts": "-avSPpq",
  "RsyncPath": "/usr/bin/rsync"
}
```

As you may guess, you will need SSH keys properly configured for this to work.

## Future Work
This code could use:

* Hysteresis. Specifically, the program needs to NOT send a log file with
  every keystroke. Same could be said for screen updates?
* Not calling exec. I'd like a SCP library to do this natively in Go, but 
  I didn't have anything that looked particularly well done. Maybe this has
  changed by the time you, dear reader, are looking at it!
