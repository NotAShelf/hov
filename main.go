package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	schema "notashelf.dev/hov/data/schema"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// get the Hyprland instance signature from the environment
	signature := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")

	// if signature is empty, we are not in Hyprland
	if signature == "" {
		fmt.Println("Hov only works in Hyprland")
	} else {
		fmt.Println("Hov is now connected to Hyprland")
		fmt.Println("Hyprland instance Signature: " + signature + "\n")
	}

	get_workspaces()

	// build_window() // will not work until fyne merges wayland support

}

func get_workspaces() {

	// receive json output from hyprctl
	cmd := exec.Command("hyprctl", "clients", "-j")
	out, err := cmd.Output()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if len(os.Args) > 1 && os.Args[1] == "--debug" {
		fmt.Println(string(out))
	}

	// unmarshal json output into a slice of structs
	var client schema.Client
	json.Unmarshal(out, &client)

	if err := json.Unmarshal(out, &client); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Print(err)
	}

	println(client)

	// loop through the slice of structs
	for i := 0; i < len(client); i++ {

		clientslice := make([]string, 3)

		clientslice = append(clientslice,
			client[i].Class+
				client[i].Title+
				client[i].Address+
				client[i].Workspaces.Name+
				strconv.Itoa(client[i].Workspaces.Id),
		)

		fmt.Println(clientslice)
	}
}

func build_window() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
