package template

import (
	"github.com/gin-gonic/gin"
)

func RouteWebPages(route *gin.RouterGroup) {
	route.GET("/", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", Indexhtml)
	})

	route.GET("/site.webmanifest", func(c *gin.Context) {
		c.Data(200, "application/manifest+json", Sitewebmanifest)
	})


	route.GET("/asset-manifest.json", func(c *gin.Context) {
		c.Data(200, "application/json", Assetmanifestjson)
	})


	route.GET("/android-chrome-512x512.png", func(c *gin.Context) {
		c.Data(200, "image/png", Androidchrome512x512png)
	})


	route.GET("/favicon.ico", func(c *gin.Context) {
		c.Data(200, "image/vnd.microsoft.icon", Faviconico)
	})


	route.GET("/index.html", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", Indexhtml)
	})


	route.GET("/mstile-150x150.png", func(c *gin.Context) {
		c.Data(200, "image/png", Mstile150x150png)
	})


	route.GET("/static/js/2.d5f08679.chunk.js", func(c *gin.Context) {
		c.Data(200, "text/javascript; charset=utf-8", Staticjs2d5f08679chunkjs)
	})


	route.GET("/static/js/2.d5f08679.chunk.js.LICENSE.txt", func(c *gin.Context) {
		c.Data(200, "text/plain; charset=utf-8", Staticjs2d5f08679chunkjsLICENSEtxt)
	})


	route.GET("/android-chrome-192x192.png", func(c *gin.Context) {
		c.Data(200, "image/png", Androidchrome192x192png)
	})


	route.GET("/browserconfig.xml", func(c *gin.Context) {
		c.Data(200, "text/xml; charset=utf-8", Browserconfigxml)
	})


	route.GET("/static/js/main.d1c06468.chunk.js", func(c *gin.Context) {
		c.Data(200, "text/javascript; charset=utf-8", Staticjsmaind1c06468chunkjs)
	})


	route.GET("/static/js/runtime-main.33603a80.js", func(c *gin.Context) {
		c.Data(200, "text/javascript; charset=utf-8", Staticjsruntimemain33603a80js)
	})


	route.GET("/static/js/runtime-main.33603a80.js.map", func(c *gin.Context) {
		c.Data(200, "application/json", Staticjsruntimemain33603a80jsmap)
	})


	route.GET("/apple-touch-icon.png", func(c *gin.Context) {
		c.Data(200, "image/png", Appletouchiconpng)
	})


	route.GET("/favicon-32x32.png", func(c *gin.Context) {
		c.Data(200, "image/png", Favicon32x32png)
	})


	route.GET("/static/js/2.d5f08679.chunk.js.map", func(c *gin.Context) {
		c.Data(200, "application/json", Staticjs2d5f08679chunkjsmap)
	})


	route.GET("/static/js/main.d1c06468.chunk.js.map", func(c *gin.Context) {
		c.Data(200, "application/json", Staticjsmaind1c06468chunkjsmap)
	})


	route.GET("/favicon-16x16.png", func(c *gin.Context) {
		c.Data(200, "image/png", Favicon16x16png)
	})

}