package utils

// we need to write
func WriteFile(dstFile string, b []byte) {

	// fp, err := os.OpenFile(snatConfigFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0664))
	// 				if err == nil {
	// 					fp.Write(jsonStr)
	// 					fp.Close()
	// 				} else {
	// 					log.Errorf("[cni-net] failed to save snat settings to %s with error: %+v", snatConfigFile, err)
	// 				}
}

// ioutil.WriteFile(dstFile, filebytes, perm)
