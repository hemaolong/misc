@echo off
SetLocal EnableDelayedExpansion
set project_path=E:\work\new_project
set target_path=hcp\
@rem default values
set needPacketResource=false
set clientCodeVersion=""
SetLocal DisableDelayedExpansion
PATH=%PATH%;C:\Program Files\TortoiseSVN\bin"

set   /p clientCodeVersion=Input the version{2013xxyyzz}:
set   /p dumpClient=Need dump client bin?{no,yes}:
set   /p dumpClientRes=Need dump client resources?{no,yes}:
@rem set   /p buildServer=是否要生成服务器{no,yes}：
@rem set   /p packetServer=是否要打包服务器资源{no,yes}：

@rem prompt
if "%dumpClient%" == "yes" (
    echo build client
)
if "%dumpClientRes%" == "yes" (
    echo packet client
    set needPacketResource=true
)

echo target path: %target_path%

@rem clean the target path
@rem del /Q %target_path%\*.hcp %target_path%*.exe

@rem Just do it
echo %needPacketResource%
if "%needPacketResource%" == "true" (
    echo "dump the svn resources..."
    @rem rd bin /s /q
    TortoiseProc  /command:dropexport /droptarget:"E:\work\webgame\qq_svn_clt_res\game" /path:"E:\work\webgame\web_game\web_bin\assets" /overwrite 
)
@echo on

@if "%dumpClient%" == "yes" (
    echo "dump client ..."
    TortoiseProc  /command:dropexport /droptarget:"E:\work\webgame\qq_svn_clt_res\game" /path:"E:\work\webgame\web_game\web_bin\wg_clt.swf" /overwrite 
)
@if "%dumpClientRes%" == "yes" (
    echo "dump client resource(client %clientCodeVersion%) ..."
    packet_tool -command client %clientCodeVersion%
)

@rem copy the files to linux server
@rem ftp -s:ftp.txt


