package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (s Animal) Eat() {
	fmt.Println(s.food)
}

func (s Animal) Move() {
	fmt.Println(s.locomotion)
}

func (s Animal) Speak() {
	fmt.Println(s.noise)
}

func main() {

	var cow, bird, snake, temp Animal
	var s1, s2 string
	var flag int
	cow.food = "grass"
	cow.locomotion = "walk"
	cow.noise = "moo"
	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noise = "peep"
	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noise = "hss"
	fmt.Println("Animal names are cow , bird , snake")
	fmt.Println("Activities available are eat, move, speak")
	for {
		flag = 0

		fmt.Print("Please enter the animal name and activity\n")
		fmt.Scanf("%s%s", &s1, &s2)
		switch s1 {
		case "cow":
			temp = cow
		case "bird":
			temp = bird
		case "snake":
			temp = snake
		default:
			fmt.Println("Enter the proper animal name")
			flag = 1
		}
		if flag != 1 {
			switch s2 {
			case "eat":
				temp.Eat()
			case "move":
				temp.Move()
			case "speak":
				temp.Speak()
			default:
				fmt.Println("Please enter the proper action")
			}

		}
	}
}
