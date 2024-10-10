package main

import c "compositedemo/compositepattern"

func main() {
	root := c.NewComposite("root")
	root.Add(c.NewLeaf("LeafA"))
	root.Add(c.NewLeaf("LeafB"))

	comp := c.NewComposite("Composite X")
	comp.Add(c.NewLeaf("LeafXA"))
	comp.Add(c.NewLeaf("LeafXB"))

	root.Add(comp)

	comp2 := c.NewComposite("Composite XY")
	comp2.Add(c.NewLeaf("LeafXYA"))
	comp2.Add(c.NewLeaf("LeafXYB"))

	comp.Add(comp2)

	leaf := c.NewLeaf("Leaf C")
	root.Add(leaf)

	leaf2 := c.NewLeaf("Leaf D")
	root.Add(leaf2)
	root.Remove(leaf2)

	root.Display(10)
}
