package main

import "fmt"

func main(){
	b := [4]int{}
	d := b[0: 3]
	
	fmt.Printf("b=%v, d=%v, len(d)=%v, cap(d)=%v\n",b, d, len(d), cap(d))
	// ① b=[0 0 0 0], d=[0 0 0], len(d)=3, cap(d)=4
	
	d[0] = 1
	fmt.Printf("b=%v, d=%v, len(d)=%v, cap(d)=%v\n",b, d, len(d), cap(d))
	// ② b=[1 0 0 0], d=[1 0 0], len(d)=3, cap(d)=4
	
	d = append(d, 2)
	fmt.Printf("b=%v, d=%v, len(d)=%v, cap(d)=%v\n",b, d, len(d), cap(d))
	// ③ b=[1 0 0 2], d=[1 0 0 2], len(d)=4, cap(d)=4
}