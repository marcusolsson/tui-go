package main

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

func main() {
	buffer := tui.NewTextEdit()
	buffer.SetSizePolicy(tui.Expanding, tui.Expanding)
	buffer.SetText(body)
	buffer.SetFocused(true)
	buffer.SetWordWrap(true)

	status := tui.NewStatusBar("lorem.txt")

	root := tui.NewVBox(buffer, status)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

const body = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque elementum dui vitae nisl scelerisque porta. Proin orci mauris, imperdiet ac venenatis id, interdum fermentum ligula. Praesent risus odio, pharetra ac maximus in, tincidunt eget nibh. Vestibulum pretium molestie fermentum. Aenean sed neque purus. Vivamus vitae nulla nec ligula ultrices lacinia. Ut in vulputate ante. Proin lacinia eleifend varius. Cras quis urna eget nisi efficitur tristique sed vitae nibh. Nam ac nisi libero. In interdum volutpat elementum. Nulla lorem magna, efficitur interdum ante at, convallis sodales nulla. Sed maximus tempor condimentum.

Nam et risus est. Cras ornare iaculis orci, sit amet fringilla nisl pharetra quis. Integer quis sem porttitor, gravida nisi eget, feugiat lacus. Aliquam aliquet quam eget ipsum ultrices, in viverra ex dapibus. Sed ullamcorper, justo sit amet feugiat faucibus, nisl sem hendrerit dui, non tincidunt lacus turpis non tortor. Aenean nisl justo, dictum non eros quis, luctus pulvinar urna. Ut finibus odio id nunc rutrum iaculis.

Nam commodo tempor augue, nec facilisis nulla pretium scelerisque. Donec eu interdum nisl. Aliquam dui nisl, venenatis id velit ac, ultrices faucibus massa. Suspendisse id condimentum augue. Sed a libero ornare, sollicitudin neque sed, blandit nunc. Quisque sed sem non erat pharetra semper. Mauris molestie leo ante, in varius elit ullamcorper at. Suspendisse sed scelerisque velit, eget rutrum nunc. Cras in ultrices risus. Aliquam maximus, purus in consequat rutrum, erat mauris pharetra lacus, nec interdum turpis metus non velit. Cras in lobortis tortor, vitae dignissim mauris. Phasellus nec massa nisi. Etiam auctor, odio egestas egestas ullamcorper, mauris risus maximus nisl, eget faucibus risus dui sed nisi. Sed ligula mi, egestas in augue vitae, tincidunt molestie sapien.

Maecenas eget tristique dolor. Quisque vel velit ante. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Pellentesque lorem diam, feugiat ut odio et, tempus consequat mauris. Phasellus lobortis sodales tellus, sed aliquam lectus lobortis id. Nulla mollis tempor elit. Etiam luctus convallis justo, sed viverra nibh sodales eget. Aliquam blandit, felis eget accumsan tempus, orci magna molestie metus, vel bibendum nibh risus at augue. Maecenas vulputate feugiat dui sit amet facilisis. In et eros vel elit vestibulum laoreet at id mauris. Nullam tincidunt suscipit diam, vel sollicitudin massa venenatis id. Fusce porttitor urna et aliquam dignissim. Maecenas mollis ligula ut ex maximus, vel feugiat metus scelerisque. Sed pharetra ac nunc in pharetra.`
