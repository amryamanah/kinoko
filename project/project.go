package project

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

const (
	configName = "kinoko.yaml"
)

var folders = map[string]string {
	"contract" : "contracts",
	"binding" : "bindings",
	"migration" : "migrations",
	"test" : "tests",
}

// Project contains name, license and paths to projects.
type Project struct{
	FullPath 	string
	Name 		string
}

func NewProject(projectPath string) {
	//goPath := build.Default.GOPATH
	pPath, pName := filepath.Split(projectPath)
	p := new(Project)
	p.FullPath = pPath
	p.Name = pName

	p.generateFolder()
	p.generateConfig()
}

func (p *Project) generateFolder() {
	for _, v := range folders {
		fPath := path.Join(p.FullPath, p.Name, v)
		log.Println("[START] Creating: " + path.Join(p.FullPath, p.Name, v))
		if err := os.MkdirAll(fPath, os.FileMode(0755));
		err != nil {
			panic(err)
		}
		log.Println("[FINISH] Creating: " + path.Join(p.FullPath, p.Name, v))
	}
}

func (p *Project) generateConfig() {
	const config = `
project: {{ .FullPath }}
license: Apache 2.0

networks:
    dev:
        url: 
        keystore: 

`
	t := template.Must(template.New("config").Parse(config))
	configPath := path.Join(p.FullPath, p.Name, configName)
	log.Println(configPath)

	f, err := os.Create(configPath)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, p)
	if err != nil {
		panic(err)
	}

	f.Close()
}