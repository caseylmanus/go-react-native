# Go React Native!

This is a small POC of integrating a Go (Golang) shared library and a React Native UI within a mobile application.

To run this demo you don't need anything go related installed just download the source, run
npm install and then run in xcode (or Android studio to run the android version).  For more
information about running a React Native Application, see the docs for that project.

![Demo Screenshot](/screenshot.png?raw=true "Demo Screenshot")

This sample illustrates 3 ways of communicating between your React Native UI 
and a Go shared library
*  Http
*  Websockets
*  Native Bridge provided by React native.

The Go source code is in the /demo folder (demo.go).  To modify it, you will need Go 1.6 and 
the go mobile tool chain.  

Once you modify the demo.go file, then you can run "gomobile bind platform=(ios|android)" to build 
either the new Demo.framework for ios or demo.aar for Android.  Once those are built,
then copy them to the appropriate platform folder (ios/Demo.framework or android/demo/demo.aar)

# NOTE
If you have something else running on PORT 8081, this demo will fail to run the http and websocket examples,
you can change the port number in the index.*.js files and the demo.go file then rebuild.
