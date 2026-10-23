/* .wardrobe/vendors/quirinux/theme/calamares/branding/ */
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
        // Fondo: la imagen cubre TODA la presentación manteniendo su
        // proporción (recorta lo que sobre). El ancla vertical al 80% conserva
        // la caja del producto, que está en la parte baja de slide1.png (1200x800).
        //
        // OJO: el Slide de Calamares no ocupa toda la presentación, sino el
        // centro (x 5%, y 20%, ancho 90%, alto 70%). Por eso este contenedor
        // se desplaza en negativo y toma el tamaño de la presentación
        // (masterWidth/masterHeight) y no el del Slide.
        Item {
            x: -parent.x
            y: -parent.y
            width: parent.masterWidth
            height: parent.masterHeight
            clip: true

            Image {
                id: slide1
                source: "slide1.png"
                property real cover: Math.max(parent.width / 1200, parent.height / 800)
                width: 1200 * cover
                height: 800 * cover
                x: (parent.width - width) / 2
                y: (parent.height - height) * 0.8
            }
        }
        Text {
            font.family: "Ubuntu"
            font.pixelSize: 13
            color: "#6d526b"
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.verticalCenter: parent.verticalCenter

text: qsTr("<h1>Quirinux GNU/Linux Versión 2.2</h1>" + 

"<h2>Basado en Devuan (Debian) GNU/Linux</h2>" + 

"<h3>https://www.quirinux.org/</h3><br/>" +

"<b>Versión:</b> 2.2 Rev. 2 - 04-09-2026 <br/><br/>" +

"<b>Autor:</b> Charlie Martínez.<br/><br/>" +

"<b>Logotipo Quirinux:</b> Thomas Gaya.<br/>"+
"<b>Sistema de Creación ISO:</b> Penguins' Eggs, de Piero Proietti.<br/>"+
"<b>Capacitación y auditoria de Seguridad:</b> Javier Obregón.<br/>"+
"<b>Pruebas de compatibilidad en Mac:</b> Sela González. <br/><br/>"+
"<b>Colaboradora en mejoras UX y de accesibilidad:</b> Noelia Gerbaudo.<br/><br/>"+

"<b>Hecho en:</b> Stgo. de Compostela (Galicia, España), Misiones (Argentina) y Roma (Italia).<br/>"+
"<b>Quirinux</b> es Marca Registrada.<b> Distribuidor oficial:</b> CREALIB TECNOLOGÍA SOSTENIBLE.<br/>"+

"<b>Dedicado a Emilio Gorini (qepd).</b>")
            wrapMode: Text.WordWrap
            width: parent.width * 0.9
            // Si el bloque no cabe en vertical (ventana pequeña), se reduce en bloque
            scale: Math.min(1, (parent.height - 16) / height)
            horizontalAlignment: Text.AlignHCenter
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
