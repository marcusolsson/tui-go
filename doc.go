/*
Package tui is a library for building user interfaces for the terminal.


Widgets

Widgets are the main building blocks of any user interface. They allow us to
present information and interact with our application. It receives keyboard and
mouse events from the terminal and draws a representation of itself.

	lbl := tui.NewLabel("Hello, World!")


Layouts

Widgets are structured using layouts. Layouts are powerful tools that let you
position your widgets without having to specify their exact coordinates.

	box := tui.NewVBox(
		tui.NewLabel("Press the button to continue ..."),
		tui.NewButton("Continue"),
	)

Here, the VBox will ensure that the Button will be placed underneath the Label.
There are currently three layouts to choose from; VBox, HBox and Grid.

Sizing

The Size method returns the space each widget occupies and is recalculated on
calling Resize. SizeHint is the minimum, or even preferred, size of the widget.

Typically, you will not have to call Resize, Size or SizeHint yourself.
Instead, SizePolicy is used to determine how a widget expands when there are
more space than it needs. Currently there are two size policies available:
Minimum and Expanding.

Minimum tells the widget not to use more space than the SizeHint, while Maximum
lets the widget expand to fill the available space.

	box := tui.NewVBox()
	box.SetSizePolicy(tui.Minimum, tui.Expanding)

Here, the layout will shrink to its minimal size along the horizontal axis but
expand along the vertical axis. An example of this would be a sidebar.
*/
package tui
