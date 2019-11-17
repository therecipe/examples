//port of: https://github.com/therecipe/examples/blob/master/basic/widgets/main.go

//import github.com/therecipe/qt/widgets

(function(){

// create a regular widget
// give it a QVBoxLayout
// and make it the central widget of the window
var widget = widgets.NewQWidget();
widget.SetLayout(widgets.NewQVBoxLayout());


// create a line edit
// with a custom placeholder text
// and add it to the central widgets layout
var input = widgets.NewQLineEdit();
input.SetPlaceholderText("Write something ...");
widget.Layout().AddWidget(input);

// create a button
// connect the clicked signal
// and add it to the central widgets layout
var button = widgets.NewQPushButton2("and click me!");
button.ConnectClicked(function(bool) {
	widgets.QMessageBox_Information(undefined, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Cancel);
})
widget.Layout().AddWidget(button);

//postamble

widget.Show();
})();
