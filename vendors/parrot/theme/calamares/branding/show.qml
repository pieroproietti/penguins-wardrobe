/* === This file is part of Calamares - <http://github.com/calamares> ===
 *
 *   Copyright 2015, Teo Mrnjavac <teo@kde.org>
 *   Copyright 2018-2019, Jonathan Carter <jcc@debian.org>
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
        interval: 20000
        running: presentation.activatedInCalamares
        repeat: true
        onTriggered: nextSlide()
    }

    Slide {
        Image {
            id: background1
            source: "background.jpg"
            width: 640; height: 480
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
        }
    }

    Slide {
        Image {
            id: security
            source: "security.png"
            width: 200; height: 200
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
        }
        Text {
            anchors.horizontalCenter: security.horizontalCenter
            anchors.top: security.bottom
            text: "<h3>Designed for Cyber Security experts.</h3><br/>"+
                  "A complete arsenal of security tools at your fingertips."
            wrapMode: Text.WordWrap
            width:650
            horizontalAlignment: Text.Center
        }
    }
    Slide {
        Image {
            id: privacy
            source: "privacy.png"
            width: 200; height: 200
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
        }
        Text {
            anchors.horizontalCenter: privacy.horizontalCenter
            anchors.top: privacy.bottom
            text: "<h3>Privacy is a first-class citizen!</h3><br/>"+
                  "A secure and hardened system ready to safely surf the web, privately communicate with others<br>and keep sensitive information safe.<br>"+
                  "No telemetry, no bullshit."
            wrapMode: Text.WordWrap
            width:650
            horizontalAlignment: Text.Center
        }
    }
    Slide {
        Image {
            id: learning
            source: "learn.png"
            width: 200; height: 200
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
        }
        Text {
            anchors.horizontalCenter: learning.horizontalCenter
            anchors.top: learning.bottom
            text: "<h3>Amazing place for beginners.</h3><br/>"+
                  "Many resources to learn from. Make your skills better every day.<br>"+
                  "The system is designed to incentivate newbies to adopt good habits and best practices from the beginning."
            wrapMode: Text.WordWrap
            width:650
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: development
            source: "dev.png"
            width: 200; height: 200
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
        }
        Text {
            anchors.horizontalCenter: development.horizontalCenter
            anchors.top: development.bottom
            text: "<h3>Software development made easy.</h3><br/>"+
                  "A full development stack with the best editors and languages out of the box.<br>"+
                  "Awesome native support for all the mainstream languages like Python, Go, Rust, NodeJS, Nim,<br>"+
                  "Java, Ruby, C, C++, .NET/MONO, Bash, Erlang/Elixir and many many more! No external repositories required."
            wrapMode: Text.WordWrap
            width:650
            horizontalAlignment: Text.Center
        }
    }
    Slide {
        Image {
            id: community
            source: "community.png"
            width: 200; height: 200
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
        }
        Text {
            anchors.horizontalCenter: community.horizontalCenter
            anchors.top: community.bottom
            text: "<h3>Strong community</h3><br/>"+
                  "A vast community of experts who want to know you and help you in your cyber security journey.<br>"+
                  "The Parrot Community speaks your language and can help you with your issues, or listen to your ideas."
            wrapMode: Text.WordWrap
            width:650
            horizontalAlignment: Text.Center
        }
    }

    Slide {
        Image {
            id: infrastructure
            source: "server.png"
            width: 200; height: 200
            fillMode: Image.PreserveAspectFit
            anchors.centerIn: parent
        }
        Text {
            anchors.horizontalCenter: infrastructure.horizontalCenter
            anchors.top: infrastructure.bottom
            text: "<h3>Rock solid infrastructure</h3><br/>"+
                  "A vendor-neutral, resilient and distributed infrastructure, worldwide mirrors and gateways to make software and information more accessible even to people in totalitarian regimes or in poorly connected locations where stable internet connection to western countries is not an option."
            wrapMode: Text.WordWrap
            width:650
            horizontalAlignment: Text.Center
        }
    }
}
