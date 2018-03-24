rem 切换到.proto协议所在的目录  
cd ./proto  
rem 将当前文件夹中的所有协议文件转换为lua文件  
for %%i in (*.proto) do (    
echo %%i  
"..\protogen.exe" -i:%%i -o:./csharp/%%~ni.cs  
"..\protoc.exe" --plugin=protoc-gen-go.exe --go_out=./golang %%i  
)  
cd ../ 
echo end  
pause  