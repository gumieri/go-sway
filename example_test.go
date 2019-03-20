package sway_test

import (
	"fmt"
	"log"
	"strings"

	// "go.i3wm.org/i3"
)

func ExampleIsUnsuccessful() {
	cr, err := sway.RunCommand("norp")
	// “norp” is not implemented, so this command is expected to fail.
	if err != nil && !sway.IsUnsuccessful(err) {
		log.Fatal(err)
	}
	log.Printf("error for norp: %v", cr[0].Error)
}

func ExampleSubscribe() {
	recv := sway.Subscribe(sway.WindowEventType)
	for recv.Next() {
		ev := recv.Event().(*sway.WindowEvent)
		log.Printf("change: %s", ev.Change)
	}
	log.Fatal(recv.Close())
}

func ExampleGetTree() {
	// Focus or start Google Chrome on the focused workspace.

	tree, err := sway.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	ws := tree.Root.FindFocused(func(n *sway.Node) bool {
		return n.Type == sway.WorkspaceNode
	})
	if ws == nil {
		log.Fatalf("could not locate workspace")
	}

	chrome := ws.FindChild(func(n *sway.Node) bool {
		return strings.HasSuffix(n.Name, "- Google Chrome")
	})
	if chrome != nil {
		_, err = sway.RunCommand(fmt.Sprintf(`[con_id="%d"] focus`, chrome.ID))
	} else {
		_, err = sway.RunCommand(`exec google-chrome`)
	}
	if err != nil {
		log.Fatal(err)
	}
}
