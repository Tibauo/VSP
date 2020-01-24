# VSP
Build
----
```
$ go build -o vsp
```

How to use
----
```
  -api
        a bool
  -file string
        Path to the file with all words (default "/tmp/dico_fr.txt")
  -nb int
        Number of word to include in the password (default 3)
  -port int
        API port (default 8080)
```
API
```
$ ./vsp -file /tmp/dico_fr.txt -api -port 8081
$ curl http://localhost:8080/getpassword
"Supplantai/Vulgarisasses/Gradients"
```

CLI
```
./vsp -nb 4
VSP
Dechagrinez/Corsetasses/Vermillonne/Confisquiez
```

