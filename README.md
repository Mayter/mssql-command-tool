# mssql-command-tool

因为wolvez老哥的没有写自定义mssql端口，我就添加了这个功能
<<<<<<< HEAD
官方链接：[http://wolvez.club/2019/09/19/mssql-command-tool/](http://wolvez.club/2019/09/19/mssql-command-tool/)
感谢老哥分享。

示例：
##### Windows环境

```
mssql-command-tool.exe	无压缩	9.47 MB
mssql-command-tool-upx.exe	upx压缩9次	4.63M

mssql-command-tool.exe -host 127.0.0.1 -u sa -p admin -c "whoami" -P 1438
mssql-command-tool.exe -host 127.0.0.1 -u sa -p admin -c "whoami" -port 1438
```

##### Linux环境

```
mssql-command-tool-linux	无压缩	9.81 MB
mssql-command-tool-linux-upx	upx压缩9次	4.91 MB
mssql-command-tool -host 127.0.0.1 -u sa -p admin -c "whoami" -P 1438
mssql-command-tool -host 127.0.0.1 -u sa -p admin -c "whoami" -port 1438	
```

##### Mac环境

```
mssql-command-tool-mac 无压缩	10.2
mssql-command-tool-mac-upx	upx压缩9次	4.9M
./mssql-command-tool --server 127.0.0.1 -u sa -p admin -c "whoami /user" -port 1438
./mssql-command-tool --server 127.0.0.1 -u sa -p admin -c "whoami /user" -P 1438	
```
