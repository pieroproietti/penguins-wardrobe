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
            source: "escritorio.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#d9fc1e"
            anchors.horizontalCenter: reproductiveSystem.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>VENDEFOUL WOLF EXCALIBUR</h1><br/>"+
                  "<h2>Es GNU/Linux basado en Devuan 6</h2>"+
                  "<h3>Y con tu entorno gráfico favorito</h3>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
        }
    }
    Slide {
        Image {
            id: startReproduction
            source: "actualizador.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#d9fc1e"
            anchors.horizontalCenter: reproductiveSystem.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>ACTUALIZADOR AUTOMATICO</h1><br/>"+
                  "<h2>Manten tus paquete al día, sin esfuerzo</h2>"+
                  "<h3>Tu sistema 100% seguro</h3>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
        }
    }
    Slide {
        Image {
            id: followPenguins
            source: "htop.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#d9fc1e"
            anchors.horizontalCenter: reproductiveSystem.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>VALIDO PARA TODO TIPO DE ORDENADORES</h1><br/>"+
                  "<h2>Optimizando el uso de RAM, un sistema muy fluido.</h2>"+
                  "<h3>Todo listo para usar y producir.</h3>")
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

