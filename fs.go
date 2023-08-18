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

	chs := make(chan struct{}, entrysLen)
	errChan := make(chan error)

	for _, entry := range entrys {
		go func(entry os.DirEntry) {
			p := path.Join(f, entry.Name())
			if entry.IsDir() {
				err := RemoveAll(p)
				if err != nil {
					errChan <- err
				}
				return
			}
			err := os.Remove(p)
			if err != nil {
				errChan <- err
				return
			}
			chs <- struct{}{}
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
