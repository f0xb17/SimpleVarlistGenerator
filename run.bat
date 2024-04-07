@echo off
REM Check if Python is installed
python --version 2>NUL
if errorlevel 1 goto errorNoPython

REM Reaching here means Python is installed
REM Execute your Python script
python main.py

REM Exit the batch file
exit /b

:errorNoPython
echo Python is not installed. Please install Python and try again.
pause