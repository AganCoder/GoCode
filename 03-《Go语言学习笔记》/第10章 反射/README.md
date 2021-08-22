## 反射

反射能让够在运行期探知对象的类型信息和内存结构, 有点像动态语言的特性。

反射是实现元编程的重要手段。

```go
func TypeOf(i interface{}) Type
func ValueOf(i interface{}) Value
```

Type: 表示真实类型（静态类型）
Kind：表示基础结构（底层类型）类别

传入的指针区分基类型和指针类型， 它们不属于同一个类型

