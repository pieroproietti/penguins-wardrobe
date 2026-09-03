import QtQuick 2.0;
import calamares.slideshow 1.0;

Presentation
{
    id: presentation

    Timer {
        id: advanceTimer
        interval: 8000
        running: true
        repeat: true
        onTriggered: presentation.goToNextSlide()
    }
    
    Slide {
        Image {
            id: image1
            source: "Base.png"
	    width: 1080
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
	    x: 0
	    y: 120
        }
    }

    Slide {

        Image {
            id: image2
            source: "Wizard.png"
	    width: 1080
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
	    x: 0
	    y: 120
        }
    }



    Slide {

        Image {
            id: image3
            source: "XFCE.png"
	    width: 1080
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
	    x: 0
	    y: 120
        }
    }
	
	
    Slide {

        Image {
            id: image4
            source: "rhinopkg.png"
	    width: 1080
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
	    x: 0
	    y: 120
        }
    }



    Component.onCompleted: advanceTimer.running = true
}
