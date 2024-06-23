# MOCKet

A pocket api mocking tool implemented using Go.


### Mock API Configuration

**hello.yml**
```yml
name: "hello-world"
port: 3000
endpoints:
  - path: /v1/hello
    method: GET
    status: 200
    delay: 2000
    headers:
      - name: Content-Type
        value: application/json
    body:
      success: true
      data:
        message: "Hello World"
```


### Run
```sh
go run github.com/jabernardo/mocket@latest hello.yml
```
