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
            source: "eagle.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#000000"
            anchors.horizontalCenter: reproductiveSystem.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Proxmox VE</h1><br/>"+
                  "<h3>Proxmox VE is a complete, open-source server management platform for enterprise virtualization</h3><br/>"+
                  "<h4>https://www.proxmox.com/en/proxmox-ve</h4>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
        }
    }
    Slide {
        Image {
            id: startReproduction
            source: "eagle.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#000000"
            anchors.horizontalCenter: startReproduction.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Proxmox VE</h1><br/>"+
                  "<h3>It tightly integrates the KVM hypervisor and Linux Containers (LXC), software-defined storage and networking functionality, on a single platform.</h3>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: itsYourSystem
            source: "eagle.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#000000"
            anchors.horizontalCenter: itsYourSystem.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Proxmox VE</h1><br/>"+
                  "<h3>With the integrated web-based user interface you can manage VMs and containers, high availability for clusters, or the integrated disaster recovery tools with ease-</h3>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: eggsPresentation
            source: "eagle.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#000000"
            anchors.horizontalCenter: eggsPresentation.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Proxmox VE</h1><br/>"+
                  "<h3>The enterprise-class features and a 100% software-based focus make Proxmox VE the perfect choice to virtualize your IT infrastructure, optimize existing resources, and increase efficiencies with minimal expense</h3>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: waitHatching
            source: "eagle.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#000000"
            anchors.horizontalCenter: waitHatching.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Proxmox VE</h1><br/>"+
                  "<h3>You can easily virtualize even the most demanding of Linux and Windows application workloads, and dynamically scale computing and storage as your needs grow, ensuring that your data center adjusts for future growth.</h3>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: followPenguins
            source: "eagle.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 22
            color: "#000000"
            anchors.horizontalCenter: slide6.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Proxmox VE</h1><br/>"+
                  "<h3>Ready to build an open and future-proof data center with Proxmox VE?</h3>")
            wrapMode: Text.WordWrap
            width: 800
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: createdBy
            source: "eagle.png"
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
            anchors.horizontalCenter: followPenguins.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Penguins' eggs</h1><br/>"+
                  "<h2>Created by Piero Proietti</h2><br/>"+
                  "<h4>issues: htts://github.com/pieroproietti/penguins-eggs</h4><br/>"+
                  "<h3>email: piero.proietti@gmail.com</h3>")
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

