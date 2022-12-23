package main

import "fmt"

func main() {
	numArr := [4]int{3, 2, 1, 5}
	var key, i, arrInI int
	fmt.Println(numArr)
	for j := 1; j <= len(numArr)-1; j++ {
		key = numArr[j]
		i = j - 1
		arrInI = numArr[i]
		fmt.Println(j, "-", key, i, "-", arrInI)
		// Let`s find out if iteration happens when
		// i > 0 and arrInI > key
		if i > 0 && arrInI > key {
			fmt.Println("In iteration", j, "condition i > 0 and arrInI > key is met")
			fmt.Println("This iteration have: j - 1 =", j-1, "key =", key, "i =", i)
			fmt.Println("And iteration", j, "have numArr[i] =", numArr[i])
			fmt.Println("Condition where A[i+1] = A[i] has the form:")
			fmt.Print("A[", i+1, "] = ", numArr[i+1], ":= A[", i, "] = ", numArr[i])
			fmt.Println("")
			numArr[i+1] = numArr[i]
			fmt.Println(numArr)
			i--
			fmt.Println(i)
		}
		fmt.Println("In iteration", j, "A[i+1] =", numArr[i+1], "key =", key)
		numArr[i+1] = key
		fmt.Println(numArr)
	}
}
