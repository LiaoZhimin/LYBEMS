package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 静态资源加载，css,js以及资源图片，修改文件自动获取最新
	//r.StaticFS("/statics", http.Dir("D:\GoDemo\gin\statics")) //加载整个静态文件夹
	//r.StaticFS("/statics",)
	r.Static("/ss", "./statics")                  //设置静态文件 html使用/ss/main.css
	r.StaticFile("/favicon.ico", "./favicon.ico") //加载单个文件

	r.LoadHTMLGlob("templates/*")
	//r.LoadHTMLGlob("templates/*/*") //Html模板，有多级结构,/t../t1/index.html

	//r.Use(Cors())  // 跨域设置

	// 位置编码 【楼号00+层号00+间号00】011203 -- 01栋12层03间
	// 设备编码 【类型号A1+购买的yyMMdd+流水码0000】 A12011230035 A1类型20年11月23日购买第35个
	// 设备,位置,责任人
	// 位置,负责人/户主,联系方式,详细位置
	// 设备信息表【编码,负责人,购买人,购买时间,保修期,点检间隔,上次点检,上次维修时间,供应商,价格,供应商电话】
	// 页面功能：【页面显示楼栋各层各设备的状态，设备异常提报，点检记录，维修记录，更换记录】
	//

	r.GET("/", homeHandler)

	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")                  //Post 传入的参数
		password := c.DefaultPostForm("username", "000000") // 可设置默认值
		//返回Json
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	p1 := r.Group("/p1")
	{
		p1.GET("/", p1IndexHandler)
		p1.GET("/Index", p1IndexHandler)
	}

	//r.Run() // listen and serve on 0.0.0.0:8080
	r.Run(":9089")
}

//
//c.Redirect(http.StatusMovedPermanently, "/index") 重定向

func homeHandler(c *gin.Context) {
	//name := c.Query("name") //没传则为空
	name := c.DefaultQuery("name", "游客")      //Get 传入的参数
	role := c.DefaultQuery("role", "teacher") //Get 传入的参数
	//返回 String
	//c.String(200, "Hello, Geektutu "+name+" "+role)
	c.HTML(http.StatusOK, "Index.html", gin.H{"name": name, "role": role})
}

/* p1 router func   */

func p1IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "lzm", "code": "asdqwe"})
}
