package main

import (
    "io"
    "os"
    "os/exec"
    "time"
    "strings"

    "github.com/kataras/iris"
    "github.com/kataras/iris/middleware/logger"
    "github.com/kataras/iris/middleware/recover"
)

func main() {
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())

    templates := iris.HTML("./views",".html")
    templates.Reload(true)
    /*templates.AddFunc("displayMessage", func (message string) string {
        return message
    })*/

    app.RegisterView(templates)
    app.StaticWeb("/static","./static")

    //init the kube class
    var kube Kube
    if !kube.load() {
        app.Logger().Infof("Cannot load your kube path config!")
        return
    }

    app.Get("/", func (ctx iris.Context) {
        ctx.ViewData("message","Welcome to KubeManager!")
        if err := ctx.View("index.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
        }
    })

    //kube get controller board, connected to board.htl
    app.Get("/board", func (ctx iris.Context) {
        ctx.ViewData("content","Result Table")
        if err := ctx.View("board.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
        }
    })

    app.Post("get", func (ctx iris.Context) {
        kubeGetCommandId := ctx.FormValue("kubegetcommand")
        content, err, log := getKubeResult(kubeGetCommandId, &kube)

        app.Logger().Infof(log)

        if err != nil {
            ctx.ViewData("content",content)
            if err = ctx.View("board.html"); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString(err.Error())
            }
            return
        }

        contentList := strings.Fields(content)
        ctx.ViewData("contentList",contentList)
        ctx.ViewData("content","Result Table")
        if err = ctx.View("board.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error()) 
        }
    })

    

    app.Post("upload", func (ctx iris.Context) {
        file, _, err := ctx.FormFile("upload")
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.ViewData("message","Upload failed!")
            ctx.View("index.html")
            return
        }

        defer file.Close()

        //filename := info.Filename
        filename := time.Now().Format("20060102150405")

        out, err := os.OpenFile("./static/pics/"+string(filename)+".jpg", os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.ViewData("message","Save failed!")
            ctx.View("index.html")
            return
        }
        defer out.Close()

        io.Copy(out,file)

        cmd := exec.Command("./pictrans.py", "--input", filename)
        _, err = cmd.Output()

        if err != nil {
            ctx.ViewData("message","Transform failed!")
            ctx.View("index.html")
            return
        }
        ctx.ViewData("filename",filename+".jpg")
        ctx.ViewData("rawpath","/static/pics/"+filename+".jpg")
        ctx.ViewData("respath","/static/pics/res"+filename+".jpg")
        ctx.ViewData("message","Upload and transform success!")
        ctx.View("index.html")
    })

    app.Run(iris.Addr(":8080"))
}