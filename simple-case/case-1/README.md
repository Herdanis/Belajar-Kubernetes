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

## 🧐 Test Case <a name = "case"></a>

1. Buatlah sebuah 'replicaset' dengan nama tugas-1-rs menggunakan image 'herdanis/simple-service:4.0' dan di dalam namespace 'aplikasi'
1. Image tersebut membutuh kan enviroment PORT dan SECRET_MESSAGE, env wajib di simpan dalam global enviroment
1. PORT dapat di simpan secara public sedang kan SECRET_MESSAGE hanya pod yg tertentu yg boleh mengakses nya
1. Buat replicaset tersebut dapat di akses oleh pod lain dengan clusterDNS
1. (Optional) Expose replicaset tersebut ke luar dari cluster lalu buka menggunakan browser biasa pastikan secretMessage dapat di lihat
1. Buat namespace baru dengan nama watcher
1. Buat sebuat pod dengan nama watcher menggunakan image 'herdanis/simple-service:checker tanpa mengexpose port apa pun
1. Deklarasikan env CHECKSITE_URL di langsung dalam object pod (tidak perlu configmap/secret)
1. CHECKSITE_URL berisi clusterDNS dari aplikasi yg kita buat sebelum nya (di awali dengan http://)
1. Coba masuk kedalam pod tersebut lalu jalan kan command ./main
1. Jika hasil nya sesuai maka aplikasi berjalan dengan sempurna
