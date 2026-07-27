package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func getBlob(dir string) ([]byte, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("stat %s: %w", dir, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", dir)
	}

	var buf bytes.Buffer
	if err := writeTar(&buf, dir); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func writeTar(w io.Writer, dir string) error {
	tw := tar.NewWriter(w)
	defer tw.Close()

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return addTarEntry(tw, dir, path, info)
	})
}

func addTarEntry(tw *tar.Writer, root, path string, info os.FileInfo) error {
	if !info.Mode().IsRegular() && !info.IsDir() {
		return nil
	}

	rel, err := filepath.Rel(root, path)
	if err != nil {
		return err
	}
	if rel == "." {
		return nil
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	header.Name = filepath.ToSlash(rel)

	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	if !info.Mode().IsRegular() {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(tw, f)
	return err
}
