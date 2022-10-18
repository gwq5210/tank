package main

import (
	"github.com/gwq5210/tank/code/core"
	"github.com/gwq5210/tank/code/support"
	_ "gorm.io/driver/mysql"
)

func main() {

	core.APPLICATION = &support.TankApplication{}
	core.APPLICATION.Start()

}
