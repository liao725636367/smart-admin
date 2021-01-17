@echo off
:: by oicu
chcp 65001
set APP_NAME=beegoTest
set APP_HOME=%cd%
:menu
echo.%APP_NAME%
echo.%APP_HOME%
echo.=-=-=-=-请选择您的操作-=-=-=-=-
echo.
echo.1: 启动(安装)服务
echo.
echo.2: 关闭服务
echo.
echo.3: 重启服务
echo.
echo.4: 卸载服务
echo.
echo.5: 退出

echo.=-=-=-=-请输入您要选择的项目序号↓-=-=-=-
set /p id=
if "%id%"=="1" goto startup
if "%id%"=="2" goto poweroff
if "%id%"=="3" goto restart
if "%id%"=="4" goto uninstall
if "%id%"=="5" goto exitTask


:exitTask
Exit

rem 安装服务
:install
REM 如果服务不存在，安装服务
echo.服务没有安装,安装服务中...
%APP_NAME%.exe -service-install
goto run


rem 启动服务
:startup
%APP_NAME%.exe -service-query |find /i "Stopped" >nul 2>nul
REM 如果服务存在，跳转至exist标签
if not errorlevel 1 (goto run) 
%APP_NAME%.exe -service-query |find /i "Running" >nul 2>nul
REM 如果服务存在，跳转至exist标签
if not errorlevel 1 (goto Running) else goto install
:run
echo.启动服务中...
%APP_NAME%.exe -service-start
pause
goto menu

:Running
echo.服务正在运行中..
pause
goto menu



rem 停止服务
:poweroff
::休眠方式 避免异常关闭
echo.停止服务中..
%APP_NAME%.exe -service-stop
rem VBoxManage controlvm centos7 poweroff
pause
goto menu



rem 重启服务
:restart
%APP_NAME%.exe -service-query |find /i "Stopped" >nul 2>nul
REM 如果服务存在，跳转至exist标签
if not errorlevel 1 (goto toRun) 
%APP_NAME%.exe -service-query |find /i "Running" >nul 2>nul
REM 如果服务存在，跳转至exist标签
if not errorlevel 1 (
	goto toStop
) else (
	echo.请先安装服务在进行操作...
	goto menu
)

:toStop
echo.停止服务中..
%APP_NAME%.exe -service-stop
:toRun
echo.启动服务中...
%APP_NAME%.exe -service-start
pause
goto menu





rem 卸载服务
:uninstall
::休眠方式 避免异常关闭

%APP_NAME%.exe -service-query |find /i "Running" >nul 2>nul
if not errorlevel 1 (
	echo.停止服务中..
	%APP_NAME%.exe -service-stop
) 
echo.卸载服务中...
%APP_NAME%.exe -service-remove
echo.卸载成功...
rem VBoxManage controlvm centos7 poweroff
pause
goto menu
