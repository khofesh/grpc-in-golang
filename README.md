# gRPC golang

## gRPC reflection and evans cli

https://github.com/grpc/grpc-go/tree/master/reflection

https://github.com/ktr0731/evans/issues/150#issuecomment-465394581

```shell
$ curl -L 'https://github.com/ktr0731/evans/releases/download/0.10.4/evans_linux_amd64.tar.gz' | tar xvzf -
$ file evans
evans: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, not stripped
$ mv evans /usr/local/bin/evans
```

the evans

```shell
[fahmad@ryzen calculator]$ evans -p 50051 -r

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


calculator.CalculatorService@127.0.0.1:50051> show pac
command show: unknown target 'pac'

calculator.CalculatorService@127.0.0.1:50051> show package
+-------------------------+
|         PACKAGE         |
+-------------------------+
| calculator              |
| grpc.reflection.v1alpha |
+-------------------------+

calculator.CalculatorService@127.0.0.1:50051> show service
+-------------------+--------------------------+---------------------------------+----------------------------------+
|      SERVICE      |           RPC            |          REQUEST TYPE           |          RESPONSE TYPE           |
+-------------------+--------------------------+---------------------------------+----------------------------------+
| CalculatorService | Sum                      | SumRequest                      | SumResponse                      |
| CalculatorService | PrimeNumberDecomposition | PrimeNumberDecompositionRequest | PrimeNumberDecompositionResponse |
| CalculatorService | ComputeAverage           | ComputeAverageRequest           | ComputeAverageResponse           |
| CalculatorService | FindMaximum              | FindMaximumRequest              | FindMaximumResponse              |
| CalculatorService | SquareRoot               | SquareRootRequest               | SquareRootResponse               |
+-------------------+--------------------------+---------------------------------+----------------------------------+

calculator.CalculatorService@127.0.0.1:50051> show message
+----------------------------------+
|             MESSAGE              |
+----------------------------------+
| ComputeAverageRequest            |
| ComputeAverageResponse           |
| FindMaximumRequest               |
| FindMaximumResponse              |
| PrimeNumberDecompositionRequest  |
| PrimeNumberDecompositionResponse |
| SquareRootRequest                |
| SquareRootResponse               |
| SumRequest                       |
| SumResponse                      |
+----------------------------------+

calculator.CalculatorService@127.0.0.1:50051> desc SumRequest
+-------+------------------------+----------+
| FIELD |          TYPE          | REPEATED |
+-------+------------------------+----------+
| data  | TYPE_MESSAGE (SumData) | false    |
+-------+------------------------+----------+

calculator.CalculatorService@127.0.0.1:50051> show package
+-------------------------+
|         PACKAGE         |
+-------------------------+
| calculator              |
| grpc.reflection.v1alpha |
+-------------------------+

calculator.CalculatorService@127.0.0.1:50051> package calculator

calculator@127.0.0.1:50051> show service
+-------------------+--------------------------+---------------------------------+----------------------------------+
|      SERVICE      |           RPC            |          REQUEST TYPE           |          RESPONSE TYPE           |
+-------------------+--------------------------+---------------------------------+----------------------------------+
| CalculatorService | Sum                      | SumRequest                      | SumResponse                      |
| CalculatorService | PrimeNumberDecomposition | PrimeNumberDecompositionRequest | PrimeNumberDecompositionResponse |
| CalculatorService | ComputeAverage           | ComputeAverageRequest           | ComputeAverageResponse           |
| CalculatorService | FindMaximum              | FindMaximumRequest              | FindMaximumResponse              |
| CalculatorService | SquareRoot               | SquareRootRequest               | SquareRootResponse               |
+-------------------+--------------------------+---------------------------------+----------------------------------+

calculator@127.0.0.1:50051> service CalculatorService

calculator.CalculatorService@127.0.0.1:50051> call Sum
data::first_number (TYPE_INT32) => 12
data::second_number (TYPE_INT32) => 32
{
  "result": 44
}

calculator.CalculatorService@127.0.0.1:50051> call PrimeNumberDecomposition
number (TYPE_INT64) => 2343234234
{
  "primeFactor": "2"
}
{
  "primeFactor": "3"
}
{
  "primeFactor": "11"
}
{
  "primeFactor": "1171"
}
{
  "primeFactor": "30319"
}

calculator.CalculatorService@127.0.0.1:50051> call ComputeAverage
number (TYPE_INT32) => 55
number (TYPE_INT32) => 32
number (TYPE_INT32) => 34
number (TYPE_INT32) => 54
number (TYPE_INT32) => 89
number (TYPE_INT32) => 29
number (TYPE_INT32) =>
{
  "average": 48.833333333333336
}

calculator.CalculatorService@127.0.0.1:50051> call FindMaximum
number (TYPE_INT32) => 8
number (TYPE_INT32) => {
  "maximum": 8
}
number (TYPE_INT32) => 9
number (TYPE_INT32) => {
  "maximum": 9
}
number (TYPE_INT32) => 40
number (TYPE_INT32) => {
  "maximum": 40
}
number (TYPE_INT32) => 38
number (TYPE_INT32) => 100
number (TYPE_INT32) => {
  "maximum": 100
}
number (TYPE_INT32) => 38
number (TYPE_INT32) => 203
number (TYPE_INT32) => {
  "maximum": 203
}
number (TYPE_INT32) =>

calculator.CalculatorService@127.0.0.1:50051> call SquareRoot
number (TYPE_INT32) => 389
{
  "numberRoot": 19.72308292331602
}

calculator.CalculatorService@127.0.0.1:50051> call SquareRoot
number (TYPE_INT32) => -38
command call: rpc error: code = InvalidArgument desc = Received a negative number: -38

calculator.CalculatorService@127.0.0.1:50051>


```

## evans blog

```shell
[fahmad@ryzen blog]$ evans -p 50051 -r

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


blog.BlogService@127.0.0.1:50051> show package
+-------------------------+
|         PACKAGE         |
+-------------------------+
| blog                    |
| grpc.reflection.v1alpha |
+-------------------------+

blog.BlogService@127.0.0.1:50051> show service
+-------------+------------+-------------------+--------------------+
|   SERVICE   |    RPC     |   REQUEST TYPE    |   RESPONSE TYPE    |
+-------------+------------+-------------------+--------------------+
| BlogService | CreateBlog | CreateBlogRequest | CreateBlogResponse |
| BlogService | ReadBlog   | ReadBlogRequest   | ReadBlogResponse   |
| BlogService | UpdateBlog | UpdateBlogRequest | UpdateBlogResponse |
| BlogService | DeleteBlog | DeleteBlogRequest | DeleteBlogResponse |
| BlogService | ListBlog   | ListBlogRequest   | ListBlogResponse   |
+-------------+------------+-------------------+--------------------+

blog.BlogService@127.0.0.1:50051> service BlogService

blog.BlogService@127.0.0.1:50051> show service
+-------------+------------+-------------------+--------------------+
|   SERVICE   |    RPC     |   REQUEST TYPE    |   RESPONSE TYPE    |
+-------------+------------+-------------------+--------------------+
| BlogService | CreateBlog | CreateBlogRequest | CreateBlogResponse |
| BlogService | ReadBlog   | ReadBlogRequest   | ReadBlogResponse   |
| BlogService | UpdateBlog | UpdateBlogRequest | UpdateBlogResponse |
| BlogService | DeleteBlog | DeleteBlogRequest | DeleteBlogResponse |
| BlogService | ListBlog   | ListBlogRequest   | ListBlogResponse   |
+-------------+------------+-------------------+--------------------+

blog.BlogService@127.0.0.1:50051> call CreateBlog
blog::id (TYPE_STRING) =>
blog::author_id (TYPE_STRING) => frog
blog::title (TYPE_STRING) => blog about grpc
blog::content (TYPE_STRING) => created using evans cli
{
  "blog": {
    "authorId": "frog",
    "content": "created using evans cli",
    "id": "6252c793667c98951aa08ca0",
    "title": "blog about grpc"
  }
}

blog.BlogService@127.0.0.1:50051> call ListBlog
{
  "blog": {
    "authorId": "Fahmi",
    "content": "This the content",
    "id": "6252acc6214746b1c240a684",
    "title": "This the content"
  }
}
{
  "blog": {
    "authorId": "Fahmi",
    "content": "This the content",
    "id": "6252b451bc97cc848e0d8b8c",
    "title": "This the content"
  }
}
{
  "blog": {
    "authorId": "Fahmi",
    "content": "This the content",
    "id": "6252b47ebc97cc848e0d8b8d",
    "title": "This the content"
  }
}
{
  "blog": {
    "authorId": "Changed author",
    "content": "This the content, lksdflksoiwer lksjdflskdjf",
    "id": "6252c078bedd458d36c81390",
    "title": "My Second Blog"
  }
}
{
  "blog": {
    "authorId": "frog",
    "content": "created using evans cli",
    "id": "6252c793667c98951aa08ca0",
    "title": "created using evans cli"
  }
}

blog.BlogService@127.0.0.1:50051> call DeleteBlog
blog_id (TYPE_STRING) => 6252acc6214746b1c240a684
{
  "blogId": "6252acc6214746b1c240a684"
}

blog.BlogService@127.0.0.1:50051> call DeleteBlog
blog_id (TYPE_STRING) => 6252acc6214746b1c240a684
command call: rpc error: code = NotFound desc = Cannot find blog in MongoDB: <nil>

blog.BlogService@127.0.0.1:50051> call ListBlog
{
  "blog": {
    "authorId": "Fahmi",
    "content": "This the content",
    "id": "6252b451bc97cc848e0d8b8c",
    "title": "This the content"
  }
}
{
  "blog": {
    "authorId": "Fahmi",
    "content": "This the content",
    "id": "6252b47ebc97cc848e0d8b8d",
    "title": "This the content"
  }
}
{
  "blog": {
    "authorId": "Changed author",
    "content": "This the content, lksdflksoiwer lksjdflskdjf",
    "id": "6252c078bedd458d36c81390",
    "title": "My Second Blog"
  }
}
{
  "blog": {
    "authorId": "frog",
    "content": "created using evans cli",
    "id": "6252c793667c98951aa08ca0",
    "title": "created using evans cli"
  }
}

blog.BlogService@127.0.0.1:50051> call ReadBlog
blog_id (TYPE_STRING) => 6252c793667c98951aa08ca0
{
  "blog": {
    "authorId": "frog",
    "content": "created using evans cli",
    "id": "6252c793667c98951aa08ca0",
    "title": "created using evans cli"
  }
}

blog.BlogService@127.0.0.1:50051> call UpdateBlog
blog::id (TYPE_STRING) => 6252c793667c98951aa08ca0
blog::author_id (TYPE_STRING) => frog 2
blog::title (TYPE_STRING) =>
blog::content (TYPE_STRING) =>
{
  "blog": {
    "authorId": "frog 2",
    "id": "6252c793667c98951aa08ca0"
  }
}

blog.BlogService@127.0.0.1:50051>
```

## references

https://kb.objectrocket.com/mongo-db/how-to-find-a-mongodb-document-by-its-bson-objectid-using-golang-452

https://www.mongodb.com/languages/golang
