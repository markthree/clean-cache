package main

import (
	"os"
	"path"
)

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func IsDir(f string) bool {
	s, err := os.Stat(f)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func RemoveAll(f string) error {
	entrys, err := os.ReadDir(f)
	if err != nil {
		return err
	}

	entrysLen := len(entrys)

	signal := struct{}{}
	errChan := make(chan error)
	chs := make(chan struct{}, entrysLen)

	for i := 0; i < entrysLen; i++ {
		go func(entry os.DirEntry) {
			p := path.Join(f, entry.Name())
			if !entry.IsDir() {
				err := os.Remove(p)
				if err != nil {
					errChan <- err
					return
				}
				chs <- signal
				return
			}
			err := RemoveAll(p)
			if err != nil {
				errChan <- err
				return
			}
			chs <- signal
		}(entrys[i])
	}

	for i := 0; i < entrysLen; i++ {
		select {
		case <-chs:
			continue
		case nestedErr := <-errChan:
			return nestedErr
		}
	}
	rootErr := os.Remove(f)
	if rootErr != nil {
		return rootErr
	}
	return nil
}
