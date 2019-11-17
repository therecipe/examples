//port of: https://github.com/therecipe/qt/blob/master/internal/examples/widgets/graphicsscene/graphicsscene.go

//import github.com/therecipe/qt/core
//import github.com/therecipe/qt/gui
//import github.com/therecipe/qt/widgets

(function(){

//preamble

var view = widgets.NewQGraphicsView();
var scene = widgets.NewQGraphicsScene3(0, 0, 300, 300);
view.SetScene(scene);

//canvas

//scene.AddRect2(5, 5, scene.Width()-10, scene.Height()-10, gui.NewQPen3(gui.NewQColor6("blue")), gui.NewQBrush());

var font = gui.NewQFont2("Meiryo", 20, 2);
var text = scene.AddText("Hello World", font);
text.SetPos2((scene.Width()-text.BoundingRect().Width())/2, (scene.Height()-text.BoundingRect().Height())/2);

/*
text = scene.AddText(Date.now().toString(), font);
text.SetPos2((scene.Width()-text.BoundingRect().Width())/3, (scene.Height()-text.BoundingRect().Height())/3);
setInterval(function() { text.SetPlainText(Date.now().toString()); }, 15);
*/

var pen = gui.NewQPen3(gui.NewQColor3(255, 0, 0, 255));
scene.AddLine2(20, scene.Height()*0.75, scene.Width()-20, scene.Height()*0.75, pen);

/*
var button = widgets.NewQPushButton2("SceneButton");
var proxy = scene.AddWidget(button, 0);
button.ConnectClicked(function() { proxy.SetX(proxy.X() + 5); });
//setInterval(function() { button.Click(); }, 15);
*/

/*
var pixmap = gui.NewQPixmap3(":/qml/dummy.png")
pixmap = pixmap.Scaled2(pixmap.Width()/2, pixmap.Height()/2, core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation);

var pixmapItem = scene.AddPixmap(pixmap);
pixmapItem.SetTransformationMode(core.Qt__SmoothTransformation)
setInterval(function() {
	if (pixmapItem.Y() >= scene.Height()){
		pixmapItem.SetY(0); 
	}else{
		pixmapItem.SetY(pixmapItem.Y()+1); 
	}
	//pixmapItem.SetRotation(pixmapItem.Rotation()+1);
}, 15);
*/

//postamble

view.Show();
})();
