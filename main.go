package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func copy(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func main() {
	n := flag.Int("n", 1, "Copy file if index % n == 0")
	flag.Parse()

	cwd := "."
	outDir := "./out/"

	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		log.Fatal(err)
	}

	for index, file := range files {
		fullPath := path.Join(cwd, file.Name())
		fi, err := os.Lstat(fullPath)
		if err != nil {
			log.Fatalf("Stat failed: %v\n", err)
		}
		if fi.IsDir() {
			log.Printf("Skip directory: %s\n", file.Name())
			continue
		}
		if !fi.Mode().IsRegular() {
			log.Fatalf("Not a regular file: %s\n", file.Name())
		}

		if index%*n == 0 {
			outPath := path.Join(cwd, outDir, file.Name())
			log.Printf("Copy file(%d): %s -> %s\n", index, fullPath, outPath)

			if err := copy(fullPath, outPath); err != nil {
				log.Fatalf("Copy failed: %v\n", err)
			}
		}
	}
}
