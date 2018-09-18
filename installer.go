package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
使用：

$ cd $GOPATH/src/github.com/zqhong
$ ./installer -project-name=demo
*/
func main() {

	go func() {
		fmt.Println("Installing Albedo(A Gin API Skeleton)...")
		consoleStr := "██"
		for i := 0; i != 25; i = i + 1 {
			fmt.Printf("%s", consoleStr)
			time.Sleep(time.Millisecond * 250)
		}
	}()

	url := "https://api.github.com/repos/zqhong/albedo/zipball/master"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/vnd.github.v3+json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	file, _ := os.Create("tmp.zip")
	io.Copy(file, res.Body)

	unzipDir("tmp.zip", "tmp")

	files, _ := ioutil.ReadDir("./tmp")

	var projectName string
	flag.StringVar(&projectName, "project-name", "albedo", "project name")
	flag.Parse()

	os.Rename("./tmp/"+files[0].Name(), "./"+files[0].Name())
	os.Remove("tmp")
	os.Remove("tmp.zip")
	os.Rename("./"+files[0].Name(), "./"+projectName)
	renameProject("./"+projectName, projectName)

	fmt.Printf("%s", "\nInstall Succeeded\n\n")
}

func unzipDir(zipFile, dir string) {

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		os.Exit(1)
	}
	defer r.Close()

	for _, f := range r.File {
		func() {
			path := dir + string(filepath.Separator) + f.Name
			os.MkdirAll(filepath.Dir(path), 0755)
			fDest, err := os.Create(path)
			if err != nil {
				return
			}
			defer fDest.Close()

			fSrc, err := f.Open()
			if err != nil {
				return
			}
			defer fSrc.Close()

			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				return
			}
		}()
	}
}

func renameProject(fileDir string, projectName string) {
	// projectName 示例：demo
	// pkgName 示例：github.com/zqhong
	projectName = strings.TrimSuffix(projectName, "/")
	projectName = strings.TrimPrefix(projectName, "/")

	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}
	pwdPath, _ := os.Getwd()
	pkgName := pwdPath
	pkgName = strings.Replace(pkgName, goPath+"/src/", "", 1)
	pkgName = strings.TrimSuffix(pkgName, "/")
	pkgName = strings.TrimPrefix(pkgName, "/")

	files, _ := ioutil.ReadDir(fileDir)
	for _, file := range files {
		if file.IsDir() {
			renameProject(fileDir+"/"+file.Name(), projectName)
		} else {
			path := fileDir + "/" + file.Name()
			buf, _ := ioutil.ReadFile(path)
			content := string(buf)

			//替换
			newContent := strings.Replace(content, "github.com/zqhong/albedo/", pkgName+"/"+projectName+"/", -1)

			//重新写入
			ioutil.WriteFile(path, []byte(newContent), 0)
		}
	}
}
