import QtQuick 2.0;
import calamares.slideshow 1.0;

Presentation
{
    id: presentation

    function nextSlide() {
        console.log("QML Component (default slideshow) Next slide");
        presentation.goToNextSlide();
    }



    Slide {
        Image {
            id: reproductiveSystem
            source: "1-reproductive-system.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
          Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#aa3333"
            anchors.horizontalCenter: reproductiveSystem.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Use arrow key to navigation</h1><br/>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
        }
    }
    Slide {
        Image {
            id: startReproduction
            source: "2-start-reproduction.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
    }

    Slide {
        Image {
            id: itsYourSystem
            source: "3-its-your-system.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
    }

    Slide {
        Image {
            id: eggsPresentation
            source: "4-eggs-presentation.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
    }

    Slide {
        Image {
            id: waitHatching
            source: "5-wait-hatching.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
    }

    Slide {
        Image {
            id: followPenguins
            source: "6-follow-penguins.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }

    }

    Slide {
        Image {
            id: createdBy
            source: "7-created-by.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }

    }
    
    
        Slide {
        Image {
            id: createdBy8
            source: "8-created-by.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }

    }
    
    
        Slide {
        Image {
            id: createdBy9
            source: "9-created-by.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }

    }
    
    
            Slide {
        Image {
            id: createdBy10
            source: "10-created-by.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         
    }
    
    
      Slide {
        Image {
            id: createdBy11
            source: "11-created-by.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
        
    }
 
 
    
      Slide {
        Image {
            id: createdBy12
            source: "12-created-by.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
        
    }
    
          Slide {
        Image {
            id: createdBy13
            source: "13-created-by.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         
    }
    

    // When this slideshow is loaded as a V1 slideshow, only
    // activatedInCalamares is set, which starts the timer (see above).
    //
    // In V2, also the onActivate() and onLeave() methods are called.
    // These example functions log a message (and re-start the slides
    // from the first).
    function onActivate() {
        console.log("QML Component (default slideshow) activated");
        presentation.currentSlide = 0;
    }

    function onLeave() {
        console.log("QML Component (default slideshow) deactivated");
    }
}

