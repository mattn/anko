@echo off

setlocal enabledelayedexpansion
set DIR=%~dp0
(cd %DIR%.. && go build)
for %%i in (%DIR%\*.t) do (
  %DIR%..\anko %DIR%\tester.ank %%i
  if !ERRORLEVEL! neq 0 goto error
)
exit /b 0
:error
exit /b 1
