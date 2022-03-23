package poker

type CLI struct {
	PlayerStore PlayerStore
}

func (cli *CLI) PlayPoker() {
	cli.PlayerStore.RecordWin("Cleo")
}
