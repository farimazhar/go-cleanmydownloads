package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	home, _ := os.UserHomeDir()
	downloadPath := filepath.Join(home, "Downloads")
	
	folders := map[string]string{
		".jpg": "Images", ".png": "Images", ".jpeg": "Images",
		".pdf": "PDFs", ".zip": "Zips", ".mp4": "Videos",
		".exe": "Apps", ".go": "Code", ".txt": "Docs",
	}

	files, _ := os.ReadDir(downloadPath)
	count := 0

	for _, file := range files {
		if file.IsDir() { continue }
		ext := filepath.Ext(file.Name())
		if folder, ok := folders[ext]; ok {
			src := filepath.Join(downloadPath, file.Name())
			destDir := filepath.Join(downloadPath, folder)
			os.MkdirAll(destDir, 0755)
			dest := filepath.Join(destDir, file.Name())
			
			os.Rename(src, dest)
			// agar rename fail ho to copy
			if _, err := os.Stat(dest); os.IsNotExist(err) {
				copyFile(src, dest)
			}
			fmt.Printf("Moved: %s -> %s\n", file.Name(), folder)
			count++
		}
	}
	fmt.Printf("\nDone! %d files cleaned by Farimazhar ✨\n", count)
}

func copyFile(src, dst string) {
	in, _ := os.Open(src)
	defer in.Close()
	out, _ := os.Create(dst)
	defer out.Close()
	io.Copy(out, in)
}
