@REM 
set "varx=%cd%" 
mkdir go_workspace go_workspace\bin go_workspace\src go_workspace\pkg
SETX GOPATH "%varx%go_workspace"  
@REM SETX PATH "%PATH%%varx%\go_workspace\bin;" 
set "varx=" 