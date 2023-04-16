//go:build !android

package audio

func checkAndSetRecorderPermission() (err error) {
	microphonePermissionGranted = true
	return nil
}

// requestPermission calls platform's native method
func requestPermission(view uintptr) (err error) {
	microphonePermissionGranted = true
	return nil
}
