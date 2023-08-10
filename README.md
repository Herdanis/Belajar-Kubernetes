<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://avatars.githubusercontent.com/u/45662503?v=4g" alt="Project logo"></a>
</p>

<h3 align="center">Belajar Kubernetes</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/Herdanis/Duat)](https://github.com/Herdanis/Duat/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/Herdanis/Duat)](https://github.com/Herdanis/Duat/pulls)

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

Ini adalah project untuk belajar bareng dan pengingat untuk diri saya sendiri tentang DevOps, Di repo ini mempelajari materi tentang Kubernetes dan command nya

## 🏁 Getting Started <a name = "getting_started"></a>

Dibutuhan pemahaman dasar tentang Docker untuk memulai pelajaran ini, dan sudah menginstall minikube untuk melakukan simulasi kubenetes, saya sendiri menggunakan minikube docker untuk menjalankannya

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

## 💉 Addon <a name = "addon"></a>

Manage Kubernetes Object

Menggunakan Declarative Management memungkinkan, untuk kubernetes akan menyimpan config sebelum nya ke dalam annotation object, tujuan nya untuk mempermudah kita untuk melalukan update/rollback

## 🫙 Minikube Docker <a name = "minikube"></a>

Beberapa command untuk memaksimalkan minikube dengan docker

- Mengeset Docker driver sebagai default

```
minikube config set driver docker
```

- Menjalankan Minikube

```
minikube start
```

- Menghapus semua minikube

```
minikube delete
```

- Mengenable Ingress

```
minikube addons enable ingress
```

- Mengexpose Internal Service

```
minikube servce <NAMA_SERVICE>
```

- Membuat tunnel ingress

```
minikube tunnel
```

## ✍️ Authors <a name = "authors"></a>

- [@Herdanis](https://github.com/Herdanis) - Idea & Initial work

See also the list of [contributors](https://github.com/Herdanis/Duat/graphs/contributors) who participated in this project.

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- https://www.postgresql.org/docs/current/app-pgdump.html
- https://stackoverflow.com/a/24158972
- https://dev.mysql.com/doc/refman/8.0/en/mysqldump.html
