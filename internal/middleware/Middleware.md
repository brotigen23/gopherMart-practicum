# Middleware


``` Go
func Logging(next http.Handler) http.Handler{
    return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request){
        next.ServerHTTP()
    })
}
```
