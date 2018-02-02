@echo off
call %~dp0\jam -s JAMFILE=flatc.jam %*
copy %~dp0\.build\win64-release\TOP\flatc\flatc.exe %~dp0\..\..\tools
