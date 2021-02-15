# Golang Trim Strings
### Middleware ini hanya berfungsi untuk request body = application/json.

### Middleware ini berjalan di framework gin, namun silahkan dimodif agar bisa berjalan di framework golang lainnya.
Middleware ini berguna untuk menghapus spasi / tab di awal dan akhir object request body

- Exception default adalah "password", jika ingin merubahnya tambahkan code berikut sebelum middleware dipanggil
```go
router := gin.Default()

gtrimstrings.Excepts = map[string]bool{
    "password" : true,
    "other except" : true,
    ...
}
router.POST("/login", gtrimstrings.Transform, loginHandler)
```
- Untuk penggunaan di framework lainnya silahkan modif sendiri

Go get :
> go get github.com/ftamhar/gtrimstrings@v0.0.3

## Contoh menggunakan postman

Pada contoh ini saya akan menghiraukan "password" dan "satu"
### Input :
![Input](/img/1.jpg)
### Output :
![Output](/img/2.jpg)
