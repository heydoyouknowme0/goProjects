package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)
var clear map[string]func() 
func init() {
    clear = make(map[string]func()) 
    clear["linux"] = func() { 
        cmd := exec.Command("clear") 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] 
    if ok { 
        value() 
    }
}
var board = [3][3]rune{}
var total =0
func setOrResetBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = ' '
		}
	}
	total = 0
}
func displayBoard() {
	fmt.Println("┌───┬───┬───┐")
	for i := 0; i < 3; i++ {
		fmt.Printf("| %c | %c | %c |\n", board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println("├───┼───┼───┤")
		}
	}
	fmt.Println("└───┴───┴───┘")
}
func Play() {
	var i,j int
	var player rune = 'X'
	
	if  input(i,j,player){
		return
	}
	total ++
	if total==5 {
		fmt.Println("Draw!")
		setOrResetBoard()
		return
	}
	player='O'
	if input(i,j,player) {
		return
	}
	Play()
}
func input(i,j int, player rune) bool {
	fmt.Printf("player %c,Enter the row and column: ", player)
	fmt.Scanln(&i,&j)
	i=i-1
	j=j-1
	for set(i,j,player) {
	    fmt.Printf("player %c,Enter the row and column: ", player)
		fmt.Scanln(&i,&j)
		i=i-1
		j=j-1
	}
	return checkWin(i, j, player)
}

func set(i int, j int, player rune)bool {
	if i>2 || i<0 || j<0 || j>2 || board[i][j] !=32 { 
		fmt.Println("Invalid Response, try again")
		return true}
	board[i][j] = player
	return false
}
func checkWin(i, j int, player rune) bool {
	CallClear()
	displayBoard()
	if (board[i][0] + board[i][1] == 2*board[i][2]) || (board[0][j] + board[1][j] ==2* board[2][j]) {
		fmt.Println("You won!")
		setOrResetBoard()
		return true
	}
	
	return false
}
func main(){
	var continueGame bool = true
    setOrResetBoard()
	displayBoard()
	for continueGame{
		Play()
		fmt.Println("Do you want to continue? (y/n)")
		var response string
		fmt.Scanln(&response)
		if response == "n" {
			continueGame = false
		}
	}
	return
}