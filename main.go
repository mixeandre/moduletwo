package moduletwo

import (
	"fmt"
	"time"
)

func GetData() string {
	return "Data from module two version: 1.0.0 " + fmt.Sprint(time.Now())
}