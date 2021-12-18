# go-notification

## Usage

```go
package main

import "github.com/anatolethien/go-notification"

func main() {
    // Instantiate the notification
    n := new(notification.Notification)

    // Add a title to the notification
    n.Title = "Timer"
    // Add a description to the notification
    n.Body = "Time's up!"
    // Change the notification icon
    n.Icon = "alarm"
    
    // Push the notification to the desktop
    n.Push()
}
```
