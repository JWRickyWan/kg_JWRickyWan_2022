# Kargo Software Engineer Internship Summer 2022 Coding Assessment
## Author: Jianghong Wan

Convert an array of integers into an array of strings representing the phonetic equivalent of the
integer.
For example:

Given an array: [3, 25, 209], print “Three,TwoFive,TwoZeroNine” into stdout.

Given an array: [10, 300, 5], print “OneZero,ThreeZeroZero,Five” into stdout.

### Approach
Every integer N, denote N[i] as its ith digit couting from right to left, and then the integer N can be expressed as 
<p align = "center">
<img src="https://latex.codecogs.com/gif.latex?\sum&space;^{number\_of\_digits-1}_{i=0}N[i]*10^i" title="\sum ^{number\_of\_digits-1}_{i=0}N[i]*10^i" />
</p>
Therefore we can apply the residual operation (N%10) to get the left-most digit of N. Afterwards, N-(N%10) is the integer with the left-most digit removed. We can repeat this procedure to keep isolating the left-most digit.

### Algorithm

We first construct a map called convtab, which has integer 0 to 9 as keys, and their phone phonetic equivalent as values, respectively. For example,convtab[5] = "Five".

The program iterates through every input integer. For each integer, the program will first check if it starts with 0 but it's not zero value. For example, 0 is valid, but 02 is not a valid input. If input it invalid, the program will panic and you need to run it again with valid input. 

Then the input is converted into integer, since Golang represents command line input as string and the program immediately checks if the integer is negative. If so, negate the number, and the program will add "Negative" to the front of the output at last.

After those checks, the program enters an infinite loop
Two conditions may apply: 
1. If N is less than 10, then directly it will be used as a key to trace its phonetic equivalentin the convtab, and append to the output list, and the loop breaks.

2. Otherwise, we do the procedure in the previous section, get the left-most digit of N, and reduce N, and go into the next iteration, until N <10 so that it enters condition 1 and will eventually break the loop

after the infinite loop, the program will print the ouputs stored in the output array. 


### How to run
The program to run is [main.go](main.go). For Golang, to run main.go, you need to first build it using 
```
go build main.go
```
Then run it with 
```
go run main.go input1 input2 input3
```
### Complexity

The time complexity of this program should be O(N*L), and space complexity should be O(N), where N is the input size and L is the number of digits
