package infrastructure

import "fmt"

type ConsoleVisualizer struct {

}

func (c *ConsoleVisualizer) Visualize(number []string) {
	for i := range number {
		fmt.Println(number[i])
	}
	fmt.Println()
}

func NewVonsoleVisualizer() *ConsoleVisualizer {
	return &ConsoleVisualizer{}
}