package main

import (
    "github.com/romanoff/fsmonitor"
    "fmt"
    "log"
    "strings"
    "os"
    "os/exec"
    "encoding/json"
    )

type Config struct {
    LogLocation string
    RsyncOpts string
    RsyncPath string
}

var config Config

func main() {
    configfile := os.Args[1]
    if len(configfile) == 0 {
      fmt.Printf("A configuration file was not specified as the first argument.\n")
      os.Exit(1)
    }

    if _, err := os.Stat(configfile); os.IsNotExist(err) {
      fmt.Printf("Configuration file does not exist: %s\n", configfile)
      os.Exit(1)
    }

    file, _ := os.Open(os.Args[1])
    decoder := json.NewDecoder(file)
    config = Config{}
    err := decoder.Decode(&config)
    if err != nil {
      fmt.Println("Error decoding Config JSON: ", err)
    }

    watcher, err := fsmonitor.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    err = watcher.Watch("/var/lib/docker/aufs/diff")
    if err != nil {
        log.Fatal(err)
    }
    for {
        select {
        case ev := <-watcher.Event:
            if strings.Contains(ev.Name, "sudosh") {
                 if ev.IsModify() {
                    // Some hysteresis would be good here to cut down on the number of 
                    // transmissions; currently this calls two rsyncs for every keystroke.
                    cmd := exec.Command(config.RsyncPath, config.RsyncOpts, 
                                        ev.Name, config.LogLocation)
                    err := cmd.Start()
                    if err != nil {
                        fmt.Println(err.Error())
                    }
                    
                }
            }
        case err := <-watcher.Error:
            fmt.Println("error:", err)
        }
    }
}
