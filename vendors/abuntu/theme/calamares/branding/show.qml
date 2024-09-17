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
            source: "slide1.png"
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
            source: "slide2.png"
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
            source: "slide3.png"
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
            source: "slide4.png"
	    width: 1080
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
	    x: 0
	    y: 120
        }
    }
    
    Slide {
        Image {
            id: image5
            source: "slide5.png"
	    width: 1080
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
	    x: 0
	    y: 120
        }
    }
    
    Slide {
        Image {
            id: image6
            source: "slide6.png"
	    width: 1080
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
	    x: 0
	    y: 120
        }
    }


    Slide {

        Image {
            id: image7
            source: "slide7.png"
	    width: 1080
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
	    x: 0
	    y: 120
        }
    }



    Component.onCompleted: advanceTimer.running = true
}
