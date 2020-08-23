package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
	"flag"
)

func Bar(games, matches int){
	fmt.Printf("########################### %d | %d ###########################\n", games, matches)
}

func Clear(){
	fmt.Println("\033[H\033[2J")
}

func Sleep(secs time.Duration){
	time.Sleep(secs * time.Second)
}


func playerPlay(choice string) string{
	switch choice{

	case "rock","r":
		choice = "rock"

	case "paper","p":
		choice = "paper"

	case "scissors","s":
		choice = "scissors"
	}

	return choice

}

func botPlay(choice int) string{
	var win string
	switch choice {
	case 0:
		win = "rock"
	case 1:
		win = "rock"
	case 2:
		win = "paper"
	case 3:
		win = "scissors"
	}
	return win

}

func WhoWon(player, bot string) string{
	var win string
	if player == "rock" && bot == "scissors"{
		win = "player"
	} else if  player == "scissors" && bot == "rock" {
		win = "bot"
	}


	if player == "paper" && bot == "rock"{
		win = "player"
	} else if player == "rock" && bot == "paper"{
		win = "bot"
	}


	if player == "scissors" && bot == "paper"{
		win = "player"
	} else if player == "paper" && bot =="scissors"{
		win = "bot"
	}



	if player == bot{
		win = "Nobody"
	}


	return win
}

func main(){
	matches := flag.Int("matches", 5, "Number of matches!")
	flag.Parse()
	//seed random with a constantly-changing number
	rand.Seed(time.Now().UnixNano())

	var name string
	var userChoice string
	var plrPoints int
	var botPoints int
	var game int = 0
	//var checkBot string
	//var checkUser string
	Clear()
	Bar(game, *matches)

	fmt.Println("Please Enter your first name:\n")
	fmt.Scanln(&name)
	Sleep(2)
	Clear()
	for game < *matches {
		var ranBotChoice int = rand.Intn(4)
		//game
		Bar(game, *matches)
		//fmt.Printf("%d", ranBotChoice) //debug
		fmt.Printf("\t %s:%d \t  ||||| \t     Bot:%d\n\n", name, plrPoints, botPoints  )
		fmt.Println("Choice (r) (p) (s)?")
		fmt.Scanln(&userChoice)
		checkUser := playerPlay(strings.ToLower(userChoice))
		fmt.Printf("%s plays: %s --> \n", name, checkUser)
		Sleep(2)
		fmt.Printf("\t\t\t\t<-- Bot plays: %s\n", botPlay(ranBotChoice))
		Sleep(2)
		checkBot := botPlay(ranBotChoice)

		//fmt.Printf("Player: %s \t || \t Bot: %s\n", checkUser, checkBot)

		if WhoWon(checkUser, checkBot) == "player"{
			fmt.Printf("\n\t\t\t%s won!\n", name)
			plrPoints+=1
		} else if WhoWon(checkUser, checkBot) == "bot"{
			fmt.Printf("\n\t\t\tBot Won!\n")
			botPoints+=1
		} else if WhoWon(checkUser, checkBot) == "Nobody"{
			fmt.Printf("\n\t\t\tNobody Won!\n")
		}
		game+=1
		Sleep(2)
		Clear()
	}

	Clear()
	Sleep(2)
	Bar(game, *matches)
	fmt.Printf("\nFINAL SCORES:\n\t\t\t%s:%d\n\t\t\t\t\t\tBot:%d\n\n Thanks For Playing\n", name, plrPoints, botPoints)
	Sleep(2)
}
