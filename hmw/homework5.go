// package main

// import (
//   "fmt"
//   "math/rand"
// )

// type Player struct{
// 	Name string
// 	Chance int64
// }

// type Game struct{
// 	Number int
// 	Player Player 
// }

// func (g Game) NewGame(player Player) Game {
// 	return Game{
// 		Number: rand.Intn(10),
// 		Player: player,
// 	}
// }

// func (p Player) NewPlayer(name string,chance int64) Player{
// 	return Player {
// 		Name: name,
// 		Chance: chance,
// 	}
// }

// func main()  {
// 	getPlayerName()
// 	player := Player{}
// 	player = player.NewPlayer(getPlayerName(),3)
// 	game := Game{}
// 	game = game.NewGame(player)
// 	game.StartGame() 
// }

// func getPlayerName() string {
// 	var playerName string 
// 	fmt.Println("Game is strting!")
// 	fmt.Printf("Please enter your name:")
// 	_,err := fmt.Scanln(&playerName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return playerName
// }

// func (g Game) StartGame(){
// 	var numb int 
// 	chance := g.Player.Chance
// 	for i := 0; i <= int(chance); i++{
// 		fmt.Printf("Your %v try.Please enter the number:",i)
// 		fmt.Scanln(&numb)
//     if g.Number == numb {
// 		fmt.Println("You are win")
// 		break
// 	}else if i == int(chance){
// 		fmt.Printf("You are a loser xD,Predicter number is %v\n",g.Number)
// 	}else{
// 		g.Player.Chance--
// 		fmt.Printf("You didnt find. You have %v chance. Try again:)\n",
// 	g.Player.Chance)
// 	}
// }
// }