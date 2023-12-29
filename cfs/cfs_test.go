package cfs

import "testing"

func TestDemo(t *testing.T) {

}

func Test_extract7z(t *testing.T) {
	// buggy now
	//extract7z(Data7z, "C:\\Users\\bigtear\\.chaos")
}

func Test_extractEmbedFS(t *testing.T) {
	return
}

func TestExtractAll(t *testing.T) {
	err := ExtractAll("C:\\Users\\bigtear\\.config\\mihomo")
	if err != nil {
		panic(err)
	}
}
