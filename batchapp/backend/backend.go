package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"syscall"
)

type Computer struct {
	Name   string `json:"Name"`
	IP     string `json:"IP"`
	Status string `json:"Status"`
}

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) GetComputers() ([]Computer, error) {
	// defines the PowerShell command as a Go string and gets all the computers on the domain
	powershellScript := `
	Get-ADComputer -Filter * -Property IPv4Address |
	Sort-Object Name |
	Select-Object Name,
				  @{Name='IP';Expression={($_.IPv4Address) -join ", "}},
				  @{Name='Status';Expression={ if ($_.IPv4Address) { "online" } else { "offline" } }} |
	ConvertTo-Json -Depth 2
	`

	// runs PowerShell silently
	cmd := exec.Command("powershell.exe", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", powershellScript)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // <--- hides the popup window

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("PowerShell error: %v\nOutput: %s", err, out.String())
	}

	// Parse output into Go struct
	var result []Computer
	err := json.Unmarshal(out.Bytes(), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// a vestigual function from my original CLI program, no longer used
func (a *App) DeploySoftware(targetPCs []string) {
	psScriptURL := "http://raspberrypi.local:8080/scripts/officeInstall.ps1"
	localScriptPath := "C:\\Windows\\Temp\\officeInstall.ps1"

	for _, targetPC := range targetPCs {
		fmt.Printf("\n✨ Deploying Office 365 to %s...\n", targetPC)

		// Step 1: Download the script to the target machine
		downloadCmd := exec.Command(
			"C:\\Windows\\PsExec64.exe",
			"\\\\"+targetPC,
			"-s",
			"powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command",
			fmt.Sprintf("Invoke-WebRequest -Uri '%s' -OutFile '%s'", psScriptURL, localScriptPath),
		)

		downloadOut, err := downloadCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("❌ Failed to download script to %s\nOutput:\n%s\n", targetPC, string(downloadOut))
			continue
		}

		// runs as SYSTEM)
		runScriptCmd := exec.Command(
			"C:\\Windows\\PsExec64.exe",
			"\\\\"+targetPC,
			"-s",
			"powershell",
			"-WindowStyle",
			"Hidden",
			"-ExecutionPolicy",
			"Bypass",
			"-File",
			localScriptPath,
		)

		runOut, err := runScriptCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("❌ Error deploying to %s\nOutput:\n%s\n", targetPC, string(runOut))
		} else {
			fmt.Printf("✅ Office 365 deployment started on %s!\n", targetPC)
		}
	}
}
