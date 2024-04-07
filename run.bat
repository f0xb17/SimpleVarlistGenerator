@echo off
REM Check if Python is installed
python --version 2>NUL
if errorlevel 1 goto errorNoPython

REM Reaching here means Python is installed
REM Execute your Python script
:ReRun
python main.py


REM After finishing just rerun the same script again.
REM jump to rerun
goto ReRun

:errorNoPython
echo Python is not installed. Please install Python and try again.
pause