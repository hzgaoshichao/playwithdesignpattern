package main

import t "translatordemo/basketballtranslator"

func main() {

	var player t.PlayerInterface

	player = t.NewForward("巴蒂尔")
	player.Attack()
	player.Defense()

	player = t.NewGuards("麦克格雷迪")
	player.Attack()
	player.Defense()

	player = t.NewTranslator("姚明")
	player.Attack()
	player.Defense()

}
