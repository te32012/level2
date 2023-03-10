Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
возвращает

<nil>
false

Значение интерфейса равно nil, только если его значение и динамический тип равны nil. 
В приведенном выше примере Foo() возвращает [nil, *os.PathError], и мы сравниваем его с [nil, nil].

Вы можете думать о значении интерфейса nil как о типизированном, и nil без типа не равен nil с типом. 
Если мы конвертируем nil в правильный тип, значения действительно равны.

fmt.Println(err == (*os.PathError)(nil)) // true