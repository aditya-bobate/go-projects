package main

import "fmt"

func main() {
	p1 := "Aditya"
	p2 := "Rahul"
	p3 := "Shravni"
	s1 := 70
	s2 := 50
	s3 := 100
	fmt.Println("\tPlayers Scores:")
	msg1 := fmt.Sprintf("%s	=%d points", p1, s1)
	msg2 := fmt.Sprintf("%s	=%d points", p2, s2)
	msg3 := fmt.Sprintf("%s	=%d points", p3, s3)
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)
}

/* ENTIRE SCOREBOARD IN ONE Sprintf() LINE
package main

import "fmt"

func main() {
    p1, s1 := "Aditya", 70
    p2, s2 := "Rahul", 50
    p3, s3 := "Shravni", 100

    scoreboard := fmt.Sprintf(
        "\tPlayer Scores:\n1. %-8s - %d points\n2. %-8s - %d points\n3. %-8s - %d points\n",
        p1, s1,
        p2, s2,
        p3, s3,
    )

    fmt.Println(scoreboard)
}
*/
