<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://avatars.githubusercontent.com/u/45662503?v=4g" alt="Project logo"></a>
</p>

<h3 align="center">Belajar Kubernetes</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/Herdanis/Belajar-Kubernetes)](https://github.com/Herdanis/Belajar-Kubernetes)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/Herdanis/Belajar-Kubernetes)](https://github.com/Herdanis/Belajar-Kubernetes/pulls)

</div>

---

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [List Command](#command)
- [Authors](#authors)
- [Addon](#addon)
- [Acknowledgements](#acknowledgement)

## 🧐 About <a name = "about"></a>

Ini adalah project untuk belajar bareng dan pengingat untuk diri saya sendiri tentang DevOps Engineer, Di repo ini mempelajari materi tentang Kubernetes dan command nya.

## 🏁 Getting Started <a name = "getting_started"></a>

Dibutuhan pemahaman dasar tentang Docker untuk memulai pelajaran ini, dan sudah menginstall minikube untuk melakukan simulasi kubenetes, saya sendiri menggunakan minikube docker untuk menjalankannya. image yg di gunakan adalah `nginx` dan `herdanis/simple-service`. Jika ingin memodifikasi image bisa ke dalam directori simple-service dan lakukan build docker image.

## 📟 List Command <a name="command"></a>

Command pod

```
kubectl run <NAMA_POD> --image=<NAMA_IMAGE>
kubectl get pod
kubectl describe pod <NAMA_POD>
kubectl delete pod <NAMA_POD>
```

Command label

```
kubectl get pod --show-label
kubectl get pod -l <KEY>
kubectl get pod -l <KEY>=<VALUE>
kubectl get pod -l <KEY>=<VALUE>,<KEY>=<VALUE>
kubectl get pod -l <KEY>!=<VALUE>
kubectl get pod -l !<KEY>
```

Command namespace

```
kubectl get namespace
kubectl get all --namespace <NAMA_NAMESPACE>
kubectl get all -n <NAMA_NAMESPACE>
```

Command replication Controller

```
kubectl get rc
kubectl get replicationcontroller
kubectl delete rc <NAMA_RC> --cascade=false
```

Command exec

```
kubectl exec -i <NAMA_POD> -- <COMMAND>
kubectl exec -it <NAMA_POD> -- sh
## untuk multi container ##
kubectl exec -it <NAMA_POD> -c <NAMA_CONTAINER> -- sh
```

Command Deployment/Statefull/Daemon

```
kubectl set image deployment/<NAMA_DEPLOYMENT> <NAMA_CONTAINER>=<NAMA_IMAGE>
kubectl rollout status deployment <NAMA_DEPLOYMENT>
kubectl rollout restart deployment/<NAMA_DEPLOYMENT>
kubectl rollout undo deployment/<NAMA_DEPLOYMENT>
kubectl rollout history deployment/<NAMA_DEPLOYMENT>
kubectl rollout pause deployment/<NAMA_DEPLOYMENT>
kubectl rollout resume deployment/<NAMA_DEPLOYMENT>
```

## 💉 Addon <a name = "addon"></a>

<h5>- Manage Kubernetes Object</h5>
Menggunakan Declarative Management memungkinkan, untuk kubernetes akan menyimpan config sebelum nya ke dalam annotation object, tujuan nya untuk mempermudah kita untuk melalukan update/rollback

## 🫙 Minikube Docker <a name = "minikube"></a>

Beberapa command untuk memaksimalkan minikube dengan docker

- \*Setup Docker driver sebagai default

```
minikube config set driver docker
```

- Running Minikube

```
minikube start
```

- Delete semua minikube

```
minikube delete
```

- Enable Ingress di Minikube

```
minikube addons enable ingress
```

- Mengexpose Internal Service Minikube

```
minikube servce <NAMA_SERVICE>
```

- \*Membuat tunnel ingress Minikube (perlu menjalankan command ini jika mengu)

```
minikube tunnel
```

\* Command perlu di jalankan jika Minikube ada di Docker

## ✍️ Authors <a name = "authors"></a>

- [@Herdanis](https://github.com/Herdanis) - Idea & Initial work

See also the list of [contributors](https://github.com/Herdanis/Belajar-Kubernetes/graphs/contributors) who participated in this project.

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- https://kubernetes.io/docs/tutorials/
- https://www.youtube.com/playlist?list=PL-CtdCApEFH8XrWyQAyRd6d_CKwxD8Ime
