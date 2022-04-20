package route

import (
	"SitesBackend/db"
	"SitesBackend/interpretor"
	"SitesBackend/tools"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"os"
)

func InitRoutes(app *fiber.App) {

	noAuth := app.Group("/noAuth")
	NoAuth(noAuth)

	route := app.Group("/auth", interpretor.MyInterpretor)
	NeedAuth(route)
}

func NeedAuth(router fiber.Router) {
	router.Post("/upload", func(ctx *fiber.Ctx) error {

		file, err := ctx.FormFile("uploadFile")

		if err != nil {
			log.Fatal(err)
			return ctx.JSON(err)
		}

		open, err := file.Open()
		buffer := make([]byte, 1024)
		//store, err := os.OpenFile("F://githubDownload//SitesBackend//"+file.Filename, os.O_CREATE|os.O_RDONLY, 0666)
		home, _ := os.UserHomeDir()
		//fmt.Println(home)
		store, err := os.OpenFile(home+"/opt/fiberStorage/"+file.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		defer store.Close()
		if err != nil {
			log.Fatal(err)
			return ctx.JSON(err)
		}
		for {
			_, err = open.Read(buffer)
			if err == io.EOF {
				break
			}
			store.Write(buffer)
		}

		return ctx.JSON(tools.ResultSet{
			Data:   "upload suc",
			Msg:    "upload suc",
			Status: 200,
		})
	})
}

type User struct {
	Username string
	Password string
}

func NoAuth(router fiber.Router) {
	router.Post("/login", func(ctx *fiber.Ctx) error {
		user := &User{}
		if err := ctx.BodyParser(user); err != nil {
			log.Fatal(err)
			return ctx.JSON(tools.ResultSet{Data: "bodyParser error", Msg: "sth wrong", Status: 500})
		}
		result := db.GetUserByName(user.Username)
		if result.Username == "" {
			return ctx.JSON(tools.ResultSet{Data: "login error", Msg: "sth wrong", Status: 200})
		} else {
			if result.Password != user.Password {
				return ctx.JSON(tools.ResultSet{Data: "login error", Msg: "wrong password", Status: 200})
			}
		}
		token, err := tools.GenerateToken(user.Username, user.Password)
		if err != nil {
			log.Fatal(err)
			return ctx.JSON(tools.ResultSet{Data: "token generate error", Msg: "sth wrong", Status: 500})

		}
		return ctx.JSON(tools.ResultSet{Data: token, Msg: "login suc", Status: 200})
	})

	router.Get("/download", func(ctx *fiber.Ctx) error {
		fileName := ctx.Query("fileName")
		//return ctx.Download("F://githubDownload//SitesBackend" + fileName)
		home, _ := os.UserHomeDir()
		return ctx.Download(home + "/opt/fiberStorage/" + fileName)
	})
}
