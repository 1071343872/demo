package main

import (
	"fmt"
	"github.com/1071343872/demo/test" //本工程的包引入
	"github.com/1071343872/foo"  //第三方的go modules包引入
	"github.com/1071343872/test/hello" //第三方非go modules包引入
	_ "github.com/cihub/seelog"
	_ "github.com/codegangsta/cli"
	_ "github.com/davyxu/tabtoy/util"
	"gitlab.taiyouxi.cn/platform/planx/version"  //gitlab私库
)

func main() {
	fmt.Println(foo.Greet("GoModule"))
	test.Hh()

	//logs.Debug("test")
	hello.SayHello()

	fmt.Println(version.GetVersion())
}
