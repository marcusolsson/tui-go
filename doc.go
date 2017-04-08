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

Size policies

Sizing of widgets is controlled by its SizePolicy. For now, you can read more
about how size policies work in the Qt docs:

http://doc.qt.io/qt-5/qsizepolicy.html#Policy-enum
*/
package tui
