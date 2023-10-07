# Saver4
CNC upload/download program

# Docs
- Description of program upload/download from the Citizen-Cincom manual can be found the `docs` folder under `CNC_IO.pdf`
- The source code for the previous upload/download program (Saver3) can be found in `docs` under `Saver3.bas`. The previous version was written in Visual Basic for Windows XP

# To run the Demo
You need to have Go and gcc installed

### On Mac/Linux*
- Note: on linux it may be neccessary to install some graphics libraries to provide the neccessary headers for compilation
```
$go build
$./Saver4
```

### On Windows
- Note that you will need to make sure mingw-gcc is installed and that `$CC` (the C compiler environment variable) is set correctly
```
$CGO_ENABLED=1 GOOS=windows CC=x86_64-w64-mingw32-gcc go build
$Saver4.exe
```

For more info on compiling `fyne` gui libraries, see [fyne docs](https://developer.fyne.io/started/cross-compiling)


# TODO for this group
Try to compile and run the program

# Useful links
- [Serial Port Basics](https://tldp.org/HOWTO/Serial-HOWTO-4.html)
- An article on debugging issues with usb to rs232 adapters can be found [here](https://www.campbellsci.com/blog/usb-rs-232-adapter-cable-issues)
