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
)

//const confPath string = "./conf/pathconf.json"

type Kube struct {
	SystemPath string `json:"systemPath"`
	KubePath string `json:"kubePath"`
}

func (kube *Kube) getNodes() {
	cmd := exec.Command(kube.KubePath+"kubectl","get","nodes")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func (kube *Kube) getPods() {
	cmd := exec.Command(kube.KubePath+"kubectl","get","pods")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func (kube *Kube) getRc() {
	cmd := exec.Command(kube.KubePath+"kubectl","get","rc")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func (kube *Kube) getImages() {
	cmd := exec.Command(kube.KubePath+"docker","images")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func (kube *Kube) getTest() {
	cmd := exec.Command(kube.SystemPath+"ls")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func (kube *Kube) load(confPath string) bool {
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
