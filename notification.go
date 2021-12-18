package notification

import (
    "fmt"
    "log"
    "os/exec"
    "runtime"
)

type Notification struct {
    Title string
    Body string
    Icon string
}

func (n *Notification) Push() {
    switch runtime.GOOS {
    case "linux":
        path, err1 := exec.LookPath("notify-send")
        if err1 != nil {
            log.Fatal(err1)
        }
        cmd := exec.Command(path, n.Title, n.Body, "-i", n.Icon)
        err2 := cmd.Run()
        if err2 != nil {
            log.Fatal(err2)
        }
    default:
        log.Fatalf("%s: not supported")
    }
    fmt.Print("\a")
}
