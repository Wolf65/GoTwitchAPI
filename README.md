# GoTwitchAPI

Installation
-----------
Install GoTwitchAPI using the "go get" command:

        go get github.com/Wolf65/GoTwitchAPI

Example
-----------
```
package main

import (
	"fmt"
	"github.com/Wolf65/GoTwitchAPI"
)

func main()  {
	t := gtapi.NewClient("twitch token")
	z := t.GetStreamByUser("monstercat")
	fmt.Println(z.Channel.ID)
	fmt.Println(z.Viewers)
	fmt.Println(z.Preview.Large)
	fmt.Println(z.Channel.ProfileBanner)
}
```

```
package main

import (
	"fmt"
	"github.com/Wolf65/GoTwitchAPI"
)

func main()  {
	t := gtapi.NewClient("twitch token")
	top,_ := t.GetTopStream()
	fmt.Println((*top.Streams)[1].Viewers)
	fmt.Println((*top.Streams)[4].Viewers)
}
```
