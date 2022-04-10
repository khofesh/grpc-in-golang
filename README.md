# gRPC golang

## gRPC reflection and evans cli

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
