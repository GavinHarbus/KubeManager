/***********************************************************
This is an Kube class, many functions which will be often used
are declared and implemented in this file. All docker and 
kubernetes' commands are packaged in these functions in order
to be used easily by class KubeController.

Author: Gavin Mandias
Date: 2019-10-20
************************************************************/
package main

import (
	"fmt"
	"encoding/json"
	"os/exec"
	"io/ioutil"
	//"regexp"
	//"strings"
)

const confPath string = "./conf/pathconf.json"

type Kube struct {
	SystemPath string `json:"systemPath"`
	KubePath string `json:"kubePath"`
	DockerPath string `json:"dockerPath"`
}

//kubectl commands
func (kube *Kube) getNodes() (output []byte, err error) {
	cmd := exec.Command(kube.KubePath+"kubectl","get","nodes")
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) getPods() (output []byte, err error) {
	cmd := exec.Command(kube.KubePath+"kubectl","get","pods")
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) getRc() (output []byte, err error) {
	cmd := exec.Command(kube.KubePath+"kubectl","get","rc")
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) getSvc() (output []byte, err error) {
	cmd := exec.Command(kube.KubePath+"kubectl","get","svc")
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) getJobs() (output []byte, err error) {
	cmd := exec.Command(kube.KubePath+"kubectl","get","jobs")
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) getClusters() (output []byte, err error) {
	cmd := exec.Command(kube.KubePath+"kubectl","get","clusters")
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) create(yamlPath string) (output []byte, err error) {
	cmd := exec.Command(kube.KubePath+"kubectl","create","-f",yamlPath)
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) delete(yamlPath string) (output []byte, err error) {
	cmd := exec.Command(kube.KubePath+"kubectl","delete","-f",yamlPath)
	output, err = cmd.CombinedOutput()
	return output, err
}

//docker commands
func (kube *Kube) getImages() (output []byte, err error) {
	cmd := exec.Command(kube.DockerPath+"docker","images")
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) getContainers() (output []byte, err error) {
	cmd := exec.Command(kube.DockerPath+"docker","ps","-a")
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) search(imageName string) (output []byte, err error) {
	cmd := exec.Command(kube.DockerPath+"docker","search",imageName)
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) pull(imageName string) (output []byte, err error) {
	cmd := exec.Command(kube.DockerPath+"docker","pull",imageName)
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) run(imageName string) (output []byte, err error) {
	cmd := exec.Command(kube.DockerPath+"docker","run","-i",imageName)
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) rm(imageName string) (output []byte, err error) {
	cmd := exec.Command(kube.DockerPath+"docker","rm",imageName)
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) rmi(imageName string) (output []byte, err error) {
	cmd := exec.Command(kube.DockerPath+"docker","rmi",imageName)
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) getYamls() (output []byte, err error) {
	cmd := exec.Command(kube.SystemPath+"ls","./yamls")
	output, err = cmd.CombinedOutput()
	return output, err
}

func (kube *Kube) load() bool {
	if file, err := ioutil.ReadFile(confPath); err == nil {
		err = json.Unmarshal(file,kube);
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	return false
}

/*func main() {
	var kube Kube
	kube.load()
	regx, _ := regexp.Compile("\\S{1,} {1}\\S{1,} {1}\\S{1,} {1}\\S{1,} {1}\\S{1,}|\\S{1,} {1}\\S{1,} {1}\\S{1,}|\\S{1,} {1}\\S{1,}|\\S{1,}")
	out,_ := kube.getContainers()
	res := regx.FindAllStringSubmatch(string(out),-1)
	fmt.Println(res)
}*/
