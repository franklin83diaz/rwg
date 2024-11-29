package pkg

import (
	"log"
	"os/exec"
	"time"
)

func CheckOut(interval time.Duration, ip string, serviceName string, wgInterface string, service bool) {
	for {
		if !isConnectionActive(ip) {
			if service {
				restartWireGuardServ(wgInterface)
			} else {
				restartWireGuard(serviceName)
			}
		}
		time.Sleep(interval)
	}
}

// execute a command ping
func executeCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error to execute command:", string(output))
	}
	return err
}

// check if the connection is active
func isConnectionActive(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", "-W", "2", ip)
	err := cmd.Run()

	if err != nil {
		time.Sleep(3 * time.Second) // wait before retry
		err = cmd.Run()
	}

	if err != nil {
		time.Sleep(2 * time.Second) // wait before retry
		err = cmd.Run()
	}

	return err == nil
}

// restart the wireguard service
func restartWireGuardServ(serviceName string) {
	log.Println("Connection down. Restarting WireGuard...")
	if err := executeCommand("systemctl", "restart", serviceName); err != nil {
		log.Println("Error to down the interface:", err)
		return
	}
}

// restart the wireguard connection
func restartWireGuard(wgInterface string) {
	log.Println("Conexión caída. Reiniciando WireGuard...")
	if err := executeCommand("wg-quick", "down", wgInterface); err != nil {
		log.Println("Error al bajar la interfaz:", err)
		return
	}
	time.Sleep(2 * time.Second) // Espera antes de reiniciar
	if err := executeCommand("wg-quick", "up", wgInterface); err != nil {
		log.Println("Error al subir la interfaz:", err)
	}
}
