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

text: qsTr("<h1>Quirinux GNU/Linux Versión 2.0</h1>"+

"<h2>Basado en Devuan (Debian) GNU/Linux</h2>"+

"<h3>https://www.quirinux.org/</h3><br/>"+

"<b>Versión:</b> 2.0 Estable 07-08-2024<br/><br/>"+

"<b>Autor:</b> Charlie Martínez.<br/><br/>"+

"<b>Logotipo Quirinux:</b> Thomas Gaya.<br/>"+
"<b>Sistema de Creación ISO:</b> Penguins' Eggs, de Piero Proietti.<br/>"+
"<b>Capacitación y Auditoria Técnica Interna:</b> Javier Obregón.<br/><br/>"+

"<b>Agradecimiento a usuarios, colaboradores y beta-testers:</b><br/>"+
"Alicia Coria, Axel Oil, Facundo Redoni (AnimaTics), Gustavo Deveze, Gustav González (Tupi),<br/>"+
"Jeremy Bullock, Leonardo Bensingor, Luis Federichi (qepd), Manong John (Tahoma2D),<br/>"+
"Manuel González, Martín Figuera, Noelia Gerbaudo, Pablo López (Etertics), Patricia Mengui,<br/>"+
"Riky Linux, Sela González.<br/><br/>"+

"<b>Hecho en:</b> Buenos Aires (Argentina), Misiones (Argentina), Roma (Italia) y Stgo. de Compostela (Galicia, España).<br/>"+
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

