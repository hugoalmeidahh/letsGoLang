# Serving HTML and JSON

Good, I copied the previus project, for continued development.

Seeing what we can build with gin, but I want to demonstrate a little more of its functionality. In particular, sending JSON responses and displaying static files.

Let's look at the JSON responses first.

``` Go
    ...
    api := r.Group("/api")

    api.GET("/starks", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "name": "John snow",
        })
    })
    ...
```

With this code in place, run the server with go run and then hit the new API endpoint:

``` bash
    curl http://localhost:3000/api/starks
```
You should get the response:

``` bash
    {"name": "John snow"}
```
