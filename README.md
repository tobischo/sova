# Sova

<img src="build/appicon.png" alt="Sova Icon" width="200"/>

Sova is an app which allows to define a time or a timer to set a computer into sleep or hibernate mode or to turn it off altogether.
It was built using https://wails.io/, which is also needed for building it.
It uses vuejs and typescript for the frontend.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

For setting up `wails` see https://wails.io/docs/gettingstarted/installation
