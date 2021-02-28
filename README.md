# icontest

systray icon test tool

We wanted a simple way to preview icons that will be used in the systray.

`icontest` creates a systray menu using <https://github.com/getlantern/systray>

It provides a menu item for each image you like to preview and sets the menu
icon when you select that image.

On Windows, if the image is a png or jpg it will be converted to an ico.

## usage

```sh
icontest image.png image.jpg image.ico...
```

