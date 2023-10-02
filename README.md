# Saver4
CNC upload/download program

# Docs
- Description of program upload/download from the Citizen-Cincom manual can be found the `docs` folder under `CNC_IO.pdf`
- The source code for the previous upload/download program (Saver3) can be found in `docs` under `Saver3.bas`. The previous version was written in Visual Basic for Windows XP

# To run the Demo
You need to have Go and gcc installed

*On Mac/linux*
```
$go build
$./Saver4
```

*On Windows*
```
$CGO_ENABLED=1 GOOS=windows go build
$Saver4.exe
```

For more info on compiling `fyne` gui libraries, see [fyne docs](https://developer.fyne.io/started/cross-compiling)


# TODO for this group
Try to compile and run the program
