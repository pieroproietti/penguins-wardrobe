/* === This file is part of Calamares - <http://github.com/calamares> ===
 *
 *   Copyright 2015, Teo Mrnjavac <teo@kde.org>
 *   Copyright 2018, Jonathan Carter <jcc@debian.org>
 *
 *   Calamares is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, or (at your option) any later version.
 *
 *   Calamares is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with Calamares. If not, see <http://www.gnu.org/licenses/>.
 */


// https://github.com/calamares/calamares-extensions


import QtQuick 2.0;
import calamares.slideshow 1.0;

Presentation
{
    id: presentation

    function nextSlide() {
        console.log("QML Component (default slideshow) Next slide");
        presentation.goToNextSlide();
    }

    Timer {
        id: advanceTimer
        interval: 7500
        running: true
        repeat: true
        onTriggered: nextSlide()
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
            color: "#ffffff"
            anchors.horizontalCenter: reproductiveSystem.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1></h1><br/>"+
                  "<h2>SysLinuxOS for System Integrators</h2>"+
                  "<h3>https://syslinuxos.com</h3>")
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
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            // color: "#002439"
            color: "#ffffff"
            anchors.horizontalCenter: startReproduction.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1></h1><br/>"+
                  "<h2></h2>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
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
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#ffffff"
            anchors.horizontalCenter: itsYourSystem.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1></h1><br/>"+
                  "<h2><br/></h2>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: syslinuxosPresentation
            source: "4-syslinuxos-presentation.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 20
            color: "#ffffff"
            anchors.horizontalCenter: syslinuxosPresentation.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1></h1><br/>"+
                  "<h3>SysLinuxOS (for System Integrators) is a Swiss army knife to always carry with us</h3>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
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
         Text {
            font.family: "Helvetica"
            font.pixelSize : 20
            color: "#ffffff"
            anchors.horizontalCenter: waitHatching.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1></h1><br/>"+
                  "<h3>SysLinuxOS was built to work right out of the box.</h3>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
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
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#ffffff"
            anchors.horizontalCenter: slide6.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1></h1><br/>"+
                  "<h3>There is no need to install anything.</h3>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
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
         Text {
            font.family: "Helvetica"
            font.pixelSize : 20
            color: "#ffffff"
            anchors.horizontalCenter: followPenguins.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1></h1><br/>"+
                  "<h3>Created by Franco Conidi</h3><br/>"+
                  "<h4>https://github.com/fconidi</h4>"+
                  "<h4>https://syslinuxos.com</h4>"+
                  "<h4>https://francoconidi.it</h4>"+
                  "<h4>fconidi@gmail.com</h4>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
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

