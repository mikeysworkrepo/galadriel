This is a WIP local batchapp RMM type application
I am creating this to not only learn windows systems, but also to continue learning go

The app uses the framework wails: https://wails.io/
This framework allows you to build the frontend in Javascript/Typescript, while creating the backend and core logic in Go

To build this on your own, after installing the dependencies from NPM, you must install the wails framework here: https://wails.io/docs/gettingstarted/installation/

To run the dev build that allows printing to the terminal just type: wails dev
To build into an executable type: wails build

It builds the exe in the src/build/bin folder. I reccomend running as admin

To do back end:
Create custom error variables for different funcitons, and possibly a debug console for the build
I may possibly decide to bring over mass deployment functionality from my CLI version, though it will only be for learning purposes, as I do not feel
like spending the time building out all the error handling for this functionality 

To do frontend:
implement search funcitonality and possibly create tabs/groups of different OU's instead of scrolling to your desired computer
