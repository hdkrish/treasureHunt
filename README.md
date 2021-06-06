# treasureHunt
treasureHunt code in GO


Treasure Hunt

A princess must cross treacherous path to reach a destination to find “elixir of life” (alias Treasure).
This path consists of n number of steps (including starting point and treasure) - each having a sign. A sign is identified by an integer number. Depending on the sign number, princess must hop those many steps forward or backward or may become a zombie forever. 
Given an input array of signs, can you write a program to find out whether the princess would ever reach her destination?
Sample scenarios

Expected Input
 (An array of numbers)	              Expected Output
                                      (An array of the step signs if she reaches destination, 0 – if becomes zombie, -1 – if caught in infinite loop)
2, 1, 2, 0, 2, -1, Treasure	          2,2,2, Treasure
2, 4, 3, 0, -1, -1, Treasure	        0
2, 10, 1, -1, 1, -2, Treasure        -1


Kindly change the in-line array and run the program
go run main.go
