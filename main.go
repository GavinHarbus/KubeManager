package main

const confPath string = "./conf/pathconf.json"

func main() {
	var kube Kube
	kube.load(confPath)
	kube.getNodes()
	kube.getTest()
}