package main

import (
	"os"
	"path"
)

func IsDir(f string) bool {
	s, _ := os.Stat(f)
	return s.IsDir()
}

func RemoveAll(f string) error {
	entrys, err := os.ReadDir(f)
	if err != nil {
		return err
	}

	entrysLen := len(entrys)

	chs := make(chan interface{}, entrysLen)
	errChan := make(chan error)

	for _, entry := range entrys {
		go func(entry os.DirEntry) {
			p := path.Join(f, entry.Name())
			if entry.IsDir() {
				nestedErr := RemoveAll(p)
				if nestedErr != nil {
					errChan <- nestedErr
				}
				return
			}
			os.Remove(p)
		}(entry)
	}

	for i := 0; i < entrysLen; i++ {
		select {
		case nestedErr := <-errChan:
			return nestedErr
		case <-chs:
		}
	}
	return nil
}
