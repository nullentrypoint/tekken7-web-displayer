package irma

import (
	"errors"
	"strings"

	"golang.org/x/sys/windows"
)

func FindProcessByName(needProcessName string) (windows.Handle, error) {

	procsIds, _, err := GetProcessesList()
	if err != nil {
		return windows.InvalidHandle, err
	}

	for _, pid := range procsIds {
		if pid == 0 {
			continue
		}
		procHandle, err := GetProcessHandle(pid, windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ)
		if err != nil {
			windows.CloseHandle(procHandle)
			continue
		}
		nameBytes, err := GetProcessImageFileName(procHandle, windows.MAX_PATH)
		if err != nil {
			windows.CloseHandle(procHandle)
			continue
		}

		nameParts := strings.Split(string(nameBytes), "\\")
		var exeName string
		if len(nameParts) > 0 {
			exeName = strings.Trim(nameParts[len(nameParts)-1], " \x00")
		}

		if exeName == needProcessName {
			return procHandle, nil
		}
		windows.CloseHandle(procHandle)
	}

	return windows.InvalidHandle, errors.New("not found process: " + needProcessName)
}