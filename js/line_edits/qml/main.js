//port of: https://github.com/therecipe/qt/blob/master/internal/examples/widgets/line_edits/line_edits.go

//import github.com/therecipe/qt/core
//import github.com/therecipe/qt/widgets

"use strict";

//widgets.NewQApplication(); //TODO

var echoGroup = widgets.NewQGroupBox2("Echo"),
	echoLabel = widgets.NewQLabel2("Mode:"),
	echoComboBox = widgets.NewQComboBox(),
	echoLineEdit = widgets.NewQLineEdit();

echoComboBox.AddItems(["Normal", "Password", "PasswordEchoOnEdit", "No Echo"]);
echoLineEdit.SetPlaceholderText("Placeholder Text");

var validatorGroup = widgets.NewQGroupBox2("Validator"),
	validatorLabel = widgets.NewQLabel2("Type:"),
	validatorComboBox = widgets.NewQComboBox(),
	validatorLineEdit = widgets.NewQLineEdit();

validatorComboBox.AddItems(["No validator", "Integer validator", "Double validator"]);
validatorLineEdit.SetPlaceholderText("Placeholder Text");

var alignmentGroup = widgets.NewQGroupBox2("Alignment"),
	alignmentLabel = widgets.NewQLabel2("Type:"),
	alignmentComboBox = widgets.NewQComboBox(),
	alignmentLineEdit = widgets.NewQLineEdit();

alignmentComboBox.AddItems(["Left", "Centered", "Right"]);
alignmentLineEdit.SetPlaceholderText("Placeholder Text");

var inputMaskGroup = widgets.NewQGroupBox2("Input mask"),
	inputMaskLabel = widgets.NewQLabel2("Type:"),
	inputMaskComboBox = widgets.NewQComboBox(),
	inputMaskLineEdit = widgets.NewQLineEdit();

inputMaskComboBox.AddItems(["No mask", "Phone number", "ISO date", "License key"]);
inputMaskLineEdit.SetPlaceholderText("Placeholder Text");

var accessGroup = widgets.NewQGroupBox2("Access"),
	accessLabel = widgets.NewQLabel2("Read-only:"),
	accessComboBox = widgets.NewQComboBox(),
	accessLineEdit = widgets.NewQLineEdit();

accessComboBox.AddItems(["False", "True"]);
accessLineEdit.SetPlaceholderText("Placeholder Text");

echoComboBox.ConnectCurrentIndexChanged(function(index) { echoChanged(echoLineEdit, index); });
validatorComboBox.ConnectCurrentIndexChanged(function(index) { validatorChanged(validatorLineEdit, index); });
alignmentComboBox.ConnectCurrentIndexChanged(function(index) { alignmentChanged(alignmentLineEdit, index); });
inputMaskComboBox.ConnectCurrentIndexChanged(function(index) { inputMaskChanged(inputMaskLineEdit, index); });
accessComboBox.ConnectCurrentIndexChanged(function(index) { accessChanged(accessLineEdit, index); });

var echoLayout = widgets.NewQGridLayout2();
echoLayout.AddWidget(echoLabel, 0, 0, 0);
echoLayout.AddWidget(echoComboBox, 0, 1, 0);
echoLayout.AddWidget3(echoLineEdit, 1, 0, 1, 2, 0);
echoGroup.SetLayout(echoLayout);

var validatorLayout = widgets.NewQGridLayout2();
validatorLayout.AddWidget(validatorLabel, 0, 0, 0);
validatorLayout.AddWidget(validatorComboBox, 0, 1, 0);
validatorLayout.AddWidget3(validatorLineEdit, 1, 0, 1, 2, 0);
validatorGroup.SetLayout(validatorLayout);

var alignmentLayout = widgets.NewQGridLayout2();
alignmentLayout.AddWidget(alignmentLabel, 0, 0, 0);
alignmentLayout.AddWidget(alignmentComboBox, 0, 1, 0);
alignmentLayout.AddWidget3(alignmentLineEdit, 1, 0, 1, 2, 0);
alignmentGroup.SetLayout(alignmentLayout);

var inputMaskLayout = widgets.NewQGridLayout2();
inputMaskLayout.AddWidget(inputMaskLabel, 0, 0, 0);
inputMaskLayout.AddWidget(inputMaskComboBox, 0, 1, 0);
inputMaskLayout.AddWidget3(inputMaskLineEdit, 1, 0, 1, 2, 0);
inputMaskGroup.SetLayout(inputMaskLayout);

var accessLayout = widgets.NewQGridLayout2();
accessLayout.AddWidget(accessLabel, 0, 0, 0);
accessLayout.AddWidget(accessComboBox, 0, 1, 0);
accessLayout.AddWidget3(accessLineEdit, 1, 0, 1, 2, 0);
accessGroup.SetLayout(accessLayout);

var layout = widgets.NewQGridLayout2();
layout.AddWidget(echoGroup, 0, 0, 0);
layout.AddWidget(validatorGroup, 1, 0, 0);
layout.AddWidget(alignmentGroup, 2, 0, 0);
layout.AddWidget(inputMaskGroup, 0, 1, 0);
layout.AddWidget(accessGroup, 1, 1, 0);

var window = widgets.NewQMainWindow();
window.SetWindowTitle("Line Edits");

var centralWidget = widgets.NewQWidget(window);
centralWidget.SetLayout(layout);
window.SetCentralWidget(centralWidget);

window.Show();

var echoChanged = function(echoLineEdit, index) {
	switch (index) {
		case 0: {
			echoLineEdit.SetEchoMode(widgets.QLineEdit__Normal);
			break;
		}

		case 1: {
			echoLineEdit.SetEchoMode(widgets.QLineEdit__Password);
			break;
		}

		case 2: {
			echoLineEdit.SetEchoMode(widgets.QLineEdit__PasswordEchoOnEdit);
			break;
		}

		case 3: {
			echoLineEdit.SetEchoMode(widgets.QLineEdit__NoEcho);
			break;
		}
	}
};

var validatorChanged = function(validatorLineEdit, index) {
	switch (index) {
		case 0: {
			validatorLineEdit.SetValidator();
			break;
		}

		case 1: {
			validatorLineEdit.SetValidator(
				widgets.NewQIntValidator(validatorLineEdit)
			);
			break;
		}

		case 2: {
			validatorLineEdit.SetValidator(
				widgets.NewQDoubleValidator2(-999.0, 999.0, 2, validatorLineEdit)
			);
			break;
		}
	}

	validatorLineEdit.Clear();
};

var alignmentChanged = function(alignmentLineEdit, index) {
	switch (index) {
		case 0: {
			alignmentLineEdit.SetAlignment(core.Qt__AlignLeft);
			break;
		}

		case 1: {
			alignmentLineEdit.SetAlignment(core.Qt__AlignCenter);
			break;
		}

		case 2: {
			alignmentLineEdit.SetAlignment(core.Qt__AlignRight);
			break;
		}
	}
};

var inputMaskChanged = function(inputMaskLineEdit, index) {
	switch (index) {
		case 0: {
			inputMaskLineEdit.SetInputMask("");
			break;
		}

		case 1: {
			inputMaskLineEdit.SetInputMask("+99 99 99 99 99;_");
			break;
		}

		case 2: {
			inputMaskLineEdit.SetInputMask("0000-00-00");
			inputMaskLineEdit.SetText("00000000");
			inputMaskLineEdit.SetCursorPosition(0);
			break;
		}

		case 3: {
			inputMaskLineEdit.SetInputMask(">AAAAA-AAAAA-AAAAA-AAAAA-AAAAA;#");
			break;
		}
	}
};

var accessChanged = function(accessLineEdit, index) {
	switch (index) {
		case 0: {
			accessLineEdit.SetReadOnly(false);
			break;
		}

		case 1: {
			accessLineEdit.SetReadOnly(true);
			break;
		}
	}
};

//widgets.QApplication_Exec() //TODO
