# modif-strings
### Middleware ini hanya berfungsi untuk request body = application/json.

### Middleware ini berjalan di framework gin, namun silahkan dimodif agar bisa berjalan di framework golang lainnya.
Middleware ini berguna untuk menghapus spasi / tab di awal dan akhir object request body

- Exception default adalah "password", jika ingin merubahnya tambahkan code berikut sebelum middleware dipanggil
```go
gtrimstrings.Excepts = map[string]bool{
    "password" : true,
    "other except" : true,
    ...
}
```
- Untuk penggunaan lainnya silahkan modif sendiri

Go get :
> go get github.com/ftamhar/gtrimstrings

## Contoh menggunakan postman

Pada contoh ini saya akan menghiraukan "password" dan "satu"
### Input :
![Input](/img/1.jpg)
### Output :
![Output](/img/2.jpg)