package tekken

import (
	"encoding/binary"
	"errors"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"

	"github.com/nullentrypoint/tekken7-web-displayer/pkg/irma"
)

func (r Api) readProcessMemory(lpBaseAddress uintptr, nSize uintptr, offsets []int) ([]byte, error) {

	if irma.GetExitCodeProcess(r.handle) != irma.STILL_ACTIVE {
		return nil, errors.New("process not active")
	}

	var size uintptr = 8
	if len(offsets) == 0 {
		size = nSize
	}

	val, err := irma.ReadProcessMemory(r.handle, lpBaseAddress, size)
	if err != nil {
		return nil, err
	}

	for i, offset := range offsets {

		if i == len(offsets)-1 {
			size = nSize
		}

		address := int(binary.LittleEndian.Uint64(val)) + offset
		val, err = irma.ReadProcessMemory(r.handle, uintptr(address), size)
		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func (r Api) findModuleByName(needModuleName string) (syscall.Handle, *irma.ModuleInfo, error) {

	modules, err := irma.EnumProcessModules(r.handle, windows.MAXIMUM_ALLOWED)
	if err != nil {
		return syscall.InvalidHandle, nil, err
	}

	for _, module := range modules {
		data, err := irma.GetModuleFileNameEx(r.handle, module, windows.MAX_PATH)
		if err != nil {
			continue
		}

		nameParts := strings.Split(string(data), "\\")
		var moduleName string
		if len(nameParts) > 0 {
			moduleName = strings.Trim(nameParts[len(nameParts)-1], " \x00")
		}

		if moduleName == needModuleName {

			info, err := irma.GetModuleInformation(r.handle, module)
			if err != nil {
				return syscall.InvalidHandle, nil, err
			}

			return module, &info, nil
		}
	}

	return syscall.InvalidHandle, nil, errors.New("module not found: " + needModuleName)
}
