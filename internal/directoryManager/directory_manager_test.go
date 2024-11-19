package directoryManager

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	powerPoint   = "test_documents/powerPoint.pptx"
	wordDocument = "test_documents/wordDocument.docx"
)

func TestReadFiles(t *testing.T) {
	manager := newTestManager(t)
	doc, err := manager.ReadFile(wordDocument)
	require.NoError(t, err)
	for _, para := range doc.Paragraphs() {
		fmt.Println(para)
	}
}
