# andrec

This is a cross-platform command line tool to record screen of Android device.

Here are advantages over
[the typical way](https://developer.android.com/studio/command-line/shell.html#screenrecord)
 by running ```adb shell screenrecord```.

* **Powerful** - thanks to [FFmpeg](https://ffmpeg.org/), recorded movie can be converted into arbitrary format like animated GIF
* **Simple** - all required tasks are automatically executed, such as recording screen, downloading file from Android to PC and converting file
* **Fast** - all tasks are executed in parallel

## Prerequisites

* [adb](https://developer.android.com/studio/command-line/adb.html)
* [FFmpeg](https://ffmpeg.org/)


## Install

TBD


## How to use

Here is a command to start recording screen and save it to the specified file.

```
$ andrec foo.gif
```

Then, you can stop recording by pressing enter key.


## License

Copyright &copy; 2016 Ikuo Terado. Released under the [MIT license](http://www.opensource.org/licenses/mit-license.php).
