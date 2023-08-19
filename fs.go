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
	signalChan := make(chan struct{}, entrysLen)
	resolve, reject, status := StatusPromise(signalChan, errChan)

	for i := 0; i < entrysLen; i++ {
		go func(entry os.DirEntry) {
			p := path.Join(f, entry.Name())
			if !entry.IsDir() {
				err := os.Remove(p)
				if err != nil {
					reject(err)
					return
				}
				resolve(signal)
				return
			}
			err := RemoveAll(p)
			if err != nil {
				reject(err)
				return
			}
			resolve(signal)
		}(entrys[i])
	}

	for i := 0; i < entrysLen; i++ {
		_, err := status()
		if err != nil {
			return err
		}
	}
	rootErr := os.Remove(f)
	if rootErr != nil {
		return rootErr
	}
	return nil
}
