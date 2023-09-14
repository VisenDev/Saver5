'             SAVER3.BAS
'
'      *************************
'      *      Copyright        *
'      * American Campro Corp. *
'      *       Aug 2003        *
'      *  All Rights Reserved  *
'      *************************
'

$COM1 2048
$COMPILE EXE "A:\SAVER3.EXE"
$CPU 80286
$DIM ARRAY
$ERROR ALL ON
$FLOAT EMULATE
$OPTIMIZE SPEED
$OPTION CNTLBREAK ON
$STRING 1

' load library modules
$INCLUDE "MATH.MOD"
$INCLUDE "CONS.MOD"
$INCLUDE "IO.MOD"
$INCLUDE "SETS.MOD"
$INCLUDE "GRAF.MOD"
$INCLUDE "MENU.MOD"

$INCLUDE "C:\PB3\CNC\SUBPROGS\CONSTANT.INC"

KEY ON
ON KEY (9) GOSUB ENDPRG
KEY (9) ON

DO
CLS
LOCATE 1,20: PRINT"            SAVER3.BAS 1.1"
LOCATE 3,20: PRINT"      *************************"
LOCATE 4,20: PRINT"      *      Copyright        *"
LOCATE 5,20: PRINT"      * American Campro Corp. *"
LOCATE 6,20: PRINT"      *       Aug 2003        *"
LOCATE 7,20: PRINT"      *  All Rights Reserved  *"
LOCATE 8,20: PRINT"      *************************"
LOCATE 25,27 : PRINT "HIT <F9> TO EXIT ANYTIME";

DIM Op$(5000)
DIM Prg$(5000)
DIM Drive$(3)
DIM File$(5000)

Drive$(1) = "C:\B12FILES\"
Drive$(2) = "C:\B20FILES\"
Drive$(3) = "C:\ML20H\"
J$ = ""
WHILE NOT IN.SET%(J$,"1,2,5,6,3,4,")
   Locate 10,20 : PRINT "Kager Industries"
   LOCATE 11,23 : PRINT "1.) Save programs from CNC to ";Drive$(1)+SPACE$(10)
   LOCATE 12,23 : PRINT "2.) Load program to CNC from ";Drive$(1)+SPACE$(10)
   LOCATE 14,20 : PRINT "Kager & Sauer CNC Machining"
   LOCATE 15,23 : PRINT "5.) Save programs from CNC to ";Drive$(2)+SPACE$(10)
   LOCATE 16,23 : PRINT "6.) Load program to CNC from ";Drive$(2)+SPACE$(10)
   LOCATE 18,20 : PRINT "Hanwha CNC"
   LOCATE 19,23 : PRINT "3.) Save programs from CNC to ";Drive$(3)+SPACE$(10)
   LOCATE 20,23 : PRINT "4.) Load program to CNC from ";Drive$(3)+SPACE$(10)



   J$ = INKEY$
WEND

Select case J$
   Case "1"
      Drive% = %TRUE
      I$ = "S"
      Progstrt$ = "O"
   Case "2"
      Drive% = %TRUE
      I$ = "L"
      Progstrt$ = "O"
   Case "5"
      Drive% = %FALSE
      I$ = "S"
      Progstrt$ = "O"
   Case "6"
      Drive% = %FALSE
      I$ = "L"
      Progstrt$ = "O"
   Case "3"
      Drive% = 1
      I$ = "S"
      Progstrt$ = ":"
   Case "4"
      Drive% = 1
      I$ = "L"
      Progstrt$ = ":"
End Select

IF I$ = "s" OR I$ = "S" THEN

   OPEN "COM1:4800,E,7,2,DS,RS,CS,CD" AS #1 LEN = 1
   CLS
   LOCATE 4,27 : PRINT "Time-out will end up-load"
   LOCATE 5,27 : PRINT "Ready..."
   LineOn% = 0 : Temp$ = "" : Start% = %FALSE : First.CR% = %FALSE
   Begin% = %FALSE : Y = TIMER

   WHILE NOT Begin% OR TIMER - Y < 5
      IF LOC(1) > 0 THEN
         Y = TIMER
         Begin% = %TRUE
         Temp$ = INPUT$(1, #1)
         IF Temp$ = ENTER$ OR Temp$ = FNLine.Feed$ THEN
            INCR LineOn%
            LOCATE 9, 27 : PRINT "Up-loading line ";LineOn%-1
            Op$(LineOn%) = Op$
            Op$ = ""
         ELSE
            Op$ = Op$ + Temp$
         END IF
      END IF
   WEND
   BEEP 1
   CLOSE #1

   LOCATE 4,27 : PRINT "Up-load complete.            "
   LOCATE 5,27 : PRINT SPACE$(8)
   DELAY 2

   CLS
   ERASE Prg$()
   FileIn% = %FALSE : K% = 1 : FileName$ = ""
   FOR I% = 1 TO LineOn%
      IF LEFT$(Op$(I%),1) = Progstrt$ OR I% = LineOn% THEN
         IF (LEFT$(Op$(I%-1),1) <> Progstrt$ AND FileName$ <> "") OR LineOn% = I% THEN
            PRINT "Saving file "+FileName$;
            PRINT USING " Length #### lines";K%
            DELAY 1
            OPEN Drive$(DRIVE%+2)+FileName$+".DL" FOR OUTPUT AS #1
               IF Drive% <> 1 THEN PRINT #1, "START"
               PRINT #1, "%";
               FOR J% = 1 TO K%-1
                  PRINT #1, Prg$(J%)
               NEXT J%
               IF LEFT$(Op$(I%),1) <> Progstrt$ THEN PRINT #1, Op$(I%)
               PRINT #1, "%"
               ERASE Prg$()
            CLOSE #1
            FileIn% = %FALSE
         END IF

         IF NOT FileIn% THEN
            FileIn% = %TRUE
            J% = 1 : FileName$ = "" : Temp$ = ""
            WHILE J% <= LEN(Op$(I%)) AND FileName$ = ""
               C$ = MID$(Op$(I%),J%,1)
               IF C$ = ":" THEN C$ = "O"
               IF NOT IN.SET%(C$,"O,1,2,3,4,5,6,7,8,9,0,") THEN
                  FileName$ = Temp$
               END IF
               Temp$ = Temp$ + C$
               INCR J%
               K% = 1
            WEND
            IF FileName$ = "" THEN FileName$ = Temp$
         END IF
      END IF
      Prg$(K%) = Op$(I%)
      INCR K%
   NEXT I%
   PRINT
   WHILE INKEY$ <> "" : WEND
   PRINT "Hit any key to return to menu..."
   WHILE NOT INSTAT : WEND
   I$ = INKEY$

ELSEIF I$ = "L" OR I$ = "l" THEN

   CLS
   I$ = "" : FileName$ = ""
   LOCATE 4,27 :  PRINT "Enter program number ->"
   WHILE I$ <> ENTER$
      LOCATE 4,51 : PRINT FileName$+"_"
      I$ = INKEY$
      IF I$ = BACK.SPACE$ AND LEN(FileName$) > 0 THEN
         FileName$ = LEFT$(FileName$,LEN(FileName$)-1)
         LOCATE 4,51 : PRINT SPACE$(15)
      END IF
      IF IN.SET%(I$,Digit.set$) AND LEN(FileName$) < 4 THEN FileName$ = FileName$ + I$
   WEND

   LOCATE 4,51 : PRINT FileName$+" "
   LOCATE 6,27: PRINT "Hit any key when ready to send..."
   WHILE NOT INSTAT : WEND : I$ = INKEY$

   OPEN Drive$(DRIVE%+2)+"O"+FileName$+".DL" FOR INPUT AS #2
   PARENTHESIS% = %FALSE
   WHILE NOT EOF(2)
      INPUT #2, Prg$
      IF NOT PARENTHESIS% THEN
         PARENTHESIS% = INSTR(Prg$,")") > 0 AND NOT IN.SET%(LEFT$(Prg$,1),"O,%,")
      END IF
   WEND
   CLOSE #2

   OPEN Drive$(DRIVE%+2)+"O"+FileName$+".DL" FOR INPUT AS #2
   OPEN "COM1:4800,E,7,2,DS,RS,CS,CD" AS #1 LEN = 1
   FIRSTO% = %FALSE : ADDNOW% = %FALSE
   WHILE NOT EOF(2)
      INPUT #2, Prg$
      IF ADDNOW% AND NOT PARENTHESIS% THEN
          Prg$ = Prg$ + "(START)"
          PARENTHESIS% = %TRUE
      END IF
      IF Prg$ = "" THEN Prg$ = ENTER$
      LOCATE 8,27:PRINT SPACE$ (40)
      LOCATE 8,27:PRINT Prg$
      DELAY .04
      PRINT #1, Prg$+FNLine.Feed$
      IF IN.SET%(LEFT$(Prg$,1),"O,%,") THEN
         IF FIRSTO% THEN ADDNOW% = %TRUE ELSE FIRSTO% = %TRUE
      END IF
   WEND
   CLOSE #1
   CLOSE #2
   LOCATE 10,27: PRINT "Hit any key to return to menu..."
   WHILE NOT INSTAT : WEND
   I$ = INKEY$

ELSEIF I$ = "B" OR I$ = "b" THEN

   CLS
   LOCATE 5,27 : PRINT "Press any key to begin backup."

   WHILE NOT INSTAT : WEND : I$=INKEY$
   LOCATE 7,27 : PRINT "Backup in progress..."
   DELAY 1
   MKDIR "e:\b12files"
   MKDIR "e:\b20files"
   LOCATE 9, 1: Print "Backing up C:\B12FILES\";
   Shell "Copy c:\b12files\*.* e:\b12files\"
   Locate 10, 1: Print "Backing up C:\B20FILES\";
   Shell "Copy c:\b20files\*.* e:\b20files\"


   Locate 12, 1: PRINT "Backup complete"
   PRINT "Hit any key to return to menu..."
   WHILE NOT INSTAT : WEND
   I$ = INKEY$

ELSE

   CLS
   I$ = "" : FileName$ = ""
   LOCATE 4,27 :  PRINT "Enter program number ->"
   WHILE I$ <> ENTER$
      LOCATE 4,51 : PRINT FileName$+"_"
      I$ = INKEY$
      IF I$ = BACK.SPACE$ AND LEN(FileName$) > 0 THEN
         FileName$ = LEFT$(FileName$,LEN(FileName$)-1)
         LOCATE 4,51 : PRINT SPACE$(15)
      END IF
      IF IN.SET%(I$,Digit.set$) AND LEN(FileName$) < 4 THEN FileName$ = FileName$ + I$
   WEND

   EXECUTE "PB /PBCONFIG.PB \B12FILES\O"+FileName$+".DL"

END IF

LOOP

ENDPRG:
CLS
