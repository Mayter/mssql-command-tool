# mssql-command-tool

#### 声明
我没有学过go语言可能代码写的很垃圾，但是工具能用就行了。

因为wolvez老哥的没有写自定义mssql端口，我就添加了这个功能 
官方链接：[http://wolvez.club/2019/09/19/mssql-command-tool/](http://wolvez.club/2019/09/19/mssql-command-tool/)
感谢老哥分享。

我又加了一个sp_oacreate组件

示例：
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

##### Mac环境

```
mssql-command-tool-mac

./mssql-command-tool-mac --server 127.0.0.1 -u sa -p admin -c "whoami /user" -port 1438
./mssql-command-tool-mac --server 127.0.0.1 -u sa -p admin -c "whoami /user" -P 1438

开启xp_cmdshell组件
./mssql-command-tool-mac -host 127.0.0.1 -u sa -p admin -P 1438 --enable/--e

开启sp_oacreate组件
./mssql-command-tool-mac -host 127.0.0.1 -u sa -p admin -P 1438 --ole/--o

执行命令：
./mssql-command-tool-mac -host 127.0.0.1 -u sa -p admin -C "c:\windows\system32\cmd.exe /c whoami >c:\whoami.txt -port 1438
```
