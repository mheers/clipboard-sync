REM read the .env line by line
for /f "tokens=1,2 delims==" %%a in ('type .env') do (
    REM set the environment variable
    set "%%a=%%b"
)

REM start the programm
clipboardsync.exe start

