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
            id: step1
            source: "slideshow/step1.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 800
            height: 520
            fillMode: Image.PreserveAspectFit
        }
    }
    Slide {
        Image {
            id: step2
            source: "slideshow/step2.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 800
            height: 520
            fillMode: Image.PreserveAspectFit
        }
    }

    Slide {
        Image {
            id: step3
            source: "slideshow/step3.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 800
            height: 520
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
