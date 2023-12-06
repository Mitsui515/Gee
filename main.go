package main

import (
	"fmt"
	"geeweb"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

// func onlyForV2() geeweb.HandlerFunc {
// 	return func(c *geeweb.Context) {
// 		// Start timer
// 		t := time.Now()
// 		// if a server error occured
// 		c.Fail(500, "Internal Server Error")
// 		// calculate resolution time
// 		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
// 	}
// }

func main() {
	// r := geeweb.New()
	// // r.GET("/index", func(c *geeweb.Context) {
	// // 	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	// // })

	// // v1 := r.Group("/v1")
	// // {
	// // 	v1.GET("/", func(c *geeweb.Context) {
	// // 		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	// // 	})
	// // 	v1.GET("/hello", func(c *geeweb.Context) {
	// // 		// expect /hello?name=mitsui
	// // 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// // 	})
	// // }

	// // v2 := r.Group("/v2")
	// // {
	// // 	v2.GET("/hello/:name", func(c *geeweb.Context) {
	// // 		// expect /hello/mitsui
	// // 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	// // 	})
	// // 	v2.POST("/login", func(c *geeweb.Context) {
	// // 		c.JSON(http.StatusOK, geeweb.H{
	// // 			"username": c.PostForm("username"),
	// // 			"password": c.PostForm("password"),
	// // 		})
	// // 	})
	// // }

	// // r.GET("/assets/*filepath", func(c *geeweb.Context) {
	// // 	c.JSON(http.StatusOK, geeweb.H{"filepath": c.Param("filepath")})
	// // })

	// r.Use(geeweb.Logger())
	// // r.GET("/", func(c *geeweb.Context) {
	// // 	c.String(http.StatusOK, "<h1>Hello Gee</h1>")
	// // })

	// // v2 := r.Group("/v2")
	// // v2.Use(onlyForV2())
	// // {
	// // 	v2.GET("/hello/:name", func(c *geeweb.Context) {
	// // 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	// // 	})
	// // }

	// r.SetFuncMap(template.FuncMap{
	// 	"FormatAsDate": FormatAsDate,
	// })
	// r.LoadHTMLGlob("templates/*")
	// r.Static("/assets", "./static")

	// stu1 := &student{Name: "Geektutu", Age: 20}
	// stu2 := &student{Name: "Jack", Age: 22}
	// r.GET("/", func(c *geeweb.Context) {
	// 	c.HTML(http.StatusOK, "css.tmpl", nil)
	// })
	// r.GET("/students", func(c *geeweb.Context) {
	// 	c.HTML(http.StatusOK, "arr.tmpl", geeweb.H{
	// 		"title":  "gee",
	// 		"stuArr": [2]*student{stu1, stu2},
	// 	})
	// })

	// r.GET("/date", func(c *geeweb.Context) {
	// 	c.HTML(http.StatusOK, "custom_func.tmpl", geeweb.H{
	// 		"title": "gee",
	// 		"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
	// 	})
	// })

	// r.Run(":9999")
	r := geeweb.Default()
	r.GET("/", func(c *geeweb.Context) {
		c.String(http.StatusOK, "Hello Mitsui\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *geeweb.Context) {
		names := []string{"mitsui"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
