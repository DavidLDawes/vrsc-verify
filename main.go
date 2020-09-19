package main

import (
	"log"

	pb "githhub.com/davidldawes/vrsc-verify/walletrpc"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func checkSignature() {
	verifyRequest := VerifyHashRequest{}
}

var check = widget.NewButton("Check signature", checkSignature)

// settings
var identity = widget.NewEntry()
var plainText = widget.NewTextGrid()
var signature = widget.NewEntry()
var result = widget.NewEntry()

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	a := app.New()
	w := a.NewWindow("Designer")

	// actions

	// settings
	identity.SetText("DavidLDawes")
	plainText.SetText("Sample text as an example")
	signature.SetText("opaque-string-goes-here-once-I-get-it")

	actions := widget.NewForm(
		widget.NewFormItem("Check the signature of the clear text using the identity", check),
		widget.NewFormItem("Clear text to check against the signature", plainText),
		widget.NewFormItem("Signature result to check", signature),
	)

	settings := widget.NewForm(
		widget.NewFormItem("Identity that signed the clear text", identity),
		widget.NewFormItem("Clear text to check against the signature", plainText),
		widget.NewFormItem("Signature result to check", signature),
	)

	ui := widget.NewHBox(widget.NewLabel("Designer"),
		actions,
		settings,
	)

	w.SetContent(ui)

	w.ShowAndRun()
}
