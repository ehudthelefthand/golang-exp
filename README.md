# Interesting Golang Resources

About array and slice

https://blog.golang.org/slices

https://blog.golang.org/slices-intro

https://github.com/golang/go/wiki/SliceTricks

https://www.calhoun.io/how-to-use-slice-capacity-and-length-in-go/

Go DST: https://research.swtch.com/godata

If you want to playaround with golang. This will be a good place. https://play.golang.org

# Install MySQL ผ่าน Docker

```sh
$ docker run --name mysql-local-5.7.28 -p 3307:3306 -v ~/container/mysql-5.7.28/data/:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=password -d mysql:5.7.28
```