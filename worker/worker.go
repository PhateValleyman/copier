// Package worker obsahuje logiku pro kopírování souborů.
package worker

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyFile zkopíruje jeden soubor ze source do dest.
func CopyFile(source, dest string) error {
	// otevřeme zdrojový soubor
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// vytvoříme cílový soubor
	dstFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// zkopírujeme obsah
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

// CopyDirectory zkopíruje obsah celého adresáře.
func CopyDirectory(srcDir, dstDir string) error {
	// Vytvoříme cílový adresář pokud neexistuje
	err := os.MkdirAll(dstDir, 0755)
	if err != nil {
		return err
	}

	// Projdeme všechny položky ve zdrojovém adresáři
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Spočítáme relativní cestu
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(dstDir, relPath)

		// Pokud je to adresář, vytvoříme jej
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		// Jinak zkopírujeme soubor
		fmt.Println("Copying:", path, "->", destPath)
		return CopyFile(path, destPath)
	})
}
