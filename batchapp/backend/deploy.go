package backend

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// to do
// eventually make more custom error variables for each function, and create a debug console for the build
type DeployFunc func(targetPC string) error

func (a *App) DeployToTargets(targetPCs []string, deploy DeployFunc) {
	for _, target := range targetPCs {
		fmt.Printf("\n Deploying to %s...\n", target)
		if err := deploy(target); err != nil {
			fmt.Printf("Failed on %s: %v\n", target, err)
		} else {
			fmt.Printf("Success on %s\n", target)
		}
	}
}

func (a *App) DeployOffice(targetPC string) error {
	return downloadAndExecuteScript(
		targetPC,
		"http://raspberrypi.local:8080/scripts/officeInstall.ps1",
		"C:\\Windows\\Temp\\officeInstall.ps1",
	)
}

func (a *App) DeploySophos(targetPC string) error {
	return downloadAndExecuteScript(targetPC,
		"http://raspberrypi.local:8080/scripts/sophos.ps1",
		"C:\\Windows\\Temp\\Sophos.ps1")

}

func (a *App) DeploySentinel(targetPC string) error {
	return downloadAndExecuteScript(targetPC,
		"http://raspberrypi.local:8080/scripts/sentinel.ps1",
		"C:\\Windows\\Temp\\sentinel.ps1")

}

func (a *App) HostName() (string, error) {
	getHost := `hostname`

	cmd := exec.Command("powershell.exe", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", getHost)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get hostname: %v\nOutput: %s", err, out.String())
	}

	// Trim whitespace and return hostname
	return strings.TrimSpace(out.String()), nil
}

func (a *App) DeployPrinters(targetPC string) error {
	return downloadAndExecuteScript(
		targetPC,
		"http://raspberrypi.local:8080/scripts/Install-Printers.ps1",
		"C:\\Windows\\Temp\\Install-Printers.ps1",
	)
}

func downloadAndExecuteScript(targetPC, remoteURL, localPath string) error {
	// this function is a blanket download and install function for the software deployments

	// Downloads the script to the target pc in the temp folder
	downloadCmd := exec.Command(
		"C:\\Windows\\PsExec64.exe",
		"\\\\"+targetPC,
		"-s",
		"powershell", "-NoProfile", "-WindowStyle",
		"Hidden", "-ExecutionPolicy", "Bypass", "-Command",
		fmt.Sprintf("Invoke-WebRequest -Uri '%s' -OutFile '%s'", remoteURL, localPath),
	)
	downloadCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	if out, err := downloadCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("download failed: %v\nOutput: %s", err, string(out))
	}

	//  Run the script
	runCmd := exec.Command(
		"C:\\Windows\\PsExec64.exe",
		"\\\\"+targetPC,
		"-s",
		"powershell.exe", "-ExecutionPolicy", "Bypass", "-NoProfile", "-File", localPath,
	)

	runCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	if out, err := runCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("execution failed: %v\nOutput: %s", err, string(out))
	}

	return nil
}
