$log = 'C:\Windows\Temp\installPrinterLog.txt'
$remoteScriptPath = 'C:\Windows\Temp\printersRemoteInstall.ps1'

try {
    $scriptBlock = @'
$script = "C:\Windows\Temp\Install-Printers.ps1"


"⏬ Downloading printer script..." | Out-File -FilePath $log
Invoke-WebRequest -Uri "http://raspberrypi.local:8080/software/Install-Printers.ps1" -OutFile $remoteScriptPath -ErrorAction Stop



"🚀 Starting Office installation..." | Out-File -Append -FilePath $log
Start-Process -FilePath $script -WindowStyle Hidden -Wait

"✅ Success" | Out-File -Append -FilePath $log
'@

    $scriptBlock | Out-File -Encoding ASCII -FilePath $remoteScriptPath
    Start-Process -FilePath "powershell.exe" -ArgumentList "-ExecutionPolicy Bypass -File `"$remoteScriptPath`"" -Wait
}
catch {
    "❌ Error: $($_.Exception.Message)" | Out-File -Append -FilePath $log
}