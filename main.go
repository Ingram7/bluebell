package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"
)

func main() {

	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	// 2.初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	// 3.初始化MySQL连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	//4.初始化Redis连接
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init InitTrans failed, err:%v\n", err)
		return
	}

	// 5.注册路由
	r := router.SetupRouter()
	// 6.启动服务
	r.Run(":9009")
	//srv := &http.Server{
	//	Addr:    ":9009",
	//	Handler: r,
	//}

	//go func() {
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()
	//
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//log.Println("Shutdown Server ...")
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown: ", err)
	//}
	//log.Println("Server exiting")

}
