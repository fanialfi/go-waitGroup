# golang sync.WaitGroup

berbeda dengan `channel`, jika _channel_ digunakan untuk sharing/kirim data antar goroutine,`sync.WaitGroup` digunakan untuk melakukan sinkronasi antar gorouting

kedua file tersebut [withoutWaitGroup.go](withoutWaitGroup.go) dan [withWaitGroup](withWaitGroup.go) besar kemungkinan mengalami race condition, untuk tracking apakah program mengalami **race condition** bisa dengan menambahkan flags _-race_ saat menjalankan / mencompile program
