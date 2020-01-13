package runner

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

//路由
type creeperRouter struct {
	routerPath  string
	method      []string
	handlerFunc gin.HandlerFunc
}

var creeperApiEngine *gin.Engine

//存储路由
var creeperRouters []*creeperRouter

func StartWebServer() {
	//启动设置端口
	cfg, err := goconfig.LoadConfigFile("etc/config.ini")
	if err != nil {
		panic(err)
	}
	mode, err := cfg.GetValue("web", "mode")
	if err != nil {
		panic(err)
	}
	gin.SetMode(mode)
	creeperApiEngine = gin.New()
	creeperApiEngine = gin.New()
	//允许使用跨域请求,全局中间件
	creeperApiEngine.Use(cors())
	httpPort, err := cfg.GetValue("web", "port")
	if err != nil {
		panic(err)
	}
	//路由加载
	loadCreeperApiEngineRouter()
	//if mode == "debug" {
	//	//swagger
	//	url := ginSwagger.URL(fmt.Sprintf("http://127.0.0.1:%s/swagger/doc.json", httpPort)) // The url pointing to API definition
	//	creeperApiEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//}
	//绑定模板
	creeperApiEngine.LoadHTMLGlob("web/tmpl/*")
	//静态目录
	creeperApiEngine.Static("/static", "web/static")
	//启动
	err = creeperApiEngine.Run(fmt.Sprintf(":%s", httpPort))
	if err != nil {
		panic(err)
	}
}

//加载已经注册的路由
func loadCreeperApiEngineRouter() {
	for _, router := range creeperRouters {
		//method空就是所有
		if len(router.method) == 0 {
			creeperApiEngine.Any(router.routerPath, router.handlerFunc)
		} else {
			for _, m := range router.method {
				creeperApiEngine.Handle(m, router.routerPath, router.handlerFunc)
			}
		}
	}
}

//给控制器注册路由使用
func RegisterCreeperApiRunner(routerPath string, method []string, handlerFunc gin.HandlerFunc) {
	creeperRouters = append(creeperRouters, &creeperRouter{
		routerPath:  routerPath,
		method:      method,
		handlerFunc: handlerFunc})
	logrus.Info("路由长度：", len(creeperRouters))
}

//跨域
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		c.Next()
	}
}
