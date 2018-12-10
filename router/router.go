package router

import (
	"html/template"
	"net/http"

	_ "github.com/1024casts/1024casts/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/1024casts/1024casts/router/middleware"

	"github.com/1024casts/1024casts/handler/api/v1/comment"
	"github.com/1024casts/1024casts/handler/api/v1/course"
	"github.com/1024casts/1024casts/handler/api/v1/order"
	"github.com/1024casts/1024casts/handler/api/v1/plan"
	"github.com/1024casts/1024casts/handler/api/v1/user"
	"github.com/1024casts/1024casts/handler/api/v1/video"
	"github.com/1024casts/1024casts/handler/qiniu"
	"github.com/1024casts/1024casts/handler/sd"

	"github.com/1024casts/1024casts/handler/web"

	"github.com/foolin/gin-template"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	// api for authentication functionalities
	g.POST("/login", user.Login)
	g.POST("/logout", user.Logout)

	// The user handlers, requiring authentication
	u := g.Group("/v1/users")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/token", user.Get)
		//u.GET("/:id", user.Get)
		u.PUT("/:id/status", user.UpdateStatus)
	}

	// The course handlers, requiring authentication
	c := g.Group("/v1/courses")
	c.Use(middleware.AuthMiddleware())
	{
		c.POST("", course.Create)
		c.DELETE("/:id", user.Delete)
		c.PUT("/:id", course.Update)
		c.GET("", course.List)
		c.GET("/:id", course.Get)
		c.GET("/:id/sections", course.Section)
	}

	// The course handlers, requiring authentication
	v := g.Group("/v1/videos")
	v.Use(middleware.AuthMiddleware())
	{
		v.GET("/:course_id", video.List)
		//v.DELETE("/:id", user.Delete)
		//v.PUT("/:id", course.Update)
		//v.GET("/:id", course.Get)
		//v.GET("/:id/sections", course.Section)
	}

	// The comment handlers, requiring authentication
	cmt := g.Group("/v1/comments")
	cmt.Use(middleware.AuthMiddleware())
	{
		cmt.GET("", comment.List)
	}

	// The order handlers, requiring authentication
	o := g.Group("/v1/orders")
	o.Use(middleware.AuthMiddleware())
	{
		o.GET("", order.List)
	}

	// The plan handlers, requiring authentication
	p := g.Group("/v1/plans")
	p.Use(middleware.AuthMiddleware())
	{
		p.GET("", plan.List)
		p.GET("/:id", plan.Get)
		//p.GET("/:alias", plan.Get)
	}

	// The plan handlers, requiring authentication
	q := g.Group("/v1/qiniu")
	q.Use(middleware.AuthMiddleware())
	{
		q.GET("", qiniu.List)
		q.POST("/upload", qiniu.Upload)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}

func InitWebRouter(g *gin.Engine) *gin.Engine {
	router := gin.Default()

	//g.Use(static.Serve("/static", static.LocalFile(viper.GetString("upload.dst"), false)))
	// static file
	//g.Static("/public", viper.GetString("upload.dst"))

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "templates",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		},
		DisableCache: true,
	})

	router.GET("/", web.Index)

	router.Run(":8099")

	return router
}
