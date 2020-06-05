## Validator
这是一个用 Golang 写的数据验证组件,主要用于验证 Struct 的属性是否满足特定的规范。

``` 
import (
	"github.com/smartwalle/validator"
	"github.com/smartwalle/errors"
)

type Human struct {
	Name string
	Age  int
}

// 为 Human 需要验证的属性添加方法
func (this Human) NameValidator(n string) error {
	if n == "" {
		// 可以只返回一个 error
		return errors.NewWithCode(1001, "请提供你的名字哦")
	}
	return nil
}

func (this Human) AgeValidator(a int) []error {
	// 也可以返回一个 error slice
	var errList = make([]error, 0, 0)
	if a <= 0 {
		errList = append(errList, errors.NewWithCode(1002, "你确定这是你的年龄？"))
	}

	if a > 100 {
		errList = append(errList, errors.NewWithCode(1003, "你也太长命了吧"))
	}
	return errList
}


// 验证
var h Human
var v = validator.Validate(h)
if !v.OK() {
	fmt.Println("抱歉，验证没有通过")
}

or

var v = validator.Validate(&h)
...

```

如上所示，验证方法命名规则为：属性名+Validator（如: NameValidator），方法有唯一的参数，即对应属性的值，方法需要返回一个满足 error 接口的对象，如果返回 nil，则表示该验证通过。

#### LazyValidate
validator.Validate() 方法将一次验证所有的属性，验证完成之后，才会返回验证的结果。为此可能会多耗费一些时间。

如果想快速一点，可以试试 validator.LazyValidate() 方法，该方法会依次验证 Struct 的属性，遇到第一个错误的时候就会停止验证，并将错误信息返回。
