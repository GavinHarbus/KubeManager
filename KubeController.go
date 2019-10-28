package main

import (
    "io"
    "os"
    "io/ioutil"
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
        if err := ctx.View("index.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
        }
    })

    //kube get controller board, connected to board.html
    app.Get("/board", func (ctx iris.Context) {
        ctx.ViewData("content","Result Table")
        if err := ctx.View("board.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
        }
    })

    app.Get("/kubestatus", func (ctx iris.Context) {
        if err := ctx.View("kubestatus.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
        }
    })

    app.Post("getkube", func (ctx iris.Context) {
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

    //docker get controller board, connected to dockerboard.html
    app.Get("/dockerboard", func (ctx iris.Context) {
        ctx.ViewData("content","Result Table")
        if err := ctx.View("dockerboard.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
        }
    })

    app.Get("/dockerstatus", func (ctx iris.Context) {
        if err := ctx.View("dockerstatus.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
        }
    })

    app.Post("getdocker", func (ctx iris.Context) {
        kubeGetCommandId := ctx.FormValue("kubegetcommand")
        content, err, log := getKubeResult(kubeGetCommandId, &kube)

        app.Logger().Infof(log)

        if err != nil {
            ctx.ViewData("content",content)
            if err = ctx.View("dockerboard.html"); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString(err.Error())
            }
            return
        }

        contentList := strings.Fields(content)
        ctx.ViewData("contentList",contentList)
        ctx.ViewData("content","Result Table")
        if err = ctx.View("dockerboard.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error()) 
        }
    })
    
    //kubectl create and delete pods
    app.Get("/podsoperation", func (ctx iris.Context) {
        if err := ctx.View("podsoperations.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
        }
    })

    app.Get("/cdpods", func (ctx iris.Context) {
        content, err, log := getKubeResult("yaml", &kube)
        app.Logger().Infof(log)

        pods, err, log := getKubeResult("2", &kube)
        app.Logger().Infof(log)

        rcs, err, log := getKubeResult("1", &kube)
        app.Logger().Infof(log)

        services, err, log := getKubeResult("3", &kube)
        app.Logger().Infof(log)

        if err != nil {
            if err = ctx.View("cdpods.html"); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString(err.Error())
            }
            return
        }

        contentList := strings.Fields(content)
        podsList := strings.Fields(pods)
        rcsList := strings.Fields(rcs)
        servicesList := strings.Fields(services)

        ctx.ViewData("textyaml","")
        ctx.ViewData("contentList",contentList)
        ctx.ViewData("podsList",podsList)
        ctx.ViewData("rcsList",rcsList)
        ctx.ViewData("servicesList",servicesList)

        if err = ctx.View("cdpods.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error()) 
        }

    })

    app.Post("uploadyaml", func (ctx iris.Context) {

        pods, err, log := getKubeResult("2", &kube)
        app.Logger().Infof(log)

        rcs, err, log := getKubeResult("1", &kube)
        app.Logger().Infof(log)

        services, err, log := getKubeResult("3", &kube)
        app.Logger().Infof(log)

        if err != nil {
            if err = ctx.View("cdpods.html"); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString(err.Error())
            }
            return
        }

        podsList := strings.Fields(pods)
        rcsList := strings.Fields(rcs)
        servicesList := strings.Fields(services)

        ctx.ViewData("textyaml","")
        ctx.ViewData("podsList",podsList)
        ctx.ViewData("rcsList",rcsList)
        ctx.ViewData("servicesList",servicesList)


        if file, info, err := ctx.FormFile("fileyaml"); err == nil {
            defer file.Close()
            filename := info.Filename
            out, _ := os.OpenFile("./yamls/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
            defer out.Close()
            io.Copy(out,file)

            content, _, log := getKubeResult("yaml", &kube)
            app.Logger().Infof(log)

            textyaml, _ := ioutil.ReadFile("./yamls/"+filename)
            ctx.ViewData("textyaml",string(textyaml))
            ctx.ViewData("content",content)

            content, err, log := getKubeResult("yaml", &kube)
            app.Logger().Infof(log)
            contentList := strings.Fields(content)
            ctx.ViewData("contentList",contentList)

            if err = ctx.View("cdpods.html"); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString(err.Error())
                return
            }
        } else if textyaml := ctx.FormValue("textyaml"); textyaml != "" {
            filename := time.Now().Format("20060102150405")
            out, _ := os.OpenFile("./yamls/"+string(filename)+".yaml", os.O_WRONLY|os.O_CREATE, 0666)
            defer out.Close()
            out.WriteString(textyaml)

            content, _, log := getKubeResult("yaml", &kube)
            app.Logger().Infof(log)

            ctx.ViewData("textyaml",string(textyaml))
            ctx.ViewData("content",content)

            content, err, log := getKubeResult("yaml", &kube)
            app.Logger().Infof(log)
            contentList := strings.Fields(content)
            ctx.ViewData("contentList",contentList)

            if err = ctx.View("cdpods.html"); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString(err.Error())
                return
            }
        } else {
            ctx.ViewData("textyaml","Input your yamls!")
            ctx.ViewData("content","Please upload yamls!")

            content, err, log := getKubeResult("yaml", &kube)
            app.Logger().Infof(log)
            contentList := strings.Fields(content)
            ctx.ViewData("contentList",contentList)

            if err = ctx.View("cdpods.html"); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString(err.Error())
                return
            }
        }
    })

    app.Post("createordelete", func (ctx iris.Context) {
        yamlName := ctx.FormValue("yamlname")
        yamlPath := "./yamls/" + yamlName

        kubeCommand := ctx.FormValue("kubecommand")

        if kubeCommand == "0" {
            app.Logger().Infof("kube create pod")
            content, err := kube.create(yamlPath)
            if err != nil {
                ctx.ViewData("status",string(content))
            }
            ctx.ViewData("status",string(content))
        } else {
            app.Logger().Infof("kube delete pod")
            content, err := kube.delete(yamlPath)
            if err != nil {
                ctx.ViewData("status",string(content))
            }
            ctx.ViewData("status",string(content))
        }

        content, err, log := getKubeResult("yaml", &kube)
        app.Logger().Infof(log)

        pods, err, log := getKubeResult("2", &kube)
        app.Logger().Infof(log)

        rcs, err, log := getKubeResult("1", &kube)
        app.Logger().Infof(log)

        services, err, log := getKubeResult("3", &kube)
        app.Logger().Infof(log)

        if err != nil {
            if err = ctx.View("cdpods.html"); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString(err.Error())
            }
            return
        }

        contentList := strings.Fields(content)
        podsList := strings.Fields(pods)
        rcsList := strings.Fields(rcs)
        servicesList := strings.Fields(services)

        ctx.ViewData("textyaml","")
        ctx.ViewData("contentList",contentList)
        ctx.ViewData("podsList",podsList)
        ctx.ViewData("rcsList",rcsList)
        ctx.ViewData("servicesList",servicesList)

        ctx.View("cdpods.html")
    })

    //docker operations
    app.Get("/dockeroperation", func (ctx iris.Context) {
        if err := ctx.View("dockeroperation.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
        }
    })

    app.Get("/cddocker", func (ctx iris.Context) {
        images, err, log := getKubeResult("d0", &kube)
        app.Logger().Infof(log)

        containiers, err, log := getKubeResult("d1", &kube)
        app.Logger().Infof(log)

        if err != nil {
            if err = ctx.View("cddocker.html"); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString(err.Error())
            }
            return
        }

        imagesList := strings.Fields(images)
        containiersList := strings.Fields(containiers)

        ctx.ViewData("imagesList",imagesList)
        ctx.ViewData("containiersList",containiersList)

        if err = ctx.View("cddocker.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error()) 
        }

    })

    app.Post("/dockersearch", func (ctx iris.Context) {
        images, err, log := getKubeResult("d0", &kube)
        app.Logger().Infof(log)

        containiers, err, log := getKubeResult("d1", &kube)
        app.Logger().Infof(log)

        if err != nil {
            if err = ctx.View("cddocker.html"); err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString(err.Error())
            }
            return
        }

        imagesList := strings.Fields(images)
        containiersList := strings.Fields(containiers)

        ctx.ViewData("imagesList",imagesList)
        ctx.ViewData("containiersList",containiersList)

        imageName := ctx.FormValue("imagename")
        if imageName != "" {
            app.Logger().Infof("docker search")
            content, _ := kube.search(imageName)
            contentList := strings.Split(string(content),"\n")

            var imageNameList []string

            for i := 0; i < len(contentList)-1; i++ {
                imageNameList = append(imageNameList,strings.Fields(contentList[i])[0])
            }

            ctx.ViewData("imagesChoose",imageNameList)
        }

        if err = ctx.View("cddocker.html"); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error()) 
        }

    })

    app.Run(iris.Addr(":8080"))
}