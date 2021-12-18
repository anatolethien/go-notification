# go-notification

A package built to send custom notifications to the desktop.

## Usage

Define a `Notification` object, edit its attributes and then use the `Push()` method.

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

This will send this notification to the desktop and trigger a terminal sound.

![notification](https://raw.githubusercontent.com/anatolethien/images/master/go-notification/notification.png)

## Supported systems and dependencies

- Linux (`notify-send`)
