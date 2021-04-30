package main

import (
	"fmt"
	"os"
	"strconv"
)
//Constructing a conversion table to translate between integer value and string representation
var convtab = map[int]string{0:"Zero",1:"One",2:"Two",3:"Three",4:"Four",
							5:"Five", 6:"Six",7:"Seven",8:"Eight",9:"Nine"}

func main(){
	//read the command line input
	args :=os.Args
	//phoneticList is the array that we store our output strings
	var phoneticList []string
	//for every input in the command line input
	for i:=0;i< len(args)-1;i++{
		negative := 0
		//convert it to integer value
		num,err:=strconv.Atoi(args[i+1])
		if err!=nil{
			panic(err)
		}
		//if num<0, we want the "Negative" in our string representation
		if num<0 {
			num  = -num
			negative = 1
		}
		for{
			//if num<10, we can directly out it into array
			if num<10{
				//for ith position in the array, if it is empty
				if len(phoneticList)<i+1 {
					//append the string representation to the list
					phoneticList = append(phoneticList, convtab[num])
				}else {
					//else, append the string representation of this digit to the front of the existing string
					phoneticList[i] = convtab[num]+phoneticList[i]
				}
				//if negative flag is true, we want to add "Negative" up front
				if negative==1{
					phoneticList[i] = "Negative"+phoneticList[i]
				}
				// if it is not the last integer in the input, we want a comma at the end of it
				if i<len(args)-2 {
					phoneticList[i] = phoneticList[i]+","
				}
				break
			}
			//if num>=10, we want to get its residual
			residual :=num%10
			//then num represents the rest of digits
			num = (num-residual)/10
			//here we convert residual to string representation and put it into array
			if len(phoneticList)<i+1 {
				phoneticList = append(phoneticList,convtab[residual])
			}else {
				phoneticList[i] = convtab[residual]+phoneticList[i]
			}

		}
	}
	//Print out the output
	for i:=0;i<len(phoneticList);i++{
		fmt.Printf(phoneticList[i])
	}
	//new line to avoid '%' warning
	fmt.Printf("\n")
}

