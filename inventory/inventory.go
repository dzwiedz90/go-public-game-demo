package inventory

import rl "github.com/gen2brain/raylib-go/raylib"

func Inventory(f func(string), a []int, b []string) (rl.Rectangle, []int, []string) {
	c := a
	d := b
	a = []int{}
	b = []string{}
	f("inventory.map")
	return rl.NewRectangle(150, 150, 48, 48), c, d
}
