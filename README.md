# Gee

Gee is a minimalistic and flexible web framework for the Go programming language, designed with inspiration from Gin. With a focus on simplicity and extensibility, Gee empowers developers to build robust web applications effortlessly. This framework incorporates essential features for routing, middleware support, HTML template rendering, and error handling.

## Directory Structure

```plaintext
Gee/
│
├── geeweb/
│   ├── gee.go
│   ├── gee_test.go
│   ├── context.go
│   ├── router.go
│   ├── router_test.go
│   ├── trie.go
│   ├── logger.go
│   ├── recovery.go
│   └── go.mod
├── static/
│   ├── css
│   │   └── mitsui.css
│   └── file1.txt
├── templates/
│   ├── arr.tmpl
│   ├── css.tmpl
│   └── custom_func.tmpl
├── main.go
├── go.mod
└── README
```

## Getting Started

To integrate Gee into your Go project, follow these steps:

1. Install Gee:
   ```bash
   go get -u github.com/Mitsui515/Gee
   ```

2. Import Gee in your Go code:
   ```go
   import "github.com/Mitsui515/Gee/geeweb"
   ```

3. Create a Gee instance:
   ```go
   r := geeweb.New()
   ```

4. Define routes using various HTTP methods:
   ```go
   r.GET("/", func(c *geeweb.Context) {
       c.HTML(http.StatusOK, "<h1>Hello, Gee!</h1>")
   })

   r.POST("/submit", func(c *geeweb.Context) {
       // Handle POST request
   })
   ```

5. Run the web server:
   ```go
   r.Run(":8080")
   ```

## Core Features

### Routing

Gee provides a straightforward routing mechanism, allowing developers to define static and dynamic routes effortlessly. Dynamic routing supports parameter matching (`:name`) and wildcard matching (`*filepath`), offering flexibility for various use cases.

### Context

The `Context` struct encapsulates HTTP request and response information, simplifying access to parameters, query data, and response rendering. Context instances enable the building of powerful and extensible web applications.

### Middleware

Gee supports middleware to enhance functionality and introduce additional features. Middleware can be easily added and customized, allowing developers to apply various layers of processing to incoming requests.

### Route Group Control

Route grouping allows for organized and controlled route management. Developers can group routes with similar characteristics, making it easier to apply middleware, manage nested groups, and improve overall code organization.

### Static File Serving

Gee enables the serving of static files, such as JavaScript, CSS, and HTML resources. This feature simplifies the handling of client-side assets, enhancing the overall user experience.

### HTML Template Rendering

The framework includes built-in support for rendering HTML templates using the `html/template` package. Developers can leverage this feature to generate dynamic and data-driven HTML content.

### Error Handling

Gee implements a robust error handling mechanism to prevent server crashes. The `Recovery` middleware catches and logs panics, ensuring the stability and reliability of web applications.

## Usage Examples

### Basic Route

```go
r.GET("/", func(c *geeweb.Context) {
    c.HTML(http.StatusOK, "<h1>Hello, Gee!</h1>")
})
```

### Dynamic Route

```go
r.GET("/hello/:name", func(c *geeweb.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello, %s!", name)
})
```

### Middleware Usage

```go
// Logger middleware
r.Use(geeweb.Logger())

// Recovery middleware
r.Use(geeweb.Recovery())
```

### Serving Static Files

Serve static files from the "/static" directory:

```go
r.Static("/static", "/usr/web")
```

### HTML Template Rendering

Render HTML templates:

```go
r.LoadHTMLGlob("templates/*")

r.GET("/render", func(c *geeweb.Context) {
    c.HTML(http.StatusOK, "arr.tmpl", geeweb.H{
        "title": "Gee",
        "arr":   [3]int{1, 2, 3},
    })
})
```