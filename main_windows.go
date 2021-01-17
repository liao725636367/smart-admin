package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/chai2010/winsvc"
	"log"
	"os"
	"path/filepath"
	_ "smartapp/initial/common"
	_ "smartapp/initial/plugins"
	_ "smartapp/routers"
)

var (
	serve string="beegoTest"
	appPath string
	flagServiceName = flag.String("service-name", serve, "Set service name")

	flagServiceDesc = flag.String("service-desc", serve+" service", "Set service description")

	flagServiceInstall   = flag.Bool("service-install", false, "Install service")

	flagServiceUninstall = flag.Bool("service-remove", false, "Remove service")

	flagServiceStart     = flag.Bool("service-start", false, "Start service")

	flagServiceStop      = flag.Bool("service-stop", false, "Stop service")
	//查询服务
	flagServiceQuery	=	flag.Bool("service-query", false, "Query service")

)

func init() {

	// change to current dir

	var err error

	if appPath, err = winsvc.GetAppPath(); err != nil {

		log.Fatal(err)

	}

	if err := os.Chdir(filepath.Dir(appPath)); err != nil {

		log.Fatal(err)

	}

}
func main() {
	flag.Parse()

	// install service

	if *flagServiceInstall {
		params:=""
		if err := winsvc.InstallService(appPath, *flagServiceName, *flagServiceDesc,params); err != nil {

			log.Fatalf("installService(%s, %s): %v\n", *flagServiceName, *flagServiceDesc, err)

		}

		fmt.Printf("Done\n")

		return

	}
	//query service
	if *flagServiceQuery {


		if serName, err := winsvc.QueryService(*flagServiceName); err != nil {

			log.Fatalln("removeService:", err)

		}else{
			fmt.Printf(serName+"\n")
		}
		fmt.Printf("Done\n")

		return

	}
	// remove service

	if *flagServiceUninstall {

		if err := winsvc.RemoveService(*flagServiceName); err != nil {

			log.Fatalln("removeService:", err)

		}

		fmt.Printf("Done\n")

		return

	}

	// start service

	if *flagServiceStart {

		if err := winsvc.StartService(*flagServiceName); err != nil {

			log.Fatalln("startService:", err)

		}

		fmt.Printf("Done\n")

		return

	}

	// stop service

	if *flagServiceStop {

		if err := winsvc.StopService(*flagServiceName); err != nil {

			log.Fatalln("stopService:", err)

		}

		fmt.Printf("Done\n")

		return

	}

	// run as service

	if !winsvc.InServiceMode() {

		log.Println("main:", "runService")

		if err := winsvc.RunAsService(*flagServiceName, StartServer, StopServer, false); err != nil {

			log.Fatalf("svc.Run: %v\n", err)

		}

		return

	}

	// run as normal

	StartServer()

}
func StartServer() {
	//log.Println("StartServer, port = 8080")

	beego.Run()


}

func StopServer() {


	log.Println("StopServer")
	panic("服务被停止")

}
