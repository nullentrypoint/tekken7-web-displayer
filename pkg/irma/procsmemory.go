package irma

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

// ProcessInformation wrap basic process information and memory dump in a structure
type ProcessInformation struct {
	PID         uint32
	ProcessName string
	ProcessPath string
	MemoryDump  []byte
}

// GetProcessMemory return a process memory dump based on its handle
func GetProcessMemory(pid uint32, handle windows.Handle, verbose bool) (ProcessInformation, []byte, error) {

	procFilename, modules, err := GetProcessModulesHandles(handle)
	if err != nil {
		return ProcessInformation{}, nil, fmt.Errorf("Unable to get PID %d memory: %s", pid, err.Error())
	}

	for _, moduleHandle := range modules {
		if moduleHandle != 0 {
			moduleRawName, err := GetModuleFileNameEx(handle, moduleHandle, 512)
			if err != nil {
				return ProcessInformation{}, nil, err
			}
			moduleRawName = bytes.Trim(moduleRawName, "\x00")
			modulePath := strings.Split(string(moduleRawName), "\\")
			moduleFileName := modulePath[len(modulePath)-1]

			memdump, err := DumpModuleMemory(handle, moduleHandle, verbose)
			if err != nil {
				return ProcessInformation{}, nil, err
			}

			if procFilename == moduleFileName {
				return ProcessInformation{PID: pid, ProcessName: procFilename, ProcessPath: string(moduleRawName)}, memdump, nil
			}
		}
	}

	return ProcessInformation{}, nil, fmt.Errorf("Unable to get PID %d memory: no module corresponding to process name", pid)
}

// KillProcessByID try to kill the specified PID
func KillProcessByID(procID uint32, verbose bool) (err error) {
	hProc, err := GetProcessHandle(procID, windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_TERMINATE)
	if err != nil && verbose {
		//logMessage(LOG_ERROR, "[ERROR]", "PID", procID, err)
		return err
	}

	exitCode := GetExitCodeProcess(hProc)
	err = windows.TerminateProcess(hProc, exitCode)
	if err != nil {
		return err
	}

	return nil
}

// GetProcessesList return PID from running processes
func GetProcessesList() (procsIds []uint32, bytesReturned uint32, err error) {
	procsIds = make([]uint32, 2048)
	err = windows.EnumProcesses(procsIds, &bytesReturned)
	return procsIds, bytesReturned, err
}

// GetProcessHandle return the process handle from the specified PID
func GetProcessHandle(pid uint32, desiredAccess uint32) (handle windows.Handle, err error) {
	handle, err = windows.OpenProcess(desiredAccess, false, pid)
	return handle, err
}

// GetProcessModulesHandles list modules handles from a process handle
func GetProcessModulesHandles(procHandle windows.Handle) (processFilename string, modules []syscall.Handle, err error) {
	var processRawName []byte
	processRawName, err = GetProcessImageFileName(procHandle, 512)
	if err != nil {
		return "", nil, err
	}
	processRawName = bytes.Trim(processRawName, "\x00")
	processPath := strings.Split(string(processRawName), "\\")
	processFilename = processPath[len(processPath)-1]

	modules, err = EnumProcessModules(procHandle, 32)
	if err != nil {
		return "", nil, err
	}

	return processFilename, modules, nil
}

// DumpModuleMemory dump a process module memory and return it as a byte slice
func DumpModuleMemory(procHandle windows.Handle, modHandle syscall.Handle, verbose bool) ([]byte, error) {
	moduleInfos, err := GetModuleInformation(procHandle, modHandle)
	if err != nil && verbose {
		//logMessage(LOG_ERROR, "[ERROR]", err)
		return nil, err
	}

	memdump, err := ReadProcessMemory(procHandle, moduleInfos.BaseOfDll, uintptr(moduleInfos.SizeOfImage))
	if err != nil && verbose {
		//logMessage(LOG_ERROR, "[ERROR]", err)
		return memdump, err
	}

	memdump = bytes.Trim(memdump, "\x00")
	return memdump, nil
}

// WriteProcessMemoryToFile try to write a byte slice to the specified directory
func WriteProcessMemoryToFile(path string, file string, data []byte) (err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0600); err != nil {
			return err
		}
	}

	if err := os.WriteFile(path+"/"+file, data, 0644); err != nil {
		return err
	}

	return nil
}