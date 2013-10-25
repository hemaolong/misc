@echo off
SetLocal EnableDelayedExpansion
set project_path=E:\work\new_project
set target_path=hcp\
@rem default values
set needPacketResource=false
set pakVersion=zh
SetLocal DisableDelayedExpansion
PATH=%PATH%;D:\program files\Microsoft Visual Studio 8\Common7\IDE
PATH=%PATH%;C:\Program Files\TortoiseSVN\bin"

set   /p pakVersion=������汾{zh, tw, english}��
set   /p buildClient=�Ƿ�Ҫ���ɿͻ���{no,yes}��
set   /p packetClient=�Ƿ�Ҫ����ͻ�����Դ{no,yes}��
set   /p buildServer=�Ƿ�Ҫ���ɷ�����{no,yes}��
set   /p packetServer=�Ƿ�Ҫ�����������Դ{no,yes}��

@rem prompt
if "%buildClient%" == "yes" (
    echo build client
)
if "%packetClient%" == "yes" (
    echo packet client
    set needPacketResource=true
)

@rem server
if "%buildServer%" == "yes" (
    echo build server
)
if "%packetServer%" == "yes" (
   echo packet server
    set needPacketResource=true
)
echo target path: %target_path%

@rem clean the target path
del /Q %target_path%\*.hcp %target_path%*.exe

@rem Just do it
echo %needPacketResource%
if "%needPacketResource%" == "true" (
    echo "dump the svn resources..."
    rd bin /s /q
    TortoiseProc  /command:dropexport /droptarget:"E:\work\pak_game" /path:"E:\work\new_project\bin"
    
    for /r E:\work\pak_game\bin\model  %%i  in (*.tga *.bmp *.jpg) do (
        texconv -m 1 -f DXT5 -d 0 -nologo -o %%~dpi "%%i"
        del /Q "%%i"
    )

    for /r E:\work\pak_game\bin\scene  %%i  in (*.tga *.bmp *.jpg) do (
        texconv -m 1 -f DXT5 -d 0 -nologo -o %%~dpi "%%i"
        del /Q "%%i"
    )

    for /r E:\work\pak_game\bin\ui  %%i  in (*.tga *.bmp *.jpg) do (
        texconv -m 1 -f DXT5 -d 0 -nologo -o %%~dpi "%%i"
        del /Q "%%i"
    )
    
    for /r E:\work\pak_game\bin\local  %%i  in (*.tga *.bmp *.jpg) do (
        texconv -m 1 -f DXT5 -d 0 -nologo -o %%~dpi "%%i"
        del /Q "%%i"
    )
)
@echo on

@if "%buildClient%" == "yes" (
    echo "build client ..."
    devenv�� %project_path%\source\win\source.sln /project game_client /build PUBLISH
)
@if "%packetClient%" == "yes" (
    echo "package(client %pakVersion%) ..."
    packet_tool -command client %pakVersion%
)
@if "%packetServer%" == "yes" (
    echo "package(server %pakVersion%) ..."
    packet_tool -command server %pakVersion%
)
@if "%buildServer%" == "yes" (
    echo "build server ..."
)

@rem move the target files
@echo "dump target files"
@if "%buildClient%" == "yes" (
    copy %project_path%\bin\client.exe  hcp\
)

@rem copy the files to linux server
ftp -s:ftp.txt


