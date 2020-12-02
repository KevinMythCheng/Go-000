package main

import (
	"fmt"
	"github.com/kevinlcheng/gotrain/Week02/service"
)

func main() {
	//Do not find
	_, err := service.Build().FindUser(-1)
	fmt.Println(err)

	//service error
	_, err = service.Build().FindUser(1)
	fmt.Println(err)

}
