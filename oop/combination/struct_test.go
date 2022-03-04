package combination

import (
	"fmt"
	"testing"
	"time"
)

type Man interface {
	SetName(string)
}

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p *Person) SetName(n string) {
	p.Name = n
}

type Student struct {
	Person //通过匿名组合的方式嵌入了Person的属性
	Score  float64
}

func (s *Student) SetScore(score float64) {
	s.Score = score
}

type Teacher struct {
	Person //通过匿名组合的方式嵌入了Person的属性。
	Course string
}

type Schoolmaster struct {
	Person   //通过匿名组合的方式嵌入了Person的属性。
	CarBrand string
}

func TestStruct(t *testing.T) {

	/**
	  第一种初始化方式:先定义后赋值
	*/
	var m Man
	s1 := Student{}
	m = &s1
	fmt.Print(m)
	s1.SetName("test name")
	s1.SetScore(88.9)
	fmt.Println(s1)
	fmt.Printf("%+v\n\n", s1) //"+v表示打印结构体的各个字段"

	/**
	  第二种初始化方式:直接初始化
	*/
	s2 := Teacher{Person{"尹正杰", 18, "boy"}, "Go并发编程"}
	fmt.Println(s2)
	fmt.Printf("%+v\n\n", s2)

	/**
	  第三种赋值方式:初始化赋值部分字段
	*/
	s3 := Schoolmaster{CarBrand: "丰田", Person: Person{Name: "JasonYin最强王者"}}
	fmt.Println(s3)
	fmt.Printf("%+v\n", s3)

}

func TestStructCombination(t *testing.T) {
	type Animal struct {
		Age int
	}

	type People struct {
		Animal
		Name   string
		Age    int
		Gender string
	}

	type IdentityCard struct {
		IdCardNO    int
		Nationality string
		Address     string
		Age         int
	}

	/*
	   此时的Students以及是多重继承
	*/
	type Students struct {
		IdentityCard
		People //多层继承
		Age    int
		Score  int
	}

	/**
	  如果子类和父类存在同名的属性,那么以就近原则为准
	*/
	s1 := Students{
		Score: 150,
		Age:   27,
		IdentityCard: IdentityCard{
			IdCardNO:    110105199003072872,
			Nationality: "中华人民共和国",
			Address:     "北京市朝阳区望京SOHO",
			Age:         8,
		},
		People: People{Name: "Jason Yin", Age: 18, Animal: Animal{Age: 20}},
	}

	/**
	  如果子类和父类存在同名的属性(如果父类还继承了其它类型，我们称之为多层继承)，那么就以就近原则为准;
	  但是如果一个子类如果继承自多个父类(我们称之为多重继承),且每个字段中都有相同的字段,此时我们无法直接在子类调用该属性;
	*/
	fmt.Printf("学生的年龄是:[%d]\n", s1.Age)
	s1.Age = 21
	fmt.Printf("学生的年龄是:[%d]\n\n", s1.Age)

	//给People类的Age赋值
	fmt.Printf("People的年龄是:[%d]\n", s1.People.Age)
	s1.People.Age = 5000
	fmt.Printf("People的年龄是:[%d]\n\n", s1.People.Age)

	//给IdentityCard类的Age赋值
	fmt.Printf("IdentityCard的年龄是:[%d]\n", s1.IdentityCard.Age)
	s1.IdentityCard.Age = 80
	fmt.Printf("IdentityCard的年龄是:[%d]\n", s1.IdentityCard.Age)

}

func TestStructPoint(t *testing.T) {
	type Vehicle struct {
		Brand string
		Wheel byte
	}

	type Car struct {
		Vehicle
		Colour string
	}

	type Driver struct {
		*Car
		DrivingTime time.Time
	}

	/**
	  对象指针匿名组合的第一种初始化方式:
	      定义时直接初始化赋值。
	*/
	d1 := Driver{&Car{
		Vehicle: Vehicle{
			Brand: "丰田",
			Wheel: 4,
		},
		Colour: "红色",
	}, time.Now(),
	}
	//打印结构体的详细信息,注意观察指针对象
	fmt.Printf("%+v\n", d1)
	//我们可以直接调用对象的属性
	fmt.Printf("品牌:%s,颜色:%s\n", d1.Brand, d1.Colour)
	fmt.Printf("驾驶时间:%+v\n\n", d1.DrivingTime)
	time.Sleep(1000000000 * 3)

	/**
	  对象指针匿名组合的第二种初始化方式:
	      先声明，再赋值
	  温馨提示:
	      遇到指针的情况一定要避免空(nil)指针,未初始化的指针默认值是nil,可以考虑使用new函数解决。
	*/
	var d2 Driver
	/**
	  由于Driver结构体中有一个对象指针匿名组合Car,因此我们需要使用new函数申请空间。
	*/
	d2.Car = new(Car)
	d2.Brand = "奔驰"
	d2.Colour = "黄色"
	d2.DrivingTime = time.Now()
	fmt.Printf("%+v\n", d2)
	fmt.Printf("品牌:%s,颜色:%s\n", d2.Brand, d2.Colour)
	fmt.Printf("驾驶时间:%+v\n", d1.DrivingTime)

}

//定义一个结构体
type Lecturer struct {
	Name string
	Age  uint8
}

//我们为Lecturer结构体封装Init成员方法
func (l *Lecturer) Init() {
	l.Name = "Jason Yin"
	l.Age = 20
}

/**
我为Lecturer结构体起一个别名
我们可以为Instructor类型添加成员方法,
通过别名和成员方法为原有类型赋值新的操作
*/
type Instructor Lecturer

/**
温馨提示:
    (1)我们为一个结构体创建成员方法时，如果成员方法有接收者,需要考虑以下两种情况:
        1>.如果这个接收者是对象的时候，是值传递;
        2>.如果这个接收者是对象指针，是引用传递;
    (2)只要函数接收者不同,哪怕函数名称相同，也不算同一个函数哟;
    (3)不管接收者变量名称是否相同,只要类型一致(包括对象和对象指针),那么我们就认为接收者是相同的,这时候不允许出现相同名称函数;
    (4)给指针添加方法的时候，不允许给指针类型添加操作(因为Go语言中指针类型是只读的);
*/
func (i *Instructor) Init() {
	i.Name = "尹正杰"
	i.Age = 18
}

func TestStructFun(t *testing.T) {
	var (
		l Lecturer
		i Instructor
	)

	//可以使用对象调用成员方法
	i.Init()
	fmt.Printf("%+v\n\n", i)

	//可以用对象指针调用成员方法
	(&l).Init()
	fmt.Printf("%+v\n", l)
}
