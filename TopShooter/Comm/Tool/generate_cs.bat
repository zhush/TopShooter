rem �л���.protoЭ�����ڵ�Ŀ¼  
cd ./proto  
rem ����ǰ�ļ����е�����Э���ļ�ת��Ϊlua�ļ�  
for %%i in (*.proto) do (    
echo %%i  
"..\protogen.exe" --csharp_out=./  %%i  
  
)  
cd ../ 
echo end  
pause  