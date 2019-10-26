/***********************************************************
This is an util file, many functions which will be often used
are declared and implemented in this file.

Author: Gavin Mandias
Date: 2019-10-21
************************************************************/

package main

func getKubeResult(kubeGetCommandId string, kube *Kube) (content string, err error, log string) {
	var contentStream []byte
	//kubevtl and docker get commands
	if kubeGetCommandId == "0" {
		contentStream, err = kube.getNodes()
		log = "getnodes"
	} else if kubeGetCommandId == "1" {
		contentStream, err = kube.getRc()
		log = "getrc"
	} else if kubeGetCommandId == "2" {
		contentStream, err = kube.getPods()
		log = "getpods"
	} else if kubeGetCommandId == "3" {
		contentStream, err = kube.getSvc()
		log = "getsvc"
	} else if kubeGetCommandId == "4" {
		contentStream, err = kube.getJobs()
		log = "getjobs"
	} else if kubeGetCommandId == "5" {
		contentStream, err = kube.getClusters()
		log = "getclusters"
	} else if kubeGetCommandId == "d0" {
		contentStream, err = kube.getImages()
		log = "getimages"
	} else if kubeGetCommandId == "d1" {
		contentStream, err = kube.getContainers()
		log = "getcontainers"
	} else if kubeGetCommandId == "yaml" {
		contentStream, err = kube.getYamls()
		log = "getyaml"
	}
	content = string(contentStream)
	return content,err,log
}
