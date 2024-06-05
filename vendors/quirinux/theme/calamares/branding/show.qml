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
            id: slide1
            source: "slide1.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Ubuntu"
            font.pixelSize : 12
            color: "#aa3333"
            anchors.horizontalCenter: slide1.horizontalCenter
text: qsTr("<h1>Quirinux GNU/Linux Versión 2.0</h1>"+
"<h2>Basado en Devuan (Debian) GNU/Linux</h2>"+
"<h3>https://www.quirinux.org/</h3><br/>"+
"Autor: Charlie Martínez. Logotipo Quirinux: Thomas Gaya.<br/><br/>"+
"Agradecimientos a: Alicia Coria, Gustav González (Tupi), Gustavo Deveze, Jeremy Bullock, <br/>"+
"Luis Federichi, Manong John (Tahoma2D), Pablo López (Etertics), Patricia Mengui, Riky Linux<br/><br/>"+
"Beta-testers: Facundo Redoni (AnimaTics),  Martín Figuera, Gustavo Deveze,<br/>Noelia Gerbaudo, Leonardo Bensingor, y Sela González<br/><br/>"+
"Capacitación interna y asesoramiento técnico al desarrollo: Javier Obregón<br/><br/>"+
"Sistema de creación ISO usado en pruebas de concepto: Penguins' Eggs, de Piero Proietti.<br/>"+
"Wallpapers: Beatriz Barcelona, Charlie Martínez, Ernesto Bazzano, Geri Ratcliffe, Leonardo Bensingor y Noelia Gerbaudo<br/><br/>"+
"Hecho en Buenos Aires (Argentina), Roma (Italia) y Santiago de Compostela (Galicia, España).<br/>"+
"Dedicado a Emilio Gorini (qepd).")
            wrapMode: Text.WordWrap
            width: kde.width
            horizontalAlignment: Text.Center
        }
    }

    function onActivate() {
        console.log("QML Component (default slideshow) activated");
        presentation.currentSlide = 0;
    }

    function onLeave() {
        console.log("QML Component (default slideshow) deactivated");
    }
}

