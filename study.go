package main

import (
	"fmt"
	//"runtime"
	//"time"
)

/*
type BaseTest struct {
}

func (b *BaseTest) Test() {
	fmt.Println("basetest")
}

type Test struct {
	BaseTest
}

func (b *Test) Test() {
	b.BaseTest.Test()
	fmt.Println("test")
}

const MAXN = 200

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func My_Sleep() {
	time.Sleep(10 * time.Second)
}
func fun(a *[]int) {
	*a = append(*a, 6)
}

*/

/*

type CppClass interface {
	Init()
	Des()
}

type BaseClass struct {
}

func (b *BaseClass) Init() {

}

func (b *BaseClass) Des() {

}

type SelfClass struct {
	*BaseClass
	Age int
}

func (sc *SelfClass) Init() {
	sc.Age = 100
}

func (sc *SelfClass) Print() {
	fmt.Println(sc.Age)
}

func (sc *SelfClass) SetAge(age int, age2 int) {
	sc.Age = age + age2
}

func (sc *SelfClass) SetAge(age int) {
	sc.Age = age
}

func SetAge_int(age int)
func SetAge_int_int(age int, age2 int)

func test() {
	fmt.Printf("abc")
	time.Sleep(10 * 10000)
}
*/
type SelfClass struct {
	Num int
	id  int
}

func (sc *SelfClass) printf() {
	fmt.Printf("%d", sc.Num)
}

func main() {
	my_sc := &SelfClass{}
	my_sc.Num = 1
	my_sc.printf()
	//sc := &SelfClass{}
	//sc.Print()

	//a := []int{1, 2, 3, 4}
	//fmt.Println(a)
	/*
		t1 := &Test{}
		t1.Test()
		var dp, w, t [MAXN]int
		var n, s int
		//fmt.Scanf("%d", &n)
		go My_Sleep()
		fmt.Printf("23")
		for i := 1; i <= n; i++ {
			fmt.Scanf("%d%d", &w[i], &t[i])
			s += t[i]
		}
		for i := 1; i <= n; i++ {
			for j := s; j >= w[i]+t[i]; j-- {
				dp[j] = Max(dp[j], dp[j-w[i]-t[i]]+t[i])
			}
		}
		fmt.Printf("%d\n", s-dp[s])
	*/
}
