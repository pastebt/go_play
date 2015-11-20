// IsTerminal returns true if the given file descriptor is a terminal.
// https://godoc.org/golang.org/x/crypto/ssh/terminal#IsTerminal
// https://github.com/golang/crypto/blob/master/ssh/terminal/util.go#L31
func IsTerminal(fd int) bool {
    var termios syscall.Termios
    _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd),
                                  ioctlReadTermios,
                                  uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
    return err == 0
}

