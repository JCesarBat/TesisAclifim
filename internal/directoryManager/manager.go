package directoryManager

import (
	"fmt"
	"io"
	"os"
)

type Manager struct {
}

func (m *Manager) copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)

	if err != nil {
		return fmt.Errorf("cannot copy file %w", err)
	}
	sourceFile.Close()
	err = os.Remove(src)
	if err != nil {
		return fmt.Errorf("error al eliminar el archivo original: %w", err)
	}

	return nil
}
