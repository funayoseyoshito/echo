package main

import (
	"net/http"

	"os"

	"io"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"naem"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/save2", save2)
	e.POST("/users", users)

	//e.POST("/users", saveUser)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))

}

func saveUser() {

}

func updateUser() {

}

func deleteUser() {

}

// e.GET("/users/:id, getUser")
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	//Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/save", save)
// curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+"email:"+email)
}

// e.POST("/save2", save2)
// curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/save
func save2(c echo.Context) error {
	// Get name
	name := c.FormValue("name")

	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you !"+name+"</b>")
}

// e.POST("/users", users)
// curl -F "name=yoshito" -F "email=funayose@gmail.com" http://localhost:1323/users
func users(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.XML(http.StatusOK, u)
	//return c.JSON(http.StatusOK, u)
}
