# Tentang
Repositori ini adalah repositori yang digunakan pada sharing session Ngalam Backend x AWS tanggal
21 Desember 2022. Repositori ini berisi tutorial dasar-dasar Docker, mulai dari perintah dasar Docker,
men-Dockerize aplikasi, hingga menjalankan banyak Docker container sekaligus menggunakan `docker-compose`.

# Referensi
Sumber bacaan serta informasi dari repositori ini bisa teman-teman akses di :
1. [https://github.com/girikuncoro/belajar-docker-pemula](https://github.com/girikuncoro/belajar-docker-pemula)
2. [https://www.youtube.com/watch?v=d2oOFasv0B4&list=PL4SGTPmSY0qkxCTe3Gd0wA-bQZChXhsNI](https://www.youtube.com/watch?v=d2oOFasv0B4&list=PL4SGTPmSY0qkxCTe3Gd0wA-bQZChXhsNI)
3. [https://docs.docker.com/reference/](https://docs.docker.com/reference/)

# Sebelum Memulai
Jika teman-teman belum memiliki Docker di komputer teman-teman, silakan mengikuti panduan instalasi Docker sesuai dengan sistem operasi yang teman-teman gunakan di [sini](https://docs.docker.com/get-docker/)

# Perintah-Perintah Dasar Docker
Ada beberapa perintah yang biasanya digunakan ketika kita menggunakan Docker, mulai dari mengunduh image dari Docker Hub, hingga mematikan container yang sedang berjalan. Berikut perintah-perintah dasar Docker yang sering digunakan.
1. `pull` : Mengunduh/pull image dari Docker Hub
2. `run` : Menjalankan image yang telah di-pull. Jika image yang di-`run` tidak ada di komputer, maka image tersebut akan langsung di-`pull` dari Docker Hub
3. `ps` : Melihat container yang berhenti atau sedang berjalan
4. `exec` : Menjalankan command tertentu di dalam container
5. `stop` : Menghentikan container yang sedang berjalan
6. `start` : Memulai/spin container dari Docker image yang telah di-`pull`
7. `rm` : Menghapus container
8. `rmi` : Menghapus image
9. `image` : Melihat Docker image yang telah di-`pull` dari Docker Hub

Contoh penggunaan perintah-perintah di atas adalah sebagai berikut:
1. Melakukan `pull` image `mongo` versi `6.0.3`
```
docker pull mongo:6.0.3
```
2. Menjalankan container dengan nama container `mongodb`, melakukan port binding dari container ke komputer/local, menjalankan dalam mode detach (supaya logging dari container tidak muncul di terminal), dari image `mongo` versi `6.0.3`
```
docker run --name mongodb -p 27017:27017 -d mongo:6.0.3
```
3. Mengakses container `mongodb` menggunakan perintah `mongosh` supaya kita bisa melakukan operasi-operasi yang ada pada `mongodb`, seperti melihat database, mengakses collection, hingga melakukan query
```
docker exec -it mongodb mongosh
```
4. Menghentikan container `mongodb`
```
docker stop mongodb
```
5. Menghapus container `mongodb`
```
docker rm mongodb
```

# Men-Dockerize Aplikasi
Kita bisa mengubah aplikasi kita menjadi Docker image yang bisa di `pull` dan `push` dari Docker Hub. Proses ini biasa dikenal dengan `Dockerize`. Kita membutuhkan `Dockerfile` untuk mengubah aplikasi kita menjadi Docker image. Di repositori ini, proses Dockerize aplikasi ada pada file `Dockerfile` yang terletak di dalam folder `app-backend`. Berikut penjelasan dari `Dockerfile` tersebut.
```Dockerfile
# Menggunakan image golang dengan versi 1.19-alpine sebagai base dari aplikasi kita
FROM golang:1.19-alpine

# Mengubah working directory di dalam container ke folder /app. Perintah ini mirip dengan
# perintah cd di windows. Setelah mengubah working directory menggunakan WORKDIR ini,
# seluruh perintah di bawah WORKDIR akan dijalankan di dalam folder /app
WORKDIR /app

# Meng-copy seluruh file dan folder dari working directory komputer/local kita ke 
# working directory container saat ini (folder /app) pada baris sebelumnya
COPY . .

# Membuat file binary golang dengan nama file binary go-demo-backend
RUN go build -o go-demo-backend

# Mengekspos port 8080 supaya ketika container ini dijalankan/di-spin, container ini bisa diakses
# lewat port 8080
EXPOSE 8080

# Mengeksekusi file binary go-demo-backend yang telah dibuat pada perintah RUN ketika container
# dijalankan/di-spin
CMD ./go-demo-backend
```

`Dockerfile` diatas harus kita compile supaya menjadi sebuah Docker image. Caranya adalah sebagai berikut
```
docker build -o go-docker-demo:v1.0 .
```

Perintah di atas berarti kita membuat docker image dengan nama `go-docker-demo` versi `v1.0`. Perhatikan argumen `.` di akhir perintah. Titik/dot tersebut adalah lokasi di mana `Dockerfile` disimpan.

Untuk men-spin container dari docker image yang telah kita build, kita bisa menggunakan perintah `docker run --name go-docker-demo -p 8080:8080 go-docker-demo:v1.0`

Jika teman-teman penasaran, teman-teman bisa langsung mem-build Docker image dari `Dockerfile` pada repositori ini.

# Mengintegrasikan Container
<em>Bisa ga saya menyambungkan satu container dengan container lain?</em>. Tentu bisa. Menghubungkan atau mengintegrasikan banyak container memungkinkan kita untuk melakukan banyak operasi ketika kita menjalankan aplikasi menggunakan container. Kita perlu menyambungkan container-container tersebut menggunakan `docker network`. Container-container yang terhubung dalam satu `docker network` yang sama bisa berkomunikasi satu sama lain, misalnya container A menjalankan web server dan container B berfungsi sebagai database, maka jika container A dan container B terhubung dalam satu `docker network`, data yang masuk dari container A bisa disimpan dalam container B. Berikut contoh mengintegrasikan dua buah container menggunakan `docker network` :

1. Kita perlu membuat `docker network` terlebih dahulu. Perintah di bawah akan membuat sebuah `docker network` dengan nama network `go-docker-demo`.
```
docker network create go-docker-demo
```
2. Kemudian, kita spin container mongo dan container yang telah kita buat docker imagenya. Kedua container tersebut perlu kita sambungkan dengan `docker network` yang telah kita buat barusan menggunakan flag `--network` sebagai berikut.

Spin container mongo
```
docker run --name mongodb -p 27017:27017 --network go-docker-demo mongo:6.0.3
```

Spin container go-docker-demo
```
docker run --name go-docker-demo -p 8080:8080 --network go-docker-demo go-docker-demo:v1.0
```

Jika tidak ada error, silakan teman-teman buka `localhost:8080/hello` di web browser teman-teman. Di halaman browser, teman-teman akan melihat tulisan `Halo Docker!`. Kemudian, teman-teman akses container `mongodb` menggunakan `docker exec -it mongodb mongosh`, dan di dalam container jalankan dua perintah berikut
```
use go-docker-demo
```
```
db.document.findOne()
```

Teman-teman akan melihat data `{"name": "rocky balboa"}` di terminal container `mongodb`. Artinya, kita sudah berhasil menghubungkan container menggunakan `docker network` yang sama.

# Menjalankan Banyak Container Sekaligus
Ketika jumlah container yang mau kita spin semakin banyak, menggunakan perintah `docker run` akan memakan banyak waktu. Belum lagi jika ada environment variable pada container yang harus kita set dan benar-benar kita cek apakah value pada environment tersebut sudah benar atau belum. Atau jika container yang kita spin ternyata berhasil ter-spin, namun tidak terhubung karena kita lupa men-set `docker network` di tiap container.

Proses menjalankan atau spin banyak container bisa kita otomasi menggunakan `docker compose`. Yang kita butuhkan adalah file `docker compose` dengan ekstensi `yml`. Di repositori ini, file tersebut bernama `docker-compose.yml`. Berikut penjelasan dari file `docker compose` pada repositori ini.

```yml
# Men-specify versi file docker compose. File ini menggunakan versi 3.8
version: "3.8"

# Melakukan spin container mongo dan app-backend
services:
  mongo:
    image: mongo:6.0.2
    container_name: mongodb
    ports:
      - 27017:27017
    networks:
      - docker-demo
  app-backend:
    build: ./app-backend
    ports:
      - 8080:8080
    depends_on: 
      - mongo
    networks:
      - docker-demo

# Membuat docker network dengan nama docker-demo
networks:
  docker-demo:
```

Perhatikan potongan kode dari file `docker compose` di atas berikut ini.

```yml
mongo:
    image: mongo:6.0.2
    container_name: mongodb
    ports:
      - 27017:27017
    networks:
      - docker-demo
```

Dari potongan kode di atas, sebenarnya kita menjalankan perintah `docker run --name mongodb -p 27017:27017 --network docker-demo mongo:6.0.3`. Kita mengotomasi perintah tersebut dalam file `docker compose`. Kemudian perhatikan potongan kode berikut.

```yml
app-backend:
    build: ./app-backend
    ports:
      - 8080:8080
    depends_on: 
      - mongo
    networks:
      - docker-demo
```

Dari potongan kode di atas, sebenarnya kita melakukan build terlebih dahulu dari docker image yang kita `dockerize`, kemudian menjalankan perintah `docker run -p 8080:8080 --network docker-demo go-docker-demo:v1.0`. Terakhir, perhatikan potongan kode berikut.

```yml
networks:
  docker-demo:
```

Potongan kode di atas berarti kita membuat `docker network` dengan nama network `docker-demo` supaya seluruh container yang terhubung dengan network tersebut bisa terhubung satu sama lain.

Untuk menjalankan file `docker compose` di atas, kita tinggal menggunakan perintah `docker compose up` di mana lokasi file `docker compose` tersimpan.

