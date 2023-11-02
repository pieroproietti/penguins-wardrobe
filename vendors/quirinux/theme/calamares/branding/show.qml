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
    from eggs author: Piero Proietti <piero.proietti@gmail.com>,
    modified by Charlie Martinez for Quirinux GNU/Linux <cmartinez@quirinux.org>
*/
import QtQuick 2.0;
import calamares.slideshow 1.0;


Presentation {
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
            source: "slide1.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 12
            color: "#1a1a1a"
            anchors.horizontalCenter: slide1.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Quirinux General</h1>" +
                  "<h2>Based on Debian GNU/Linux</h2>" +
                  "<h3>https://www.quirinux.org/</h3>" +
                  "<p>The UX and UI experience of Quirinux Pro, the operating system of animated movies, now also in this edition designed for general users, which includes everything except programs for animation.</p>")
            wrapMode: Text.WordWrap
            width: kde.width
            horizontalAlignment: Text.Center
        }
    }
    Slide {
        Image {
            id: slide2
            source: "slide2.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 12
            color: "#1a1a1a"
            anchors.horizontalCenter: slide2.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Quirinux GNU/Linux</h1>" +
                  '<h2>The "Chameleon":</h2>' +
                  "<h3>The only operating system that adapts to you</h3>"+
                  "<p>Choose from 12 screen themes including retro or vintage styles.</p>")
            wrapMode: Text.WordWrap
            width: neon.width
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: slide3
            source: "slide3.png"
            anchors.centerIn: parent
            anchors.top: background.bottom
            width: 810
            height: 485
            fillMode: Image.PreserveAspectFit
        }
         Text {
            font.family: "Helvetica"
            font.pixelSize : 12
            color: "#1a1a1a"
            anchors.horizontalCenter: slide3.horizontalCenter
            anchors.top: background.top
            text: qsTr("<h1>Quirinux GNU/Linux</h1>" +
                  "<h2>Community support</h2>"+
                  "<p>You can create a free account at http://foro.quirinux.org</p>")
            wrapMode: Text.WordWrap
            width: plasma.width
            horizontalAlignment: Text.Center
        }
    }

    Slide {
		Image {
			id: slide4
			source: "slide4.png"
			anchors.centerIn: parent
			anchors.top: background.bottom
			width: 810
			height: 485
			fillMode: Image.PreserveAspectFit
		}
			Text {
			font.family: "Helvetica"
			font.pixelSize : 12
			color: "#1a1a1a"
			anchors.horizontalCenter: slide4.horizontalCenter
			anchors.top: background.top
			text: qsTr("<h1>Quirinux GNU/Linux</h1>" +
				"<h2>Created by: Charlie Martinez</h2>" +
				"<h3>email: cmartinez@quirinux.org</h3>"+
				"<p>Logo created by: Thomas Gaya</p>"+
				"<p>Beta Tester: Leonardo Besingor</p>")
			wrapMode: Text.WordWrap
			width: secure.width
			horizontalAlignment: Text.Center
		}
    } 		
}
