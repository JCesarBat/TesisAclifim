package directoryManager

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMoveFiles(t *testing.T) {
	m := newTestManager(t)
	err := m.copyFile("test_documents/mio.pptx", "../../files/Habana/mio.pptx")
	require.NoError(t, err)
}
