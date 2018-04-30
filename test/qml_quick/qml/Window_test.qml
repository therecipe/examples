import QtQuick 2.7
import GoTest 1.0	//TestRunner

Window {
	TestRunner {
		objectName: "localTestRunner"

		onCallQml: {
			//console.log(obj, sig)
			obj[sig]()
		}

		onCallQmlWithArgs: {
			//console.log(obj, sig, args)
			obj[sig].apply(null, args)
		}
	}
}