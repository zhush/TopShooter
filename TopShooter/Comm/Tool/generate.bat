rem �л���.protoЭ�����ڵ�Ŀ¼  
cd ./proto  
rem ����ǰ�ļ����е�����Э���ļ�ת��Ϊlua�ļ�  
for %%i in (*.proto) do (    
echo %%i  
"..\protogen.exe" -i:%%i -o:./csharp/%%~ni.cs  
"..\protoc.exe" --plugin=protoc-gen-go.exe --go_out=./golang %%i  
)  
cd ../ 
echo end  
pause  