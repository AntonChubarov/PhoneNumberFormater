package infrastructure

import "fmt"

type ConsoleVisualizer struct {

}

func (c *ConsoleVisualizer) Visualize(number []string) {
	fmt.Println()
	for i := range number {
		fmt.Println(number[i])
	}
	fmt.Println()
}

func NewConsoleVisualizer() *ConsoleVisualizer {
	return &ConsoleVisualizer{}
}