/*
    Copyright 2019 Harald Sitter <sitter@kde.org>

    This program is free software; you can redistribute it and/or
    modify it under the terms of the GNU General Public License as
    published by the Free Software Foundation; either version 3 of
    the License or any later version accepted by the membership of
    KDE e.V. (or its successor approved by the membership of KDE
    e.V.), which shall act as a proxy defined in Section 14 of
    version 3 of the license.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

    PLEASE NOTE:
    This is not the original neon theme, but an adpment made for eggs 
    from eggs author: Piero Proietti <piero.proietti@gmail.com> 
*/


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
        id: timer
        interval: 5000
        running: false
        repeat: true
        onTriggered: nextSlide()
    }

    Slide {
        Image {
            id: slide1
            source: "neon.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 20
            color: "#fcfcfc"
            anchors.horizontalCenter: slide1.horizontalCenter
            anchors.top: background.top
                  text: qsTr("<h1>LinuxFX</h1><br/>" +
                  "<h2>Fast, stable and very safe.<h2>" +
                  "<h3>https://linuxfx.org/</h3>")
            wrapMode: Text.WordWrap
            width: kde.width
            horizontalAlignment: Text.Center
        }
    }
    Slide {
        Image {
            id: slide2
            source: "kde.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 20
            color: "#fcfcfc"
            anchors.horizontalCenter: slide2.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>LinuxFX</h1><br/>" +
                  "<h2>The Microsoft Windows interface with the speed<br/>and security of Linux</h2>")
            wrapMode: Text.WordWrap
            width: neon.width
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: slide3
            source: "plasma.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 20
            color: "#fcfcfc"
            anchors.horizontalCenter: slide3.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>LinuxFX</h1><br/>" +
                  "<h2>Please wait, we are hatching...<br/>Don't disturb the process,<br/>in few time You will have a new penguin.</h2>")
            wrapMode: Text.WordWrap
            width: plasma.width
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: slide4
            source: "secure.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 20
            color: "#fcfcfc"
            anchors.horizontalCenter: slide4.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>LinuxFX</h1><br/>" +
                  "<h2>Created by Rafael Rachid</h2><br/>" +
                  "<h3>email: rafael@linuxfx.org</h3>")
            wrapMode: Text.WordWrap
            width: secure.width
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
