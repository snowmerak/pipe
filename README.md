# pipe

pipe는 고에서 함수 파이프라이닝을 구현하는 라이브러리입니다.

## install

```bash
go get github.com/snomwerak/pipe
```

일반적인 방법과 동일하게 go get으로 디펜던시를 추가합니다.

## using

### Link

```go
func Link[T any](funs ...interface{}) func(in ...interface{}) T
```

Link 함수는 함수를 받고 순차적으로 실행시켜주며 `T` 타입에 값을 넣어 반환합니다.

특이한 점은 고랭 컨벤션에 따라 마지막에 `error`가 반환되고 `nil`이 아닐 경우에는 함수 콜을 중단하고 `T` 타입이 `error`를 가지거나 `error`일 경우 에러를 반환하고 그렇지 않다면 제로값을 반환합니다.

주의할 점은 `panic`에 대한 처리가 하나도 되어 있지 않습니다. `Link` 함수를 호출할 때 혹시 모를 사태에 대비하기 위해 `recover`를 처리하는 코드를 작성해주세요.

## example 

### struct

```go
package main

import (
	"fmt"

	"github.com/snowmerak/pipe"
)

func main() {
	p := pipe.Link[struct {
		A   int
		B   int
		Err error
	}](
		func(a int, b int) (int, int, error) {
			if a > b {
				return 0, 0, fmt.Errorf("a > b")
			}
			return a + b, a - b, nil
		},
		func(a int, b int) int {
			return a - b
		},
		func(a int) (int, int) {
			return a * 2, a * 3
		},
	)
	fmt.Println(p(8, 2))
	fmt.Println(p(4, 6))
}
```

```bash
{0 0 a > b}
{24 36 <nil>}
```

### int

```go
package main

import (
	"fmt"

	"github.com/snowmerak/pipe"
)

func main() {
	p := pipe.Link[int](
		func(a int) int {
			return a * 3
		},
		func(a int) int {
			return a + 1
		},
		func(a int) int {
			return a * 2
		},
	)
	fmt.Println(p(8))
}
```

```bash
50
```
