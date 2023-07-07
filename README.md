# Ogredock
This is a Docker front-end built in GO.

It's designed to be super simple, easy to extend, and a viable alternative to the proprietary/paid Docker front-ends out there. 
I use Docker for low-level Ethernet mac stuff, so one noticable thing that's missing is port configuration. You want it, add it! 

This front-end uses the GO Docker API (located here: github.com/docker/docker) along with the GO net/http and html/template packages.

## Getting Started

Just pull the repo and run: <br />
`go run main.go`<br />

The web server will start at your loopback address on port 8080. <br /> 
Feel free to change this! Look in webmod/webmod.go! <br />

## Multi Window View

![Image](/docs/OgreDockView.png)
