# GO语言笔记

## 变量和常量

### 标识符与关键字

#### 标识符

​	在编程语言中标识符就是程序员定义的具有特殊意义的词，比如变量名、常量名、函数名等等。GO语言中标识符由字母和_(下划线)组成，并且只能以字母和 _开头。例子：abc, _, _123,a123.

#### 关键字

​	关键字是指编程语言中预先定义好的具有特殊含义的标识符。关键字和保留字都不建议用作变量名。

​	GO语言中的25个关键字：

```go
break		default	  func			 interfacce	   select 	 case  	    defer		
go		    map		  struct   		 chan	       else	     goto	    package	  
switch	    const	  fallthrough	 if			   range  	 type		continue	
for 	 	import 	  return		 var
```

​	此外，GO语言中还有37个保留字：

```go
Constants:	true		false		iota		nil	
     Type:	int			int8		int16		int32		int64
			uint		uint8		uint16		uint32		uint64		uintptr
			float32		float64		complex64	complex128	
			bool		byte		rune		string		error
Functions:	make		len			cap			new			append		copy	
			close		delete		complex		real		imag
			panic		recover
```

### 变量

#### 变量的来历

​	程序运行过程中的数据都是保存在内存中，我们想要在代码中操作某个数据时就需要去内存上找到这个变量，但是如果我们直接在代码中通过内存地址去操作变量的话，代码的可读性就会非常差而且容易出错，所以我们就利用变量将这个数据的内存地址保存起来，以后通过这个变量就可以找到内存上对应的数据了。

#### 变量类型

​	变量(Variable)的功能就是存储数据。不同的变量保存的数据类型可能会不一样。经过半个多世纪的发展，编程语言已经基本形成了一套固定的类型，常见的变量数据类型有：整形、浮点型、布尔型等。

​	GO语言中的每一个变量都有自己的类型，并且变量必须经过声明才可以开始使用。

#### 变量声明

​	GO语言中的变量需要声明才能使用，同一作用域内不支持重复声明。便且GO语言的变量声明后必须使用。

##### 	标准声明：

```go
var 变量名	变量类型
```

​	变量声明以关键字`var`开头，变量类型放在变量的后面，行尾无需分号。例子：

```go
var name string
var age  int
var	isOk bool	
```

##### 批量声明：

​	每声明一个变量就需要写`var`关键字会比较繁琐，go语言中还支持批量变量声明：

```go
var (
	a string 
	b int
	c bool
	d float32
)
```

##### 变量的初始化：

​	GO语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量都会被初始化成其类型的默认值，例如：整型和浮点型变量的默认值为`0`。字符串变量的默认值为`空字符串`。布尔值变量默认为`false`。切片、函数、指针变量的默认值为`nil`。

​	我们也可以在变量声明的时候为其指定初始值。变量初始化的标准格式如下：

```go
var 变量名 类型 = 表达式
```

​	举个例子：

```go
var name string = "Q1mi"
var age int = 18	
```

​	或者一次初始化多个变量：

```go
var name, age = "Q1mi",18	
```

##### 类型推导：

​	有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。

```go
var name = "Q1mi"
var age  = 18
```

##### 短变量声明：

​	在函数内部，可以使用更简略的`:=`方式声明并初始化变量。

```go
package main

import "fmt"
var m = 100 //定义全部变量m
func main(){
	n := 10
	m := 200 //此处声明局部变量m
	fmt.Println(m,n)
}
```

##### 匿名变量：

​	在使用多重赋值时，如果想忽略某个值，可以使用`匿名变量(anonymous varible)`。匿名变量用一个下划线`_`表示，如：

```go
func foo()(int,string){
    return 10,"Q1mi"
}

func main(){
    x,_ := foo()
    _,y := foo()
    fmp.Println("x=",x)
    fmt.Println("y=",y)
}
```

​	匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。(在`Lua`等编程语言中，匿名变量也被叫做哑元变量)

​	注意事项：

​			1.函数外的每个语句都必须以关键字开始(var、const、func等)

​			2.`:=`不能使用在函数外。

​			3.`_`多用于占位，表示忽略值。

### 常量

​	相对于变量，常量是恒定不变的值，多用于程序运行期间不会改变的那些之。常量的声明和变量的声明非常类似，只是把`var`换成了`const`,在定义常量的时候必须赋值。

```go
const pi = 3.1415
const e  = 2.7182
```

​	声明了`pi`和`e`两个常量之后，在整个程序运行期间他们的值都不会再发生变化。

​	多个常量也可以一起声明：

```go
const (
	const pi = 3.1415
	const e  = 2.7182
)
```

​	const同时声明多个常量，如果省略了值则表示和上面一行的值相同。例如：

```go
const (
	n1 = 100
    n2
    n3
)
```

​	上面例子中，常量`n1`、`n2`、`n3`的值都是100。

#### iota

​	`iota`是go语言的常量计数器，只能在常量的表达式中使用。

​	`iota`在const关键字出现时将被重置为0。const中每新增一行常量声明将使`iota`计数一次(iota可以理解为const语句块中的行索引)。使用iota能简化定义，在定义枚举时很有用。示例：

```go
const (
	n1 = iota //0
    n2		  //1
    n3        //2
    n4		  //3
)
```

​	几种iota示例：

​	使用`_`跳过某些值		

```go
const (
	n1 = iota	//0
	n2			//1
	_
	n4			//3
)
```

​	`iota`中间声明插队

```go
const (
	n1 = iota	//0
	n2 = 100 	//100
    n3			//2
	n4			//3
)
const n5 = iota //0
```

​	定义数量级(这里的`<<`表示左移操作，`1<<10`表示将1的二进制表示向左移10位，也就是由`1`变成了`10000000000`，也就是十进制的1024。同理`2<<2`表示将2的二进制表示向左移2为，也就是由`10`变成了`1000`,也就是10进制的8。)

```GO
const (
	_ =iota
	KB = 1 << (10*iota)
	MB = 1 << (10*iota)
	GB = 1 << (10*iota)
	TB = 1 << (10*iota)
	PB = 1 << (10*iota)
)
```

​	多个`iota`定义在一行

```go
const (
	a,b = iota + 1,iota +2	//1,2
	c,d						//2,3
	e,f						//3,4
)
```

​	

## 基本数据类型

### 整型

​	整型分为两大类：按长度分为：int8,int16,int32,int64	对应的无符号整型：uint8,uint16,uint32,uint64

​	其中，`uint8`就是`byte`类型，`int16`对应C语言中的`short`型，`int64`对应C语言中的`long`型

​	

|  类型  | 描述                                                     |
| :----: | :------------------------------------------------------- |
| uint8  | 无符号8位整型(0~255)                                     |
| uint16 | 无符号16为整型(0~65535)                                  |
| uint32 | 无符号32位整型(0~4294967295)                             |
| uint64 | 无符号64位整型(0~18446744073709551615)                   |
|  int8  | 有符号8位整型(0-128-127)                                 |
| int16  | 有符号16位整型(-32768~32767)                             |
| int32  | 有符号32位整型(-2147483648~2147483647)                   |
| int64  | 有符号64位整型(-9223372036854775808~9223372036854775807) |

#### 特殊类型

|  类型   | 描述                                               |
| :-----: | -------------------------------------------------- |
|  uint   | 32位操作系统上就是uint32，64位操作系统上就是uint64 |
|   int   | 32位操作系统上就是int32，64位操作系统上就是int64   |
| uintptr | 无符号整型，用于存放一个指针                       |

​	**注意**：在使用`int`和`uint`类型时，不能假定他是32位或者64位的类型，而是考虑`int`和`uint`可能在不同平台上的差异。

​	**注意事项**：获取对象长度的内建`len()`函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中，切片或者map的元素数量等都可以用`int`来表示。在设计二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同平台字节长度的影响，不要使用`int`和`uint`。

#### 	数字字面量语法

​	(Number literals syntax)

​	Go1.13版本之后引入了数字字面量语法，这样便于开发者以二进制、八进制或者十六进制浮点数的格式定义数字。例如：`v :=0b00101101`，代表二进制的101101，相当于十进制的45。`v :=0o377`，代表八进制的377，相当于十进制的255。`v := 0x1p-2`，代表十六进制的1除以*2²*,也就是0.25。

​	而且还允许用`_`来分隔数字，比如说：`v := 123_456`表示v的值等于123456。

​	可以借助fmt函数来讲一个整数以不同进制形式展示。

```go
package main	

import "fmt"

func main(){
    var a int = 10			//十进制
    fmt.Printf("%d\n",a)	//10
    fmt.Printf("%b\n",a)	//1010,占位符%b表示二进制
    
    var b int = 077			//八进制，以0开头
    fmt.Printf("%o\n",b)	//77
    
    var c int = 0xff		//十六进制，以0x开头
    fmt.Printf("%x\n")		//ff
    fmt.Printf("%X\n")		//FF
}
```

### 浮点型

​	Go语言支持两种浮点型数：`float32`和`float64`。

​	这两种浮点型数据格式遵循`IEEE 754`标准：

​		`float32`的浮点数的最大范围约位`3.4e38`，可以使用常量定义：`math.MaxFloat32`。

​		`float64`的浮点数的最大范围约为`1.8e308`，可以使用一个常量定义`math.MaxFloat64`。	

​	打印浮点数时，可以使用`fmt`包配合动词`%f`，代码如下：

```go
package	main
import (
	"fmt"
    "math"
)

func main(){
    fmt.Printf("%f\n",math.Pi)
    fmt.Printf("%.2f\n",math.Pi)
}
```

### 复数

​	complex64和complex128

```go
var c1 complex64
c1 = 1 + 2i
var c2 cpmplex128
c2 = 2 + 3i
fmt.Println(c1)
fmt.Println(c2)
```

​	复数由实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部位64位。

### 布尔值

​	Go语言中以`bool`类型进行声明布尔数据类型，布尔类型数据只有`true(真)`和`false(假)`两个值、

​	**注意**：

​			1.布尔类型变量的默认值为`false`

​			2.Go语言中不允许将整型强制转换为布尔类型

​			3.布尔类型无法参与值运算，也无法与其他类型进行转换

### 字符串

​	Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型(int,bool,float32,float64等)一样。Go语言的字符串的内部实现使用`UTF-8`编码。字符串的值为`双引号("")`中的内容，可以在Go的源码中直接添加非ASCII码字符，例如：

```go
s1 := "hello"
s2 := "你好"
```

#### 	字符串转义

​		Go语言的字符串常见转义符为包含回车、换行、单双引号、制表符等，如下：

| 转义符 | 含义                             |
| :----: | -------------------------------- |
|   \r   | 回车符(返回行首)                 |
|   \n   | 换行符(直接跳到下一行的同列位置) |
|   \t   | 制表符                           |
|  \\'   | 单引号                           |
|  \\"   | 双引号                           |
|  \\\   | 反斜杠                           |

​	例子：打印一个Windows平台下的文件路径

```go
package main

import "fmt"

func main(){
    fmt.Println("str := \"c:\\Code\\lesson1\\go,exe\"")
}
```

#### 多行字符串

​	Go语言中要定义多行字符串时，就必须使用`反引号`字符：

```go
s1 := `第一行
	   第二行
	   第三行`
fmt.Println(s1)
```

​	反引号间换行将被作为字符串中的换行，但是所有转义字符均无效，文本会原样输出。

#### 字符串的常用操作

|                 方法                 | 介绍           |
| :----------------------------------: | -------------- |
|               len(str)               | 求长度         |
|            +或fmt.Sprintf            | 拼接字符串     |
|            strings.Split             | 分割           |
|           strings.contains           | 判断是否包含   |
| strings.HasPrefix, strings.HasSuffix | 前缀/后缀判断  |
| strings.Index(), strings.LastIndex() | 子串出现的位置 |
| strings.Join(a[]string, sep string)  | join操作       |

### byte和rune类型

​	组成每个字符串的元素叫做"字符"，可以通过遍历或者单个获取字符串元素获得字符。字符用单引号(\')包裹起来，如：

```go
var a = '中'
var b = 'x'
```

   Go语言中的字符有以下两种：

​			1.`uint8`类型，或者叫byte类型，代表一个`ASCII`码字符。

​			2.`rune`类型，代表一个`UTF-8`字符。

​	当需要处理中文、日文或者其他复合字符时，则需要用到`rune`类型。`rune`类型实际是一个`int32`。

​	Go使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 类型进行默认字符串处理，性能和扩展性都有照顾。

```go
//遍历字符串
func traversalString(){
	s := "非洲小王子"
    for i := 0; i < len(s) ; i++{
        fmt.Printf("%v(%c)",s[i],s[i])
    }
    fmt.Println()
    for _,r := range s {
        fmt.Printf("%v(%c)",r,r)
    }
    fmt.Println()
}
```

输出：

```
233(é)157()158()230(æ)180(´)178(²)229(å)176(°)143()231(ç)142()139()229(å)173(­)144()
38750(非)27954(洲)23567(小)29579(王)23376(子)
```

### 类型转换

​	Go语言中只有强制类型转换，没有隐士类型转换。该语法只能在两个类型之间支持互相转换的时候使用。

​	强制类型转换的基本语法如下：

```go
T(表达式)
```

​	其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值。

​	比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，而变量a和b都是int类型，这个时候就需要将a和b强制类型转换为float64类型。

```go
func sqrtDemo(){
	var a,b = 3,4
    var c int		//int
    c = int(math.Sqrt(float64( a*a + b*b ))) //math.sqrt()接受的参数是float64类型，所以需要强制转换
    fmt.Println(c)
}
```



## 运算符

​	Go语言内置的运算符有：

​			1.算数运算符

​			2.关系运算符

​			3.逻辑预算符

​			4.位运算符

​			5.赋值运算符

### 	算数运算符

​	

| 运算符 | 描述 |
| :----: | ---- |
|   +    | 相加 |
|   -    | 相减 |
|   *    | 相乘 |
|   /    | 相除 |
|   %    | 取余 |

注意：`++`(自增)和`--`(自减)在Go语言中是单独的语句，并不是运算符

### 关系运算符

| 运算符 | 描述                                                      |
| :----: | --------------------------------------------------------- |
|   ==   | 检查两个值是否相等，如果相等返回True否则返回False         |
|   !=   | 检查两个值是否不等，如果不等返回True否则返回False         |
|   >    | 检查左边值是否大于右边值，如果是返回True否则返回False     |
|   >=   | 检查左边值是否大于等于右边值，如果是返回True否则返回False |
|   <    | 检查左边值是否小于右边值，如果是返回True否则返回False1    |
|   <=   | 检查左边值是否小于等于右边值，如果是返回True否则返回False |

### 逻辑运算符

| 运算符 | 描述                                                         |
| :----: | ------------------------------------------------------------ |
|   &&   | 逻辑与(AND)运算符。如果两边的操作数都为True，则结果位True，否则为False |
|  \|\|  | 逻辑或(OR)运算符。如果两边的操作数有一个True，则结果为True，否则为False |
|   !    | 逻辑非(NOT)运算符。如果条件为False，则结果为True，否则为False |

### 位运算符

| 运算符 | 描述                                                         |
| :----: | ------------------------------------------------------------ |
|   &    | 参与运算的两数个对应的二进位相与                                                                                     `两位均为1才为1` |
|   \|   | 参与运算的两数各对应的二进位相或                                                                                     `两位有一个1就为1` |
|   ^    | 参与运算的两数个对应的二进位相异或，当两对应的二进位相异时，结果为1                  `两位不一样则为1` |
|   <<   | 左移n位就是乘以2的n次方                                                       `a << b 就是把a的各二进位全部左移b位，高位丢弃，低位补0` |
|   >>   | 右移n位就是除以2的n次方                                                       `a >> b 就是把a的各二进位全部右移b位` |

### 赋值运算符

| 运算符 | 描述                                           |
| :----: | ---------------------------------------------- |
|   =    | 简单的赋值运算符，将一个表达式的值赋给一个左值 |
|   +=   | 相加后再赋值                                   |
|   -+   | 相减后再赋值                                   |
|   *=   | 相乘后再赋值                                   |
|   /=   | 相除后再赋值                                   |
|   %=   | 取余后再赋值                                   |
|  <<=   | 左移后赋值                                     |
|  >>=   | 右移后赋值                                     |
|   &=   | 按位与后赋值                                   |
|  \|=   | 按位或后赋值                                   |
|   ^=   | 按位异或后赋值                                 |



## 流程控制

​	Go语言中常用的流程控制有`if`和`for`，而`switch`和`goto`主要是为了简化代码、降低重复代码而生的结构，属于扩展类的流程控制。

### if else分支结构

#### 	if条件判断的基本写法

​	  Go语言中`if`条件判断的格式如下：

```go
if 表达式1{
    分支1
}eles if 表达式2{
    分支2
}else{
    分支3
}
```

​	当表达式1的结果为`true`时，执行分支1，否则判断表达式2，如果满足则执行分支2，都不满足时，则执行分支3。if 判断中的`else if`和`else`都是可选的，可以根据实际需要进行选择。

​	Go语言中规定与`if`匹配的左括号`{`也必须与`else`写在同一行，`else`也必须与上一个`if`或`else if`右边的大括号在同一行。

​	举个例子：

```go
func ifDemo01(){
    score := 50
    if score >= 90 {
        fmt.Println("A")
    }else if score >75 {
        fmt.Println("B")
    }else{
        fmt.Println("C")
    }
}
```

#### if条件判断的特殊写法

​	if 条件判断还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值来进行判断，例如：

```go
func ifDemo02(){
    if score := 65;score >= 90{
        fmt.Println("A")
    }else if score >75{
        fmt.Println("B")
    }else{
        fmt.Println("C")
    }
}
```

### for循环结构

​	Go语言中的所有循环类型均可以使用`for`关键字来完成

​	for循环的基本格式如下：

```go
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```

​	条件表达式返回`true`时循环体不停的进行循环，知道条件表达式返回`false`时自动退出循环。

```go
func forDemo(){
    for i := 0;i < 10;i++{
        fmt.Println(i)
    }
}
```

​	for循环的初始语句可以被忽略，但是初始语句后的分号必须要写，例如：

```go
func forDemo02(){
    i := 0 
    for ; i < 10; i++ {
        fmt.Println(i)
    }   
}

```

​	for 循环的初始语句和结束语句都可以省略，例如：

```go
func forDemo03(){
    i := 0
    for i < 10 {
        fmt.Println(i)
    }
}
```

​	这种写法类似于其他变成语言中的`while`，在`while`后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。

#### 无线循环

```go
for {
    循环体语句
}

```

### for range键值循环

​	Go语言中可以使用`for range`遍历数组、切片、字符串、map以及通道(channel)。通过`for range`遍历的返回值具有以下规律：

​				1.数组、切片、字符串返回索引和值

​				2.map返回键和值

​				3.通道(channel)只返回通道内的值

### switch case

​	使用`switch`语句可以方便的对大量的值进行条件判断

```go
func switchDemo01(){
    finger := 3
    switch finger {
    case 1:
        fmt.Println("大拇指")
    case 2:
        fmt.Println("食指")
    case 3:
        fmt.Println("中指")
    case 4:
        fmt.Println("无名指")
    case 5:
        fmt.Println("小拇指")
    default:
        fmt.Println("无效的输入！")
    }
}
```

​	Go语言规定每个`switch`只能有一个`default`分支。

​	一个分支可以有多个值，多个`case`的值中间用英文逗号分隔。

```go
func switchDemo03(){
    switch n := 7;n{
    case 1,3,5,7,9:
    	fmt.Println("奇数")
    case 2,4,6,8:
    	fmt.Println("偶数")   
    default:
        fmt.Println("n")
    }
}
```

​	分支还可以使用表达式，这时候`switch`语句后面不需要再跟判断变量。例如：

```go
func switchDemo04() {
    age := 30
    switch {
    case age > 5:
    	fmt.Println("好好学习吧")
    case age >25 && age < 35:
    	fmt.Println("好好工作吧")
    case age > 60:
        fmt.Println("好好等死吧")
    default:
        fmt.Println("或者真好")
    }
}
```

​	`fallthrough`语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。

```go
func switchDemo05(){
    s := "a"
    switch {
    case s == "a":
        fmt.Println("a")
        fallthrough
    case s == "b":
        fmt.Println("b")
    case s == "c"
        fmt.Println("c")
    default:
        fmt.Println("...")
    }
}
```

​	输出：

```
a
b
```

### goto 跳转到指定标签

​	`goto`语句通过标签进行代码中的无条件跳转。`goto`语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用`goto`能简化一些代码的实现过程。例如双层嵌套的for循环要退出时：

```go
func gotoDemo01(){
    var breakFlag bool
    for i := 0; i < 10; i++ {
        for j := 0 ;j < 10; j++ {
            if j == 2 {
                breakFlag = true	//设置推出标签
                break
            }
            fmt.Println("%v-%v\n",i,j)
        }
        if breakFlag {	//外层for循环判断
            break
        }
    }
}
```

​	使用`goto`语句简化代码：

```go
func gotoDemo02(){
    for i := 0 ; i < 10; i++ {
        for j := 0 ; j < 10 ; j++ {
            if j == 2{
                goto breakTag	//设置退出标签
            }
        }
    }
    return 
    breakTag :	//标签
    	fmt.Println("结束循环")
}
```

### break跳出循环

​	`break`语句可以结束`for`、`switch`和`select`代码块。

​	`break`语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的`for`、`switch`和`select`的代码块上，例如：

```go
func breakDemo01(){
    BREAKDEMO01:
    for i := 0;i < 10;i++{
        for j := 0 ; j < 10; j++ {
            break BREAKDEMO01
        }
        fmt.Println("%v-%v\n",i,j)
    }
    fmt.Println("...")
}
```

### continue继续下次循环

​	`continue`语句可以结束当前循环，开始下一次的循环迭代过程，仅限`for`循环内使用。

​	在`continue`语句后面添加标签时，表示开始标签对应的循环。例如：

```go
func continueDemo01(){
    forloop1:
    for i := 0; i < 5 ; i++ {
        for j := 0 ; j < 5; j++{
            if i == 2 && j == 2 {
                continue forloop1
            }
            fmt.Println("%v-%v\n",i,j)
        }
    }
}
```



## Array(数组)

数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。 基本语法：

```go
var a [3]int    // 定义一个长度为3元素类型为int的数组a
```

### 	数组的定义

```go
var 数组变量名 [元素数量]T
```

​	比如：`var a [5]int`， 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 `[5]int`和`[10]int`是不同的类型。

```go
var a [3]int
var b [4]int
a = b //不可以这样做，因为此时a和b是不同的类
```

​	数组可以通过下标进行访问，下标是从`0`开始，最后一个元素下标是：`len-1`，访问越界（下标在合法范围之外），则触发访问越界，会panic。

### 数组的初始化

#### 方法一

​	初始化数组时可以使用初始化列表来设置数组元素的值。

```go
func main() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}
```

#### 方法二

​	按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：

```go
func main() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}
```

#### 方法三

我们还可以使用指定索引值的方式来初始化数组，例如:

```go
func main() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}
```

### 数组的遍历

遍历数组a有以下两种方法：

```go
func main() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
```

### 多维数组

Go语言是支持多维数组的，我们这里以二维数组为例（数组中又嵌套数组）。

#### 二维数组的定义

```go
func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}
```

#### 二维数组的遍历

```go
func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
```

输出：

```bash
北京	上海	
广州	深圳	
成都	重庆	
```

**注意：** 多维数组**只有第一层**可以使用`...`来让编译器推导数组长度。例如：

```go
//支持的写法
a := [...][2]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}
//不支持多维数组的内层使用...
b := [3][...]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}
```

### 数组是值类型

数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

```go
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
func main() {
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x[100 2 3]
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}
```

**注意：**

1. 数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
2. `[n]*T`表示指针数组，`*[n]T`表示数组指针 。



## 切片

切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。

### 切片的定义

声明切片类型的基本语法如下：

```go
var name []T
```

其中，

- name:表示变量名
- T:表示切片中的元素类型

举个例子：

```go
func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}
```

#### 切片的长度和容量

切片拥有自己的长度和容量，我们可以通过使用内置的`len()`函数求长度，使用内置的`cap()`函数求切片的容量。

#### 切片表达式

切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。它有两种变体：一种指定low和high两个索引界限值的简单的形式，另一种是除了low和high索引界限值外还指定容量的完整的形式。

##### 简单切片表达式

切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。 切片表达式中的`low`和`high`表示一个索引范围（左包含，右不包含），也就是下面代码中从数组a中选出`1<=索引值<4`的元素组成切片s，得到的切片`长度=high-low`，容量等于得到的切片的底层数组的容量。

```go
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
}
```

输出：

```bash
s:[2 3] len(s):2 cap(s):4
```

为了方便起见，可以省略切片表达式中的任何索引。省略了`low`则默认为0；省略了`high`则默认为切片操作数的长度:

```go
a[2:]  // 等同于 a[2:len(a)]
a[:3]  // 等同于 a[0:3]
a[:]   // 等同于 a[0:len(a)]
```

**注意：**

对于数组或字符串，如果`0 <= low <= high <= len(a)`，则索引合法，否则就会索引越界（out of range）。

对切片再执行切片表达式时（切片再切片），`high`的上限边界是切片的容量`cap(a)`，而不是长度。**常量索引**必须是非负的，并且可以用int类型的值表示;对于数组或常量字符串，常量索引也必须在有效范围内。如果`low`和`high`两个指标都是常数，它们必须满足`low <= high`。如果索引在运行时超出范围，就会发生运行时`panic`。

```go
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}
```

输出：

```bash
s:[2 3] len(s):2 cap(s):4
s2:[5] len(s2):1 cap(s2):1
```

##### 完整切片表达式

对于数组，指向数组的指针，或切片a(**注意不能是字符串**)支持完整切片表达式：

```go
a[low : high : max]
```

上面的代码会构造与简单切片表达式`a[low: high]`相同类型、相同长度和元素的切片。另外，它会将得到的结果切片的容量设置为`max-low`。在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0。

```go
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}
```

输出结果：

```bash
t:[2 3] len(t):2 cap(t):4
```

完整切片表达式需要满足的条件是`0 <= low <= high <= max <= cap(a)`，其他条件和简单切片表达式相同。

#### 使用make()函数构造切片

我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的`make()`函数，格式如下：

```bash
make([]T, size, cap)
```

其中：

- T:切片的元素类型
- size:切片中元素的数量
- cap:切片的容量

举个例子：

```go
func main() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}
```

上面代码中`a`的内部存储空间已经分配了10个，但实际上只用了2个。 容量并不会影响当前元素的个数，所以`len(a)`返回2，`cap(a)`则返回该切片的容量。

#### 切片的本质

切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

举个例子，现在有一个数组`a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}`，切片`s1 := a[:5]`，相应示意图如下。![slice_01](https://www.liwenzhou.com/images/Go/slice/slice_01.png)切片`s2 := a[3:6]`，相应示意图如下：![slice_02](https://www.liwenzhou.com/images/Go/slice/slice_02.png)

#### 判断切片是否为空

要检查切片是否为空，请始终使用`len(s) == 0`来判断，而不应该使用`s == nil`来判断。

### 切片不能直接比较

切片之间是不能比较的，我们不能使用`==`操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和`nil`比较。 一个`nil`值的切片并没有底层数组，一个`nil`值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是`nil`，例如下面的示例：

```go
var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
```

所以要判断一个切片是否是空的，要是用`len(s) == 0`来判断，不应该使用`s == nil`来判断。

### 切片的赋值拷贝

下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。

```go
func main() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
}
```

### 切片遍历

切片的遍历方式和数组是一致的，支持索引遍历和`for range`遍历。

```go
func main() {
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
```

### append()方法为切片添加元素

Go语言的内建函数`append()`可以为切片动态添加元素。 可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。

```go
func main(){
	var s []int
	s = append(s, 1)        // [1]
	s = append(s, 2, 3, 4)  // [1 2 3 4]
	s2 := []int{5, 6, 7}  
	s = append(s, s2...)    // [1 2 3 4 5 6 7]
}
```

**注意：**通过var声明的零值切片可以在`append()`函数直接使用，无需初始化。

```go
var s []int
s = append(s, 1, 2, 3)
```

没有必要像下面的代码一样初始化一个切片再传入`append()`函数使用，

```go
s := []int{}  // 没有必要初始化
s = append(s, 1, 2, 3)

var s = make([]int)  // 没有必要初始化
s = append(s, 1, 2, 3)
```

每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在`append()`函数调用时，所以我们通常都需要用原变量接收append函数的返回值。

举个例子：

```go
func main() {
	//append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}
```

输出：

```bash
[0]  len:1  cap:1  ptr:0xc0000a8000
[0 1]  len:2  cap:2  ptr:0xc0000a8040
[0 1 2]  len:3  cap:4  ptr:0xc0000b2020
[0 1 2 3]  len:4  cap:4  ptr:0xc0000b2020
[0 1 2 3 4]  len:5  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5]  len:6  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5 6]  len:7  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5 6 7]  len:8  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5 6 7 8]  len:9  cap:16  ptr:0xc0000b8000
[0 1 2 3 4 5 6 7 8 9]  len:10  cap:16  ptr:0xc0000b8000
```

从上面的结果可以看出：

1. `append()`函数将元素追加到切片的最后并返回该切片。
2. 切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。

append()函数还支持一次性追加多个元素。 例如：

```go
var citySlice []string
// 追加一个元素
citySlice = append(citySlice, "北京")
// 追加多个元素
citySlice = append(citySlice, "上海", "广州", "深圳")
// 追加切片
a := []string{"成都", "重庆"}
citySlice = append(citySlice, a...)
fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]
```

### 切片的扩容策略

可以通过查看`$GOROOT/src/runtime/slice.go`源码，其中扩容相关代码如下：

```go
newcap := old.cap
doublecap := newcap + newcap
if cap > doublecap {
	newcap = cap
} else {
	if old.len < 1024 {
		newcap = doublecap
	} else {
		// Check 0 < newcap to detect overflow
		// and prevent an infinite loop.
		for 0 < newcap && newcap < cap {
			newcap += newcap / 4
		}
		// Set newcap to the requested cap when
		// the newcap calculation overflowed.
		if newcap <= 0 {
			newcap = cap
		}
	}
}
```

从上面的代码可以看出以下内容：

- 首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
- 否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
- 否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
- 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。

需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如`int`和`string`类型的处理方式就不一样。

### 使用copy()函数复制切片

首先我们来看一个问题：

```go
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]
}
```

由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。

Go语言内建的`copy()`函数可以迅速地将一个切片的数据复制到另外一个切片空间中，`copy()`函数的使用格式如下：

```bash
copy(destSlice, srcSlice []T)
```

其中：

- srcSlice: 数据来源切片
- destSlice: 目标切片

举个例子：

```go
func main() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}
```

### 从切片中删除元素

Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。 代码如下：

```go
func main() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}
```

总结一下就是：要从切片a中删除索引为`index`的元素，操作方法是`a = append(a[:index], a[index+1:]...)`

## map

​	map是一种无序的基于`key-value`的数据结构，Go语言中的map是引用类型，必须初始化才能使用。

### map定义

Go语言中 `map`的定义语法如下：

```go
map[KeyType]ValueType
```

其中，

- KeyType:表示键的类型。
- ValueType:表示键对应的值的类型。

map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

```go
make(map[KeyType]ValueType, [cap])
```

其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

### map基本使用

map中的数据都是成对出现的，map的基本使用示例代码如下：

```go
func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)
}
```

输出：

```bash
map[小明:100 张三:90]
100
type of a:map[string]int
```

map也支持在声明的时候填充元素，例如：

```go
func main() {
	userInfo := map[string]string{
		"username": "helloworld",
		"password": "123456",
	}
	fmt.Println(userInfo) //
}
```

### 判断某个键是否存在

Go语言中有个判断map中键是否存在的特殊写法，格式如下:

```go
value, ok := map[key]
```

举个例子：

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100	
	v, ok := scoreMap["张三"]	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
```

### map的遍历

Go语言中使用`for range`遍历map。

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}
```

但我们只想遍历key的时候，可以按下面的写法：

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k := range scoreMap {
		fmt.Println(k)
	}
}
```

**注意：** 遍历map时的元素顺序与添加键值对的顺序无关。

### 使用delete()函数删除键值对

使用`delete()`内建函数从map中删除一组键值对，`delete()`函数的格式如下：

```go
delete(map, key)
```

其中，

- map:表示要删除键值对的map
- key:表示要删除的键值对的键

示例代码如下：

```go
func main(){
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	delete(scoreMap, "小明")//将小明:100从map中删除
	for k,v := range scoreMap{
		fmt.Println(k, v)
	}
}
```

### 按照指定顺序遍历map

```go
func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```

### 元素为map类型的切片

下面的代码演示了切片中的元素为map类型时的操作：

```go
func main() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
```

### 值为切片类型的map

下面的代码演示了map中值为切片类型的操作：

```go
func main() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
```

## 函数

​	Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”。

### 函数定义

Go语言中定义函数使用`func`关键字，具体格式如下：

```go
func 函数名(参数)(返回值){
    函数体
}
```

其中：

- 函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
- 参数：参数由参数变量和参数变量的类型组成，多个参数之间使用`,`分隔。
- 返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用`()`包裹，并用`,`分隔。
- 函数体：实现指定功能的代码块。

我们先来定义一个求两个数之和的函数：

```go
func intSum(x int, y int) int {
	return x + y
}
```

函数的参数和返回值都是可选的，例如我们可以实现一个既不需要参数也没有返回值的函数：

```go
func sayHello() {
	fmt.Println("Hello 沙河")
}
```

### 函数的调用

定义了函数之后，我们可以通过`函数名()`的方式调用函数。 例如我们调用上面定义的两个函数，代码如下：

```go
func main() {
	sayHello()
	ret := intSum(10, 20)
	fmt.Println(ret)
}
```

注意，调用有返回值的函数时，可以不接收其返回值。

### 参数

#### 类型简写

函数的参数中如果相邻变量的类型相同，则可以省略类型，例如：

```go
func intSum(x, y int) int {
	return x + y
}
```

上面的代码中，`intSum`函数有两个参数，这两个参数的类型均为`int`，因此可以省略`x`的类型，因为`y`后面有类型说明，`x`参数也是该类型。

#### 可变参数

可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加`...`来标识。

注意：可变参数通常要作为函数的最后一个参数。

举个例子：

```go
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
```

调用上面的函数：

```go
ret1 := intSum2()
ret2 := intSum2(10)
ret3 := intSum2(10, 20)
ret4 := intSum2(10, 20, 30)
fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60
```

固定参数搭配可变参数使用时，可变参数要放在固定参数的后面，示例代码如下：

```go
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}
```

调用上述函数：

```go
ret5 := intSum3(100)
ret6 := intSum3(100, 10)
ret7 := intSum3(100, 10, 20)
ret8 := intSum3(100, 10, 20, 30)
fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160
```

本质上，函数的可变参数是通过切片来实现的。

### 返回值

Go语言中通过`return`关键字向外输出返回值。

#### 多返回值

Go语言中函数支持多返回值，函数如果有多个返回值时必须用`()`将所有返回值包裹起来。

举个例子：

```go
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
```

#### 返回值命名

函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过`return`关键字返回。

例如：

```go
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}
```

#### 返回值补充

当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。

```go
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	...
}
```

### 变量作用域

#### 全局变量

全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。

```go
package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}
func main() {
	testGlobalVar() //num=10
}
```

#### 局部变量

局部变量又分为两种： 函数内定义的变量无法在该函数外使用，例如下面的示例代码main函数中无法使用testLocalVar函数中定义的变量x：

```go
func testLocalVar() {
	//定义一个函数局部变量x,仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

func main() {
	testLocalVar()
	fmt.Println(x) // 此时无法使用变量x
}
```

如果局部变量和全局变量重名，优先访问局部变量。

```go
package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量
}
func main() {
	testNum() // num=100
}
```

接下来我们来看一下语句块定义的变量，通常我们会在if条件判断、for循环、switch语句上使用这种定义变量的方式。

```go
func testLocalVar2(x, y int) {
	fmt.Println(x, y) //函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 //变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)//此处无法使用变量z
}
```

还有我们之前讲过的for循环语句中定义的变量，也是只在for语句块中生效：

```go
func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}
```

### 函数类型与变量

#### 定义函数类型

我们可以使用`type`关键字来定义一个函数类型，具体格式如下：

```go
type calculation func(int, int) int
```

上面语句定义了一个`calculation`类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。

简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。

```go
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
```

add和sub都能赋值给calculation类型的变量。

```go
var c calculation
c = add
```

#### 函数类型变量

我们可以声明函数类型的变量并且为该变量赋值：

```go
func main() {
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}
```

### 高阶函数

高阶函数分为函数作为参数和函数作为返回值两部分。

#### 函数作为参数

函数可以作为参数：

```go
func add(x, y int) int {
	return x + y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main() {
	ret2 := calc(10, 20, add)
	fmt.Println(ret2) //30
}
```

#### 函数作为返回值

函数也可以作为返回值：

```go
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
```

### 匿名函数和闭包

#### 匿名函数

函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

```go
func(参数)(返回值){
    函数体
}
```

匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

```go
func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
```

匿名函数多用于实现回调函数和闭包。

#### 闭包

闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，`闭包=函数+引用环境`。 首先我们来看一个例子：

```go
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}
```

变量`f`是一个函数并且它引用了其外部作用域中的`x`变量，此时`f`就是一个闭包。 在`f`的生命周期内，变量`x`也一直有效。 闭包进阶示例1：

```go
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder2(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f1 := adder2(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110
}
```

闭包进阶示例2：

```go
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}
```

闭包进阶示例3：

```go
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
```

闭包其实并不复杂，只要牢记`闭包=函数+引用环境`。

### defer语句

Go语言中的`defer`语句会将其后面跟随的语句进行延迟处理。在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行，也就是说，先被`defer`的语句最后被执行，最后被`defer`的语句，最先被执行。

举个例子：

```go
func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
```

输出结果：

```go
start
end
3
2
1
```

由于`defer`语句延迟调用的特性，所以`defer`语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。

#### defer执行时机

在Go语言的函数中`return`语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而`defer`语句执行的时机就在返回值赋值操作后，RET指令执行前。具体如下图所示：![defer执行时机](https://www.liwenzhou.com/images/Go/func/defer.png)

#### defer经典案例

阅读下面的代码，写出最后的打印结果。

```go
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
```

#### defer面试题

```go
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
```

问，上面代码的输出结果是？（提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值）

### 内置函数介绍

|    内置函数    | 介绍                                                         |
| :------------: | :----------------------------------------------------------- |
|     close      | 主要用来关闭channel                                          |
|      len       | 用来求长度，比如string、array、slice、map、channel           |
|      new       | 用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针 |
|      make      | 用来分配内存，主要用来分配引用类型，比如chan、map、slice     |
|     append     | 用来追加元素到数组、slice中                                  |
| panic和recover | 用来做错误处理                                               |

#### panic/recover

Go语言中目前（Go1.12）是没有异常机制，但是使用`panic/recover`模式来处理错误。 `panic`可以在任何地方引发，但`recover`只有在`defer`调用的函数中有效。 首先来看一个例子：

```go
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
```

输出：

```bash
func A
panic: panic in B

goroutine 1 [running]:
main.funcB(...)
        .../code/func/main.go:12
main.main()
        .../code/func/main.go:20 +0x98
```

程序运行期间`funcB`中引发了`panic`导致程序崩溃，异常退出了。这个时候我们就可以通过`recover`将程序恢复回来，继续往后执行。

```go
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
```

**注意：**

1. `recover()`必须搭配`defer`使用。
2. `defer`一定要在可能引发`panic`的语句之前定义。

## 指针

任何程序数据载入内存后，在内存都有他们的地址，这就是指针。而为了保存一个数据在内存中的地址，我们就需要指针变量。

比如，“永远不要高估自己”这句话是我的座右铭，我想把它写入程序中，程序一启动这句话是要加载到内存（假设内存地址0x123456），我在程序中把这段话赋值给变量`A`，把内存地址赋值给变量`B`。这时候变量`B`就是一个指针变量。通过变量`A`和变量`B`都能找到我的座右铭。

Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：`&`（取地址）和`*`（根据地址取值）。



### 指针地址和指针类型

每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用`&`字符放在变量前面对变量进行“取地址”操作。 Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：`*int`、`*int64`、`*string`等。

取变量指针的语法如下：

```go
ptr := &v    // v的类型为T
```

其中：

- v:代表被取地址的变量，类型为`T`
- ptr:用于接收地址的变量，ptr的类型就为`*T`，称做T的指针类型。*代表指针。

举个例子：

```go
func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
}
```

我们来看一下`b := &a`的图示：![取变量地址图示](https://www.liwenzhou.com/images/Go/pointer/ptr.png)

### 指针取值

在对普通变量使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用*操作，也就是指针取值，代码如下。

```go
func main() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
```

输出如下：

```go
type of b:*int
type of c:int
value of c:10
```

**总结：** 取地址操作符`&`和取值操作符`*`是一对互补操作符，`&`取出地址，`*`根据地址取出地址指向的值。

变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：

- 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
- 指针变量的值是指针地址。
- 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。

**指针传值示例：**

```go
func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}
```

### new和make

我们先来看一个例子：

```go
func main() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
```

执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存。

#### new

new是一个内置的函数，它的函数签名如下：

```go
func new(Type) *Type
```

其中，

- Type表示类型，new函数只接受一个参数，这个参数是一个类型
- *Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。举个例子：

```go
func main() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}	
```

本节开始的示例代码中`var a *int`只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。应该按照如下方式使用内置的new函数对a进行初始化之后就可以正常对其赋值了：

```go
func main() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)
}
```

#### make

make也是用于内存分配的，区别于new，它只用于slice、map以及channel的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make函数的函数签名如下：

```go
func make(t Type, size ...IntegerType) Type
```

make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。这个我们在上一章中都有说明，关于channel我们会在后续的章节详细说明。

本节开始的示例中`var b map[string]int`只是声明变量b是一个map类型的变量，需要像下面的示例代码一样使用make函数进行初始化操作之后，才能对其进行键值对赋值：

```go
func main() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
```

#### new与make的区别

1. 二者都是用来做内存分配的。
2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3. 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

## 类型别名和自定义类型

### 自定义类型

在Go语言中有一些基本的数据类型，如`string`、`整型`、`浮点型`、`布尔`等数据类型， Go语言中可以使用`type`关键字来定义自定义类型。

自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。例如：

```go
//将MyInt定义为int类型
type MyInt int
```

通过`type`关键字的定义，`MyInt`就是一种新的类型，它具有`int`的特性。

### 类型别名

类型别名是`Go1.9`版本添加的新功能。

类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。

```go
type TypeAlias = Type
```

我们之前见过的`rune`和`byte`就是类型别名，他们的定义如下：

```go
type byte = uint8
type rune = int32
```

### 类型定义和类型别名的区别

类型别名与类型定义表面上看只有一个等号的差异，我们通过下面的这段代码来理解它们之间的区别。

```go
//类型定义
type NewInt int

//类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt
	
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}
```

结果显示a的类型是`main.NewInt`，表示main包下定义的`NewInt`类型。b的类型是`int`。`MyInt`类型只会在代码中存在，编译完成时并不会有`MyInt`类型。

## 结构体

### 结构体的定义

使用`type`和`struct`关键字来定义结构体，具体代码格式如下：

```go
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
```

其中：

- 类型名：标识自定义结构体的名称，在同一个包内不能重复。
- 字段名：表示结构体字段名。结构体中的字段名必须唯一。
- 字段类型：表示结构体字段的具体类型。

举个例子，我们定义一个`Person`（人）结构体，代码如下：

```go
type person struct {
	name string
	city string
	age  int8
}
```

同样类型的字段也可以写在一行，

```go
type person1 struct {
	name, city string
	age        int8
}
```

这样我们就拥有了一个`person`的自定义类型，它有`name`、`city`、`age`三个字段，分别表示姓名、城市和年龄。这样我们使用这个`person`结构体就能够很方便的在程序中表示和存储人信息了。

语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型

### 结构体实例化

只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。

结构体本身也是一种类型，我们可以像声明内置类型一样使用`var`关键字声明结构体类型。

```go
var 结构体实例 结构体类型
```

#### 基本实例化

举个例子：

```go
type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person
	p1.name = "沙河娜扎"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}
}
```

我们通过`.`来访问结构体的字段（成员变量）,例如`p1.name`和`p1.age`等。

#### 匿名结构体

在定义一些临时数据结构等场景下还可以使用匿名结构体。

```go
package main
     
import (
    "fmt"
)
     
func main() {
    var user struct{Name string; Age int}
    user.Name = "小王子"
    user.Age = 18
    fmt.Printf("%#v\n", user)
}
```

#### 结构体的匿名字段

结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

```go
//Person 结构体Person类型
type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"小王子",
		18,
	}
	fmt.Printf("%#v\n", p1)        //main.Person{string:"北京", int:18}
	fmt.Println(p1.string, p1.int) //北京 18
}
```

**注意：**这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。



#### 创建指针类型结构体

我们还可以通过使用`new`关键字对结构体进行实例化，得到的是结构体的地址。 格式如下：

```go
var p2 = new(person)
fmt.Printf("%T\n", p2)     //*main.person
fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
```

从打印的结果中我们可以看出`p2`是一个结构体指针。

需要注意的是在Go语言中支持对结构体指针直接使用`.`来访问结构体的成员。

```go
var p2 = new(person)
p2.name = "小王子"
p2.age = 28
p2.city = "上海"
fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"小王子", city:"上海", age:28}
```

#### 取结构体的地址实例化

使用`&`对结构体进行取地址操作相当于对该结构体类型进行了一次`new`实例化操作。

```go
p3 := &person{}
fmt.Printf("%T\n", p3)     //*main.person
fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
p3.name = "七米"
p3.age = 30
p3.city = "成都"
fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}
```

`p3.name = "七米"`其实在底层是`(*p3).name = "七米"`，这是Go语言帮我们实现的语法糖。

### 结构体初始化

没有初始化的结构体，其成员变量都是对应其类型的零值。

```go
type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
}
```

#### 使用键值对初始化

使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。

```go
p5 := person{
	name: "小王子",
	city: "北京",
	age:  18,
}
fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}
```

也可以对结构体指针进行键值对初始化，例如：

```go
p6 := &person{
	name: "小王子",
	city: "北京",
	age:  18,
}
fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}
```

当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

```go
p7 := &person{
	city: "北京",
}
fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}
```

#### 使用值的列表初始化

初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：

```go
p8 := &person{
	"沙河娜扎",
	"北京",
	28,
}
fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}
```

使用这种格式初始化时，需要注意：

1. 必须初始化结构体的所有字段。
2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
3. 该方式不能和键值初始化方式混用。

### 嵌套结构体

一个结构体中可以嵌套包含另一个结构体或结构体指针，就像下面的示例代码那样。

```go
//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

func main() {
	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City:     "威海",
		},
	}
	fmt.Printf("user1=%#v\n", user1)//user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
}
```

#### 嵌套匿名字段

上面user结构体中嵌套的`Address`结构体也可以采用匿名字段的方式，例如：

```go
//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address //匿名字段
}

func main() {
	var user2 User
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
	user2.City = "威海"                // 匿名字段可以省略
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
}
```

当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。

#### 嵌套结构体的字段名冲突

嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。

```go
//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	Address
	Email
}

func main() {
	var user3 User
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
}
```

### 结构体内存布局

结构体占用一块连续的内存。

```go
type test struct {
	a int8
	b int8
	c int8
	d int8
}
n := test{
	1, 2, 3, 4,
}
fmt.Printf("n.a %p\n", &n.a)
fmt.Printf("n.b %p\n", &n.b)
fmt.Printf("n.c %p\n", &n.c)
fmt.Printf("n.d %p\n", &n.d)
```

输出：

```bash
n.a 0xc0000a0060
n.b 0xc0000a0061
n.c 0xc0000a0062
n.d 0xc0000a0063
```

【进阶知识点】关于Go语言中的内存对齐推荐阅读:[Go结构体的内存对齐](https://www.liwenzhou.com/posts/Go/struct-memory-layout/)

#### 空结构体

空结构体是不占用空间的。

```go
var v struct{}
fmt.Println(unsafe.Sizeof(v))  // 0
```

### 构造函数

Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个`person`的构造函数。 因为`struct`是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。

```go
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
```

调用构造函数

```go
p9 := newPerson("张三", "沙河", 90)
fmt.Printf("%#v\n", p9) //&main.person{name:"张三", city:"沙河", age:90}
```

### 结构体的“继承”

Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。

```go
//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
```

### 结构体字段的可见性

结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。



## 方法和接收者

Go语言中的`方法（Method）`是一种作用于特定类型变量的函数。这种特定类型变量叫做`接收者（Receiver）`。接收者的概念就类似于其他语言中的`this`或者 `self`。

方法的定义格式如下：

```go
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

其中，

- 接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是`self`、`this`之类的命名。例如，`Person`类型的接收者变量应该命名为 `p`，`Connector`类型的接收者变量应该命名为`c`等。
- 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
- 方法名、参数列表、返回参数：具体格式与函数定义相同。

举个例子：

```go
//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
}
```

方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。

### 指针类型的接收者

指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的`this`或者`self`。 例如我们为`Person`添加一个`SetAge`方法，来修改实例变量的年龄。

```go
// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}
```

调用该方法：

```go
func main() {
	p1 := NewPerson("小王子", 25)
	fmt.Println(p1.age) // 25
	p1.SetAge(30)
	fmt.Println(p1.age) // 30
}
```

### 值类型的接收者

当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。

```go
// SetAge2 设置p的年龄
// 使用值接收者
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
	fmt.Println(p1.age) // 25
	p1.SetAge2(30) // (*p1).SetAge2(30)
	fmt.Println(p1.age) // 25
}
```

### 什么时候应该使用指针类型接收者

1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

### 任意类型添加方法

在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的`int`类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

```go
//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}
func main() {
	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}
```

**注意事项：** 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。





## 包与依赖管理

**本章学习目标**

- 掌握包的定义和使用
- 掌握init初始化函数的使用
- 掌握依赖管理工具go module的使用

在工程化的Go语言开发项目中，Go语言的源码复用是建立在包（package）基础之上的。本文介绍了Go语言中如何定义包、如何导出包的内容及如何引入其他包。同时也将介绍如何在项目中使用go module管理依赖。

### 包（package）

#### 包介绍

Go语言中支持模块化的开发理念，在Go语言中使用`包（package）`来支持代码模块化和代码复用。一个包是由一个或多个Go源码文件（.go结尾的文件）组成，是一种高级的代码复用方案，Go语言为我们提供了很多内置包，如`fmt`、`os`、`io`等。

例如，在之前的章节中我们频繁使用了`fmt`这个内置包。

```go
package main

import "fmt"

func main(){
  fmt.Println("Hello world!")
}
```

上面短短的几行代码就涉及到了如何定义包以及如何引入其它包两个内容，接下来我们依次介绍一下它们。

#### 定义包

我们可以根据自己的需要创建自定义包。一个包可以简单理解为一个存放`.go`文件的文件夹。该文件夹下面的所有`.go`文件都要在非注释的第一行添加如下声明，声明该文件归属的包。

```go
package packagename
```

其中：

- package：声明包的关键字
- packagename：包名，可以不与文件夹的名称一致，不能包含 `-` 符号，最好与其实现的功能相对应。

另外需要注意一个文件夹下面直接包含的文件只能归属一个包，同一个包的文件不能在多个文件夹下。包名为`main`的包是应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含`main`包的源代码则不会得到可执行文件。

#### 标识符可见性

在同一个包内部声明的标识符都位于同一个命名空间下，在不同的包内部声明的标识符就属于不同的命名空间。想要在包的外部使用包内部的标识符就需要添加包名前缀，例如`fmt.Println("Hello world!")`，就是指调用`fmt`包中的`Println`函数。

如果想让一个包中的标识符（如变量、常量、类型、函数等）能被外部的包使用，那么标识符必须是对外可见的（public）。在Go语言中是通过标识符的首字母大/小写来控制标识符的对外可见（public）/不可见（private）的。在一个包内部只有首字母大写的标识符才是对外可见的。

例如我们定义一个名为`demo`的包，在其中定义了若干标识符。在另外一个包中并不是所有的标识符都能通过`demo.`前缀访问到，因为只有那些首字母是大写的标识符才是对外可见的。

```go
package demo

import "fmt"

// 包级别标识符的可见性

// num 定义一个全局整型变量
// 首字母小写，对外不可见(只能在当前包内使用)
var num = 100

// Mode 定义一个常量
// 首字母大写，对外可见(可在其它包中使用)
const Mode = 1

// person 定义一个代表人的结构体
// 首字母小写，对外不可见(只能在当前包内使用)
type person struct {
	name string
	Age  int
}

// Add 返回两个整数和的函数
// 首字母大写，对外可见(可在其它包中使用)
func Add(x, y int) int {
	return x + y
}

// sayHi 打招呼的函数
// 首字母小写，对外不可见(只能在当前包内使用)
func sayHi() {
	var myName = "七米" // 函数局部变量，只能在当前函数内使用
	fmt.Println(myName)
}
```

同样的规则也适用于结构体，结构体中可导出字段的字段名称必须首字母大写。

```go
type Student struct {
	Name  string // 可在包外访问的方法
	class string // 仅限包内访问的字段
}
```

#### 包的引入

要在当前包中使用另外一个包的内容就需要使用`import`关键字引入这个包，并且import语句通常放在文件的开头，`package`声明语句的下方。完整的引入声明语句格式如下:

```go
import importname "path/to/package"
```

其中：

- importname：引入的包名，通常都省略。默认值为引入包的包名。
- path/to/package：引入包的路径名称，必须使用双引号包裹起来。
- Go语言中禁止循环导入包。

一个Go源码文件中可以同时引入多个包，例如：

```go
import "fmt"
import "net/http"
import "os"
```

当然可以使用批量引入的方式。

```go
import (
    "fmt"
  	"net/http"
    "os"
)
```

当引入的多个包中存在相同的包名或者想自行为某个引入的包设置一个新包名时，都需要通过`importname`指定一个在当前文件中使用的新包名。例如，在引入`fmt`包时为其指定一个新包名`f`。

```go
import f "fmt"
```

这样在当前这个文件中就可以通过使用`f`来调用`fmt`包中的函数了。

```go
f.Println("Hello world!")
```

如果引入一个包的时候为其设置了一个特殊`_`作为包名，那么这个包的引入方式就称为匿名引入。一个包被匿名引入的目的主要是为了加载这个包，从而使得这个包中的资源得以初始化。 被匿名引入的包中的`init`函数将被执行并且仅执行一遍。

```go
import _ "github.com/go-sql-driver/mysql"
```

匿名引入的包与其他方式导入的包一样都会被编译到可执行文件中。

需要注意的是，Go语言中不允许引入包却不在代码中使用这个包的内容，如果引入了未使用的包则会触发编译错误。

#### init初始化函数

在每一个Go源文件中，都可以定义任意个如下格式的特殊函数：

```go
func init(){
  // ...
}
```

这种特殊的函数不接收任何参数也没有任何返回值，我们也不能在代码中主动调用它。当程序启动的时候，init函数会按照它们声明的顺序自动执行。

一个包的初始化过程是按照代码中引入的顺序来进行的，所有在该包中声明的`init`函数都将被串行调用并且仅调用执行一次。每一个包初始化的时候都是先执行依赖的包中声明的`init`函数再执行当前包中声明的`init`函数。确保在程序的`main`函数开始执行时所有的依赖包都已初始化完成。![包初始化函数执行顺序示意图](https://www.liwenzhou.com/images/Go/package/package01.png)

每一个包的初始化是先从初始化包级别变量开始的。例如从下面的示例中我们就可以看出包级别变量的初始化会先于`init`初始化函数。

```go
package main

import "fmt"

var x int8 = 10

const pi = 3.14

func init() {
  fmt.Println("x:", x)
  fmt.Println("pi:", pi)
  sayHi()
}

func sayHi() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("你好，世界！")
}
```

输出结果：

```bash
x: 10
pi: 3.14
Hello World!
你好，世界！
```

在上面的代码中，我们了解了Go语言中包的定义及包的初始化过程，这让我们能够在开发时按照自己的需要定义包。同时我们还学到了如何在我们的代码中引入其它的包，不过在本小节的所有示例中我们都是引入Go内置的包。现代编程语言大多都允许开发者对外发布包/库，也支持开发者在自己的代码中引入第三方库。这样的设计能够让广大开发者一起参与到语言的生态环境建设当中，把生态建设的更加完善。

### go module

在Go语言的早期版本中，我们编写Go项目代码时所依赖的所有第三方包都需要保存在GOPATH这个目录下面。这样的依赖管理方式存在一个致命的缺陷，那就是不支持版本管理，同一个依赖包只能存在一个版本的代码。可是我们本地的多个项目完全可能分别依赖同一个第三方包的不同版本。

#### go module介绍

Go module 是 Go1.11 版本发布的依赖管理方案，从 Go1.14 版本开始推荐在生产环境使用，于Go1.16版本默认开启。Go module 提供了以下命令供我们使用：

go module相关命令

|      命令       |                   介绍                   |
| :-------------: | :--------------------------------------: |
|   go mod init   |      初始化项目依赖，生成go.mod文件      |
| go mod download |          根据go.mod文件下载依赖          |
|   go mod tidy   | 比对项目文件中引入的依赖与go.mod进行比对 |
|  go mod graph   |              输出依赖关系图              |
|   go mod edit   |              编辑go.mod文件              |
|  go mod vendor  |     将项目的所有依赖导出至vendor目录     |
|  go mod verify  |        检验一个依赖包是否被篡改过        |
|   go mod why    |          解释为什么需要某个依赖          |

Go语言在 go module 的过渡阶段提供了 `GO111MODULE` 这个环境变量来作为是否启用 go module 功能的开关，考虑到 Go1.16 之后 go module 已经默认开启，所以本书不再介绍该配置，对于刚接触Go语言的读者而言完全没有必要了解这个历史包袱。

**GOPROXY**

这个环境变量主要是用于设置 Go 模块代理（Go module proxy），其作用是用于使 Go 在后续拉取模块版本时能够脱离传统的 VCS 方式，直接通过镜像站点来快速拉取。

GOPROXY 的默认值是：`https://proxy.golang.org,direct`，由于某些原因国内无法正常访问该地址，所以我们通常需要配置一个可访问的地址。目前社区使用比较多的有两个`https://goproxy.cn`和`https://goproxy.io`，当然如果你的公司有提供GOPROXY地址那么就直接使用。设置GOPAROXY的命令如下：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

GOPROXY 允许设置多个代理地址，多个地址之间需使用英文逗号 “,” 分隔。最后的 “direct” 是一个特殊指示符，用于指示 Go 回源到源地址去抓取（比如 GitHub 等）。当配置有多个代理地址时，如果第一个代理地址返回 404 或 410 错误时，Go 会自动尝试下一个代理地址，当遇见 “direct” 时触发回源，也就是回到源地址去抓取。

**GOPRIVATE**

设置了GOPROXY 之后，go 命令就会从配置的代理地址拉取和校验依赖包。当我们在项目中引入了非公开的包（公司内部git仓库或 github 私有仓库等），此时便无法正常从代理拉取到这些非公开的依赖包，这个时候就需要配置 GOPRIVATE 环境变量。GOPRIVATE用来告诉 go 命令哪些仓库属于私有仓库，不必通过代理服务器拉取和校验。

GOPRIVATE 的值也可以设置多个，多个地址之间使用英文逗号 “,” 分隔。我们通常会把自己公司内部的代码仓库设置到 GOPRIVATE 中，例如：

```bash
$ go env -w GOPRIVATE="git.mycompany.com"
```

这样在拉取以`git.mycompany.com`为路径前缀的依赖包时就能正常拉取了。

此外，如果公司内部自建了 GOPROXY 服务，那么我们可以通过设置 `GONOPROXY=none`，允许通内部代理拉取私有仓库的包。

#### 使用go module引入包

接下来我们将通过一个示例来演示如何在开发项目时使用 go module 拉取和管理项目依赖。

**初始化项目** 我们在本地新建一个名为`holiday`项目，按如下方式创建一个名为`holiday`的文件夹并切换到该目录下：

```bash
$ mkdir holiday
$ cd holiday
```

目前我们位于`holiday`文件夹下，接下来执行下面的命令初始化项目。

```bash
$ go mod init holiday
go: creating new go.mod: module holiday
```

该命令会自动在项目目录下创建一个`go.mod`文件，其内容如下。

```go
module holiday

go 1.16
```

其中：

- module holiday：定义当前项目的导入路径
- go 1.16：标识当前项目使用的 Go 版本

`go.mod`文件会记录项目使用的第三方依赖包信息，包括包名和版本，由于我们的`holiday`项目目前还没有使用到第三方依赖包，所以`go.mod`文件暂时还没有记录任何依赖包信息，只有当前项目的一些信息。

接下来，我们在项目目录下新建一个`main.go`文件，其内容如下：

```go
// holiday/main.go

package main

import "fmt"

func main() {
	fmt.Println("现在是假期时间...")
}
```

然后，我们的`holiday`项目现在需要引入一个第三方包`github.com/q1mi/hello`来实现一些必要的功能。类似这样的场景在我们的日常开发中是很常见的。我们需要先将依赖包下载到本地同时在`go.mod`中记录依赖信息，然后才能在我们的代码中引入并使用这个包。下载依赖包主要有两种方法。

第一种方法是在项目目录下执行`go get`命令手动下载依赖的包：

```bash
holiday $ go get -u github.com/q1mi/hello
go get: added github.com/q1mi/hello v0.1.1
```

这样默认会下载最新的发布版本，你也可以指定想要下载指定的版本号的。

```bash
holiday $ go get -u github.com/q1mi/hello@v0.1.0
go: downloading github.com/q1mi/hello v0.1.0
go get: downgraded github.com/q1mi/hello v0.1.1 => v0.1.0
```

如果依赖包没有发布任何版本则会拉取最新的提交，最终`go.mod`中的依赖信息会变成类似下面这种由默认v0.0.0的版本号和最新一次commit的时间和hash组成的版本格式：

```go
require github.com/q1mi/hello v0.0.0-20210218074646-139b0bcd549d
```

如果想指定下载某个commit对应的代码，可以直接指定commit hash，不过没有必要写出完整的commit hash，一般前7位即可。例如：

```bash
holiday $ go get github.com/q1mi/hello@2ccfadd
go: downloading github.com/q1mi/hello v0.1.2-0.20210219092711-2ccfaddad6a3
go get: added github.com/q1mi/hello v0.1.2-0.20210219092711-2ccfaddad6a3
```

此时，我们打开`go.mod`文件就可以看到下载的依赖包及版本信息都已经被记录下来了。

```go
module holiday

go 1.16

require github.com/q1mi/hello v0.1.0 // indirect
```

行尾的`indirect`表示该依赖包为间接依赖，说明在当前程序中的所有 import 语句中没有发现引入这个包。

另外在执行`go get`命令下载一个新的依赖包时一般会额外添加`-u`参数，强制更新现有依赖。

第二种方式是我们直接编辑`go.mod`文件，将依赖包和版本信息写入该文件。例如我们修改`holiday/go.mod`文件内容如下：

```go
module holiday

go 1.16

require github.com/q1mi/hello latest
```

表示当前项目需要使用`github.com/q1mi/hello`库的最新版本，然后在项目目录下执行`go mod download`下载依赖包。

```bash
holiday $ go mod download
```

如果不输出其它提示信息就说明依赖已经下载成功，此时`go.mod`文件已经变成如下内容。

```go
module holiday

go 1.16

require github.com/q1mi/hello v0.1.1
```

从中我们可以知道最新的版本号是`v0.1.1`。如果事先知道依赖包的具体版本号，可以直接在`go.mod`中指定需要的版本然后再执行`go mod download`下载。

这种方法同样支持指定想要下载的commit进行下载，例如直接在`go.mod`文件中按如下方式指定commit hash，这里只写出来了commit hash的前7位。

```go
require github.com/q1mi/hello 2ccfadda
```

执行`go mod download`下载完依赖后，`go.mod`文件中对应的版本信息会自动更新为类似下面的格式。

```go
module holiday

go 1.16

require github.com/q1mi/hello v0.1.2-0.20210219092711-2ccfaddad6a3
```

下载好要使用的依赖包之后，我们现在就可以在`holiday/main.go`文件中使用这个包了。

```go
package main

import (
	"fmt"

	"github.com/q1mi/hello"
)

func main() {
	fmt.Println("现在是假期时间...")

	hello.SayHi() // 调用hello包的SayHi函数
}
```

将上述代码编译执行，就能看到执行结果了。

```bash
holiday $ go build
holiday $ ./holiday
现在是假期时间...
你好，我是七米。很高兴认识你。
```

当我们的项目功能越做越多，代码越来越多的时候，通常会选择在项目内部按功能或业务划分成多个不同包。Go语言支持在一个项目（project）下定义多个包（package）。

例如，我们在`holiday`项目内部创建一个新的package——`summer`，此时新的项目目录结构如下：

```bash
holidy
├── go.mod
├── go.sum
├── main.go
└── summer
    └── summer.go
```

其中`holiday/summer/summer.go`文件内容如下：

```go
package summer

import "fmt"

// Diving 潜水...
func Diving() {
	fmt.Println("夏天去诗巴丹潜水...")
}
```

此时想要在当前项目目录下的其他包或者`main.go`中调用这个`Diving`函数需要如何引入呢？这里以在`main.go`中演示详细的调用过程为例，在项目内其他包的引入方式类似。

```go
package main

import (
	"fmt"

	"holiday/summer" // 导入当前项目下的包

	"github.com/q1mi/hello" // 导入github上第三方包
)

func main() {
	fmt.Println("现在是假期时间...")
	hello.SayHi()

	summer.Diving()
}
```

从上面的示例可以看出，项目中定义的包都会以项目的导入路径为前缀。

如果你想要导入本地的一个包，并且这个包也没有发布到到其他任何代码仓库，这时候你可以在`go.mod`文件中使用`replace`语句将依赖临时替换为本地的代码包。例如在我的电脑上有另外一个名为`liwenzhou.com/overtime`的项目，它位于`holiday`项目同级目录下：

```bash
├── holiday
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── summer
│       └── summer.go
└── overtime
    ├── go.mod
    └── overtime.go
```

由于`liwenzhou.com/overtime`包只存在于我本地，并不能通过网络获取到这个代码包，这个时候应该如何在`holidy`项目中引入它呢？

我们可以在`holidy/go.mod`文件中正常引入`liwenzhou.com/overtime`包，然后像下面的示例那样使用`replace`语句将这个依赖替换为使用相对路径表示的本地包。

```go
module holiday

go 1.16

require github.com/q1mi/hello v0.1.1
require liwenzhou.com/overtime v0.0.0

replace liwenzhou.com/overtime  => ../overtime
```

这样，我们就可以在`holiday/main.go`下正常引入并使用`overtime`包了。

```go
package main

import (
	"fmt"

	"holiday/summer" // 导入当前项目下的包

	"liwenzhou.com/overtime" // 通过replace导入的本地包

	"github.com/q1mi/hello" // 导入github上第三方包
)

func main() {
	fmt.Println("现在是假期时间...")
	hello.SayHi()

	summer.Diving()

	overtime.Do()
}
```

我们也经常使用`replace`将项目依赖中的某个包，替换为其他版本的代码包或我们自己修改后的代码包。

**go.mod文件**

`go.mod`文件中记录了当前项目中所有依赖包的相关信息，声明依赖的格式如下：

```bash
require module/path v1.2.3
```

其中：

- require：声明依赖的关键字
- module/path：依赖包的引入路径
- v1.2.3：依赖包的版本号。支持以下几种格式：
  - latest：最新版本
  - v1.0.0：详细版本号
  - commit hash：指定某次commit hash

引入某些没有发布过`tag`版本标识的依赖包时，`go.mod`中记录的依赖版本信息就会出现类似`v0.0.0-20210218074646-139b0bcd549d`的格式，由版本号、commit时间和commit的hash值组成。

![go module生成的版本信息组成示意图](https://www.liwenzhou.com/images/Go/package/module_version_info.png)

**go.sum文件**

使用go module下载了依赖后，项目目录下还会生成一个`go.sum`文件，这个文件中详细记录了当前项目中引入的依赖包的信息及其hash 值。`go.sum`文件内容通常是以类似下面的格式出现。

```go
<module> <version>/go.mod <hash>
```

或者

```go
<module> <version> <hash>
<module> <version>/go.mod <hash>
```

不同于其他语言提供的基于中心的包管理机制，例如 npm 和 pypi等，Go并没有提供一个中央仓库来管理所有依赖包，而是采用分布式的方式来管理包。为了防止依赖包被非法篡改，Go module 引入了`go.sum`机制来对依赖包进行校验。

**依赖保存位置**

Go module 会把下载到本地的依赖包会以类似下面的形式保存在 `$GOPATH/pkg/mod`目录下，每个依赖包都会带有版本号进行区分，这样就允许在本地存在同一个包的多个不同版本。

```bash
mod
├── cache
├── cloud.google.com
├── github.com
    	└──q1mi
          ├── hello@v0.0.0-20210218074646-139b0bcd549d
          ├── hello@v0.1.1
          └── hello@v0.1.0
...
```

如果想清除所有本地已缓存的依赖包数据，可以执行 `go clean -modcache` 命令。

#### 使用go module发布包

在上面的小节中我们学习了如何在项目中引入别人提供的依赖包，那么当我们想要在社区发布一个自己编写的代码包或者在公司内部编写一个供内部使用的公用组件时，我们该怎么做呢？接下来，我们就一起编写一个代码包并将它发布到`github.com`仓库，让它能够被全球的Go语言开发者使用。

我们首先在自己的 github 账号下新建一个项目，并把它下载到本地。我这里就以创建和发布一个名为`hello`的项目为例进行演示。这个`hello`包将对外提供一个名为`SayHi`的函数，它的作用非常简单就是向调用者发去问候。

```bash
$ git clone https://github.com/q1mi/hello
$ cd hello
```

我们当前位于`hello`项目目录下，执行下面的命令初始化项目，创建`go.mod`文件。需要注意的是这里定义项目的引入路径为`github.com/q1mi/hello`，读者在自行测试时需要将这部分替换为自己的仓库路径。

```bash
hello $ go mod init github.com/q1mi/hello
go: creating new go.mod: module github.com/q1mi/hello
```

接下来我们在该项目根目录下创建 `hello.go` 文件，添加下面的内容：

```go
package hello

import "fmt"

func SayHi() {
	fmt.Println("你好，我是七米。很高兴认识你。")
}
```

然后将该项目的代码 push 到仓库的远端分支，这样就对外发布了一个Go包。其他的开发者可以通过`github.com/q1mi/hello`这个引入路径下载并使用这个包了。

一个设计完善的包应该包含开源许可证及文档等内容，并且我们还应该尽心维护并适时发布适当的版本。github 上发布版本号使用git tag为代码包打上标签即可。

```bash
hello $ git tag -a v0.1.0 -m "release version v0.1.0"
hello $ git push origin v0.1.0
```

经过上面的操作我们就发布了一个版本号为`v0.1.0`的版本。

Go modules中建议使用语义化版本控制，其建议的版本号格式如下：

![语义化版本号示意图](https://www.liwenzhou.com/images/Go/package/version_number.png)

其中：

- 主版本号：发布了不兼容的版本迭代时递增（breaking changes）。
- 次版本号：发布了功能性更新时递增。
- 修订号：发布了bug修复类更新时递增。

**发布新的主版本**

现在我们的`hello`项目要进行与之前版本不兼容的更新，我们计划让`SayHi`函数支持向指定人发出问候。更新后的`SayHi`函数内容如下：

```go
package hello

import "fmt"

// SayHi 向指定人打招呼的函数
func SayHi(name string) {
	fmt.Printf("你好%s，我是七米。很高兴认识你。\n", name)
}
```

由于这次改动巨大（修改了函数之前的调用规则），对之前使用该包作为依赖的用户影响巨大。因此我们需要发布一个主版本号递增的`v2`版本。在这种情况下，我们通常会修改当前包的引入路径，像下面的示例一样为引入路径添加版本后缀。

```go
// hello/go.mod

module github.com/q1mi/hello/v2

go 1.16
```

把修改后的代码提交：

```bash
hello $ git add .
hello $ git commit -m "feat: SayHi现在支持给指定人打招呼啦"
hello $ git push
```

打好 tag 推送到远程仓库。

```bash
hello $ git tag -a v2.0.0 -m "release version v2.0.0"
hello $ git push origin v2.0.0
```

这样在不影响使用旧版本的用户的前提下，我们新的版本也发布出去了。想要使用`v2`版本的代码包的用户只需按修改后的引入路径下载即可。

```bash
go get github.com/q1mi/hello/v2@v2.0.0
```

在代码中使用的过程与之前类似，只是需要注意引入路径要添加 v2 版本后缀。

```go
package main

import (
	"fmt"

	"github.com/q1mi/hello/v2" // 引入v2版本
)

func main() {
	fmt.Println("现在是假期时间...")

	hello.SayHi("张三") // v2版本的SayHi函数需要传入字符串参数
}
```

**废弃已发布版本**

如果某个发布的版本存在致命缺陷不再想让用户使用时，我们可以使用`retract`声明废弃的版本。例如我们在`hello/go.mod`文件中按如下方式声明即可对外废弃`v0.1.2`版本。

```go
module github.com/q1mi/hello

go 1.16


retract v0.1.2
```

用户使用go get下载`v0.1.2`版本时就会收到提示，催促其升级到其他版本。



## 接口

**本章学习目标**

- 了解为什么需要接口以及接口的特点
- 掌握接口的声明和使用
- 掌握接口值的概念
- 掌握空接口的特点及其使用场景

在Go语言中接口（interface）是一种类型，一种抽象的类型。相较于之前章节中讲到的那些具体类型（字符串、切片、结构体等）更注重“我是谁”，接口类型更注重“我能做什么”的问题。接口类型就像是一种约定——概括了一种类型应该具备哪些方法，在Go语言中提倡使用面向接口的编程方式实现解耦。

### 接口类型

接口是一种由程序员来定义的类型，一个接口类型就是一组方法的集合，它规定了需要实现的所有方法。

相较于使用结构体类型，当我们使用接口类型说明相比于它是什么更关心它能做什么。

#### 接口的定义

每个接口类型由任意个方法签名组成，接口的定义格式如下：

```go
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

其中：

- 接口类型名：Go语言的接口在命名时，一般会在单词后面添加`er`，如有写操作的接口叫`Writer`，有关闭操作的接口叫`closer`等。接口名最好要能突出该接口的类型含义。
- 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
- 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

举个例子，定义一个包含`Write`方法的`Writer`接口。

```go
type Writer interface{
    Write([]byte) error
}
```

当你看到一个`Writer`接口类型的值时，你不知道它是什么，唯一知道的就是可以通过调用它的`Write`方法来做一些事情。

#### 实现接口的条件

接口就是规定了一个**需要实现的方法列表**，在 Go 语言中一个类型只要实现了接口中规定的所有方法，那么我们就称它实现了这个接口。

我们定义的`Singer`接口类型，它包含一个`Sing`方法。

```go
// Singer 接口
type Singer interface {
	Sing()
}
```

我们有一个`Bird`结构体类型如下。

```go
type Bird struct {}
```

因为`Singer`接口只包含一个`Sing`方法，所以只需要给`Bird`结构体添加一个`Sing`方法就可以满足`Singer`接口的要求。

```go
// Sing Bird类型的Sing方法
func (b Bird) Sing() {
	fmt.Println("汪汪汪")
}
```

这样就称为`Bird`实现了`Singer`接口。

#### 为什么要使用接口？

现在假设我们的代码世界里有很多小动物，下面的代码片段定义了猫和狗，它们饿了都会叫。

```go
package main

import "fmt"

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

func main() {
	c := Cat{}
	c.Say()
	d := Dog{}
	d.Say()
}
```

这个时候又跑来了一只羊，羊饿了也会发出叫声。

```go
type Sheep struct{}

func (s Sheep) Say() {
	fmt.Println("咩咩咩~")
}
```

我们接下来定义一个饿肚子的场景。

```go
// MakeCatHungry 猫饿了会喵喵喵~
func MakeCatHungry(c Cat) {
	c.Say()
}

// MakeSheepHungry 羊饿了会咩咩咩~
func MakeSheepHungry(s Sheep) {
	s.Say()
}
```

接下来会有越来越多的小动物跑过来，我们的代码世界该怎么拓展呢？

在饿肚子这个场景下，我们可不可以把所有动物都当成一个“会叫的类型”来处理呢？当然可以！使用接口类型就可以实现这个目标。 我们的代码其实并不关心究竟是什么动物在叫，我们只是在代码中调用它的`Say()`方法，这就足够了。

我们可以约定一个`Sayer`类型，它必须实现一个`Say()`方法，只要饿肚子了，我们就调用`Say()`方法。

```go
type Sayer interface {
    Say()
}
```

然后我们定义一个通用的`MakeHungry`函数，接收`Sayer`类型的参数。

```go
// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
	s.Say()
}
```

我们通过使用接口类型，把所有会叫的动物当成`Sayer`类型来处理，只要实现了`Say()`方法都能当成`Sayer`类型的变量来处理。

```go
var c cat
MakeHungry(c)
var d dog
MakeHungry(d)
```

在电商系统中我们允许用户使用多种支付方式（支付宝支付、微信支付、银联支付等），我们的交易流程中可能不太在乎用户究竟使用什么支付方式，只要它能提供一个实现支付功能的`Pay`方法让调用方调用就可以了。

再比如我们需要在某个程序中添加一个将某些指标数据向外输出的功能，根据不同的需求可能要将数据输出到终端、写入到文件或者通过网络连接发送出去。在这个场景下我们可以不关注最终输出的目的地是什么，只需要它能提供一个`Write`方法让我们把内容写入就可以了。

Go语言中为了解决类似上面的问题引入了接口的概念，接口类型区别于我们之前章节中介绍的那些具体类型，让我们专注于该类型提供的方法，而不是类型本身。使用接口类型通常能够让我们写出更加通用和灵活的代码。

#### 面向接口编程

PHP、Java等语言中也有接口的概念，不过在PHP和Java语言中需要显式声明一个类实现了哪些接口，在Go语言中使用隐式声明的方式实现接口。只要一个类型实现了接口中规定的所有方法，那么它就实现了这个接口。

Go语言中的这种设计符合程序开发中抽象的一般规律，例如在下面的代码示例中，我们的电商系统最开始只设计了支付宝一种支付方式：

```go
type ZhiFuBao struct {
	// 支付宝
}

// Pay 支付宝的支付方法
func (z *ZhiFuBao) Pay(amount int64) {
  fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}

// Checkout 结账
func Checkout(obj *ZhiFuBao) {
	// 支付100元
	obj.Pay(100)
}

func main() {
	Checkout(&ZhiFuBao{})
}
```

随着业务的发展，根据用户需求添加支持微信支付。

```go
type WeChat struct {
	// 微信
}

// Pay 微信的支付方法
func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信付款：%.2f元。\n", float64(amount/100))
}
```

在实际的交易流程中，我们可以根据用户选择的支付方式来决定最终调用支付宝的Pay方法还是微信支付的Pay方法。

```go
// Checkout 支付宝结账
func CheckoutWithZFB(obj *ZhiFuBao) {
	// 支付100元
	obj.Pay(100)
}

// Checkout 微信支付结账
func CheckoutWithWX(obj *WeChat) {
	// 支付100元
	obj.Pay(100)
}
```

实际上，从上面的代码示例中我们可以看出，我们其实并不怎么关心用户选择的是什么支付方式，我们只关心调用Pay方法时能否正常运行。这就是典型的“不关心它是什么，只关心它能做什么”的场景。

在这种场景下我们可以将具体的支付方式抽象为一个名为`Payer`的接口类型，即任何实现了`Pay`方法的都可以称为`Payer`类型。

```go
// Payer 包含支付方法的接口类型
type Payer interface {
	Pay(int64)
}
```

此时只需要修改下原始的`Checkout`函数，它接收一个`Payer`类型的参数。这样就能够在不修改既有函数调用的基础上，支持新的支付方式。

```go
// Checkout 结账
func Checkout(obj Payer) {
	// 支付100元
	obj.Pay(100)
}

func main() {
	Checkout(&ZhiFuBao{}) // 之前调用支付宝支付

	Checkout(&WeChat{}) // 现在支持使用微信支付
}
```

像类似的例子在我们编程过程中会经常遇到：

- 比如一个网上商城可能使用支付宝、微信、银联等方式去在线支付，我们能不能把它们当成“支付方式”来处理呢？
- 比如三角形，四边形，圆形都能计算周长和面积，我们能不能把它们当成“图形”来处理呢？
- 比如满减券、立减券、打折券都属于电商场景下常见的优惠方式，我们能不能把它们当成“优惠券”来处理呢？

接口类型是Go语言提供的一种工具，在实际的编码过程中是否使用它由你自己决定，但是通常使用接口类型可以使代码更清晰易读。

#### 接口类型变量

那实现了接口又有什么用呢？一个接口类型的变量能够存储所有实现了该接口的类型变量。

例如在上面的示例中，`Dog`和`Cat`类型均实现了`Sayer`接口，此时一个`Sayer`类型的变量就能够接收`Cat`和`Dog`类型的变量。

```go
var x Sayer // 声明一个Sayer类型的变量x
a := Cat{}  // 声明一个Cat类型变量a
b := Dog{}  // 声明一个Dog类型变量b
x = a       // 可以把Cat类型变量直接赋值给x
x.Say()     // 喵喵喵
x = b       // 可以把Dog类型变量直接赋值给x
x.Say()     // 汪汪汪
```

### 值接收者和指针接收者

在结构体那一章节中，我们介绍了在定义结构体方法时既可以使用值接收者也可以使用指针接收者。那么对于实现接口来说使用值接收者和使用指针接收者有什么区别呢？接下来我们通过一个例子看一下其中的区别。

我们定义一个`Mover`接口，它包含一个`Move`方法。

```go
// Mover 定义一个接口类型
type Mover interface {
	Move()
}
```

#### 值接收者实现接口

我们定义一个`Dog`结构体类型，并使用值接收者为其定义一个`Move`方法。

```go
// Dog 狗结构体类型
type Dog struct{}

// Move 使用值接收者定义Move方法实现Mover接口
func (d Dog) Move() {
	fmt.Println("狗会动")
}
```

此时实现`Mover`接口的是`Dog`类型。

```go
var x Mover    // 声明一个Mover类型的变量x

var d1 = Dog{} // d1是Dog类型
x = d1         // 可以将d1赋值给变量x
x.Move()

var d2 = &Dog{} // d2是Dog指针类型
x = d2          // 也可以将d2赋值给变量x
x.Move()
```

从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。

#### 指针接收者实现接口

我们再来测试一下使用指针接收者实现接口有什么区别。

```go
// Cat 猫结构体类型
type Cat struct{}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}
```

此时实现`Mover`接口的是`*Cat`类型，我们可以将`*Cat`类型的变量直接赋值给`Mover`接口类型的变量`x`。

```go
var c1 = &Cat{} // c1是*Cat类型
x = c1          // 可以将c1当成Mover类型
x.Move()
```

但是不能给将`Cat`类型的变量赋值给`Mover`接口类型的变量`x`。

```go
// 下面的代码无法通过编译
var c2 = Cat{} // c2是Cat类型
x = c2         // 不能将c2当成Mover类型
```

由于Go语言中有对指针求值的语法糖，对于值接收者实现的接口，无论使用值类型还是指针类型都没有问题。但是我们并不总是能对一个值求址，所以对于指针接收者实现的接口要额外注意。

### 类型与接口的关系

#### 一个类型实现多个接口

一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。例如狗不仅可以叫，还可以动。我们完全可以分别定义`Sayer`接口和`Mover`接口，具体代码示例如下。

```go
// Sayer 接口
type Sayer interface {
	Say()
}

// Mover 接口
type Mover interface {
	Move()
}
```

`Dog`既可以实现`Sayer`接口，也可以实现`Mover`接口。

```go
type Dog struct {
	Name string
}

// 实现Sayer接口
func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪\n", d.Name)
}

// 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}
```

同一个类型实现不同的接口互相不影响使用。

```go
var d = Dog{Name: "旺财"}

var s Sayer = d
var m Mover = d

s.Say()  // 对Sayer类型调用Say方法
m.Move() // 对Mover类型调用Move方法
```

#### 多种类型实现同一接口

Go语言中不同的类型还可以实现同一接口。例如在我们的代码世界中不仅狗可以动，汽车也可以动。我们可以使用如下代码体现这个关系。

```go
// 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

// Car 汽车结构体类型
type Car struct {
	Brand string
}

// Move Car类型实现Mover接口
func (c Car) Move() {
	fmt.Printf("%s速度70迈\n", c.Brand)
}
```

这样我们在代码中就可以把狗和汽车当成一个会动的类型来处理，不必关注它们具体是什么，只需要调用它们的`Move`方法就可以了。

```go
var obj Mover

obj = Dog{Name: "旺财"}
obj.Move()

obj = Car{Brand: "宝马"}
obj.Move()
```

上面的代码执行结果如下：

```go
旺财会跑
宝马速度70迈
```

一个接口的所有方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

```go
// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}
```

### 接口组合

接口与接口之间可以通过互相嵌套形成新的接口类型，例如Go标准库`io`源码中就有很多接口之间互相组合的示例。

```go
// src/io/io.go

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// ReadWriter 是组合Reader接口和Writer接口形成的新接口类型
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloser 是组合Reader接口和Closer接口形成的新接口类型
type ReadCloser interface {
	Reader
	Closer
}

// WriteCloser 是组合Writer接口和Closer接口形成的新接口类型
type WriteCloser interface {
	Writer
	Closer
}
```

对于这种由多个接口类型组合形成的新接口类型，同样只需要实现新接口类型中规定的所有方法就算实现了该接口类型。

接口也可以作为结构体的一个字段，我们来看一段Go标准库`sort`源码中的示例。

```go
// src/sort/sort.go

// Interface 定义通过索引对元素排序的接口类型
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}


// reverse 结构体中嵌入了Interface接口
type reverse struct {
    Interface
}
```

通过在结构体中嵌入一个接口类型，从而让该结构体类型实现了该接口类型，并且还可以改写该接口的方法。

```go
// Less 为reverse类型添加Less方法，重写原Interface接口类型的Less方法
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}
```

`Interface`类型原本的`Less`方法签名为`Less(i, j int) bool`，此处重写为`r.Interface.Less(j, i)`，即通过将索引参数交换位置实现反转。

在这个示例中还有一个需要注意的地方是`reverse`结构体本身是不可导出的（结构体类型名称首字母小写），`sort.go`中通过定义一个可导出的`Reverse`函数来让使用者创建`reverse`结构体实例。

```go
func Reverse(data Interface) Interface {
	return &reverse{data}
}
```

这样做的目的是保证得到的`reverse`结构体中的`Interface`属性一定不为`nil`，否者`r.Interface.Less(j, i)`就会出现空指针panic。

此外在Go内置标准库`database/sql`中也有很多类似的结构体内嵌接口类型的使用示例，各位读者可自行查阅。

### 空接口

#### 空接口的定义

空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。

```go
package main

import "fmt"

// 空接口

// Any 不包含任何方法的空接口类型
type Any interface{}

// Dog 狗结构体
type Dog struct{}

func main() {
	var x Any

	x = "你好" // 字符串型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = 100 // int型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = true // 布尔型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = Dog{} // 结构体类型
	fmt.Printf("type:%T value:%v\n", x, x)
}
```

通常我们在使用空接口类型时不必使用`type`关键字声明，可以像下面的代码一样直接使用`interface{}`。

```go
var x interface{}  // 声明一个空接口类型变量x
```

#### 空接口的应用

##### 空接口作为函数的参数

使用空接口实现可以接收任意类型的函数参数。

```go
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
```

##### 空接口作为map的值

使用空接口实现可以保存任意值的字典。

```go
// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
```



### 接口值

由于接口类型的值可以是任意一个实现了该接口的类型值，所以接口值除了需要记录具体**值**之外，还需要记录这个值属于的**类型**。也就是说接口值由“类型”和“值”组成，鉴于这两部分会根据存入值的不同而发生变化，我们称之为接口的`动态类型`和`动态值`。

![接口值示例](https://www.liwenzhou.com/images/Go/interface/interface01.png)

我们接下来通过一个示例来加深对接口值的理解。

下面的示例代码中，定义了一个`Mover`接口类型和两个实现了该接口的`Dog`和`Car`结构体类型。

```go
type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

func (d *Dog) Move() {
	fmt.Println("狗在跑~")
}

type Car struct {
	Brand string
}

func (c *Car) Move() {
	fmt.Println("汽车在跑~")
}
```

首先，我们创建一个`Mover`接口类型的变量`m`。

```go
var m Mover
```

此时，接口变量`m`是接口类型的零值，也就是它的类型和值部分都是`nil`，就如下图所示。

![接口值示例](https://www.liwenzhou.com/images/Go/interface/interface02.png)

我们可以使用`m == nil`来判断此时的接口值是否为空。

```go
fmt.Println(m == nil)  // true
```

**注意：**我们不能对一个空接口值调用任何方法，否则会产生panic。

```go
m.Move() // panic: runtime error: invalid memory address or nil pointer dereference
```

接下来，我们将一个`*Dog`结构体指针赋值给变量`m`。

```go
m = &Dog{Name: "旺财"}
```

此时，接口值`m`的动态类型会被设置为`*Dog`，动态值为结构体变量的拷贝。

![接口值示例](https://www.liwenzhou.com/images/Go/interface/interface03.png)

然后，我们给接口变量`m`赋值为一个`*Car`类型的值。

```go
var c *Car
m = c
```

这一次，接口值`m`的动态类型为`*Car`，动态值为`nil`。

![接口值示例](https://www.liwenzhou.com/images/Go/interface/interface04.png)

**注意：**此时接口变量`m`与`nil`并不相等，因为它只是动态值的部分为`nil`，而动态类型部分保存着对应值的类型。

```go
fmt.Println(m == nil) // false
```

接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等时才相等。

```go
var (
	x Mover = new(Dog)
	y Mover = new(Car)
)
fmt.Println(x == y) // false
```

但是有一种特殊情况需要特别注意，如果接口值保存的动态类型相同，但是这个动态类型不支持互相比较（比如切片），那么对它们相互比较时就会引发panic。

```go
var z interface{} = []int{1, 2, 3}
fmt.Println(z == z) // panic: runtime error: comparing uncomparable type []int
```

### 类型断言

接口值可能赋值为任意类型的值，那我们如何从接口值获取其存储的具体数据呢？

我们可以借助标准库`fmt`包的格式化打印获取到接口值的动态类型。

```go
var m Mover

m = &Dog{Name: "旺财"}
fmt.Printf("%T\n", m) // *main.Dog

m = new(Car)
fmt.Printf("%T\n", m) // *main.Car
```

而`fmt`包内部其实是使用反射的机制在程序运行时获取到动态类型的名称。关于反射的内容我们会在后续章节详细介绍。

而想要从接口值中获取到对应的实际值需要使用类型断言，其语法格式如下。

```go
x.(T)
```

其中：

- x：表示接口类型的变量
- T：表示断言`x`可能是的类型。

该语法返回两个参数，第一个参数是`x`转化为`T`类型后的变量，第二个值是一个布尔值，若为`true`则表示断言成功，为`false`则表示断言失败。

举个例子：

```go
var n Mover = &Dog{Name: "旺财"}
v, ok := n.(*Dog)
if ok {
	fmt.Println("类型断言成功")
	v.Name = "富贵" // 变量v是*Dog类型
} else {
	fmt.Println("类型断言失败")
}
```

如果对一个接口值有多个实际类型需要判断，推荐使用`switch`语句来实现。

```go
// justifyType 对传入的空接口类型变量x进行类型断言
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
```

由于接口类型变量能够动态存储不同类型值的特点，所以很多初学者会滥用接口类型（特别是空接口）来实现编码过程中的便捷。只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。切记不要为了使用接口类型而增加不必要的抽象，导致不必要的运行时损耗。

在 Go 语言中接口是一个非常重要的概念和特性，使用接口类型能够实现代码的抽象和解耦，也可以隐藏某个功能的内部实现，但是缺点就是在查看源码的时候，不太方便查找到具体实现接口的类型。

相信很多读者在刚接触到接口类型时都会有很多疑惑，请牢记接口是一种类型，一种抽象的类型。区别于我们在之前章节提到的那些具体类型（整型、数组、结构体类型等），它是一个只要求实现特定方法的抽象类型。

**小技巧：** 下面的代码可以在程序编译阶段验证某一结构体是否满足特定的接口类型。

```go
// 摘自gin框架routergroup.go
type IRouter interface{ ... }

type RouterGroup struct { ... }

var _ IRouter = &RouterGroup{}  // 确保RouterGroup实现了接口IRouter
```

上面的代码中也可以使用`var _ IRouter = (*RouterGroup)(nil)`进行验证。

### 练习题

1. 使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库。



## Error接口和错误处理

### Error 接口

Go 语言中把错误当成一种特殊的值来处理，不支持其他语言中使用`try/catch`捕获异常的方式。

#### Error 接口

Go 语言中使用一个名为 `error` 接口来表示错误类型。

```go
type error interface {
    Error() string
}
```

`error` 接口只包含一个方法——`Error`，这个函数需要返回一个描述错误信息的字符串。

当一个函数或方法需要返回错误时，我们通常是把错误作为最后一个返回值。例如下面标准库 os 中打开文件的函数。

```go
func Open(name string) (*File, error) {
	return OpenFile(name, O_RDONLY, 0)
}
```

由于 error 是一个接口类型，默认零值为`nil`。所以我们通常将调用函数返回的错误与`nil`进行比较，以此来判断函数是否返回错误。例如你会经常看到类似下面的错误判断代码。

```go
file, err := os.Open("./xx.go")
if err != nil {
	fmt.Println("打开文件失败,err:", err)
	return
}
```

**注意**

当我们使用`fmt`包打印错误时会自动调用 error 类型的 Error 方法，也就是会打印出错误的描述信息。

#### 创建错误

我们可以根据需求自定义 error，最简单的方式是使用`errors` 包提供的`New`函数创建一个错误。

**errors.New**

函数签名如下，

```go
func New(text string) error
```

它接收一个字符串参数返回包含该字符串的错误。我们可以在函数返回时快速创建一个错误。

```go
func queryById(id int64) (*Info, error) {
	if id <= 0 {
		return nil, errors.New("无效的id")
	}

	// ...
}
```

或者用来定义一个错误变量，例如标准库`io.EOF`错误定义如下。

```go
var EOF = errors.New("EOF")
```

#### fmt.Errorf

当我们需要传入格式化的错误描述信息时，使用`fmt.Errorf`是个更好的选择。

```go
fmt.Errorf("查询数据库失败，err:%v", err)
```

但是上面的方式会丢失原有的错误类型，只拿到错误描述的文本信息。

为了不丢失函数调用的错误链，使用`fmt.Errorf`时搭配使用特殊的格式化动词`%w`，可以实现基于已有的错误再包装得到一个新的错误。

```go
fmt.Errorf("查询数据库失败，err:%w", err)
```

对于这种二次包装的错误，`errors`包中提供了以下三个方法。

```go
func Unwrap(err error) error                 // 获得err包含下一层错误
func Is(err, target error) bool              // 判断err是否包含target
func As(err error, target interface{}) bool  // 判断err是否为target类型
```

#### 错误结构体类型

此外我们还可以自己定义结构体类型，实现``error`接口。

```go
// OpError 自定义结构体类型
type OpError struct {
	Op string
}

// Error OpError 类型实现error接口
func (e *OpError) Error() string {
	return fmt.Sprintf("无权执行%s操作", e.Op)
}
```

## 反射

### 变量的内在机制

Go语言中的变量是分为两部分的:

- 类型信息：预先定义好的元信息。
- 值信息：程序运行过程中可动态变化的。

### 反射介绍

反射是指在程序运行期间对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。

支持反射的语言可以在程序编译期间将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期间获取类型的反射信息，并且有能力修改它们。

Go程序在运行期间使用reflect包访问程序的反射信息。

在上一篇博客中我们介绍了空接口。 空接口可以存储任意类型的变量，那我们如何知道这个空接口保存的数据是什么呢？ 反射就是在运行时动态的获取一个变量的类型信息和值信息。

### reflect包

在Go语言的反射机制中，任何接口值都由是`一个具体类型`和`具体类型的值`两部分组成的(我们在上一篇接口的博客中有介绍相关概念)。 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由`reflect.Type`和`reflect.Value`两部分组成，并且reflect包提供了`reflect.TypeOf`和`reflect.ValueOf`两个函数来获取任意对象的Value和Type。

#### TypeOf

在Go语言中，使用`reflect.TypeOf()`函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。

```go
package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}
func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}
```

##### type name和type kind

在反射中关于类型还划分为两种：`类型（Type）`和`种类（Kind）`。因为在Go语言中我们可以使用type关键字构造很多自定义类型，而`种类（Kind）`就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到`种类（Kind）`。 举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。

```go
package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType(d) // type:person kind:struct
	reflectType(e) // type:book kind:struct
}
```

Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的`.Name()`都是返回`空`。

在`reflect`包中定义的Kind类型如下：

```go
type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
```

#### ValueOf

`reflect.ValueOf()`返回的是`reflect.Value`类型，其中包含了原始值的值信息。`reflect.Value`与原始值之间可以互相转换。

`reflect.Value`类型提供的获取原始值的方法如下：

|           方法           |                             说明                             |
| :----------------------: | :----------------------------------------------------------: |
| Interface() interface {} | 将值以 interface{} 类型返回，可以通过类型断言转换为指定类型  |
|       Int() int64        |     将值以 int 类型返回，所有有符号整型均可以此方式返回      |
|      Uint() uint64       |     将值以 uint 类型返回，所有无符号整型均可以此方式返回     |
|     Float() float64      | 将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回 |
|       Bool() bool        |                     将值以 bool 类型返回                     |
|     Bytes() []bytes      |               将值以字节数组 []bytes 类型返回                |
|     String() string      |                     将值以字符串类型返回                     |

##### 通过反射获取值

```go
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
func main() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
}
```

##### 通过反射设置变量的值

想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的`Elem()`方法来获取指针对应的值。

```go
package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func main() {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}
```

##### isNil()和isValid()

###### isNil()

```go
func (v Value) IsNil() bool
```

`IsNil()`报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。

###### isValid()

```go
func (v Value) IsValid() bool
```

`IsValid()`返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。

###### 举个例子

`IsNil()`常被用于判断指针是否为空；`IsValid()`常被用于判定返回值是否有效。

```go
func main() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}
```

### 结构体反射

#### 与结构体相关的方法

任意值通过`reflect.TypeOf()`获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（`reflect.Type`）的`NumField()`和`Field()`方法获得结构体成员的详细信息。

`reflect.Type`中与获取结构体成员相关的的方法如下表所示。

| 方法                                                        | 说明                                                         |
| :---------------------------------------------------------- | ------------------------------------------------------------ |
| Field(i int) StructField                                    | 根据索引，返回索引对应的结构体字段的信息。                   |
| NumField() int                                              | 返回结构体成员字段数量。                                     |
| FieldByName(name string) (StructField, bool)                | 根据给定字符串返回字符串对应的结构体字段的信息。             |
| FieldByIndex(index []int) StructField                       | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。 |
| FieldByNameFunc(match func(string) bool) (StructField,bool) | 根据传入的匹配函数匹配需要的字段。                           |
| NumMethod() int                                             | 返回该类型的方法集中方法的数目                               |
| Method(int) Method                                          | 返回该类型方法集中的第i个方法                                |
| MethodByName(string)(Method, bool)                          | 根据方法名返回该类型方法集中的方法                           |

#### StructField类型

`StructField`类型用来描述结构体中的一个字段的信息。

`StructField`的定义如下：

```go
type StructField struct {
    // Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
    // 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
    Name    string
    PkgPath string
    Type      Type      // 字段的类型
    Tag       StructTag // 字段的标签
    Offset    uintptr   // 字段在结构体中的字节偏移量
    Index     []int     // 用于Type.FieldByIndex时的索引切片
    Anonymous bool      // 是否匿名字段
}
```

#### 结构体反射示例

当我们使用反射得到一个结构体数据之后可以通过索引依次获取其字段信息，也可以通过字段名去获取指定的字段信息。

```go
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
```

接下来编写一个函数`printMethod(s interface{})`来遍历打印s包含的方法。

```go
// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
```

### 反射是把双刃剑

反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。

1. 基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
2. 大量使用反射的代码通常难以理解。
3. 反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级。

### 练习题

1. 编写代码利用反射实现一个ini文件的解析器程序。



# Go语言神经练习题整理。

```
练习 1.1： 修改echo程序，使其能够打印os.Args[0]。  13
练习 1.2： 修改echo程序，使其打印value和index，每个value和index显示一行。13
练习 1.3： 上手实践前面提到的strings.Join和直接Println，并观察输出结果的区别。13
练习 1.4： 修改dup2，使其可以分别打印重复的行出现在哪些文件。
练习 1.5： 修改前面的Lissajous程序里的调色板，由绿色改为黑色。我们可以用color.RGBA{0xRR,0xGG, 0xBB}来得到#RRGGBB这个色值，三个十六进制的字符串分别代表红、绿、蓝像素。	20
练习 1.6： 修改Lissajous程序，修改其调色板来生成更丰富的颜色，然后修改SetColorIndex的第三个参数，看看显示结果吧。20
练习 1.7： 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。22
练习 1.8： 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。22
练习 1.9： 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。22
练习 1.10： 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一致，修改本节中的程序，将响应结果输出，以便于进行对比。	24
练习 1.12： 修改Lissajour服务，从URL读取变量，比如你可以访问 http://localhost:8000/?cycles=20 这个URL，这样访问可以将程序里的cycles默认的5修改为20。字符串转换为数字可以调用strconv.Atoi函数。你可以在dodoc里查看strconv.Atoi的详细说明。	29
```

```
练习 2.1： 向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是样。48
练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。50
练习 2.3： 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。（11.4节将展示如何系统地比较两个不同实现的性能。）52
练习 2.4： 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。52
练习 2.5： 表达式x&(x-1)用于将x的最低的一个非零的bit位清零。使用这个算法重写PopCount函数，然后比较性能52
```

```
练习 3.1： 如果f函数返回的是无限制的float64值，那么SVG文件可能输出无效的多边形元素（虽然许多SVG渲染器会妥善处理这类问题）。修改程序跳过无效的多边形。65
练习 3.2： 试验math包中其他函数的渲染图形。你是否能输出一个egg box、moguls或a saddle图案?65
练习 3.3： 根据高度给每个多边形上色，那样峰值部将是红色(#ff0000)，谷部将是蓝色(#0000ff)。65
练习 3.4： 参考1.7节Lissajous例子的函数，构造一个web服务器，用于计算函数曲面然后返回SVG数据给客户端。服务器必须设置Content-Type头部：65
练习 3.5： 实现一个彩色的Mandelbrot图像，使用image.NewRGBA创建图像，使用color.RGBA或color.YCbCr生成颜色。68
练习 3.6： 升采样技术可以降低每个像素对计算颜色值和平均值的影响。简单的方法是将每个像素分层四个子像素，实现它。68
练习 3.7： 另一个生成分形图像的方式是使用牛顿法来求解一个复数方程，例如z^4-1=04−1=0。每个起点到四个根的迭代次数对应阴影的灰度。方程根对应的点用颜色表示。68
练习 3.8： 通过提高精度来生成更多级别的分形。使用四种不同精度类型的数字实现相同的分形：complex64、complex128、big.Float和big.Rat。（后面两种类型在math/big包声明。Float是有指定限精度的浮点数；Rat是无效精度的有理数。）它们间的性能和内存使用对比如何？当渲染图可见时缩放的级别是多少？68
练习 3.9： 编写一个web服务器，用于给客户端生成分形的图像。运行客户端用过HTTP参数参数指定x,y和zoom参数68
```

```
练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)89
练习 4.2： 编写一个程序，默认打印标准输入的以SHA256哈希码，也可以通过命令行标准参数选择SHA384或SHA512哈希算法。89
练习 4.3： 重写reverse函数，使用数组指针代替slice。99
练习 4.4： 编写一个rotate函数，通过一次循环完成镟转。99
练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。99
练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回99
练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？99
练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。106
练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。106
练习 4.10： 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。121
练习 4.11： 编写一个工具，允许用户在命令行创建、读取、更新和删除GitHub上的issue，当必要的时候自动打开用户默认的编辑器用于输入文本信息。121
练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。例如，一个https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。下载每个链接（只下载一次）然后创建一个离线索引。编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。121
练习 4.13： 使用开放电影数据库的JSON服务接口，允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。编写一个poster工具，通过命令行输入的电影名字，下载对应的海报。121
练习 4.14： 创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。126
```

```
练习 5.1： 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。131
练习 5.2： 编写函数，记录在HTML树中出现的同名元素的次数。131
练习 5.3： 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。131
练习 5.4： 扩展vist函数，使其能够处理其他类型的结点，如images、scripts和style sheets。131
练习 5.5： 实现countWordsAndImages。（参考练习4.9如何分词）134
练习 5.6： 修改gopl.io/ch3/surface (§3.2) 中的corner函数，将返回值命名，并使用bare return。134
练习 5.7： 完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）。编写测试，验证程序输出的格式正确。（详见11章）142
练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍历。使用修改后的代码编写ElementByID函数，根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。142
func ElementByID(doc *html.Node, id string) *html.Node

练习 5.9： 编写函数expand，将s中的"foo"替换为f("foo")的返回值。142
func expand(s string, f func(string) string) string
练习5.10： 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。148
练习5.11： 现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图中的环。148
练习5.12： gopl.io/ch5/outline2（5.5节）的startElement和endElement共用了全局变量depth，将它们修改为匿名函数，使其共享outline中的局部变量。148
练习5.13： 修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。只保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。148
练习5.14： 使用breadthFirst遍历其他数据结构。比如，topoSort例子中的课程依赖关系（有向图）,个人计算机的文件层次结构（树），你所在城市的公交或地铁线路（无向图）。148
练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。150
练习5.16：编写多参数版本的strings.Join。150
练习5.17：编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。下面给出了2个例子：150
func ElementsByTagName(doc *html.Node, name...string) []*html.Node
images := ElementsByTagName(doc, "img")
headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
练习5.18：不修改fetch的行为，重写fetch函数，要求使用defer机制关闭文件	156
练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。161
```

```

```



# Go语言圣经笔记

## 2	程序结构

### 2.1 命名

​		Go语言的函数名、变量名、常量名、类型名、语句标号和包名等所有的命名，都遵循一个简单的命名规则：一个名字必须以一个字（Unicode字母）或者下划线开头，后面可以跟任意数量的字母、数字或下划线。大写字母和小写字母是不同的。

​		Go中的关键字有25个：关键字不能用于自定义名字，只能在特定语法结构中使用。

| **break**    | default         | func       | interface   | select     |
| ------------ | --------------- | ---------- | ----------- | ---------- |
| **case**     | **defer**       | **go**     | **map**     | **struct** |
| **chan**     | **else**        | **goto**   | **package** | **switch** |
| **const**    | **fallthrough** | **if**     | **range**   | **type**   |
| **continue** | **for**         | **import** | **return**  | **var**    |

​		还有30个预定义的名字，比如int和true等，主要对应内建的常量、类型和函数。

```go
内建常量:	true	 false	 iota		 nil
内建类型:	int		 int8	 int16		 int32	 	 int64
		   uint		uint8	uint16	    uint32		uint64
		   float32	float64	complex128	complex64
		   bool		byte	rune		string		error
内置函数:	make	 len		cap		new		append		copy	close	delete
		   complex		real	imag
		   panic		recover
```

​		内部预先定义的名字并不是关键字，可以再定义重新使用。在一些特殊的场景中重新定义他们也是有意义的，但是也要注意避免过度而引起语义混乱。

​		如果一个名字是在函数内部定义的，那么他就只在函数内部生效。如果是在函数外部定义，那么将在当前包的所有文件中都可以访问。名字的开头字母的大小写决定了名字在包外的可见性。如果一个名字是大写的并且是在函数外部定义的包级名字（包级函数名也属于包级名字），那么它将是到处的，也就是说可以被外部的包访问。例如fmt包的Printf函数就是导出的，可以在fmt包外部访问。包本身的名字一般总是用小写字母。

​		名字的长度一般没有逻辑限制。如果一个名字的作用域比较大，生命周期也比较长，那么用长的名字会更有意义。

### 2.2	声明

​		声明语句定义了程序的各种实体对象以及部分或全部的属性。

​		go语言主要有四种类型的声明：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明。

​		一个go语言编写的程序对应一个或多个以.go为文件后缀名的源文件中。每个源文件以包的声明语句开始，说明该源文件属于哪个包。包声明语句之后是import语句导入依赖的其他包，然后是包一级的类型、变量、常量和函数的声明语句，包一级的各种类型的声明语句的顺序无关紧要（函数内部的名字则必须先声明后才可以使用）。

​		在包以及声明语句	声明的名字可以在整个包对应的每个源文件中访问，而不是仅仅在其声明语句所在的源文件中访问。局部声明的名字就只能在函数内部很小的范围进行访问。

​		一个函数的声明由一个函数的名字、参数列表（由函数的调用者提供参数变量的具体值）、一个可选的返回值列表和包含函数定义的函数体组成。如果函数没有返回值，那么返回值列表可以省略。执行函数从函数的第一个语句开始，依次顺序执行知道遇到return语句，如果没有返回语句则是执行到函数末尾，然后返回到函数调用者。

### 2.3	变量

​		var声明语句可以创建一个特定类型的变量，然后给变量附加一个名字，并且设置变量的初始值。变量声明语法如下：

```go
var 变量名称 类型 = 表达式
```

​		其中“类型“或”=表达式“两个部分可以省略其中一个。如果省略的是类型信息，那么将根据初始化表达式来推导变量的类型信息。如果初始化表达式被省略，那么将用零值来初始化该变量。

​		数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型变量对应的零值为空字符串，接口或引用类型（包括slice、map、chan和函数）变量对应的零值为nil。数组或结构体等聚合类型对应的零值是每个元素或者字段都是对应该字段类型的零值。

​		零值初始化机制可以确保每个声明的变量总是有一个良好定义的值，因此在go中不存在未初始化的变量，也可以保证不管任何类型的变量总有一个合理有效的零值状态。

​		也可以在同一个声明语句中声明一组变量，或用一组初始化表达式声明并初始化一组变量。如果省略每个变量的类型，也可以声明多个不同类型的变量（类型由初始化表达式推导）：

```go
var i,j,k int					//int,int,int
var b,f,s = true,2.3,"four"		//bool,float64,string
```

​		初始化表达式可以是字面量或者任意表达式。在包级别声明的变量会在main入口函数执行前完成初始化，局部变量将在声明语句被执行到的时候完成初始化。

​		一组变量也可以通过调用一个函数，由函数返回的多个值完成初始化：

```go
var f , err = os.Open(name) 	//os.Open返回一个文件和一个错误
```

#### 2.3.1	简短声明

​		以"名字 := 表达式"形式声明变量，变量的类型根据表达式自动进行推导。

​		简短声明被广泛应用于大部分的局部变量的声明和初始化。var声明的形式往往应用于需要显示指定变量类型的地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方。

​		简短变量声明语句也可以用来声明和初始化一组变量：":="是一个变量声明语句，"="是一个变量赋值操作

```go
i , j := 0 , 1
```

​		简短变量声明左边的变量可能并不都是刚刚声明的。如果有一些已经在相同的词法域（词法域 == 作用域）声明过，那么简短变量声明语句对这些已经声明过的变量就只有赋值行为。

​		**简短变量声明语句中至少需要有一个新的变量。**

​		简短变量声明语句只有对已经在同级词法域声明过的变量具有赋值操作，如果变量是在外部词法域声明的，那么简短变量声明语句将会在当前词法域声明一个新的变量。

#### 2.3.2	指针

​		一个变量对应了一个保存了变量类型对应值的内存空间。普通变量在声明语句创建时被绑定到了一个变量名，比如叫x的变量，但还有很多变量始终以表达式的方式引入，例如 x[i] 、x.f 变量。所有这些表达式一般都是读取一个变量的值，除非它们是出现在赋值语句的左边，这个时候是给对应变量赋予一个新的值。

​		一个指针的值是另一个变量的地址。一个指针对应 变量 在内存中存储的位置。并不是每一个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址。通过指针，我们可以直接读取或更新对应变量的值，而不需要知道该变量的名字。

​		如果使用” var x int ”声明语句声明一个x变量，那么&x表达式（取x变量的内存地址）将产生一个指向该整数变量的指针，指针对应的数据类型是 *int ，指针被称之为“ 指向int类型的指针 ”。如果指针名字为p，那么可以说p指针指向x，或者说“ p指针保存了变量x的内存地址 ”。同时 \*p 表达式对应p指针指向变量的值。一般\*p表达式读取指针指向变量的值，这里为int类型的值，同时因为\*p对应一个变量，所以该表达式也可以出现在赋值语句的左边，表示更新指针所指向的变量的值。

```
x := 1			
p := &x				//p，指针类型，指向变量x
fmt.Println(*p)		//"1"
*p = 2 				//相当于x = 2
fmt.Println(x)		//"2"
```

​		对于聚合类型的每个成员，比如结构体的每个字段、或者是数组的每个元素，也都是对应一个变量，因此可以被取地址。

​		变量有时候被称为可寻址的值。即使变量由表达式临时生成，那么表达式也必须能接受 &取地址 操作。

​		任何类型的指针的零值都是nil。如果 p != nil 为真，那么p是指向某个有效变量。指针之间也是可以进行相等测试的，只有当他们指向同一个变量或全部为nil时才相等

```go
var x, y int
fmt.Println(&x==&x,&x==&y,&x==nil)	//true false true
```

​		在go语言中，返回函数中局部变量的地址也是安全的。

```go
//调用f函数时创建变量v，在局部变量地址被返回后依然有效，因为指针p依然应用这个变量
var p = f()

func f () *int{
	v := 1
	return &v	
}
//每次调用f函数都将返回不同的结果
fmt.Println(f() == f())	//false
```

​		因为指针包含了一个变量的地址，因此如果将指针作为参数调用函数，那将可以在函数中通过该指针来更新变量。

```go
func incr(p *int) int{
	*p++				//非常重要：只是增加p指向的变量的值，并不改变p指针
    return *p
}

v := 1
incr(&v)				//副作用：v现在值为2

fmt.Println(incr(&v))	//"3" (v现在是3)

```

​		每次我们对一个变量取地址，或者复制指针，都是为原变量创建了新的别名。例如，*p就是变量v的别名。指针特别有价值的地方在于我们可以不用名字而访问变量，但是要找到一个变量的所有访问者并不容易，我们必须知道变量全部的别名。不仅仅是指针会创建别名，很多其他引用类型也会创建别名，例如slice、map和chan，甚至结构体、数组和接口都会创建所引用变量的别名。

#### 2.3.3	new函数

​		另一个创建变量的方法是调用内建的new函数。表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T。

```go
p := new(int)		//p,*int类型，指向匿名变量p
fmt.Println(*p)		//“0”
*p = 2				//设置int匿名变量值为2
fmt.Println(*p)		//“2”
```

​		用new函数创建变量和普通变量声明语句创建变量没什么区别，除了不需要声明一个临时的变量名字外，我们还可以在表达式中使用new(T)。换言之，new函数类似是一种语法糖，而不是一个新的基础概念。语法糖：简而言之，语法糖就是一种简便写法。

​		例子中两个newInt函数有着相同的行为：

```go
func newInt() *int{
	return new(int)
}

func newInt() *int{
	var d int
	return &d
}
```

​		每次调用new函数都是返回一个新的变量的地址，因此两次调用的地址是不同的：	

```go
p := new(int)
q := new(int)
fmt.Println(p == q)		//false
```

​		当然也可能有特殊情况：如果两个类型都是空的，也就是说类型的大小是0，例如struct{} 和[0]int，有可能有相同的地址（以来具体的语言实现）（请谨慎使用大小为0的类型，因为如果类型的大小为0的话，可能导致go语言的自动垃圾回收器有不同的行为）

​		由于new只是一个预定义的函数，并不是一个关键字，因此我们可以将new名字重新定义为别的类型。

```go
func delta(old, new int)int{ return new - old}
```

​		由于new被定义为int类型的变量名，因此在delta函数的内部是无法使用内置的new函数的。

#### 2.3.4	变量的生命周期

​		变量的生命周期是指在程序运行期间变量有效存在的时间间隔。对于包一级声明的变量来说，他们的生命周期和整个程序的运行周期是一致的。而相比之下，局部变量的生命周期则是动态的：从每次创建一个新变量的声明语句开始，知道该变量不再被引用为止，然后变量的存储空间可能会被回收。函数的参数变量和返回值变量都是局部变量。他们在函数每次被调用的时候创建。

​		Go语言自动垃圾回收器基本思路：从每个包级的变量和每个当前运行函数的每一个局部变量开始，通过指针或应用的访问路径遍历，是否能找到该变量。如果不存在这样的访问路径，则说明变量是不可达的，也就是说他是否存在并不会影响程序后续的计算结果。

​		因为一个变量的有效期只取决于是否可达，因此一个循环迭代内部的局部变量的生命周期可能超出其局部作用域。同时，局部变量可能在函数返回后依然存在。

​		编译器会自动选择在栈上还是在堆上分配局部变量的存储空间，这个选择并不是由用var还是new函数声明变量的方式决定的。

```go
var global *int

func f(){
	var x int
	x = 1
	global =  &x
}

func g(){
	y := new(int)
	*y = 1
}
```

​		f函数中的x变量必须分配在堆上，因为它在函数推出后依然可以通过包一级的global变量找到，虽然他是在函数内部定义的：也就是说局部变量x从f函数内部逃逸了。相反，但g函数返回时，变量*y是不可达的，是可以马上被回收的，因此\*y并没有从g函数中逃逸，编译器可以选择在栈上分配\*y的存储空间。

​		如果将指向短生命周期对象的指针保存到具有长生命周期的对象中，特别是保存到全局变量时，会阻止对短声明周期对象的垃圾回收(从而可能会影响性能)。



### 2.4	赋值

​		使用赋值语句可以更新一个变量的值，将赋值变量放在 “ = ” 左边，新的表达式放在 “ = ” 右边

```go
x = 1 							//变量的赋值
*p = true						//通过指针间接赋值
person.name = "bob"				//结构体字段赋值
count[x] = count[x] * scale 	//数组、slice、map的元素赋值
```

​		特定的二元算数运算符和赋值语句的复合操作有一个更简洁的写法，可以省去对变量表达式的重复计算：

```go
count[x] *= scale	
```

​		数值变量也可以支持 ++ 自增 和 -- 自减 语句（自增和自建是语句，而不是表达式，因此 x = i++之类的表达式是错误的）

```go
v := 1
v++		//等价于v = v + 1 	v 为 2
v--		//等价于v = v - 1	v 变成 1
```

#### 2.4.1	元组赋值

​		元组赋值是另一种形式的赋值语句，它允许更新多个变量的值。在赋值前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应的变量的值。这对于同时出现在元组赋值语句左右两边的变量很有帮助，例如交换两个变量的值：

```go
x , y = y , x
a[i] , a[j] = a[j] , a[i]
```

​		但如果表达式太复杂的话应该尽量避免使用元组赋值，因为可能会影响写法可读性。

​		有些表达式会产生多个值，比如调用一个由多个返回值的函数。当这样一个函数调用出现在元组赋值右边的表达式中时（右边不能再有其他的表达式），左侧变量的数目和右边一致。

​		和声明变量一样，可以用下划线空白标识符来丢弃不需要的值

```go
_ , err = io.Copy(dst,src)		//丢弃字节数
_ , ok  = x.(T)					//只检测类型，忽略具体值
```

#### 2.4.2	可赋值性

​		赋值语句是显式的赋值形式，但程序中很多地方会发生隐式的赋值行为：函数调用会隐式的将调用参数的值赋值给函数的参数变量，一个返回语句将隐式的将返回操作的值赋值给结果变量，一个复合类型的字面量也会产生赋值行为，例如：

```go
medals := []string{"glod","silver","bronze"}
```

​		隐式的对slice的每个元素进行赋值操作，类似于：

```go
medals[0] = "glod"
medals[1] = "silver"
medals[2] = "bronze"
```

​		map和channel的元素，虽然不是普通变量，但也有类似的隐式赋值行为，只有右边的值对于左边变量是可赋值的，赋值语句才被允许。

### 2.5	类型

​		变量或表达式的类型定义了对应存储值的属性特征，在任何程序中都会存在一些变量有着相同的内部结构，但是却表示完全不同的概念。

​		一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构。新命名的类型提供了一个方法，用来隔绝不同概念的类型，这样即使他们底层类型相同也是不兼容的。

```go
type 类型名字 底层类型
```

​		类型声明语句一般出现在包一级，因此如果创建的新类型首字母是大写的，则在外部包也可以使用。

​		类型转换不会改变自身的值，但是会使他们的语义发生变化。

​		对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转换为T类型（如果T是指针类型，可能需要用小括号包装T，比如(*int)(0)）。只有当两个类型的底层数据类型相同时，才允许这种转型操作，或者是两者都是指向相同底层结构的指针类型，这些转换只会改变值类型而不会影响值本身。如果x是可赋值给T类型的值，那么x也必然可以转换为T类型。

​		数值类型之间的转换也是允许的，并且字符串和一些特定的slice之前也是可以转换的，这类转换可能改变值的表现。

​		底层数据类型决定了内部结构和表达方式，也决定是否可以像底层类型一样对内置运算符的支持。

​		命名类型还可以为该类型的值定义新的行为。这些行为表示为一组关联到该类型的函数集合，称之为类型的方法集。

### 2.7	作用域

​		一个声明语句将程序中的实体和一个名字关联，比如一个函数或一个变量。声明语句的作用域是指源代码中可以有效使用这个名字的范围。

​		作用域与生命周期不同:声明语句的作用域对应的是源代码中的文本区域：它是一个编译时的属性

​												一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内，它可以被程序的其他部分所引用：是运行时的概念

​		语法快是由花括号所包含的一系列语句，就像函数体或循环体花括号对应的语法块那样。语法块内部声明的名字时无法被外部语法块所访问的。语法决定了内部声明的名字的作用范围。

​		语法块可以包含其他类似组批量声明等没有用花括号所包含的代码，我们称之为语法块。有一个语法块为整个源代码，成为全局语法块；然后是每个包的包级语法块；每个for、if和switch的语法块；每个switch和select分支也有独立的语法块；当然也包括显示书写的语法块（花括号包含的语句）

​		声明语句对应的词法域决定了作用域范围的大小。对于内置的类型、函数和变量，比如int、len和true等实在全局作用域的，因此可以在整个程序中直接使用。任何在函数外部（也就是包级语法域）声明的名字可以在同一个包的任何源文件中访问。对于导入的包，则是对应源文件级的作用域，因此只能在当前的文件中访问，当前包的其他源文件无法访问在当前源文件导入的包。还有许多声明则是局部作用域，他只能在函数内部（甚至只能是局部的某些部分）访问。

​		控制流标号，就是break、continue或goto语句后面跟着的那种标号，则是函数级的作用域。

​		一个程序可能包含多个同名的声明，只要它们在不同的词法域就没有关系。

​		当编译器遇到一个名字引用时，如果它看起来像一个声明，它首先从最内层的词法域向全局的作用域查找。如果查找失败，则报告“未声明的名字”这样的错。如果该名字在内部和外部的块分别声明过，则内部块的声明首先被找到。在这种情况下，内部声明屏蔽了外部同名的声明，让外部的声明的名字无法被访问。

```go
func f(){}

var g = "g"

func main(){
	f := "f"
	fmt.Println(f) // "f"; local var f shadows package-level func f  局部变量f屏蔽包级函数f
	fmt.Println(g) // "g"; package-level var						 包级变量
	fmt.Println(h) // compile error: undefined: h					 编译错误：未定义：h
}
```

​		在函数中词法域可以深度嵌套，因此内部的一个声明可能屏蔽外部的声明。还有许多语法块是if或for等控制流语句构造的。下面的代码有三个不同的变量x，因为它们是定义在不同的词法域

```go
func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)	每次迭代一个字母
		}
	}
}
```

​		在x[i]和x + 'A' - 'a'声明语句的初始化的表达式中都引用了外部作用域声明的x变量。

​		并不是所有的词法域都显式地对应到由花括弧包含的语句；还有一些隐含的规则。上面的for语句创建了两个词法域：花括弧包含的是显式的部分是for的循环体部分词法域，另外一个隐式的部分则是循环的初始化部分，比如用于迭代变量i的初始化。隐式的词法域部分的作用域还包含条件测试部分和循环后的迭代部分（i++），当然也包含循环体词法域。

​		下面的例子同样有三个不同的x变量，每个声明在不同的词法域，一个在函数体词法域，一个在for隐式的初始化词法域，一个在for循环体词法域；只有两个块是显式创建的：

```go
func main() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
}
```

​		和for循环类似，if和switch语句也会在条件部分创建隐式词法域，还有它们对应的执行体词法域。下面的if-else测试链演示了x和y的有效作用域范围

```go
if x := f(); x == 0 {
	fmt.Println(x)
} else if y := g(x); x == y {
	fmt.Println(x, y)
} else {
	fmt.Println(x, y)
}
fmt.Println(x, y) // compile error: x and y are not visible here
```

​		第二个if语句嵌套在第一个内部，因此第一个if语句条件初始化词法域声明的变量在第二个if中也可以访问。switch语句的每个分支也有类似的词法域规则：条件部分为一个隐式词法域，然后每个是每个分支的词法域。

​		在包级别，声明的顺序并不会影响作用域范围，因此一个先声明的可以引用它自身或者是引用后面的一个声明，这可以让我们定义一些相互嵌套或递归的类型或函数。但是如果一个变量或常量递归引用了自身，则会产生编译错误

```go
if f, err := os.Open(fname); err != nil { // compile error: unused: f
	return err
}
f.ReadByte() // compile error: undefined f
f.Close() // compile error: undefined f
```

​		变量f的作用域只有在if语句内，因此后面的语句将无法引入它，这将导致编译错误。你可能会收到一个局部变量f没有声明的错误提示，具体错误信息依赖编译器的实现。通常需要在if之前声明变量，这样可以确保后面的语句依然可以访问变量

```go
f, err := os.Open(fname)
if err != nil {
	return err
}
f.ReadByte()
f.Close()
```

​		你可能会考虑通过将ReadByte和Close移动到if的else块来解决这个问题：

```go
if f, err := os.Open(fname); err != nil {
	return err
} else {
	// f and err are visible here too
	f.ReadByte()
	f.Close()
}
```

​		要特别注意短变量声明语句的作用域范围，考虑下面的程序，它的目的是获取当前的工作目录然后保存到一个包级的变量中。这可以本来通过直接调用os.Getwd完成，但是将这个从主逻辑中分离出来可能会更好，特别是在需要处理错误的时候。函数log.Fatalf用于打印日志信息，然后调用os.Exit(1)终止程序。

```go
var cwd string
func init() {
	cwd, err := os.Getwd() // compile error: unused: cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
```

​		虽然cwd在外部已经声明过，但是 := 语句还是将cwd和err重新声明为新的局部变量。因为内部声明的cwd将屏蔽外部的声明，因此上面的代码并不会正确更新包级声明的cwd变量。

​		由于当前的编译器会检测到局部声明的cwd并没有本使用，然后报告这可能是一个错误，但是这种检测并不可靠。因为一些小的代码变更，例如增加一个局部cwd的打印语句，就可能导致这种检测失效。

```go
var cwd string
func init() {
	cwd, err := os.Getwd() // NOTE: wrong!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
```

​		全局的cwd变量依然是没有被正确初始化的，而且看似正常的日志输出更是让这个BUG更加隐晦。有许多方式可以避免出现类似潜在的问题。最直接的方法是通过单独声明err变量，来避免使用 := 的简短声明方式：

```go
var cwd string
func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
```

## 3	基础数据类型

Go语言将数据类型分为四类：基础类型、复合类型、引用类型和接口类型。本章介绍基础类型，包括： 数字、字符串和布尔型。复合数据类型——数组和结构体——是通过组合简单类 型，来表达更加复杂的数据结构。引用类型包括指针、切片字典、函 数、通道，虽然数据种类很多，但它们都是对程序中一个变量或状态的间接引用。这意 味着对任一引用类型数据的修改都会影响所有该引用的拷贝

### 3.1	整型

​		Go语言的数值类型包括几种不同大小的整形数、浮点数和复数。每种数值类型都决定了对应的大小范围 和是否支持正负符号。

​		Go语言同时提供了有符号和无符号类型的整数运算。这里有int8、int16、int32和int64四种截然不同大 小的有符号整形数类型，分别对应8、16、32、64bit大小的有符号整形数，与此对应的是uint8、 uint16、uint32和uint64四种无符号整形数类型。

​		还有两种一般对应特定CPU平台机器字大小的有符号和无符号整数int和uint；其中int是应用最广泛 的数值类型。这两种类型都有同样的大小，32或64bit，为不同的编译器即使在相同的硬件平台上可能产生不同的大小。

​		Unicode字符rune类型是和int32等价的类型，通常用于表示一个Unicode码点。这两个名称可以互换使 用。同样byte也是uint8类型的等价类型，byte类型一般用于强调数值是一个原始的数据而不是一个小的整数。

​		还有一种无符号的整数类型uintptr，没有指定具体的bit大小但是足以容纳指针。uintptr类型只有在底层编程是才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。

​		int、uint和uintptr是不同类型的兄弟类型。其中int和int32也是不同的类型， 即使int的大小也是32bit，在需要将int当作int32类型的地方需要一个显式的类型转换操作，反之亦然。

​		其中有符号整数采用2的补码形式表示，也就是最高bit位用作表示符号位，一个n-bit的有符号数的值域 是从 -2^{n-1} 到 −1。例如，int8类型整数的值域是从-128到127，而uint8类型整数的值域是从0到255。

​		下面是Go语言中关于算术运算、逻辑运算和比较运算的二元运算符，它们按照先级递减的顺序的排列：

```
* 	/	 %	 << 	>>	 &	 &^
+	-	 |	 ^ 
==  != 	 <   <= 	>	 >= 
&& 
||
```

​		二元运算符有五种优先级。在同一个优先级，使用左优先结合规则，但是使用括号可以明确优先顺序， 使用括号也可以用于提升优先级，例如mask & (1 << 28)。 对于上表中前两行的运算符，例如+运算符还有一个与赋值相结合的对应运算符+=，可以用于简化赋值语 句。

​		算术运算符+、-、*和/可以适用与于整数、浮点数和复数，但是取模运算符%仅用于整数间的运算。对于 不同编程语言，%取模运算的行为可能并不相同。在Go语言中，%取模运算符的符号和被取模数的符号总 是一致的，因此-5%3和-5%-3结果都是-2。除法运算符/的行为则依赖于操作数是否为全为整数，比如 5.0/4.0的结果是1.25，但是5/4的结果是1，因为整数除法会向着0方向截断余数。

​		如果一个算术运算的结果，不管是有符号或者是无符号的，如果需要更多的bit位才能正确表示的话，就 说明计算结果是溢出了。超出的高位的bit位部分将被丢弃。如果原始的数值是有符号类型，而且最左边 的bit为是1的话，那么最终结果可能是负的，例如：

```go
var u uint8 = 255 
fmt.Println(u, u+1, u*u) // "255 0 1" var i int8 = 127
var i int8 = 127
fmt.Println(i, i+1, i*i) // "127 -128 1"
```

​		两个相同的整数类型可以使用下面的二元比较运算符进行比较；比较表达式的结果是布尔类型。

```go
== 		equal to 
!= 		not equal to
< 		less than 
<= 		less than or equal to 
>		greater than 
>=	    greater than or equal to
```

​		布尔型、数字类型和字符串等基本类型都是可比较的，也就是说两个相同类型的值可以用 ==  和 != 进行比较。整数、浮点数和字符串可以根据比较结果排序。许多其它类型的值可能是不可比 较的，因此也就可能是不可排序的。对于我们遇到的每种类型，我们需要保证规则的一致性。

​		这里是一元的加法和减法运算符：

```go
+	 一元加法 (无效果)
-	 负数
```

​		对于整数，+x是0+x的简写，-x则是0-x的简写；对于浮点数和复数，+x就是x，-x则是x 的负数。

​		 Go语言还提供了以下的bit位操作运算符，前面4个操作运算符并不区分是有符号还是无符号数：

```go
& 		位运算 AND 
|		位运算 OR 
^ 		位运算 XOR 
&^ 		位清空 (AND NOT) 
<< 		左移 
>> 		右移
```

​		位操作运算符^作为二元运算符时是按位异或（ XOR），当用作一元运算符时表示按位取反；也就是说， 它返回一个每个bit位都取反的数。位操作运算符&^用于按位置零（ AND NOT）：表达式z = x &^ y结果z 的bit位为0，如果对应y中bit位为1的话，否则对应的bit位等于x相应的bit位的值。

​		下面的代码演示了如何使用位操作解释uint8类型值的8个独立的bit位。它使用了Printf函数的%b参数打 印二进制格式的数字；其中%08b中08表示打印至少8个字符宽度，不足的前缀部分用0填充。

```go
var x uint8 = 1<<1 | 1<<5 
var y uint8 = 1<<1 | 1<<2 
fmt.Printf("%08b\n", x)		 // "00100010", the set {1, 5} 
fmt.Printf("%08b\n", y)		 // "00000110", the set {1, 2}
fmt.Printf("%08b\n", x&y)	 // "00000010", the intersection {1} 
fmt.Printf("%08b\n", x|y) 	 // "00100110", the union {1, 2, 5}
fmt.Printf("%08b\n", x^y) 	 // "00100100", the symmetric difference {2, 5} 
fmt.Printf("%08b\n", x&^y)	 // "00100000", the difference {5}
for i := uint(0); i < 8; i++ { 
    if x&(1<<i) != 0 { // membership test 
        fmt.Println(i) // "1", "5" 
    } 
}
fmt.Printf("%08b\n", x<<1)	 // "01000100", the set {2, 6} 
fmt.Printf("%08b\n", x>>1)	 // "00010001", the set {0, 4}
```

​		在x<<n和x>>n移位运算中，决定了移位操作bit数部分必须是无符号数；被操作的x数可以是有符号或无 符号数。算术上，一个x<<n左移运算等价于乘以 2^n 。

​		左移运算用零填充右边空缺的bit位，无符号数的右移运算也是用0填充左边空缺的bit位，但是有符号数 的右移运算会用符号位的值填充左边空缺的bit位。

​		内置的len函数 返回一个有符号的int，我们可以像下面例子那样处理逆序循环。

```go
medals := []string{"gold", "silver", "bronze"}
for i := len(medals) - 1; i >= 0; i-- { 
	fmt.Println(medals[i]) // "bronze", "silver", "gold" 
}
```

​		如果len函数返回一个无符号数，那么i也将是无符号的 uint类型，然后条件i >= 0则永远为真。在三次迭代之后，也就是i == 0时，i--语句将不会产生-1，而 是变成一个uint类型的最大值（可能是 2^64-16 4−1），然后medals[i]表达式将发生运行时panic异常（§5.9），也就是试图访问一个slice范围以外的元素。

​		出于这个原因，无符号数往往只有在位运算或其它特殊的运算场景才会使用，就像bit集合、分析二进制 文件格式或者是哈希和加密操作等。它们通常并不用于仅仅是表达非负数量的场合。

​		需要一个显式的转换将一个值从一种类型转化位另一种类型，并且算术和逻辑运算的二元操 作中必须是相同的类型。虽然这偶尔会导致需要很长的表达式，但是它消除了所有和类型相关的问题， 而且也使得程序容易理解。

```go
var apples int32 = 1 
var oranges int16 = 2 
var compote int = apples + oranges // compile error：	invalid operation: apples + oranges (mismatched types int32 and int16)
```

​		这种类型不匹配的问题可以有几种不同的方法修复，最常见方法是将它们都显式转型为一个常见类型

```go
var compote = int(apples) + int(oranges)
```

​		对于每种类型T，如果转换允许的话，类型转换操作T(x)将x转换为T类型。许多整形数之 间的相互转换并不会改变数值；它们只是告诉编译器如何解释这个值。但是对于将一个大尺寸的整数类 型转为一个小尺寸的整数类型，或者是将一个浮点数转为整数，可能会改变数值或丢失精度：

```go
f := 3.141				 // a float64 
i := int(f) 
fmt.Println(f, i)		 // "3.141 3" f = 1.99
fmt.Println(int(f)) 	 // "1"
```

​		浮点数到整数的转换将丢失任何小数部分，然后向数轴零方向截断。你应该避免对可能会超出目标类型 表示范围的数值类型转换，因为截断的行为可能依赖于具体的实现：

```go
f := 1e100 // a float64 
i := int(f) // 结果依赖于具体实现
```

​		任何大小的整数字面值都可以用以0开始的八进制格式书写，例如0666；或用以0x或0X开头的十六进制格 式书写，例如0xdeadbeef。十六进制数字可以用大写或小写字母。如今八进制数据通常用于POSIX操作系 统上的文件访问权限标志，十六进制数字则更强调数字值的bit位模式。

​		当使用fmt包打印一个数值时，我们可以用%d、%o或%x参数控制输出的进制格式：

```go
o := 0666
fmt.Printf("%d %[1]o %#[1]o\n", o)		 // "438 	666  0666" 
x := int64(0xdeadbeef)
fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x) 
// Output: 
// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
```

​		通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操 作数，但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。第二，%后的#副词告诉Printf在 用%o、%x或%X输出时生成0、0x或0X前缀。

​		字符面值通过一对单引号直接包含对应字符。最简单的例子是ASCII中类似'a'写法的字符面值，但是我 们也可以通过转义的数值来表示任意的Unicode码点对应的字符，字符使用%c参数打印，或者是用%q参数打印带单引号的字符：

```go
ascii := 'a' 
unicode := '国' 
newline := '\n' 
fmt.Printf("%d %[1]c %[1]q\n", ascii)		 // "97 a 'a'" 
fmt.Printf("%d %[1]c %[1]q\n", unicode) 	 // "22269 国 '国'" 
fmt.Printf("%d %[1]q\n", newline) 			 // "10 '\n'"
```

### 3.2	浮点数

​		Go语言提供了两种精度的浮点数，float32和float64。它们的算术规范由IEEE754浮点数国际标准定义， 该浮点数规范被所有现代的CPU支持。

​		这些浮点数类型的取值范围可以从很微小到很巨大。浮点数的范围极限值可以在math包找到。常量 math.MaxFloat32表示float32能表示的最大数值，大约是 3.4e38；对应的math.MaxFloat64常量大约是 1.8e308。它们分别能表示的最小值近似为1.4e-45和4.9e-324。

​		一个float32类型的浮点数可以提供大约6个十进制数的精度，而float64则可以提供约15个十进制数的精 度；通常应该优先使用float64类型，因为float32类型的累计计算误差很容易扩散，并且float32能精确 表示的正整数并不是很大（因为float32的有效bit位只有23个，其它的bit位用于指数和符号；当 整数大于23bit能表达的范围时，float32的表示将出现误差）：

```go
var f float32 = 16777216 	// 1 << 24 
fmt.Println(f == f+1) 		// "true"!
```

​		浮点数的字面值可以直接写小数部分：

```go
const e = 2.71828 // (approximately)
```

​		小数点前面或后面的数字都可能被省略（例如.707或1.）。很小或很大的数最好用科学计数法书写，通 过e或E来指定指数部分：

```go
const Avogadro = 6.02214129e23 	// 阿伏伽德罗常数 
const Planck = 6.62606957e-34 	// 普朗克常数
```

​		用Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，但是对应表格的数据，使用%e（带指数）或%f的形式打印可能更合适。所有的这三个打印形式都可以指定打印的宽度和控制打印精度。

```go
for x := 0; x < 8; x++ {
	fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
}
```

​		上面代码打印e的幂，打印精度是小数点后三个小数精度和8个字符宽度：

```go
x = 0	 e^x = 1.000
x = 1 	 e^x = 2.718
x = 2	 e^x = 7.389
x = 3	 e^x = 20.086
x = 4	 e^x = 54.598
x = 5 	 e^x = 148.413
x = 6 	 e^x = 403.429
x = 7 	 e^x = 1096.633
```

​		math包中除了提供大量常用的数学函数外，还提供了IEEE754浮点数标准中定义的特殊值的创建和测试：正无穷大和负无穷大，分别用于表示太大溢出的数字和除零的结果；还有NaN非数，一般用于表示无效的除法操作结果0/0或Sqrt(-1).		

```GO
var z float64
fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
```

​		函数math.IsNaN用于测试一个数是否是非数NaN，math.NaN则返回非数对应的值。虽然可以用math.NaN来表示一个非法的结果，但是测试一个结果是否是非数NaN则是充满风险的，因为NaN和任何数都是不相等的（在浮点数中，NaN、正无穷大和负无穷大都不是唯一的，每个都有非常多种的bit模式表示）：

```GO
nan := math.NaN()
fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"

NA： 表示缺失值(Missing value)，是“Not Available”的缩写
Inf：表示无穷大，是“Infinite”的缩写
NaN：表示非数值，是“Not a Number”的缩写
```

​		如果一个函数返回的浮点数结果可能失败，最好的做法是用单独的标志报告失败：

```go
func compute() (value float64, ok bool) {
	// ...
	if failed {
		return 0, false
	}
	return result, true
}
```

### 3.3	复数

​		Go语言提供了两种精度的复数类型：complex64和complex128，分别对应float32和float64两种浮点数精度。内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部：

```go
var x complex128 = complex(1, 2)	 // 1+2i
var y complex128 = complex(3, 4)	 // 3+4i
fmt.Println(x*y)					 // "(-5+10i)"
fmt.Println(real(x*y)) 				 // "-5"
fmt.Println(imag(x*y))				 // "10"
```

​		如果一个浮点数面值或一个十进制整数面值后面跟着一个i，例如3.141592i或2i，它将构成一个复数的虚部，复数的实部是0：

```go
fmt.Println(1i * 1i) 		// "(-1+0i)", i^2 = -1
```

​		在常量算术规则下，一个复数常量可以加到另一个普通数值常量（整数或浮点数、实部或虚部），我们可以用自然的方式书写复数，就像1+2i或与之等价的写法2i+1。上面x和y的声明语句还可以简化：

```go
x := 1 + 2i
y := 3 + 4i
```

​		复数也可以用==和!=进行相等比较。只有两个复数的实部和虚部都相等的时候它们才是相等的

​		math/cmplx包提供了复数处理的许多函数，例如求复数的平方根函数和求幂函数。

```go
fmt.Println(cmplx.Sqrt(-1))		 // "(0+1i)"
```

### 3.4	bool类型

​		一个布尔类型的值只有两种：true和false。if和for语句的条件部分都是布尔类型的值，并且==和<等比较操作也会产生布尔型的值。一元操作符!对应逻辑非操作，因此!true的值为false。

​		布尔值可以和&&（AND）和||（OR）操作符结合，并且可能会有短路行为：如果运算符左边值已经可以确定整个布尔表达式的值，那么运算符右边的值将不在被求值，因此下面的表达式总是安全的：

```go
s != "" && s[0] == 'x
其中s[0]操作如果应用于空字符串将会导致panic异常。
```

​		因为&&的优先级比||高（助记：&&对应逻辑乘法，||对应逻辑加法，乘法比加法优先级要高），下面形式的布尔表达式是不需要加小括弧的：

```go
if 'a' <= c && c <= 'z' ||
	'A' <= c && c <= 'Z' ||
	'0' <= c && c <= '9' {
	// ...ASCII letter or digit...
}
```

​		布尔值并不会隐式转换为数字值0或1，反之亦然。必须使用一个显式的if语句辅助转换：

```go
i := 0
if b {
	i = 1
}
```

```go
// btoi returns 1 if b is true and 0 if false.
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
```

```go
// itob reports whether i is non-zero.
func itob(i int) bool { 
	return i != 0 
	}
```

### 3.5	字符串

​		一个字符串是一个不可改变的字节序列。字符串可以包含任意的数据，包括byte值0，但是通常是用来包含人类可读的文本。文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。

​		内置的len函数可以返回一个字符串中的字节数目（不是rune字符数目），索引操作s[i]返回第i个字节的字节值，i必须满足0 ≤ i< len(s)条件约束。

```go
s := "hello, world"
fmt.Println(len(s))				 // "12"
fmt.Println(s[0], s[7]) 		 // "104 119" ('h' and 'w')
```

​		如果试图访问超出字符串索引范围的字节将会导致panic异常：

```go
c := s[len(s)]     // panic: index out of range
```

​		第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个字节

​		子字符串操作s[i:j]基于原始的s字符串的第i个字节开始到第j个字节（并不包含j本身）生成一个新字符串。生成的新字符串将包含j-i个字节。

```go
fmt.Println(s[0:5]) 		// "hello"
```

​		同样，如果索引超出字符串范围或者j小于i的话将导致panic异常。

​		不管i还是j都可能被忽略，当它们被忽略时将采用0作为开始位置，采用len(s)作为结束的位置。

```go
fmt.Println(s[:5])  	// "hello"
fmt.Println(s[7:])	    // "world"
fmt.Println(s[:])	    // "hello, world"
```

​		其中+操作符将两个字符串链接构造一个新字符串：

```go
fmt.Println("goodbye" + s[5:]) // "goodbye, world"
```

​		字符串可以用==和<进行比较；比较通过逐个字节比较完成的，因此比较的结果是字符串自然编码的顺序。

​		字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变，当然我们也可以给一个字符串变量分配一个新字符串值。可以像下面这样将一个字符串追加到另一个字符串：

```go
s := "left foot"
t := s
s += ", right foot"
```

​		这并不会导致原始的字符串值被改变，但是变量s将因为+=语句持有一个新的字符串值，但是t依然是包含原先的字符串值。

```go
fmt.Println(s) 	// "left foot, right foot"
fmt.Println(t) 	// "left foot"
```

​		因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的：

```go
s[0] = 'L'	 // compile error: cannot assign to s[0]
```

​		不变性意味如果两个字符串共享相同的底层数据的话也是安全的，这使得复制任何长度的字符串代价是低廉的。同样，一个字符串s和对应的子字符串切片s[7:]的操作也可以安全地共享相同的内存，因此字符串切片操作代价也是低廉的。在这两种情况下都没有必要分配新的内存。

#### 3.5.1	字符串面值

​		字符串值也可以用字符串面值方式编写，只要将一系列字节序列包含在双引号即可：

```
"Hello, 世界"
```

​		在一个双引号包含的字符串面值中，可以用以反斜杠\开头的转义序列插入任意的数据。下面的换行、回车和制表符等是常见的ASCII控制代码的转义方式：

```go
\a 	 响铃
\b 	 退格
\f	 换页
\n	 换行
\r	 回车
\t 	 制表符
\v	 垂直制表符
\'   单引号 (只用在 '\'' 形式的rune符号面值中)
\"	 双引号 (只用在 "..." 形式的字符串面值中)
\\	 反斜杠
```

​		可以通过十六进制或八进制转义在字符串面值包含任意的字节。一个十六进制的转义形式是\xhh，其中两个h表示十六进制数字（大写或小写都可以）。一个八进制转义形式是\ooo，包含三个八进制的o数字（0到7），但是不能超过\377（译注：对应一个字节的范围，十进制为255）。每一个单一的字节表达一个特定的值。

​		一个原生的字符串面值形式是...，使用反引号代替双引号。在原生的字符串面值中，没有转义操作；全部的内容都是字面的意思，包含退格和换行，因此一个程序中的原生字符串面值可能跨越多行（译注：在原生字符串面值内部是无法直接写字符的，可以用八进制或十六进制转义或+"```"链接字符串常量完成）。唯一的特殊处理是会删除回车以保证在所有平台上的值都是一样的，包括那些把回车也放入文本文件的系统。

​		原生字符串面值用于编写正则表达式会很方便，因为正则表达式往往会包含很多反斜杠。原生字符串面值同时被广泛应用于HTML模板、JSON面值、命令行提示信息以及那些需要扩展到多行的场景。

```go
const GoUsage = `Go is a tool for managing Go source code.
Usage:
go command [arguments]
...`
```

#### 3.5.2	Unicode

​		ASCII字符集：美国信息交换标准代码。使用7bit来表示128个字符：包含英文字母的大小写、数字、各种标点符号和设置控制符

​		Unicode:它收集了这个世界上所有的符号系统，包括重音符号和其它变音符号，制表符和回车符，还有很多神秘的符号，每个符号都分配一个唯一的Unicode码点，Unicode码点对应Go语言中的rune整数类型（rune是int32等价类型）,通用的表示一个Unicode码点的数据类型是int32,每个Unicode码点都使用同样的大小32bit来表示。这种方式比较简单统一，但是它会浪费很多存储空间。

#### 3.5.3	UTF-8

​		UTF8是一个将Unicode码点编码为字节序列的变长编码。UTF8编码使用1到4个字节来表示每个Unicode码点，ASCII部分字符只使用1个字节，常用字符部分使用2或3个字节表示。每个符号编码后第一个字节的高端bit位用于表示总共有多少编码个字节。如果第一个字节的高端bit为0，则表示对应7bit的ASCII字符，ASCII字符每个字符依然是一个字节，和传统的ASCII编码兼容。如果第一个字节的高端bit是110，则说明需要2个字节；后续的每个高端bit都以10开头。更大的Unicode码点也是采用类似的策略处理。

```go
0xxxxxxx		  runes 0-127 			   			  (ASCII)
110xxxxx 10xxxxxx 128-2047							  (values <128 unused)
1110xxxx 10xxxxxx 10xxxxxx 2048-65535				  (values <2048 unused)
11110xxx 10xxxxxx 10xxxxxx 10xxxxxx 65536-0x10ffff 	  (other values unused)
```

​		变长的编码无法直接通过索引来访问第n个字符，但是UTF8编码获得了很多额外的优点。首先UTF8编码比较紧凑，完全兼容ASCII码，并且可以自动同步：没有任何字符的编码是其它字符编码的子串，或是其它编码序列的字串，因此搜索一个字符时只要搜索它的字节编码序列即可，不用担心前后的上下文会对搜索结果产生干扰。同时UTF8编码的顺序和Unicode码点的顺序一致，因此可以直接排序UTF8编码序列。同时因为没有嵌入的NUL(0)字节，可以很好地兼容那些使用NUL作为字符串结尾的编程语言。

​		Go语言字符串面值中的Unicode转义字符让我们可以通过Unicode码点输入特殊的字符。有两种形式：\uhhhh对应16bit的码点值，\Uhhhhhhhh对应32bit的码点值，其中h是一个十六进制数字；一般很少需要使用32bit的形式。每一个对应码点的UTF8编码。例如：下面的字母串面值都表示相同的值：

```
"世界"
"\xe4\xb8\x96\xe7\x95\x8c"
"\u4e16\u754c"
"\U00004e16\U0000754c"
```

​		上面三个转义序列都为第一个字符串提供替代写法，但是它们的值都是相同的。

​		Unicode转义也可以使用在rune字符中。下面三个字符是等价的：

```
'世' '\u4e16' '\U00004e16'
```

​		对于小于256码点值可以写在一个十六进制转义字节中，例如'\x41'对应字符'A'，但是对于更大的码点则必须使用\u或\U转义形式。因此，'\xe4\xb8\x96'并不是一个合法的rune字符，虽然这三个字节对应一个有效的UTF8编码的码点。

​		测试一个字符串是否是另一个字符串的前缀：

```go
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
```

​		或者是后缀测试：

```go
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
```

​		或者是包含子串测试：

```go
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
```

​		对于UTF8编码后文本的处理和原始的字节处理逻辑是一样的。但是对应很多其它编码则并不是这样的。

​		字符串包含13个字节，以UTF8形式编码，但是只对应9个Unicode字符：

```go
import "unicode/utf8"
s := "Hello, 世界"
fmt.Println(len(s))					   // "13"
fmt.Println(utf8.RuneCountInString(s)) // "9"
```

​		为了处理这些真实的字符，我们需要一个UTF8解码器。unicode/utf8包提供了该功能，我们可以这样使用：

```go
for i := 0; i < len(s); {
	r, size := utf8.DecodeRuneInString(s[i:])
	fmt.Printf("%d\t%c\n", i, r)
	i += size
}
```

​		每一次调用DecodeRuneInString函数都返回一个r和长度，r对应字符本身，长度对应r采用UTF8编码后的编码字节数目。长度可以用于更新第i个字符在字符串中的字节索引位置。Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串。

```go
for i, r := range "Hello, 世界" {
	fmt.Printf("%d\t%q\t%d\n", i, r, r)
}
```

​		我们可以使用一个简单的循环来统计字符串中字符的数目，像这样：

```go
n := 0
for _, _ = range s {
	n++
}
```

​		像其它形式的循环那样，我们也可以忽略不需要的变量：

```go
n := 0
for range s {
	n++
}	
```

​		或者我们可以直接调用utf8.RuneCountInString(s)函数。

​		每一个UTF8字符解码，不管是显式地调用utf8.DecodeRuneInString解码或是在range循环中隐式地解码，如果遇到一个错误的UTF8编码输入，将生成一个特别的Unicode字符'\uFFFD'，在印刷中这个符号通常是一个黑色六角或钻石形状，里面包含一个白色的问号（?）。当程序遇到这样的一个字符，通常是一个危险信号，说明输入并不是一个完美没有错误的UTF8字符串。

​		UTF8字符串作为交换格式是非常方便的，但是在程序内部采用rune序列可能更方便，因为rune大小一致，支持数组索引和方便切割。

​		string接受到[]rune的类型转换，可以将一个UTF8编码的字符串解码为Unicode字符序列：

```go
// "program" in Japanese katakana
s := "プログラム"
fmt.Printf("% x\n", s) 		 // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
r := []rune(s)
fmt.Printf("%x\n", r)		 // "[30d7 30ed 30b0 30e9 30e0]
```

​		如果是将一个[]rune类型的Unicode字符slice或数组转为string，则对它们进行UTF8编码：

```
fmt.Println(string(r)) // "プログラム"
```

​		将一个整数转型为字符串意思是生成以只包含对应Unicode码点字符的UTF8字符串：

```
fmt.Println(string(65)) 		 // "A", not "65"
fmt.Println(string(0x4eac))	 	 // "京
```

​		如果对应码点的字符是无效的，则用'\uFFFD'无效字符作为替换：

```
fmt.Println(string(1234567))	 // "(?)"
```

#### 3.5.4	字符串和Byte切片

​		标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。strings包提供了许 多如字符串的查询、替换、比较、截断、拆分和合并等功能。

​		bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。在这种情况下，使用bytes.Buffer类型将会更有效。

​		strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。

​		unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。每个函数有一个单一的rune类型的参数，然后返回一个布尔值。而像ToUpper和ToLower之类的转换函数将用于rune字符的大小写转换。所有的这些函数都是遵循Unicode标准定义的字母、数字等分类规范。strings包也有类似的函数，它们是ToUpper和ToLower，将原始字符串的每个字符都做相应的转换，然后返回新的字符串。

​		一个字符串是包含的只读字节数组，一旦创建，是不可变的。相比之下，一个字节slice的元素则可以自由地修改。

​		字符串和字节slice之间可以相互转换：

```go
s := "abc"
b := []byte(s)
s2 := string(b)
```

​		从概念上讲，一个[]byte(s)转换是分配了一个新的字节数组用于保存字符串数据的拷贝，然后引用这个底层的字节数组。编译器的优化可以避免在一些场景下分配和复制字符串数据，但总的来说需要确保在变量b被修改的情况下，原始的s字符串也不会改变。将一个字节slice转到字符串的string(b)操作则是构造一个字符串拷贝，以确保s2字符串是只读的。

​		为了避免转换中不必要的内存分配，bytes包和strings同时提供了许多实用函数。下面是strings包中的六个函数：

```go
func Contains(s, substr string) bool			//包含
func Count(s, sep string) int					//统计包含
func Fields(s string) []string					//转换为string类型切片
func HasPrefix(s, prefix string) bool			
func Index(s, sep string) int					//统计第一次出现位置
func Join(a []string, sep string) string
```

​		bytes包中也对应的六个函数：

```go
func Contains(b, subslice []byte) bool
func Count(s, sep []byte) int
func Fields(s []byte) [][]byte
func HasPrefix(s, prefix []byte) bool
func Index(s, sep []byte) int
func Join(s [][]byte, sep []byte) []byte
```

​		它们之间唯一的区别是字符串类型参数被替换成了字节slice类型的参数。

​		bytes包还提供了Buffer类型用于字节slice的缓存。一个Buffer开始是空的，但是随着string、byte或[]byte等类型数据的写入可以动态增长，一个bytes.Buffer变量并不需要处理化，因为零值也是有效的：

```go
gopl.io/ch3/printints
// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
```

​		当向bytes.Buffer添加任意字符的UTF8编码时，最好使用bytes.Buffer的WriteRune方法，但是WriteByte方法对于写入类似'['和']'等ASCII字符则会更加有效。

#### 3.5.5	 字符串和数字的转换

​		除了字符串、字符、字节之间的转换，字符串和数值之间的转换也比较常见。由strconv包提供这类转换功能。

​		将一个整数转为字符串，一种方法是用fmt.Sprintf返回一个格式化的字符串；另一个方法是用strconv.Itoa(“整数到ASCII”)：

```go
x := 123
y := fmt.Sprintf("%d", x)
fmt.Println(y, strconv.Itoa(x))			    // "123 123"
```

​		FormatInt和FormatUint函数可以用不同的进制来格式化数字：

```go
fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
```

​		fmt.Printf函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多，特别是在需要包含附加额外信息的时候：

```go
s := fmt.Sprintf("x=%b", x) 			    // "x=1111011"
```

​		如果要将一个字符串解析为整数，可以使用strconv包的Atoi或ParseInt函数，还有用于解析无符号整数的ParseUint函数：

```go
x, err := strconv.Atoi("123") 			  // x is an int
y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
```

​		ParseInt函数的第三个参数是用于指定整型数的大小；例如16表示int16，0则表示int。在任何情况下，返回的结果y总是int64类型，你可以通过强制类型转换将它转为更小的整数类型。

​		有时候也会使用fmt.Scanf来解析输入的字符串和数字，特别是当字符串和数字混合在一行的时候，它可以灵活处理不完整或不规则的输入。

### 3.6	常量

​		常量表达式的值在编译期计算，而不是在运行期。每种常量的潜在类型都是基础类型：boolean、string或数字。

​		一个常量的声明语句定义了常量的名字，和变量的声明语法类似，常量的值不可修改，这样可以防止在运行期被意外或恶意的修改。例如，常量比变量更适合用于表达像π之类的数学常数，因为它们的值不会发生变化：

```go
const pi = 3.14159		 // approximately; math.Pi is a better approximation
```

​		和变量声明一样，可以批量声明多个常量；这比较适合声明一组相关的常量：

```go
const (
	e = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)
```

​		所有常量的运算都可以在编译期完成，当操作数是常量时，一些运行时的错误也可以在编译时被发现，例如整数除零、字符串索引越界、任何导致无效浮

点数的操作等。

​		常量间的所有算术运算、逻辑运算和比较运算的结果也是常量，对常量的类型转换操作或以下函数调用都是返回常量结果：len、cap、real、imag、complex和unsafe.Sizeof。

​		因为它们的值是在编译期就确定的，因此常量可以是构成类型的一部分，例如用于指定数组类型的长度：

```go
const IPv4Len = 4
// parseIPv4 parses an IPv4 address (d.d.d.d).
func parseIPv4(s string) IP {
	var p [IPv4Len]byte
	// ...
}
```

​		一个常量的声明也可以包含一个类型和一个值，但是如果没有显式指明类型，那么将从右边的表达式推断类型。在下面的代码中，time.Duration是一个命名类型，底层类型是int64，time.Minute是对应类型的常量。下面声明的两个常量都是time.Duration类型，可以通过%T参数打印类型信息：

```go
const noDelay time.Duration = 0
const timeout = 5 * time.Minute
fmt.Printf("%T %[1]v\n", noDelay) 			// "time.Duration 0"
fmt.Printf("%T %[1]v\n", timeout) 			// "time.Duration 5m0s"
fmt.Printf("%T %[1]v\n", time.Minute)	    // "time.Duration 1m0s"
```

​		如果是批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略，如果省略初始化表达式则表示使用前面常量的初始化表达式写法，对应的常量类型也一样的。例如：

```go
const (
	a = 1
	b
	c = 2
	d
)
fmt.Println(a, b, c, d)		 // "1 1 2 2"
```

​		如果只是简单地复制右边的常量表达式，其实并没有太实用的价值。但是它可以带来其它的特性，那就是iota常量生成器语法。

#### 3.6.1	iota 常量生成器

​		常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一。

```go
type Weekday int
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
//周日将对应0，周一为1，如此等等。		
```

​		我们也可以在复杂的常量表达式中使用iota，下面是来自net包的例子，用于给一个无符号整数的最低5bit的每个bit指定一个名字：

```go
type Flags uint
const (
	FlagUp Flags = 1 << iota		 // is up
	FlagBroadcast 					 // supports broadcast access capability
	FlagLoopback 					 // is a loopback interface
	FlagPointToPoint 				 // belongs to a point-to-point link
	FlagMulticast 					 // supports multicast access capability
)
```

​		随着iota的递增，每个常量对应表达式1 << iota，是连续的2的幂，分别对应一个bit位置。使用这些常量可以用于测试、设置或清除对应的bit位的值：

```go
gopl.io/ch3/netflag

func IsUp(v Flags) bool 		 { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)			 { *v &^= FlagUp }
func SetBroadcast(v *Flags)		 { *v |= FlagBroadcast }
func IsCast(v Flags) bool 		 { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) 		// "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) 		// "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) 		// "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v))	    // "10010 true"
}
```

​		下面是一个更复杂的例子，每个常量都是1024的幂：

```go
const (
	_ = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)
```

​		不过iota常量生成规则也有其局限性。例如，它并不能用于产生1000的幂（KB、MB等），因为Go语言并没有计算幂的运算符。

#### 3.6.2	 无类型常量

​		虽然一个常量可以有任意有一个确定的基础类型，例如int或float64，或者是类似time.Duration这样命名的基础类型，但是许多常量并没有一个明确的基础类型。编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算；你可以认为至少有256bit的运算精度。这里有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。

​		通过延迟明确常量的具体类型，无类型的常量不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换。例如，例子中的ZiB和YiB的值已经超出任何Go语言中整数类型能表达的范围，但是它们依然是合法的常量，而且可以像下面常量表达式依然有效（译注：YiB/ZiB是在编译期计算出来的，并且结果常量是1024，是Go语言int变量能有效表示的）：

```go
fmt.Println(YiB/ZiB) // "1024"
```

​		另一个例子，math.Pi无类型的浮点数常量，可以直接用于任意需要浮点数或复数的地方：

```go
var x float32 = math.Pi
var y float64 = math.Pi
var z complex128 = math.Pi
```

​		如果math.Pi被确定为特定类型，比如float64，那么结果精度可能会不一样，同时对于需要float32或complex128类型值的地方则会强制需要一个明确的类型转换：

```go
const Pi64 float64 = math.Pi

var x float32 = float32(Pi64)
var y float64 = Pi64
var z complex128 = complex128(Pi64)
```

​		对于常量面值，不同的写法可能会对应不同的类型。例如0、0.0、0i和'\u0000'虽然有着相同的常量值，但是它们分别对应无类型的整数、无类型的浮点数、无类型的复数和无类型的字符等不同的常量类型。同样，true和false也是无类型的布尔类型，字符串面值常量是无类型的字符串类型。

​		前面说过除法运算符/会根据操作数的类型生成对应类型的结果。因此，不同写法的常量除法表达式可能对应不同的结果：

```go
var f float64 = 212
fmt.Println((f - 32) * 5 / 9)			 // "100"; (f - 32) * 5 is a float64
fmt.Println(5 / 9 * (f - 32))			 // "0"; 5/9 is an untyped integer, 0
fmt.Println(5.0 / 9.0 * (f - 32)) 		 // "100"; 5.0/9.0 is an untyped float
```

​		只有常量可以是无类型的。当一个无类型的常量被赋值给一个变量的时候，就像上面的第一行语句，或者是像其余三个语句中右边表达式中含有明确类型的值，无类型的常量将会被隐式转换为对应的类型，如果转换合法的话。

```go
var f float64 = 3 + 0i 			// untyped complex -> float64
f = 2						    // untyped integer -> float64
f = 1e123					    // untyped floating-point -> float64
f = 'a' 						// untyped rune -> float64
```

​		上面的语句相当于:

```go
var f float64 = float64(3 + 0i)
f = float64(2)
f = float64(1e123)
f = float64('a')
```

​		无论是隐式或显式转换，将一种类型转换为另一种类型都要求目标可以表示原始值。对于浮点数和复数，可能会有舍入处理：

```go
const (
	deadbeef = 0xdeadbeef	  // untyped int with value 3735928559
	a = uint32(deadbeef) 	  // uint32 with value 3735928559
	b = float32(deadbeef)	  // float32 with value 3735928576 (rounded up)
	c = float64(deadbeef)	  // float64 with value 3735928559 (exact)
	d = int32(deadbeef) 	  // compile error: constant overflows int32
	e = float64(1e309)  	  // compile error: constant overflows float64
	f = uint(-1)        	  // compile error: constant underflows uint
)
```

​		对于一个没有显式类型的变量声明语法（包括短变量声明语法），无类型的常量会被隐式转为默认的变量类型，就像下面的例子：

```go
i := 0		 // untyped integer; implicit int(0)
r := '\000'  // untyped rune; implicit rune('\000')
f := 0.0     // untyped floating-point; implicit float64(0.0)
c := 0i      // untyped complex; implicit complex128(0i)
```

​		注意默认类型是规则的：无类型的整数常量默认转换为int，对应不确定的内存大小，但是浮点数和复数常量则默认转换为float64和complex128。Go语言本身并没有不确定内存大小的浮点数和复数类型，而且如果不知道浮点数类型的话将很难写出正确的数值算法。

​		如果要给变量一个不同的类型，我们必须显式地将无类型的常量转化为所需的类型，或给声明的变量指定明确的类型，像下面例子这样：

```go
var i = int8(0)
var i int8 = 0
```

​		当尝试将这些无类型的常量转为一个接口值时，这些默认类型将显得尤为重要，因为要靠它们明确接口对应的动态类型。

```go
fmt.Printf("%T\n", 0) 			// "int"
fmt.Printf("%T\n", 0.0) 		// "float64"
fmt.Printf("%T\n", 0i) 			// "complex128"
fmt.Printf("%T\n", '\000')	    // "int32" (rune)
```

## 4	复合数据类型

​		复合数据类型，是以不同的方式组合基本类型可以构造出来的复合数据类型。我们主要讨论四种类型——数组、slice、map和结构体。

​		数组和结构体是聚合类型；它们的值由许多元素或成员字段的值组成。数组是由同构的元素组成——每个数组元素都是完全相同的类型——结构体则是由异构的元素组成的。数组和结构体都是有固定内存大小的数据结构。相比之下，slice和map则是动态的数据结构，它们将根据需要动态增长。

### 4.1	数组

​		数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。和数组对应的类型是Slice（切片），它是可以增长和收缩动态序列，slice功能也更灵活。

​		数组的每个元素可以通过索引下标来访问，索引下标的范围是从0开始到数组长度减1的位置。内置的len函数将返回数组中元素的个数。

```go
var a [3]int 							// array of 3 integers	3个元素的整数数组
fmt.Println(a[0]) 						// print the first element	打印第一个元素
fmt.Println(a[len(a)-1]) 				// print the last element, a[2]	打印最后一个元素，下标a[2]
										
for i, v := range a {
	fmt.Printf("%d %d\n", i, v)			// Print the indices and elements.	打印索引和元素
}
										
for _, v := range a {
	fmt.Printf("%d\n", v)				// Print the elements only.	只打印元素
}
```

​		默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就是0。我们也可以使用数组字面值语法用一组值来初始化数组：

```go
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
fmt.Println(r[2]) 					// "0"
```

​		在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算。因此，上面q数组的定义可以简化为

```go
q := [...]int{1, 2, 3}
fmt.Printf("%T\n", q)			   // "[3]int"
```

​		数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。

```go
q := [3]int{1, 2, 3}
q = [4]int{1, 2, 3, 4} 		// compile error: cannot assign [4]int to [3]int
```

​		数组、slice、map和结构体字面值的写法都很相似。上面的形式是直接提供顺序初始化值序列，但是也可以指定一个索引和对应值列表的方式初始化，就像下面这样：

```go
type Currency int
const (
	USD Currency = iota // 美元
	EUR 				// 欧元
	GBP 				// 英镑
	RMB				    // 人民币
)
symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
fmt.Println(RMB, symbol[RMB])			 // "3 ￥"
```

​		在这种形式的数组字面值形式中，初始化索引的顺序是无关紧要的，而且没用到的索引可以省略，和前面提到的规则一样，未指定初始值的元素将用零值初始化。例如：	

```go
r := [...]int{99: -1}
```

​		定义了一个含有100个元素的数组r，最后一个元素被初始化为-1，其它元素都是用0初始化。

​		如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的，可以直接通过==比较运算符来比较两个数组，只有当两个数组的所有元素都是相等的时候数组才是相等的。不相等比较运算符!=遵循同样的规则。

```go
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}
fmt.Println(a == b, a == c, b == c) 		// "true false false"
d := [3]int{1, 2}
fmt.Println(a == d)						    // compile error: cannot compare [2]int == [3]int
```

​		当调用一个函数的时候，函数的每个调用参数将会被赋值给函数内部的参数变量，所以函数参数变量接收的是一个复制的副本，并不是原始调用的变量。因为函数参数传递的机制导致传递大的数组类型将是低效的，并且对数组参数的任何的修改都是发生在复制的数组上，并不能直接修改调用时原始的数组变量。

​		可以显式地传入一个数组指针，那样的话函数通过指针对数组的任何修改都可以直接反馈到调用者。下面的函数用于给[32]byte类型的数组清零：

```go
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
```

​		其实数组字面值[32]byte{}就可以生成一个32字节的数组。而且每个数组的元素都是零值初始化，也就是0。因此，我们可以将上面的zero函数写的更简洁一点：

```go
func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
```

​		虽然通过指针来传递数组参数是高效的，而且也允许在函数内部修改数组的值，但是数组依然是僵化的类型，因为数组的类型包含了僵化的长度信息。上面的zero函数并不能接收指向[16]byte类型数组的指针，而且也没有任何添加或删除数组元素的方法。除了像SHA256这类需要处理特定大小数组的特例外，数组依然很少用作函数参数；相反，我们一般使用slice来替代数组。

### 4.2	Slice

​		Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。一个slice类型一般写作[]T，其中**T**代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已。

​		一个slice是一个轻量级的数据结构，提供了访问数组子序列（或者全部）元素的功能，而且slice的底层确实引用一个数组对象。一个slice由三个部分构成：指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址，要注意的是**slice的第一个元素并不一定就是数组的第一个元素**。长度对应slice中元素的数目；长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分别返回slice的长度和容量。

​		多个slice之间可以共享底层的数据，并且引用的数组部分区间可能重叠。

​		slice的切片操作s[i:j]，其中0 ≤ i≤ j≤ cap(s)，用于创建一个新的slice，引用s的从第i个元素开始到第j-1个元素的子序列。新的slice将只有j-i个元素。如果i位置的索引被省略的话将使用0代替，如果j位置的索引被省略的话将使用len(s)代替。s[:]切片操作则是引用整个数组。

​		如果切片操作超出cap(s)的上限将导致一个panic异常，但是超出len(s)则是意味着扩展了slice，因为新slice的长度会变大：	

```go
fmt.Println(summer[:20])			 // panic: out of range
endlessSummer := summer[:5] 		 // extend a slice (within capacity)
fmt.Println(endlessSummer)			 // "[June July August September October]"
```

​		字符串的切片操作和[]byte字节类型切片的切片操作是类似的。它们都写作x[m:n]，并且都是返回一个原始字节系列的子序列，底层都是共享之前的底层数组，因此切片操作对应常量时间复杂度。x[m:n]切片操作对于字符串则生成一个新字符串，如果x是[]byte的话则生成一个新的[]byte。

​		因为slice值包含指向第一个slice元素的指针，因此向函数传递slice将允许在函数内部修改底层数组的元素。换句话说，复制一个slice只是对底层的数组创建了一个新的slice别名。下面的reverse函数在原内存空间将[]int类型的slice反转，而且它可以用于任意长度的slice。

```go
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
```

​		这里我们反转数组的应用：

```go
a := [...]int{0, 1, 2, 3, 4, 5}
reverse(a[:])
fmt.Println(a) // "[5 4 3 2 1 0]"
```

​		一种将slice元素循环向左镟转n个元素的方法是三次调用reverse反转函数，第一次是反转开头的n个元素，然后是反转剩下的元素，最后是反转整个slice的元素。（如果是向右循环镟转，则将第三个函数调用移到第一个调用位置就可以了。）

```go
s := []int{0, 1, 2, 3, 4, 5}
// Rotate s left by two positions.
reverse(s[:2])
reverse(s[2:])
reverse(s)
fmt.Println(s) // "[2 3 4 5 0 1]"
```

​		要注意的是slice类型的变量s和数组类型的变量a的初始化语法的差异。slice和数组的字面值语法很类似，它们都是用花括弧包含一系列的初始化元素，但是对于slice并没有指明序列的长度。这会隐式地创建一个合适大小的数组，然后slice的指针指向底层的数组。就像数组字面值一样，slice的字面值也可以按顺序指定初始化值序列，或者是通过索引和元素值指定，或者的两种风格的混合语法初始化。

​		和数组不同的是，slice之间不能比较，因此我们不能使用==操作符来判断两个slice是否含有全部相等元素。不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等（[]byte），但是对于其他类型的slice，我们必须自己展开每个元素进行比较：

```go
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
```

​		上面关于两个slice的深度相等测试，运行的时间并不比支持==操作的数组或字符串更多，为何slice不直接支持比较运算符呢？

​		第一个原因，一个slice的元素是间接引用的，一个slice甚至可以包含自身。虽然有很多办法处理这种情形，但是没有一个是简单有效的。

​		第二个原因，因为slice的元素是间接引用的，一个固定值的slice在不同的时间可能包含不同的元素，因为底层数组的元素可能会被修改。并且Go语言中map等哈希表之类的数据结构的key只做简单的浅拷贝，它要求在整个声明周期中相等的key必须对相同的元素。对于像指针或chan之类的引用类型，==相等测试可以判断两个是否是引用相同的对象。一个针对slice的浅相等测试的==操作符可能是有一定用处的，也能临时解决map类型的key问题，但是slice和数组不同的相等测试行为会让人困惑。安全的做法是直接禁止slice之间的比较操作。

​		slice唯一合法的比较操作是和nil比较，例如：

```
if summer == nil { /* ... */ }
```

​		一个零值的slice等于nil。一个nil值的slice并没有底层数组。一个nil值的slice的长度和容量都是0，但是也有非nil值的slice的长度和容量也是0的，例如[]int{}或make([]int, 3)[3:]。与任意类型的nil值一样，我们可以用[]int(nil)类型转换表达式来生成一个对应类型slice的nil值。	

```go
var s []int 	// len(s) == 0, s == nil
s = nil 		// len(s) == 0, s == nil
s = []int(nil)  // len(s) == 0, s == nil
s = []int{} 	// len(s) == 0, s != nil
```

​		如果需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断。除了和nil相等比较外，一个nil值的slice的行为和其它任意0产长度的slice一样；例如reverse(nil)也是安全的。

​		内置的make函数创建一个指定元素类型、长度和容量的slice。容量部分可以省略，在这种情况下，容量将等于长度。

```go
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
```

​		在底层，make创建了一个匿名的数组变量，然后返回一个slice；只有通过返回的slice才能引用底层匿名的数组变量。在第一种语句中，slice是整个数组的view。在第二个语句中，slice只引用了底层数组的前len个元素，但是容量将包含整个的数组。额外的元素是留给未来的增长用的。

#### 4.2.1	append函数

​		内置的append函数用于向slice追加元素：

```go
var runes []rune
for _, r := range "Hello, 世界" {
	runes = append(runes, r)
}
fmt.Printf("%q\n", runes) 			// "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
```

​		在循环中使用append函数构建一个由九个rune字符构成的slice，当然对应这个特殊的问题我们可以通过Go语言内置的[]rune("Hello, 世界")转换操作完成。

​		append函数对于理解slice底层是如何工作的非常重要，所以让我们仔细查看究竟是发生了什么。下面是第一个版本的appendInt函数，专门用于处理[]int类型的slice：

```go
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

```

​		每次调用appendInt函数，必须先检测slice底层数组是否有足够的容量来保存新添加的元素。如果有足够空间的话，直接扩展slice（依然在原有的底层数组之上），将新添加的y元素复制到新扩展的空间，并返回slice。因此，输入的x和输出的z共享相同的底层数组。

​		如果没有足够的增长空间的话，appendInt函数则会先分配一个足够大的slice用于保存新的结果，先将输入的x复制到新的空间，然后添加y元素。结果z和输入的x引用的将是不同的底层数组。

​		虽然通过循环复制元素更直接，不过内置的copy函数可以方便地将一个slice复制另一个相同类型的slice。copy函数的第一个参数是要复制的目标slice，第二个参数是源slice，目标和源的位置顺序和dst = src赋值语句是一致的。两个slice可以共享同一个底层数组，甚至有重叠也没有问题。copy函数将返回成功复制的元素的个数，等于两个slice中较小的长度，所以我们不用担心覆盖会超出目标slice的范围。

​		为了提高内存使用效率，新分配的数组一般略大于保存x和y所需要的最低大小。通过在每次扩展数组时直接将长度翻倍从而避免了多次内存分配，也确保了添加单个元素操的平均时间是一个常数时间。这个程序演示了效果：

```go
func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
```

​		每一次容量的变化都会导致重新分配内存和copy操作：

```
0 cap=1 [0]
1 cap=2 [0 1]
2 cap=4 [0 1 2]
3 cap=4 [0 1 2 3]
4 cap=8 [0 1 2 3 4]
5 cap=8 [0 1 2 3 4 5]
6 cap=8 [0 1 2 3 4 5 6]
7 cap=8 [0 1 2 3 4 5 6 7]
8 cap=16 [0 1 2 3 4 5 6 7 8]
9 cap=16 [0 1 2 3 4 5 6 7 8 9]
```

​		让我们仔细查看i=3次的迭代。当时x包含了[0 1 2]三个元素，但是容量是4，因此可以简单将新的元素添加到末尾，不需要新的内存分配。然后新的y的长度和容量都是4，并且和x引用着相同的底层数组

![image-20230702195105784](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230702195105784.png)

​		在下一次迭代时i=4，现在没有新的空余的空间了，因此appendInt函数分配一个容量为8的底层数组，将x的4个元素[0123]复制到新空间的开头，然后添加新的元素i，新元素的值是4。新的y的长度是5，容量是8；后面有3个空闲的位置，三次迭代都不需要分配新的空间。当前迭代中，y和x是对应不同底层数组的view。

![image-20230702195217366](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230702195217366.png)

​		内置的append函数可能使用比appendInt更复杂的内存扩展策略。通常我们并不知道append调用是否导致了内存的重新分配，不能确认新的slice和原始的slice是否引用的是相同的底层数组空间。同样，不能确认在原先的slice上的操作是否会影响到新的slice。因此，通常是将append返回的结果直接赋值给输入的slice变量：

```
runes = append(runes, r)
```

​		更新slice变量不仅对调用append函数是必要的，实际上对应任何可能导致长度、容量或底层数组变化的操作都是必要的。要正确地使用slice，需要记住尽管底层数组的元素是间接访问的，但是slice对应结构体本身的指针、长度和容量部分是直接访问的。要更新这些信息需要像上面例子那样一个显式的赋值操作。从这个角度看，slice并不是一个纯粹的引用类型，它实际上是一个类似下面结构体的聚合类型：

```go
type IntSlice struct {
	ptr *int
	len, cap int
}
```

​		我们的appendInt函数每次只能向slice追加一个元素，但是内置的append函数则可以追加多个元素，甚至追加一个slice。

```go
var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, 4, 5, 6)
x = append(x, x...) // append the slice x
fmt.Println(x) 		// "[1 2 3 4 5 6 1 2 3 4 5 6]"
```

​		通过下面的小修改可以达到append函数类似的功能

```go
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	// ...expand z to at least zlen...
	copy(z[len(x):], y)
	return z
}
```

### 4.3	map

​		哈希表是一种巧妙并且实用的数据结构。它是一个无序的key/value对的集合，其中所有的key都是不同 的，然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value。

​		在Go语言中，一个map就是一个哈希表的引用，map类型可以写为map[K]V，其中K和V分别对应key和 value。map中所有的key都有相同的类型，所以的value也有着相同的类型，但是key和value之间可以是 不同的数据类型。其中K对应的key必须是支持==比较运算符的数据类型，所以map可以通过测试key是否 相等来判断是否已经存在。虽然浮点数类型也是支持相等运算符比较的，但是将浮点数用做key类型则是 一个坏的想法，最坏的情况是可能出现的NaN和任何浮点数都不相等。对于V对应的 value数据类型则没有任何的限制。

​		内置的make函数可以创建一个map：

```go
ages := make(map[string]int) // mapping from strings to ints
```

​		我们也可以用map字面值的语法创建map，同时还可以指定一些最初的key/value：

```go
ages := map[string]int{
	"alice": 31, 
	"charlie": 34,
}
```

​		相当于

```go
ages := make(map[string]int) 
ages["alice"] = 31 
ages["charlie"] = 34
```

​		另一种创建空的map的表达式是map[string]int{}。 Map中的元素通过key对应的下标语法访问：

```
ages["alice"] = 32
fmt.Println(ages["alice"]) // "32"
```

​		使用内置的delete函数可以删除元素：

```go
delete(ages, "alice") // remove element ages["alice"]
```

​		所有这些操作是安全的，即使这些元素不在map中也没有关系；如果一个查找失败将返回value类型对应 的零值，例如，即使map中不存在“bob”下面的代码也可以正常工作，因为ages["bob"]失败时将返回 0。

```go
ages["bob"] = ages["bob"] + 1 // happy birthday!
```

​		而且x += y和x++等简短赋值语法也可以用在map上，所以上面的代码可以改写成

```go
ages["bob"] += 1
更简单的写法：
ages["bob"]++
```

​		但是map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作：

```go
_ = &ages["bob"] // compile error: cannot take address of map element
```

​		禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之 前的地址无效。

​		要想遍历map中全部的key/value对的话，可以使用range风格的for循环实现，和之前的slice遍历语法类 似。下面的迭代语句将在每次迭代时设置name和age变量，它们对应下一个键/值对：

```go
for name, age := range ages { 
	fmt.Printf("%s\t%d\n", name, age) 
}
```

​		Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。遍历的顺序 是随机的，每一次遍历的顺序都不相同。如果要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使 用sort包的Strings函数对字符串slice进行排序。下面是常见的处理方式：

```go
import "sort"
var names []string
for name := range ages {
	names = append(names, name)
}
sort.Strings(names)
for _, name := range names {
	fmt.Printf("%s\t%d\n", name, ages[name])
}
```

​		因为我们一开始就知道names的最终大小，因此给slice分配一个合适的大小将会更有效。下面的代码创建了一个空的slice，但是slice的容量刚好可以放下map中全部的key：

```
names := make([]string, 0, len(ages))
```

​		在上面的第一个range循环中，我们只关心map中的key，所以我们忽略了第二个循环变量。在第二个循环中，我们只关心names中的名字，所以我们使用“_”空白标识符来忽略第一个循环变量，也就是迭代slice时的索引。

​		map类型的零值是nil，也就是没有引用任何哈希表

```go
var ages map[string]int
fmt.Println(ages == nil) 		// "true"
fmt.Println(len(ages) == 0) 	// "true"
```

​		map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。但是向一个nil值的map存入元素将导致一个panic异常：

```go
ages["carol"] = 21 		// panic: assignment to entry in nil map
```

​		在向map存数据前必须先创建map。

通过key作为索引下标来访问map将产生一个value。如果key在map中是存在的，那么将得到与key对应的value；如果key不存在，那么将得到value对应类型的零值，正如我们前面看到的ages["bob"]那样。这个规则很实用，但是有时候可能需要知道对应的元素是否真的是在map之中。例如，如果元素类型是一个数字，可以需要区分一个已经存在的0，和不存在而返回零值的0，可以像下面这样测试：

```go
age, ok := ages["bob"]
if !ok { /* "bob" is not a key in this map; age == 0. */ }
```

```go
if age, ok := ages["bob"]; !ok { /* ... */ }
```

​		在这种场景下，map的下标语法将产生两个值；第二个是一个布尔值，用于报告元素是否真的存在。布尔变量一般命名为ok，特别适合马上用于if条件判断部分。

​		和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较。要判断两个map是否包含相同的key和value，我们必须通过一个循环实现：

```go
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
```

​		要注意我们是如何用!ok来区分元素缺失和元素不同的。我们不能简单地用xv != y[k]判断，那样会导致在判断下面两个map时产生错误的结果：

```go
// True if equal is written incorrectly.
equal(map[string]int{"A": 0}, map[string]int{"B": 42})
```

​		Go语言中并没有提供一个set类型，但是map中的key也是不相同的，可以用map实现类似set的功能。为了说明这一点，下面的dedup程序读取多行输入，但是只打印第一次出现的行。dedup程序通过map来表示所有的输入行所对应的set集合，以确保已经在集合存在的行不会被重复打印

```go
func main() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
```

​		将这种忽略value的map当作一个字符串集合，并非所有map[string]bool类型value都是无关紧要的；有一些则可能会同时包含tue和false的值。

​		有时候我们需要一个map或set的key是slice类型，但是map的key必须是可比较的类型，但是slice并不满足这个条件。不过，我们可以通过两个步骤绕过这个限制。第一步，定义一个辅助函数k，将slice转为map对应的string类型的key，确保只有x和y相等时k(x) == k(y)才成立。然后创建一个key为string类型的map，在每次对map操作时先用k辅助函数将slice转化为string类型。

​		下面的例子演示了如何使用map来记录提交相同的字符串列表的次数。它使用了fmt.Sprintf函数将字符串列表转换为一个字符串以用于map的key，通过%q参数忠实地记录每个字符串元素的信息：

```go
var m = make(map[string]int)
func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
```

​		使用同样的技术可以处理任何不可比较的key类型，而不仅仅是slice类型。这种技术对于想使用自定义key比较函数的时候也很有用，例如在比较字符串的时候忽略大小写。同时，辅助函数k(x)也不一定是字符串类型，它可以返回任何可比较的类型，例如整数、数组或结构体等。

​		这是map的另一个例子，下面的程序用于统计输入中每个Unicode码点出现的次数。虽然Unicode全部码点的数量巨大，但是出现在特定文档中的字符种类并没有多少，使用map可以用比较自然的方式来跟踪那些出现过字符的次数

```go
// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

```

​		ReadRune方法执行UTF-8解码并返回三个值：解码的rune字符的值，字符UTF-8编码后的长度，和一个错误值。我们可预期的错误值只有对应文件结尾的io.EOF。如果输入的是无效的UTF-8编码的字符，返回的将是unicode.ReplacementChar表示无效字符，并且编码长度是1。

​		charcount程序同时打印不同UTF-8编码长度的字符数目。对此，map并不是一个合适的数据结构；因为UTF-8编码的长度总是从1到utf8.UTFMax（最大是4个字节），使用数组将更有效。

​		Map的value类型也可以是一个聚合类型，比如是一个map或slice。在下面的代码中，图graph的key类型是一个字符串，value类型map[string]bool代表一个字符串集合。从概念上将，graph将一个字符串类型的key映射到一组相关的字符串集合，它们指向新的graph的key。

```go
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}
```

​		其中addEdge函数惰性初始化map是一个惯用方式，也就是说在每个值首次作为key时才初始化。addEdge函数显示了如何让map的零值也能正常工作；即使from到to的边不存在，graph[from][to]依然可以返回一个有意义的结果。

### 4.4	结构体

​		结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。

​		下面两个语句声明了一个叫Employee的命名的结构体类型，并且声明了一个Employee类型的变量dilbert：

```go
type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}
var dilbert Employee
```

​		dilbert结构体变量的成员可以通过点操作符访问，比如dilbert.Name和dilbert.DoB。因为dilbert是一个变量，它所有的成员也同样是变量，我们可以直接对每个成员赋值：

```go
dilbert.Salary -= 5000					   // demoted, for writing too few lines of code
```

​		或者是对成员取地址，然后通过指针访问：

```go
position := &dilbert.Position
*position = "Senior " + *position 			// promoted, for outsourcing to Elbonia
```

​		点操作符也可以和指向结构体的指针一起工作：

```go
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
```

​		相当于下面语句

```go
(*employeeOfTheMonth).Position += " (proactive team player)"
```

​		下面的EmployeeByID函数将根据给定的员工ID返回对应的员工信息结构体的指针。我们可以使用点操作符来访问它里面的成员：

```go
func EmployeeByID(id int) *Employee { /* ... */ }

fmt.Println(EmployeeByID(dilbert.ManagerID).Position)			 // "Pointy-haired boss"

id := dilbert.ID
EmployeeByID(id).Salary = 0										 // fired for... no real reason
```

​		后面的语句通过EmployeeByID返回的结构体指针更新了Employee结构体的成员。如果将EmployeeByID函数的返回值从*Employee指针类型改为Employee值类型，那么更新语句将不能编译通过，因为在赋值语句的左边并不确定是一个变量（调用函数返回的是值，并不是一个可取地址的变量）。

​		通常一行对应一个结构体成员，成员的名字在前类型在后，不过如果相邻的成员类型如果相同的话可以被合并到一行，就像下面的Name和Address成员那样：

```go
type Employee struct {
	ID int
	Name, Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}
```

​		结构体成员的输入顺序也有重要的意义。我们也可以将Position成员合并（因为也是字符串类型），或者是交换Name和Address出现的先后顺序，那样的话就是定义了不同的结构体类型。

​		如果结构体成员名字是以大写字母开头的，那么该成员就是导出的；这是Go语言导出规则决定的。一个结构体可能同时包含导出和未导出的成员。

​		结构体类型往往是冗长的，因为它的每个成员可能都会占一行。虽然我们每次都可以重写整个结构体成员，但是重复会令人厌烦。因此，完整的结构体写法通常只在类型声明语句的地方出现，就像Employee类型声明语句那样。

​		一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身。（该限制同样适应于数组。）但是S类型的结构体可以包含*S指针类型的成员，这可以让我们创建递归的数据结构，比如链表和树结构等。在下面的代码中，我们使用一个二叉树来实现一个插入排序：

```go
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
```

​		结构体类型的零值是每个成员都对是零值。通常会将零值作为最合理的默认值。例如，对于bytes.Buffer类型，结构体初始值就是一个随时可用的空缓存，sync.Mutex的零值也是有效的未锁定状态。有时候这种零值可用的特性是自然获得的，但是也有些类型需要一些额外的工作。

​		如果结构体没有任何成员的话就是空结构体，写作struct{}。它的大小为0，也不包含任何信息，但是有时候依然是有价值的。有些用map带模拟set数据结构时，用它来代替map中布尔类型的value，只是强调key的重要性，但是因为节约的空间有限，而且语法比较复杂，所有我们通常避免避免这样的用法。

```go
seen := make(map[string]struct{}) // set of strings
// ...
if _, ok := seen[s]; !ok {
	seen[s] = struct{}{}
	// ...first time seeing s...
}
```

#### 4.4.1	结构体面值

​		结构体值也可以用结构体面值表示，结构体面值可以指定每个成员的值。

```go
type Point struct{ X, Y int }
p := Point{1, 2}
```

​		这里有两种形式的结构体面值语法，上面的是第一种写法，要求以结构体成员定义的顺序为每个结构体成员指定一个面值。它要求写代码和读代码的人要记住结构体的每个成员的类型和顺序，不过结构体成员有细微的调整就可能导致上述代码不能编译。因此，上述的语法一般只在定义结构体的包内部使用，或者是在较小的结构体中使用，这些结构体的成员排列比较规则，比如image.Point{x, y}或color.RGBA{red, green, blue, alpha}。

​		其实更常用的是第二种写法，以成员名字和相应的值来初始化，可以包含部分或全部的成员。

```go
anim := gif.GIF{LoopCount: nframes}
```

​		在这种形式的结构体面值写法中，如果成员被忽略的话将默认用零值。

​		两种不同形式的写法不能混合使用。而且，你不能企图在外部包中用第一种顺序赋值的技巧来偷偷地初始化结构体中未导出的成员。

```go
package p
type T struct{ a, b int }		 // a and b are not exported

package q
import "p"
var _ = p.T{a: 1, b: 2} 		 // compile error: can't reference a, b
var _ = p.T{1, 2}				 // compile error: can't reference a, b
```

​		虽然上面最后一行代码的编译错误信息中并没有显式提到未导出的成员，但是这样企图隐式使用未导出成员的行为也是不允许的。

​		结构体可以作为函数的参数和返回值。例如，这个Scale函数将Point类型的值缩放后返回：

```go
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}
fmt.Println(Scale(Point{1, 2}, 5)) // "{5 10}"
```

​		如果考虑效率的话，较大的结构体通常会用指针的方式传入和返回

```go
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}
```

​		如果要在函数内部修改结构体成员的话，用指针传入是必须的；因为在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。

```go
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
```

​		因为结构体通常通过指针处理，可以用下面的写法来创建并初始化一个结构体变量，并返回结构体的地址：

```
pp := &Point{1, 2}
```

​		它是下面的语句是等价的

```
pp := new(Point)
*pp = Point{1, 2}
```

​		不过&Point{1, 2}写法可以直接在表达式中使用，比如一个函数调用。

#### 4.4.2	结构体比较

​		如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用==或!=运算符进行比较。相等比较运算符==将比较两个结构体的每个成员，因此下面两个比较的表达式是等价的：

```go
type Point struct{ X, Y int }
p := Point{1, 2}
q := Point{2, 1}
fmt.Println(p.X == q.X && p.Y == q.Y)		 // "false"
fmt.Println(p == q) 						 // "false"
```

​		可比较的结构体类型和其他可比较的类型一样，可以用于map的key类型。

```go
type address struct {
	hostname string
	port int
}
hits := make(map[address]int)
hits[address{"golang.org", 443}]++
```

#### 4.4.3	结构体嵌套和你名成员

​		使用Go语言提供的不同寻常的结构体嵌入机制让一个命名的结构体包含另一个结构体类型的匿名成员，这样就可以通过简单的点运算符x.f来访问匿名成员链中嵌套的x.d.e.f成员。

​		考虑一个二维的绘图程序，提供了一个各种图形的库，例如矩形、椭圆形、星形和轮形等几何形状。这里是其中两个的定义：

```go
type Circle struct {
	X, Y, Radius int
}
type Wheel struct {
	X, Y, Radius, Spokes int
}
```

​		一个Circle代表的圆形类型包含了标准圆心的X和Y坐标信息，和一个Radius表示的半径信息。一个Wheel轮形除了包含Circle类型所有的全部成员外，还增加了Spokes表示径向辐条的数量。我们可以这样创建一个wheel变量：

```go
var w Wheel
w.X = 8
w.Y = 8
w.Radius = 5
w.Spokes = 20
```

​		随着库中几何形状数量的增多，我们一定会注意到它们之间的相似和重复之处，所以我们可能为了便于维护而将相同的属性独立出来：

```go
type Point struct {
	X, Y int
}
type Circle struct {
	Center Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}
```

​		这样改动之后结构体类型变的清晰了，但是这种修改同时也导致了访问每个成员变得繁琐：

```
var w Wheel
w.Circle.Center.X = 8
w.Circle.Center.Y = 8
w.Circle.Radius = 5
w.Spokes = 20
```

​		Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。下面的代码中，Circle和Wheel各自都有一个匿名成员。我们可以说Point类型被嵌入到了Circle结构体，同时Circle类型被嵌入到了Wheel结构体。

```go
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}
```

​		得意于匿名嵌入的特性，我们可以直接访问叶子属性而不需要给出完整的路径：

```go
var w Wheel
w.X = 8 			 // equivalent to w.Circle.Point.X = 8
w.Y = 8				 // equivalent to w.Circle.Point.Y = 8
w.Radius = 5		 // equivalent to w.Circle.Radius = 5
w.Spokes = 20
```

​		在右边的注释中给出的显式形式访问这些叶子成员的语法依然有效，因此匿名成员并不是真的无法访问了。其中匿名成员Circle和Point都有自己的名字——就是命名的类型名字——但是这些名字在点操作符中是可选的。我们在访问子成员的时候可以忽略任何匿名成员部分。

​		不幸的是，结构体字面值并没有简短表示匿名成员的语法， 因此下面的语句都不能编译通过：

```go
w = Wheel{8, 8, 5, 20}							   	    // compile error: unknown fields
w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} 			// compile error: unknown fields
```

​		结构体字面值必须遵循形状类型声明时的结构，所以我们只能用下面的两种语法，它们彼此是等价的：

```go
w = Wheel{Circle{Point{8, 8}, 5}, 20}
w = Wheel{
	Circle: Circle{
		Point:  Point{X: 8, Y: 8},
		Radius: 5,
	},
	Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
}
fmt.Printf("%#v\n", w)
// Output:
// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}
w.X = 42
fmt.Printf("%#v\n", w)
// Output:
// Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}
```

​		需要注意的是Printf函数中%v参数包含的#副词，它表示用和Go语言类似的语法打印值。对于结构体类型来说，将包含每个成员的名字。

​		因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员，这会导致名字冲突。同时，因为成员的名字是由其类型隐式地决定的，所有匿名成员也有可见性的规则约束。在上面的例子中，Point和Circle匿名成员都是导出的。即使它们不导出（比如改成小写字母开头的point和circle），我们依然可以用简短形式访问匿名成员嵌套的成员

```go
w.X = 8 		// equivalent to w.circle.point.X = 8
```

​		但是在包外部，因为circle和point没有导出不能访问它们的成员，因此简短的匿名成员访问语法也是禁止的。

​		实任何命令的类型都可以作为结构体的匿名成员。但是为什么要嵌入一个没有任何子成员类型的匿名成员类型呢？答案是匿名类型的方法集。简短的点运算符语法可以用于选择匿名成员嵌套的成员，也可以用于访问它们的方法。实际上，外层的结构体不仅仅是获得了匿名成员类型的所有成员，而且也获得了该类型导出的全部的方法。这个机制可以用于将一个有简单行为的对象组合成有复杂行为的对象。

#### 4.5	JSON

​		JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。在类似的协议中，JSON并不是唯一的一个标准协议。 XML）、ASN.1和Google的Protocol Buffers都是类似的协议，并且有各自的特色，但是由于简洁性、可读性和流行程度等原因，JSON是应用最广泛的一个。

​		Go语言对于这些标准格式的编码和解码都有良好的支持，由标准库中的encoding/json、encoding/xml、encoding/asn1等包提供支持（Protocol Buffers的支持由 github.com/golang/protobuf 包提供），并且这类包都有着相似的API接口。

​		JSON是对JavaScript中各种类型的值——字符串、数字、布尔值和对象——Unicode本文编码。它可以用有效可读的方式表示基础数据类型和本章的数组、slice、结构体和map等聚合数据类型。

​		基本的JSON类型有数字（十进制或科学记数法）、布尔值（true或false）、字符串，其中字符串是以双引号包含的Unicode字符序列，支持和Go语言类似的反斜杠转义特性，不过JSON使用的是\Uhhhh转义数字来表示一个UTF-16编码（UTF-16和UTF-8一样是一种变长的编码，有些Unicode码点较大的字符需要用4个字节表示；而且UTF-16还有大端和小端的问题），而不是Go语言的rune类型。

​		这些基础类型可以通过JSON的数组和对象类型进行递归组合。一个JSON数组是一个有序的值序列，写在一个方括号中并以逗号分隔；一个JSON数组可以用于编码Go语言的数组和slice。一个JSON对象是一个字符串到值的映射，写成以系列的name:value对形式，用花括号包含并以逗号分隔；JSON的对象类型可以用于编码Go语言的map类型（key类型是字符串）和结构体。例如：

```json
boolean true
number -273.15
string "She said \"Hello, BF\""
array ["gold", "silver", "bronze"]
object {"year": 1980,
		"event": "archery",
		"medals": ["gold", "silver", "bronze"]}
```

​		考虑一个应用程序，该程序负责收集各种电影评论并提供反馈功能。它的Movie数据类型和一个典型的表示电影的值列表如下所示。（在结构体声明中，Year和Color成员后面的字符串面值是结构体成员Tag；我们稍后会解释它的作用。）

```go
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}
```

​		这样的数据结构特别适合JSON格式，并且在两种之间相互转换也很容易。将一个Go语言中类似movies的结构体slice转为JSON的过程叫编组（marshaling）。编组通过调用json.Marshal函数完成：

```go
data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
fmt.Printf("%s\n", data)
```

​		Marshal函数返还一个编码后的字节slice，包含很长的字符串，并且没有空白缩进；我们将它折行以便于显示：

```
[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingr
id Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Ac
tors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"
Actors":["Steve McQueen","Jacqueline Bisset"]}]
```

​		这种紧凑的表示形式虽然包含了全部的信息，但是很难阅读。为了生成便于阅读的格式，另一个json.MarshalIndent函数将产生整齐缩进的输出。该函数有两个额外的字符串参数用于表示每一行输出的前缀和每一个层级的缩进：

```go
data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
fmt.Printf("%s\n", data)
```

​		上面的代码将产生这样的输出（在最后一个成员或元素后面并没有逗号分隔符）：

```json
[
	{
		"Title": "Casablanca",
		"released": 1942,
		"Actors": [
			"Humphrey Bogart",
			"Ingrid Bergman"
		]
	},
	{
		"Title": "Cool Hand Luke",
		"released": 1967,
		"color": true,
		"Actors": [
			"Paul Newman"
		]
	},
	{
		"Title": "Bullitt",
		"released": 1968,
		"color": true,
		"Actors": [
			"Steve McQueen",
			"Jacqueline Bisset"
		]
	}
]
```

​		在编码时，默认使用Go语言结构体的成员名字作为JSON的对象（通过reflect反射技术),只有导出的结构体成员才会被编码，这也就是我们为什么选择用大写字母开头的成员名称。Year名字的成员在编码后变成了released，还有Color成员编码后变成了小写字母开头的color。这是因为构体成员Tag所导致的。一个构体成员Tag是和在编译阶段关联到该成员的元信息字符串：

```
Year int `json:"released"`
Color bool `json:"color,omitempty"`
```

​		结构体的成员Tag可以是任意的字符串面值，但是通常是一系列用空格分隔的key:"value"键值对序列；因为值中含义双引号字符，因此成员Tag一般用原生字符串面值的形式书写。json开头键名对应的值用于控制encoding/json包的编码和解码的行为，并且encoding/...下面其它的包也遵循这个约定。成员Tag中json对应值的第一部分用于指定JSON对象的名字，比如将Go语言中的TotalCount成员对应到JSON中的total_count对象。Color成员的Tag还带了一个额外的omitempty选项，表示当Go语言结构体成员为空或零值时不生成JSON对象（这里false为零值）。果然，Casablanca是一个黑白电影，并没有输出Color成员.

​		编码的逆操作是解码，对应将JSON数据解码为Go语言的数据结构，Go语言中一般叫unmarshaling，通过json.Unmarshal函数完成。下面的代码将JSON格式的电影数据解码为一个结构体slice，结构体中只有Title成员。通过定义合适的Go语言数据结构，我们可以选择性地解码JSON中感兴趣的成员。当Unmarshal函数调用返回，slice将被只含有Title信息值填充，其它JSON成员将被忽略。

```go
var titles []struct{ Title string }
if err := json.Unmarshal(data, &titles); err != nil {
	log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
```

​		许多web服务都提供JSON接口，通过HTTP接口发送JSON格式请求并返回JSON格式的信息。为了说明这一点，我们通过Github的issue查询服务来演示类似的用法。首先，我们要定义合适的类型和常量：

```go
package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

```

​		和前面一样，即使对应的JSON对象名是小写字母，每个结构体的成员名也是声明为大写字母开头的。因为有些JSON成员名字和Go结构体成员名字并不相同，因此需要Go语言结构体成员Tag来指定对应的JSON名字。同样，在解码的时候也需要做同样的处理，GitHub服务返回的信息比我们定义的要多很多。

​		SearchIssues函数发出一个HTTP请求，然后解码返回的JSON格式的结果。因为用户提供的查询条件可能包含类似?和&之类的特殊字符，为了避免对URL造成冲突，我们用url.QueryEscape来对查询中的特殊字符进行转义操作。

```go
import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
```

​		在早些的例子中，我们使用了json.Unmarshal函数来将JSON格式的字符串解码为字节slice。但是这个例子中，我们使用了基于流式的解码器json.Decoder，它可以从一个输入流解码JSON数据，还有一个针对输出流的json.Encoder编码对象。

​		调用Decode方法来填充变量。这里有多种方法可以格式化结构。下面是最简单的一种，以一个固定宽度打印每个issue，但是在下一节我们将看到如果利用模板来输出复杂的格式。

```go
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

```

​		通过命令行参数指定检索条件。下面的命令是查询Go语言项目中和JSON解码相关的问题，还有查询返回的结果：

```go
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680 eaigner encoding/json: set key converter on en/decoder
#6050 gopherbot encoding/json: provide tokenizer
#8658 gopherbot encoding/json: use bufio
#8462 kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901 rsc encoding/json: allow override type marshaling
#9812 klauspost encoding/json: string tag not symmetric
#7872 extempora encoding/json: Encoder internally buffers full output
#9650 cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716 gopherbot encoding/json: include field name in unmarshal error me
#6901 lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384 joeshaw encoding/json: encode precise floating point integers u
#6647 btracey x/tools/cmd/godoc: display type kind of each named type
#4237 gjemiller encoding/base64: URLEncoding padding is optional
```

## 5	函数

​		函数可以让我们将一个语句序列打包为一个单元，然后可以从程序中其它地方多次调用。函数的机制可以让我们将一个大的工作分解为小的任务，这样的小任务可以让不同程序员在不同时间、不同地方独立完成。一个函数同时对用户隐藏了其实现细节。

### 5.1	函数声明

​		函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。

```go
func name(parameter-list) (result-list) {
	body
}
```

​		形式参数列表描述了函数的参数名以及参数类型。这些参数作为局部变量，其值由参数调用者提供。返回值列表描述了函数返回值的变量名以及类型。如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。如果一个函数声明不包括返回值列表，那么函数体执行完毕后，不会返回任何值。 在hypot函数中：

```go
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
fmt.Println(hypot(3,4)) 			// "5"
```

​		x和y是形参名,3和4是调用时的传入的实数，函数返回了一个float64类型的值。返回值也可以像形式参数一样被命名。在这种情况下，每个返回值被声明成一个局部变量，并根据该返回值的类型，将其初始化为0。 如果一个函数在声明时，包含返回值列表，该函数必须以 return语句结尾，除非函数明显无法运行到结尾处。例如函数在结尾时调用了panic异常或函数中存在无限循环。

​		如果一组形参或返回值有相同的类型，我们不必为每个形参都写出参数类型。下面2个声明是等价的：

```go
func f(i, j, k int, s, t string) 				{ /* ... */ }
func f(i int, j int, k int, s string, t string) { /* ... */ }
```

​		我们给出4种方法声明拥有2个int型参数和1个int型返回值的函数.blank identifier(即下文的_符号)可以强调某个参数未被使用。

```go
func add(x int, y int) int		 {return x + y}
func sub(x, y int) (z int) 		 { z = x - y; return}
func first(x int, _ int) int 	 { return x }
func zero(int, int) int			 { return 0 }

fmt.Printf("%T\n", add)			 // "func(int, int) int"
fmt.Printf("%T\n", sub)			 // "func(int, int) int"
fmt.Printf("%T\n", first)		 // "func(int, int) int"
fmt.Printf("%T\n", zero)		 // "func(int, int) int"
```

​		函数的类型被称为函数的标识符。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型和标识符。形参和返回值的变量名不影响函数标识符也不影响它们是否可以以省略参数类型的形式表示。

​		每一次函数调用都必须按照声明顺序为所有参数提供实参（参数值）。在函数调用时，Go语言没有默认参数值，也没有任何方法可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言没有意义。

​		在函数体中，函数的形参作为局部变量，被初始化为调用者提供的值。函数的形参和有名返回值作为函数最外层的局部变量，被存储在相同的词法块中。

​		实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的简介引用被修改。

​		可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现的。这样的声明定义了函数标识符

```go
package math

func Sin(x float64) float //implemented in assembly language
```

### 5.2	递归

​		函数可以是递归的，这意味着函数可以直接或间接的调用自身。

​		下文的示例代码使用了非标准包 golang.org/x/net/html ，解析HTML。golang.org/x/... 目录下存储了一些由Go团队设计、维护，对网络编程、国际化文件处理、移动平台、图像处理、加密解密、开发者工具提供支持的扩展包。

​		例子中调用golang.org/x/net/html的部分api如下所示。html.Parse函数读入一组bytes.解析后，返回html.node类型的HTML页面树状结构根节点。HTML拥有很多类型的结点如text（文本）,commnets（注释）类型，在下面的例子中，我们 只关注< name key='value' >形式的结点。

```go
package html

import "io"

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)

```

​		main函数解析HTML标准输入，通过递归函数visit获得links（链接），并打印出这些links：

```go
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

```

​		visit函数遍历HTML的节点树，从每一个anchor元素的href属性获得link,将这些links存入字符串数组中，并返回这个字符串数组。

```go
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
```

​		为了遍历结点n的所有后代结点，每次遇到n的孩子结点时，visit递归的调用自身。这些孩子结点存放在FirstChild链表中。让我们以Go的主页（golang.org）作为目标，运行findlinks。我们以fetch的输出作为findlinks的输入。下面的输出做了简化处理。

```
$ go build gopl.io/ch1/fetch
$ go build gopl.io/ch5/findlinks1
$ ./fetch https://golang.org | ./findlinks1
#
/doc/
/pkg/
/help/
/blog/
http://play.golang.org/
//tour.golang.org/
https://golang.org/dl/
//blog.golang.org/
/LICENSE
/doc/tos.html
http://www.google.com/intl/en/policies/privacy/
```

​		注意在页面中出现的链接格式，在之后我们会介绍如何将这些链接，根据根路径（ https://golang.org）生成可以直接访问的url。

​		在函数outline中，我们通过递归的方式遍历整个HTML结点树，并输出树的结构。在outline内部，每遇到一个HTML元素标签，就将其入栈，并输出。

```go
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
```

​		有一点值得注意：outline有入栈操作，但没有相对应的出栈操作。当outline调用自身时，被调用者接收的是stack的拷贝。被调用者的入栈操作，修改的是stack的拷贝，而不是调用者的stack,因对当函数返回时,调用者的stack并未被修改。

下面是 https://golang.org 页面的简要结构:

```
$ go build gopl.io/ch5/outline
$ ./fetch https://golang.org | ./outline
[html]
[html head]
[html head meta]
[html head title]
[html head link]
[html body]
[html body div]
[html body div]
[html body div div]
[html body div div form]
[html body div div form div]
[html body div div form div a]
...
```

​		正如你在上面实验中所见，大部分HTML页面只需几层递归就能被处理，但仍然有些页面需要深层次的递归。

​		大部分编程语言使用固定大小的函数调用栈，常见的大小从64KB到2MB不等。固定大小栈会限制递归的深度，当你用递归处理大量数据时，需要避免栈溢出；除此之外，还会导致安全性问题。与相反,Go语言使用可变栈，栈的大小按需增加(初始时很小)。这使得我们使用递归时不必考虑溢出和安全问题。

### 5.3	多返回值

​		一个函数可以返回多个值。是期望得到的返回值，另一个是函数出错时的错误信息。

​		下面的程序是findlinks的改进版本。修改后的findlinks可以自己发起HTTP请求，这样我们就不必再运 行fetch。因为HTTP请求和解析操作可能会失败，因此findlinks声明了2个返回值：链接列表和错误信 息。一般而言，HTML的解析器可以处理HTML页面的错误结点，构造出HTML页面结构，所以解析HTML很少 失败。这意味着如果findlinks函数失败了，很可能是由于I/O的错误导致的。

```go
func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}
```

​		在findlinks中，有4处return语句，每一处return都返回了一组值。前三处return，将http和html包中 的错误信息传递给findlinks的调用者。第一处return直接返回错误信息，其他两处通过 fmt.Errorf输出详细的错误信息。如果findlinks成功结束，最后的return语句将一组解析获 得的连接返回给用户。

​		在finallinks中，我们必须确保resp.Body被关闭，释放网络资源。虽然Go的垃圾回收机制会回收不被使 用的内存，但是这不包括操作系统层面的资源，比如打开的文件、网络连接。因此我们必须显式的释放 这些资源。

​		调用多返回值函数时，返回给调用者的是一组值，调用者必须显式的将这些值分配给变量:

```
links, err := findLinks(url)
```

​		如果某个值不被使用，可以将其分配给blank identifier:

```
links, _ := findLinks(url) // errors ignored
```

​		一个函数内部可以将另一个有多返回值的函数作为返回值，下面的例子展示了与findLinks有相同功能的 函数，两者的区别在于下面的例子先输出参数：

```go
func findLinksLog(url string) ([]string, error) {
	log.Printf("findLinks %s", url)
	return findLinks(url)
}
```

​		当你调用接受多参数的函数时，可以将一个返回多参数的函数作为该函数的参数。虽然这很少出现在实 际生产代码中，但这个特性在debug时很方便，我们只需要一条语句就可以输出所有的返回值。下面的代 码是等价的：

```go
log.Println(findLinks(url)) 

links, err := findLinks(url) 
log.Println(links, err)
```

​		准确的变量名可以传达函数返回值的含义。尤其在返回值的类型都相同时，就像下面这样：

```go
func Size(rect image.Rectangle) (width, height int)
func Split(path string) (dir, file string)
func HourMinSec(t time.Time) (hour, minute, second int)
```

​		虽然良好的命名很重要，但也不必为每一个返回值都取一个适当的名字。比如，按照惯例，函数的最后一个bool类型的返回值表示函数是否运行成功，error类型的返回值代表函数的错误信息，对于这些类似的惯例，我们不必思考合适的命名，它们都无需解释。

​		如果一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数。这称之为bare return。

```go
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) { /* ... */ }
```

​		按照返回值列表的次序，返回所有的返回值，在上面的例子中，每一个return语句等价于：

```
return words, images, err
```

​		当一个函数有多处return语句以及许多返回值时，bare return 可以减少代码的重复，但是使得代码难以被理解。

### 5.4	错误

​		在Go中有一部分函数总是能成功的运行。比如strings.Contains和strconv.FormatBool函数，对各种可 能的输入都做了良好的处理，使得运行时几乎不会失败，除非遇到灾难性的、不可预料的情况，比如运 行时的内存溢出。导致这种错误的原因很复杂，难以处理，从错误中恢复的可能性也很低。

​		还有一部分函数只要输入的参数满足一定条件，也能保证运行成功。比如time.Date函数，该函数将年月 日等参数构造成time.Time对象，除非最后一个参数（时区）是nil。这种情况下会引发panic异常。 panic是来自被调函数的信号，表示发生了某个已知的bug。一个良好的程序永远不应该发生panic异常。

​		在Go的错误处理中，错误是软件包API和应用程序用户界面的一个重要组成部分，程序运行失败仅被认为 是几个预期的结果之一。

​		对于那些将运行失败看作是预期结果的函数，它们会返回一个额外的返回值，通常是最后一个，来传递 错误信息。如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为ok。比如， cache.Lookup失败的唯一原因是key不存在，那么代码可以按照下面的方式组织：

```go
value, ok := cache.Lookup(key)
if !ok {
	// ...cache[key] does not exist…
}
```

​		通常，导致失败的原因不止一种，尤其是对I/O操作而言，用户需要了解更多的错误信息。因此，额外的 返回值不再是简单的布尔类型，而是error类型。

​		内置的error是接口类型。error类型可能是nil或者non-nil。nil意味着函数运行成功，non-nil表示失败。对于non-nil 的error类型,我们可以通过调用error的Error函数或者输出函数获得字符串类型的错误信息。

```
fmt.Println(err)
fmt.Printf("%v", err)
```

​		，当函数返回non-nil的error时，其他的返回值是未定义的(undefined),这些未定义的返回值应该被忽略。然而，有少部分函数在发生错误时，仍然会返回一些有用的返回值。比如，当读取文件发生错误时，Read函数会返回可以读取的字节数以及错误信息。对于这种情况，正确的处理方式应该是先处理这些不完整的数据，再处理错误。因此对函数的返回值要有清晰的说明，以便于其他人使用。

​		在Go中，函数运行失败时会返回错误信息，这些错误信息被认为是一种预期的值而非异常（exception）

#### 5.4.1	错误处理策略

​		是最常用的方式是传播错误。这意味着函数中某个子程序的失败，会变成该函数的失败。下面节的findLinks函数作为例子。如果findLinks对http.Get的调用失败，findLinks会直接将这个HTTP错误返回给调用者：

```go
resp, err := http.Get(url)
if err != nil{
	return nill, err
}
```

​		当对html.Parse的调用失败时，findLinks不会直接返回html.Parse的错误，因为缺少两条重要信息：

​		1、错误发生在解析器；2、url已经被解析。这些信息有助于错误的处理，findLinks会构造新的错误信息返回给调用者：

```go
doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
	return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}
```

​		fmt.Errorf函数使用fmt.Sprintf格式化错误信息并返回。我们使用该函数前缀添加额外的上下文信息到原始错误信息。当错误最终由main函数处理时，错误信息应提供清晰的从原因到后果的因果链，就像美国宇航局事故调查时做的那样：

```go
genesis: crashed: no parachute: G-switch failed: bad relay orientation
```

​		由于错误信息经常是以链式组合在一起的，所以错误信息中应避免大写和换行符。

​		编写错误信息时，我们要确保错误信息对问题细节的描述是详尽的。尤其是要注意错误信息表达的一致 性，即相同的函数或同包内的同一组函数返回的错误在构成和处理方式上是相似的。

​		以OS包为例，OS包确保文件操作（如os.Open、Read、Write、Close）返回的每个错误的描述不仅仅包含错误的原因（如无权限，文件目录不存在）也包含文件名，这样调用者在构造新的错误信息时无需再添加这些信息。

​		一般而言，被调函数f(x)会将调用信息和参数信息作为发生错误时的上下文放在错误信息中并返回给调 用者，调用者需要添加一些错误信息中不包含的信息，比如添加url到html.Parse返回的错误中。

​		的第二种策略。如果错误的发生是偶然性的，或由不可预知的问题导致的。一个 明智的选择是重新尝试失败的操作。在重试时，我们需要限制重试的时间间隔或重试的次数，防止无限 制的重试

```go
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
```

​		用第三种策略：输出错误信息并结束程序。需要注 意的是，这种策略只应在main中执行。对库函数而言，应仅向上传播错误，除非该错误意味着程序内部 包含不一致性，即遇到了bug，才能在库函数中结束程序。

```go
// (In function main.)
if err := WaitForServer(url); err != nil {
	fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
	os.Exit(1)
}
```

​		调用log.Fatalf可以更简洁的代码达到与上文相同的效果。log中的所有函数，都默认会在错误信息之前输出时间信息。

```go
if err := WaitForServer(url); err != nil {
	log.Fatalf("Site is down: %v\n", err)
}
```

​		长时间运行的服务器常采用默认的时间格式，而交互式工具很少采用包含如此多信息的格式。

```
2006/01/02 15:04:05 Site is down: no such domain:
bad.gopl.io
```

​		我们可以设置log的前缀信息屏蔽时间信息，一般而言，前缀信息会被设置成命令名。

```
log.SetPrefix("wait: ")
log.SetFlags(0)
```

​		第四种策略：有时，我们只需要输出错误信息就足够了，不需要中断程序的运行。我们可以通过log包提供函数

```
if err := Ping(); err != nil {
	log.Printf("ping failed: %v; networking disabled",err)
}
```

​		或者标准错误流输出错误信息。

```
if err := Ping(); err != nil {
	fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
}
```

​		log包中的所有函数会为没有换行符的字符串增加换行符。

​		第五种，也是最后一种策略：我们可以直接忽略掉错误。

```go
dir, err := ioutil.TempDir("", "scratch")
if err != nil {
	return fmt.Errorf("failed to create temp dir: %v", err)
}
// ...use temp dir…
os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically
```

​		在Go中，错误处理有一套独特的编码风格。检查某个子函数是否失败后，我们通常将处理失败的逻辑代 码放在处理成功的代码之前。如果某个错误会导致函数返回，那么成功时的逻辑代码不应放在else语句 块中，而应直接放在函数体中。Go中大部分函数的代码结构几乎相同，首先是一系列的初始检查，防止 错误发生，之后是函数的实际逻辑。

#### 5.4.2. 文件结尾错误（EOF）

​		io包保证任何由文件结束引起的读取失败都返回同一个错误——io.EOF，该错误在io包中定义：

```go
package io
import "errors"
// EOF is the error returned by Read when no more input is available.
var EOF = errors.New("EOF")
```

​		调用者只需通过简单的比较，就可以检测出这个错误.下面的例子展示了如何从标准输入中读取字符， 以及判断文件结束。

```go
in := bufio.NewReader(os.Stdin)
for {
	r, _, err := in.ReadRune()
	if err == io.EOF {
		break // finished reading
	}
	if err != nil {
		return fmt.Errorf("read failed:%v", err)
	}
	// ...use r…
}
```

​		因为文件结束这种错误不需要更多的描述，所以io.EOF有固定的错误信息——“EOF”。对于其他错误，我们可能需要在错误信息中描述错误的类型和数量，这使得我们不能像io.EOF一样采用固定的错误信息。

### 5.5	函数值

​		在Go中，函数被看作第一类值（first-class values）：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。对函数值（function value）的调用类似函数调用。例子如下：

```go
func square(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }

f := square
fmt.Println(f(3)) // "9"

f = negative
fmt.Println(f(3)) // "-3"
fmt.Printf("%T\n", f) // "func(int) int"

f = product // compile error: can't assign func(int, int) int to func(int) int
```

​		函数类型的零值是nil。调用值为nil的函数值会引起panic错误：

```go
var f func(int) int
f(3) 					// 此处f的值为nil, 会引起panic错误
```

​		函数值可以与nil比较：

```go
var f func(int) int
if f != nil {
	f(3)
}
```

​		但是函数值之间是不可比较的，也不能用函数值作为map的key。

​		函数值使得我们不仅仅可以通过数据来参数化函数，亦可通过行为。标准库中包含许多这样的例子。strings.Map对字符串中的每个字符调用add1函数，并将每个add1函数的返回值组成一个新的字符串返回给调用者。

```go
func add1(r rune) rune { return r + 1 }
fmt.Println(strings.Map(add1, "HAL-9000"))   // "IBM.:111"
fmt.Println(strings.Map(add1, "VMS")) 		 // "WNT"
fmt.Println(strings.Map(add1, "Admix"))		 // "Benjy"
```

​		findLinks函数使用了辅助函数visit,遍历和操作了HTML页面的所有结点。使用函数值，我们可以将遍历结点的逻辑和操作结点的逻辑分离，使得我们可以复用遍历的逻辑，从而对结点进行不同的操作。

```go
// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
```

​		该函数接收2个函数作为参数，分别在结点的孩子被访问前和访问后调用。这样的设计给调用者更大的灵活性。举个例子，现在我们有startElemen和endElement两个函数用于输出HTML元素的开始标签和结束标签<b>...</b>：

```go
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
```

​		上面的代码利用fmt.Printf的一个小技巧控制输出的缩进。%*s中的*会在字符串之前填充一些空格。在例子中,每次输出会先填充depth*2数量的空格，再输出""，最后再输出HTML标签。

​		如果我们像下面这样调用forEachNode：

```
forEachNode(doc, startElement, endElement)
```

​		与之前的outline程序相比，我们得到了更加详细的页面结构：

```html
$ go build gopl.io/ch5/outline2
$ ./outline2 http://gopl.io
<html>
<head>
<meta>
</meta>
<title>
</title>
<style>
</style>
</head>
<body>
<table>
<tbody>
<tr>
<td>
<a>
<img>
</img>
...
```

### 5.6	匿名函数

​		拥有函数名的函数只能在包级语法块中被声明，通过函数字面量（function literal），我们可绕过这一限制，在任何表达式中表示一个函数值。函数字面量的语法和函数声明相似，区别在于func关键字后没有函数名。函数值字面量是一种表达式，它的值被成为匿名函数（anonymous function）。

​		函数字面量允许我们在使用时函数时，再定义它。

```go
strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
```

​		更为重要的是，通过这种方式定义的函数可以访问完整的词法环境（lexical environment），这意味着在函数中定义的内部函数可以引用该函数的变量，如下例所示：

```go
// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
```

​		函数squares返回另一个类型为 func() int 的函数。对squares的一次调用会生成一个局部变量x并返回一个匿名函数。每次调用时匿名函数时，该函数都会先使x的值加1，再返回x的平方。第二次调用squares时，会生成第二个x变量，并返回一个新的匿名函数。新匿名函数操作的是第二个x变量。

​		squares的例子证明，函数值不仅仅是一串代码，还记录了状态。在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。这就是函数值属于引用类型和函数值不可比较的原因。

​		通过这个例子，我们看到变量的生命周期不由它的作用域决定：squares返回后，变量x仍然隐式的存在于f中。

​		给定一些计算机课程，每个课程都有前置课程，只有完成了前置课程才可以开始当前课程的学习；我们的目标是选择出一组课程，这组课程必须确保按顺序学习时，能全部被完成。每个课程的前置课程如下：

```go
var prereqs = map[string][]string{
	"算法":  {"数据结构"},
	"微积分": {"线性代数"},
	"编译器": {
		"数据结构",
		"形式语言",
		"计算机组织",
	},
	"数据结构": {"离散数学"},
	"数据库":  {"数据结构"},
	"离散数学": {"编程入门"},
	"形式语言": {"离散数学"},
	"网络":   {"操作系统"},
	"操作系统": {"数据结构", "计算机组织"},
	"编程语言": {"数据结构", "计算机组织"},
}
```

​		这类问题被称作拓扑排序。从概念上说，前置条件可以构成有向图。图中的顶点表示课程，边表示课程间的依赖关系。显然，图中应该无环，这也就是说从某点出发的边，最终不会回到该点。下面的代码用深度优先搜索了整张图，获得了符合要求的课程序列。

```go
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
```

​		当匿名函数需要被递归调用时，我们必须首先声明一个变量，再将匿名函数赋值给这个变量。如果不分成两部，函数字面量无法与visitAll绑定，我们也无法递归调用该匿名函数。

```go
visitAll := func(items []string) {
	// ...
	visitAll(m[item]) // compile error: undefined: visitAll
	// ...
}
```

​		在topsort中，首先对prereqs中的key排序，再调用visitAll。因为prereqs映射的是切片而不是更复杂的map，所以数据的遍历次序是固定的，这意味着你每次运行topsort得到的输出都是一样的。 topsort的输出结果如下:

```
1:	编程入门
2:	离散数学
3:	形式语言
4:	线性代数
5:	微积分
6:	数据结构
7:	计算机组织
8:	操作系统
9:	数据库
10:	算法
11:	编程语言
12:	编译器
13:	网络
```

​		让我们回到findLinks这个例子。新的匿名函数被引入，用于替换原来的visit函数。该匿名函数负责将新连接添加到切片中。在Extract中，使用forEachNode遍历HTML页面，由于Extract只需要在遍历结点前操作结点，所以forEachNode的post参数被传入nil。

```go
package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}


```

​		上面的代码对之前的版本做了改进，现在links中存储的不是href属性的原始值，而是通过resp.Request.URL解析后的值。解析后，这些连接以绝对路径的形式存在，可以直接被http.Get访问。

​		网页抓取的核心问题就是如何遍历图。在topoSort的例子中，已经展示了深度优先遍历，在网页抓取中，我们会展示如何用广度优先遍历图。

​		下面的函数实现了广度优先算法。调用者需要输入一个初始的待访问列表和一个函数f。待访问列表中的每个元素被定义为string类型。广度优先算法会为每个元素调用一次f。每次f执行完毕后，会返回一组待访问元素。这些元素会被加入到待访问列表中。当待访问列表中的所有元素都被访问后，breadthFirst函数运行结束。为了避免同一个元素被访问两次，代码中维护了一个map。

```go
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
```

​		将f返回的一组元素一个个添加到worklist中。

​		在我们网页抓取器中，元素的类型是url。crawl函数会将URL输出，提取其中的新链接，并将这些新链接返回。我们会将crawl作为参数传递给breadthFirst。

```go
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
```

​		为了使抓取器开始运行，我们用命令行输入的参数作为初始的待访问url。

```go
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
```

​		让我们从 https://golang.org 开始，下面是程序的输出结果：

```
$ go build gopl.io/ch5/findlinks3
$ ./findlinks3 https://golang.org
https://golang.org/
https://golang.org/doc/
https://golang.org/pkg/
https://golang.org/project/
https://code.google.com/p/go-tour/
https://golang.org/doc/code.html
https://www.youtube.com/watch?v=XCsL89YtqCs
http://research.swtch.com/gotour
```

​		当所有发现的链接都已经被访问或电脑的内存耗尽时，程序运行结束。

#### 5.6.1	警告：捕获迭代变量

​		考虑这个样一个问题：你被要求首先创建一些目录，再将目录删除。在下面的例子中我们用函数值来完成删除操作。下面的示例代码需要引入os包。为了使代码简单，我们忽略了所有的异常处理。

```go
var rmdirs []func()
for _, d := range tempDirs() {
	dir := d               // NOTE: necessary!
	os.MkdirAll(dir, 0755) // creates parent directories too
	rmdirs = append(rmdirs, func() {
		os.RemoveAll(dir)
	})
}
// ...do some work…
for _, rmdir := range rmdirs {
	rmdir() // clean up
}
```

​		为什么要在循环体中用循环变量d赋值一个新的局部变量，而不是像下面的代码一样直接使用循环变量dir。需要注意，下面的代码是错误的。

```go
var rmdirs []func()
for _, dir := range tempDirs() {
	os.MkdirAll(dir, 0755)
	rmdirs = append(rmdirs, func() {
		os.RemoveAll(dir) // NOTE: incorrect!
	})
}
```

​		问题的原因在于循环变量的作用域。在上面的程序中，for循环语句引入了新的词法块，循环变量dir在这个词法块中被声明。在该循环中生成的所有函数值都共享相同的循环变量。需要注意，函数值中记录的是循环变量的内存地址，而不是循环变量某一时刻的值。以dir为例，后续的迭代会不断更新dir的值，当删除操作执行时，for循环已完成，dir中存储的值等于最后一次迭代的值。这意味着，每次对os.RemoveAll的调用删除的都是相同的目录。

​		通常，为了解决这个问题，我们会引入一个与循环变量同名的局部变量，作为循环变量的副本。比如下面的变量dir，虽然这看起来很奇怪，但却很有用。

```go
for _, dir := range tempDirs() {
	dir := dir // declares inner dir, initialized to outer dir
	// ...
}
```

​		这个问题不仅存在基于range的循环，在下面的例子中，对循环变量i的使用也存在同样的问题：

```go
var rmdirs []func()
dirs := tempDirs()
for i := 0; i < len(dirs); i++ {
	os.MkdirAll(dirs[i], 0755) // OK
	rmdirs = append(rmdirs, func() {
		os.RemoveAll(dirs[i]) // NOTE: incorrect!
	})
}
```

### 5.7	可变参数

​		参数数量可变的函数称为为可变参数函数。典型的例子就是fmt.Printf和类似函数。Printf首先接收一 个的必备参数，之后接收任意个数的后续参数。

​		在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号“...”，这表示该函数 会接收任意数量的该类型参数。

```go
func sum(vals...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
```

​		sum函数返回任意个int型参数的和。在函数体中,vals被看作是类型为[] int的切片。sum可以接收任意 数量的int型参数

```go
fmt.Println(sum()) 			 // "0"
fmt.Println(sum(3))			 // "3"
fmt.Println(sum(1, 2, 3, 4)) // "10
```

​		在上面的代码中，调用者隐式的创建一个数组，并将原始参数复制到数组中，再把数组的一个切片作为 参数传给被调函数。如果原始参数已经是切片类型，我们该如何传递给sum？只需在最后一个参数后加上 省略符。下面的代码功能与上个例子中最后一条语句相同。

```go
values := []int{1, 2, 3, 4}
fmt.Println(sum(values...)) 	// "10"
```

​		虽然在可变参数函数内部，...int 型参数的行为看起来很像切片类型，但实际上，可变参数函数和以切片作为参数的函数是不同的。

```
func f(...int) {}
func g([]int) {}
fmt.Printf("%T\n", f) // "func(...int)"
fmt.Printf("%T\n", g) // "func([]int)"
```

​		可变参数函数经常被用于格式化字符串。下面的errorf函数构造了一个以行号开头的，经过格式化的错误信息。函数名的后缀f是一种通用的命名规范，代表该可变参数函数可以接收Printf风格的格式化字符串。

```go
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
linenum, name := 12, "count"
errorf(linenum, "undefined: %s", name) // "Line 12: undefined: count"
```

### 5.8	Deferred函数

​		在findLinks的例子中，我们用http.Get的输出作为html.Parse的输入。只有url的内容的确是HTML格式的，html.Parse才可以正常工作，但实际上，url指向的内容很丰富，可能是图片，纯文本或是其他。将这些格式的内容传递给html.parse，会产生不良后果。

​		下面的例子获取HTML页面并输出页面的标题。title函数会检查服务器返回的Content-Type字段，如果发现页面不是HTML，将终止函数运行，返回错误。

```go
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// Check Content-Type is HTML (e.g., "text/html;charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}
```

​		下面展示了运行效果：

```
$ go build gopl.io/ch5/title1
$ ./title1 http://gopl.io
The Go Programming Language
$ ./title1 https://golang.org/doc/effective_go.html
Effective Go - The Go Programming Language
$ ./title1 https://golang.org/doc/gopher/frontpage.png
title: https://golang.org/doc/gopher/frontpage.png has type image/png, not text/html
```

​		resp.Body.close调用了多次，这是为了确保title在所有执行路径下（即使函数运行失败）都关闭了网络连接。随着函数变得复杂，需要处理的错误也变多，维护清理逻辑变得越来越困难。而Go语言独有的defer机制可以让事情变得简单。

​		你只需要在调用普通函数或方法前加上关键字defer，就完成了defer所需要的语法。当defer语句被执行时，跟在defer后面的函数会被延迟执行。直到包含该defer语句的函数执行完毕时，defer后的函数才会被执行，不论包含defer语句的函数是通过return正常结束，还是由于panic导致的异常结束。你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。

​		defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。释放资源的defer应该直接跟在请求资源的语句后。在下面的代码中，一条defer语句替代了之前的所有resp.Body.Close

​		

```go
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	// ...print doc's title element…
	return nil
}
```

​		在处理其他资源时，也可以采用defer机制，比如对文件的操作：

```go
package ioutil

import "os"

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadAll(f)
}

```

​		或是处理互斥锁

```go
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}
```

​		调试复杂程序时，defer机制也常被用于记录何时进入和退出函数。下例中的bigSlowOperation函数，直接调用trace记录函数的被调情况。bigSlowOperation被调时，trace会返回一个函数值，该函数值会在bigSlowOperation退出时被调用。通过这种方式， 我们可以只通过一条语句控制函数的入口和所有的出口，甚至可以记录函数的运行时间，如例子中的start。需要注意一点：不要忘记defer语句后的圆括号，否则本该在进入时执行的操作会在退出时执行，而本该在退出时执行的，永远不会被执行。

```go
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the	extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow	operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

```

​		每一次bigSlowOperation被调用，程序都会记录函数的进入，退出，持续时间。（我们用time.Sleep模拟一个耗时的操作）

```
$ go build gopl.io/ch5/trace
$ ./trace
2015/11/18 09:53:26 enter bigSlowOperation
2015/11/18 09:53:36 exit bigSlowOperation (10.000589217s)
```

​		我们知道，defer语句中的函数会在return语句更新返回值变量后再执行，又因为在函数中定义的匿名函数可以访问该函数包括返回值变量在内的所有变量，所以，对匿名函数采用defer机制，可以使其观察函数的返回值。

​		以double函数为例：

```go
func double(x int) int {
	return x + x
}
```

​		我们只需要首先命名double的返回值，再增加defer语句，我们就可以在double每次被调用时，输出参数以及返回值。

```go
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}
_ = double(4)
// Output:
// "double(4) = 8"

```

​		可能doulbe函数过于简单，看不出这个小技巧的作用，但对于有许多return语句的函数而言，这个技巧很有用。

​		被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值：

```go
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
fmt.Println(triple(4)) // "12
```

​		在循环体中的defer语句需要特别注意，因为只有在函数执行完毕后，这些被延迟的函数才会执行。下面的代码会导致系统的文件描述符耗尽，因为在所有文件都被处理之前，没有文件会被关闭。

```go
for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close() // NOTE: risky; could run out of file
		descriptors
		// ...process f…
	}
```

​		一种解决方法是将循环体中的defer语句移至另外一个函数。在每次循环时，调用这个函数。

```go
func main() {
	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			return err
		}
	}
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...process f…
}
```

​		下面的代码是fetch的改进版，我们将http响应信息写入本地文件而不是从标准输出流输出。我们通过path.Base提出url路径的最后一段作为文件名。

```go
// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
```

​		通过os.Create打开文件进行写入，在关闭文件时，我们没有对f.close采用defer机制，因为这会产生一些微妙的错误。许多文件系统，尤其是NFS，写入文件时发生的错误会被延迟到文件关闭时反馈。如果没有检查文件关闭时的反馈信息，可能会导致数据丢失，而我们还误以为写入操作成功。如果io.Copy和f.close都失败了，我们倾向于将io.Copy的错误信息反馈给调用者，因为它先于f,close发生，更有可能接近问题的本质。

### 		5.9. Panic异常

​		Go的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用 等。这些运行时错误会引起painc异常。

​		一般而言，当panic异常发生时，程序会中断运行，并立即执行在该goroutine，随后，程序崩溃并输出日志信息。日志信息包括 panic value和函数调用的堆栈跟踪信息。panic value通常是某种错误信息。对于每个goroutine，日志 信息中都会有与之相对的，发生panic时的函数调用堆栈跟踪信息。通常，我们不需要再次运行程序去定位问题，日志信息已经提供了足够的诊断依据。因此，在我们填写问题报告时，一般会将panic异常和日 志信息一并记录

​		不是所有的panic异常都来自运行时，直接调用内置的panic函数也会引发panic异常；panic函数接受任 何值作为参数。当某些不应该发生的场景发生时，我们就应该调用panic。比如，当程序到达了某条逻辑 上不可能到达的路径

```go
switch s := suit(drawCard()); s {
case "Spades": // ...
case "Hearts": // ...
case "Diamonds": // ...
case "Clubs": // ...
default:
	panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
}
```

​		断言函数必须满足的前置条件是明智的做法，但这很容易被滥用。除非你能提供更多的错误信息，或者能更快速的发现错误，否则不需要使用断言，编译器在运行时会帮你检查代码。

```go
func Reset(x *Buffer) {
	if x == nil {
		panic("x is nil") // unnecessary!
	}
	x.elements = nil
}
```

​		虽然Go的panic机制类似于其他语言的异常，但panic的适用场景有一些不同。由于panic会引起程序的崩溃，因此panic一般用于严重错误，如程序内部的逻辑不一致。所以对于大部分漏洞，我们应该使用Go提供的错误机制，而不是panic，尽量避免程序的崩溃。在健壮的程序中，任何可以预料到的错误，如不正确的输入、错误的配置或是失败的I/O操作都应该被优雅的处理，最好的处理方式，就是使用Go的错误机制。

​		考虑regexp.Compile函数，该函数将正则表达式编译成有效的可匹配格式。当输入的正则表达式不合法时，该函数会返回一个错误。当调用者明确的知道正确的输入不会引起函数错误时，要求调用者检查这个错误是不必要和累赘的。我们应该假设函数的输入一直合法，就如前面的断言一样：当调用者输入了不应该出现的输入时，触发panic异常。

​		在程序源码中，大多数正则表达式是字符串字面值（string literals），因此regexp包提供了包装函数regexp.MustCompile检查输入的合法性。

```go
package regexp

func Compile(expr string) (*Regexp, error) { /* ... */ }
func MustCompile(expr string) *Regexp {
	re, err := Compile(expr)
	if err != nil {
		panic(err)
	}
	return re
}

```

​		包装函数使得调用者可以便捷的用一个编译后的正则表达式为包级别的变量赋值：

```
var httpSchemeRE = regexp.MustCompile(`^https?:`) //"http:" or "https:"
```

​		显然，MustCompile不能接收不合法的输入。函数名中的Must前缀是一种针对此类函数的命名约定，比如template.Must

```go
func main() {
	f(3)
}
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
```

​		上例中的运行输出如下：

```
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
```

​		当f(0)被调用时，发生panic异常，之前被延迟执行的的3个fmt.Printf被调用。程序中断执行后，panic信息和堆栈信息会被输出（下面是简化的输出）：

```
panic: runtime error: integer divide by zero
main.f(0)
src/gopl.io/ch5/defer1/defer.go:14
main.f(1)
src/gopl.io/ch5/defer1/defer.go:16
main.f(2)
src/gopl.io/ch5/defer1/defer.go:16
main.f(3)
src/gopl.io/ch5/defer1/defer.go:16
main.main()
src/gopl.io/ch5/defer1/defer.go:10
```

​		为了方便诊断问题，runtime包允许程序员输出堆栈信息。在下面的例子中，我们通过在main函数中延迟调用printStack输出堆栈信息。

```go
func main() {
	defer printStack()
	f(3)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
```

​		printStack的简化输出如下（下面只是printStack的输出，不包括panic的日志信息）：

```
goroutine 1 [running]:
main.printStack()
src/gopl.io/ch5/defer2/defer.go:20
main.f(0)
src/gopl.io/ch5/defer2/defer.go:27
main.f(1)
src/gopl.io/ch5/defer2/defer.go:29
main.f(2)
src/gopl.io/ch5/defer2/defer.go:29
main.f(3)
src/gopl.io/ch5/defer2/defer.go:29
main.main()
src/gopl.io/ch5/defer2/defer.go:15
```

​		将panic机制类比其他语言异常机制的读者可能会惊讶，rumtime.Stack为何能输出已经被释放函数的信息？在Go的panic机制中，延迟函数的调用在释放堆栈信息之前。

### 5.10. Recover捕获异常

​		通常来说，不应该对panic异常做任何处理，但有时，也许我们可以从异常中恢复，至少我们可以在程序崩溃前，做一些操作。举个例子，当web服务器遇到不可预料的严重问题时，在崩溃前应该将所有的连接 关闭；如果不做任何处理，会使得客户端一直处于等待状态。如果web服务器还在开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助调试。

​		如果在deferred函数中调用了内置函数recover，并且定义该defer语句的函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value。导致panic异常的函数不会继续运行，但能正常返回。在未发生panic时调用recover，recover会返回nil。

​		让我们以语言解析器为例，说明recover的使用场景。考虑到语言解析器的复杂性，即使某个语言解析器目前工作正常，也无法肯定它没有漏洞。因此，当某个异常出现时，我们不会选择让解析器崩溃，而是会将panic异常当作普通的解析错误，并附加额外信息提醒用户报告此错误。

```go
func Parse(input string) (s *Syntax, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	// ...parser...
}
```

​		deferred函数帮助Parse从panic中恢复。在deferred函数内部，panic value被附加到错误信息中；并用err变量接收错误信息，返回给调用者。我们也可以通过调用runtime.Stack往错误信息中添加完整的堆栈调用信息。不加区分的恢复所有的panic异常，不是可取的做法；因为在panic之后，无法保证包级变量的状态仍然和我们预期一致。比如，对数据结构的一次重要更新没有被完整完成、文件或者网络连接没有被关闭、获得的锁没有被释放。此外，如果写日志时产生的panic被不加区分的恢复，可能会导致漏洞被忽略。

​		虽然把对panic的处理都集中在一个包下，有助于简化对复杂和不可以预料问题的处理，但作为被广泛遵守的规范，你不应该试图去恢复其他包引起的panic。公有的API应该将函数的运行失败作为error返回，而不是panic。同样的，你也不应该恢复一个由他人开发的函数引起的panic，比如说调用者传入的回调函数，因为你无法确保这样做是安全的。有时我们很难完全遵循规范，举个例子，net/http包中提供了一个web服务器，将收到的请求分发给用户提供的处理函数。很显然，我们不能因为某个处理函数引发的panic异常，杀掉整个进程；web服务器遇到处理函数导致的panic时会调用recover，输出堆栈信息，继续运行。这样的做法在实践中很便捷，但也会引起资源泄漏，或是因为recover操作，导致其他问题。

​		基于以上原因，安全的做法是有选择性的recover。换句话说，只恢复应该被恢复的panic异常，此外， 这些异常所占的比例应该尽可能的低。为了标识某个panic是否应该被恢复，我们可以将panic value设置成特殊类型。在recover时对panic value进行检查，如果发现panic value是特殊类型，就将这个panic作为errror处理，如果不是，则按照正常的panic进行处理（在下面的例子中，我们会看到这种方式）。

下面的例子是title函数的变形，如果HTML页面包含多个<title>，该函数会给调用者返回一个错误（error）。在soleTitle内部处理时，如果检测到有多个<title>，会调用panic，阻止函数继续递归，并将特殊类型bailout作为panic的参数。

```go
// soleTitle returns the text of the first non-empty title element
// in doc, and an error if there was not exactly one.
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
		// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()
	// Bail out of recursion if we find more than one nonempty title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple titleelements
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}
```

​		在上例中，deferred函数调用recover，并检查panic value。当panic value是bailout{}类型时，deferred函数生成一个error返回给调用者。当panic value是其他non-nil值时，表示发生了未知的panic异常，deferred函数将调用panic函数并将当前的panic value作为参数传入；此时，等同于recover没有做任何操作。（请注意：在例子中，对可预期的错误采用了panic，这违反了之前的建议，我们在此只是想向读者演示这种机制。）

​		有些情况下，我们无法恢复。某些致命错误会导致Go在运行时终止程序，如内存不足。







































