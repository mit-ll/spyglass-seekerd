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

## Disclaimer

DISTRIBUTION STATEMENT A. Approved for public release: distribution unlimited.

This material is based upon work supported by the Under Secretary of Defense for Research and Engineering under Air Force Contract No. FA8721-05-C-0002. Any opinions, findings, conclusions or recommendations expressed in this material are those of the author(s) and do not necessarily reflect the views of the Under Secretary of Defense for Research and Engineering.

The software/firmware is provided to you on an As-Is basis
