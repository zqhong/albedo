package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"go/build"
)

/*
使用：

$ cd $GOPATH/src/github.com/zqhong
$ ./installer --project-name=demo
 */
func main() {

	go func() {
		fmt.Println("installing...")
		fmt.Printf("%s", "[")
		consoleStr := "█"
		for i := 0; i != 10; i = i + 1 {
			fmt.Printf("%s", consoleStr)
			time.Sleep(time.Second)
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
	flag.StringVar(&projectName, "project-projectName", "albedo", "project projectName")
	flag.Parse()

	os.Rename("./tmp/"+files[0].Name(), "./"+files[0].Name())
	os.Remove("tmp")
	os.Remove("tmp.zip")
	os.Rename("./"+files[0].Name(), "./"+projectName)

	renameProject("./"+projectName, projectName)
}

func unzipDir(zipFile, dir string) {

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Fatalf("Open zip file failed: %s\n", err.Error())
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
	projectName = strings.TrimSuffix(projectName, "/")
	projectName = strings.TrimPrefix(projectName, "/")

	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}
	pwdDir, _ := os.Getwd()

	fmt.Println("go path: " + goPath)
	fmt.Println("pwd path: " + pwdDir)

	files, _ := ioutil.ReadDir(fileDir)
	for _, file := range files {
		if file.IsDir() {
			renameProject(fileDir+"/"+file.Name(), projectName)
		} else {
			path := fileDir + "/" + file.Name()
			buf, _ := ioutil.ReadFile(path)
			content := string(buf)

			//替换
			newContent := strings.Replace(content, "github.com/zqhong/albedo/", projectName+"/", -1)

			//重新写入
			ioutil.WriteFile(path, []byte(newContent), 0)
		}
	}
}
