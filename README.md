# KubeManager
---

#### About

This is a ZJU PaaS homework project. It is aimmed to help users to operate K8S and docker more easily. On another way, it is just a visualization managing web tool for K8S and docker. Although it is fiinshed on a very low level.

#### Requirements

1. Go v1.12.5
2. Go Iris v1.11.1

#### Installation

```shell
git clone https://github.com/GavinHarbus/KubeManager.git
``` 

#### Usage

1. set the path
	
	```shell
	cd ~/KubeManager
	vim conf/pathconf.json
	```
	
	please replace the rawpath with your own path
	
	```json
	{
		"systemPath": "/bin/",
		"kubePath": "/usr/local/bin/",
		"dockerPath": "/usr/local/bin/"
	}

	```
2. start the service

	```shell
	cd ~/KubeManager
	go build KubeController.go Kube.go util.go
	chmod +x Kube
	./Kube
	```

#### Preview

1. Index
	![](https://tva1.sinaimg.cn/large/006y8mN6ly1g8hdvwdnwsj30jj0afdgs.jpg)
2. Kube Status
	![](https://tva1.sinaimg.cn/large/006y8mN6ly1g8hdx188u9j30ib0bbaau.jpg)
3. Docker Status
	
4. Pods Operation
	![](https://tva1.sinaimg.cn/large/006y8mN6ly1g8hdyn3vogj30u10e2abq.jpg)
	![](https://tva1.sinaimg.cn/large/006y8mN6ly1g8hdz3476zj30i208agmf.jpg)
5. Docker Operation
	![](https://tva1.sinaimg.cn/large/006y8mN6ly1g8hdzroixpj30ui0gstc0.jpg)
	![](https://tva1.sinaimg.cn/large/006y8mN6ly1g8he0b1uq9j30t10grjtz.jpg)

#### License

Copywrite 2019 KubeManager  
[**MIT**](https://github.com/GavinHarbus/LICENSE)

