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
            color: "#6d526b"
            anchors.horizontalCenter: slide1.horizontalCenter

text: qsTr("<h1>Quirinux GNU/Linux Versión 2.0</h1>" + 

"<h2>Basado en Devuan (Debian) GNU/Linux</h2>" + 

"<h3>https://www.quirinux.org/</h3><br/>" +

"<b>Versión:</b> 2.0 ESTABLE 09-11-2024 | Rev. 8 [02-02-2025]<br/><br/>" +

"<b>Autor:</b> Charlie Martínez.<br/><br/>" +

"<b>Logotipo Quirinux:</b> Thomas Gaya.<br/>"+
"<b>Sistema de Creación ISO:</b> Penguins' Eggs, de Piero Proietti.<br/>"+
"<b>Capacitación y Auditoria Técnica Interna:</b> Javier Obregón.<br/>"+
"<b>Pruebas de compatibilidad en Mac:</b> Sela González. <br/><br/>"+
"Agradecimiento a toda la comunidad que colaboró<br/>"+ 
"como beta-tester y en la difusión del proyecto.<br/><br/>" +

"<b>Hecho en:</b> Stgo. de Compostela (Galicia, España), Misiones (Argentina) y Roma (Italia).<br/>"+
"<b>Quirinux</b> es Marca Registrada.<b> Distribuidor oficial:</b> CREALIB TECNOLOGÍA SOSTENIBLE.<br/>"+

"<b>Dedicado a Emilio Gorini (qepd).</b>")
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

