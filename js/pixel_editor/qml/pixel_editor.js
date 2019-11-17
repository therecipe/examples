//port of: https://github.com/therecipe/qt/blob/master/internal/examples/widgets/pixel_editor/pixel_editor.go

//import github.com/therecipe/qt/core
//import github.com/therecipe/qt/gui
//import github.com/therecipe/qt/widgets

(function(){

//preamble

var view = widgets.NewQGraphicsView();
var scene = widgets.NewQGraphicsScene();
view.SetScene(scene);

view.ConnectResizeEvent(function(event) {
	view.FitInView(scene.ItemsBoundingRect(), core.Qt__KeepAspectRatio);
});

//canvas

var img = gui.NewQImage3(16, 32, gui.QImage__Format_ARGB32);

var x;
for (x = 0; x < img.Width(); x++) {
	var y;
	for (y = 0; y < img.Height(); y++) {
		img.SetPixelColor2(x, y, gui.NewQColor3(x*3, y*6, x*9, 255));
	}
}

var item = widgets.NewQGraphicsPixmapItem2(gui.QPixmap_FromImage(img, 0));

var color = gui.NewQColor3(255, 255, 255, 255); //TODO: the QColor object is garbage collected 
var drawPixel = function(x, y) {
	x = Math.trunc(x);
	y = Math.trunc(y);
	var pixmap = item.Pixmap();
	if (x >= 1 && x < pixmap.Width()-1 &&
		y >= 1 && y < pixmap.Height()-1) {

		var img = item.Pixmap().ToImage();
		img.SetPixelColor2(x, y, color);
		item.SetPixmap(gui.QPixmap_FromImage(img, 0));
	}
}

item.ConnectMouseMoveEvent(function(event) {
	var mousePosition = event.Pos();
	drawPixel(mousePosition.X(), mousePosition.Y());
});
item.ConnectMousePressEvent(function(event) {
	var mousePosition = event.Pos();
	drawPixel(mousePosition.X(), mousePosition.Y());
});
scene.AddItem(item);

//postamble

view.Show();
})();
