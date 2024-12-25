package zipp

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	purgeTestData = true
	dataDir       = "test_data"
	dataZip       = "x.zip"
)

func TestUnpackAndPackWithoutPrefix(t *testing.T) {
	err := unpackAndPackArchive("test_unpack")
	if err != nil {
		t.Fatalf("Failed to unpack and pack archive: %v", err)
	}
}

func TestPackAndUnpackDirectory(t *testing.T) {
	err := packAndUnpackDir("x", "test_pack")
	if err != nil {
		t.Fatalf("Failed to pack and unpack directory: %v", err)
	}
}

func unpackAndPackArchive(testPrefix string) error {
	srcZip := filepath.Join(dataDir, dataZip)
	dstDir := filepath.Join(dataDir, testPrefix)

	err := Unpack(srcZip, dstDir)
	if err != nil {
		return fmt.Errorf("unpack archive: %v", err)
	}
	if purgeTestData {
		defer os.RemoveAll(dstDir)
	}

	zip := dstDir + ".zip"
	err = Pack(dstDir, zip)
	if err != nil {
		return fmt.Errorf("pack archive: %v", err)
	}
	if purgeTestData {
		defer os.Remove(zip)
	}

	return nil
}

func packAndUnpackDir(srcDir, testPrefix string) error {
	srcDir = filepath.Join(dataDir, srcDir)
	zip := filepath.Join(dataDir, testPrefix+".zip")

	err := Pack(srcDir, zip)
	if err != nil {
		return fmt.Errorf("pack archive: %v", err)
	}
	if purgeTestData {
		defer os.Remove(zip)
	}

	dstDir := filepath.Join(dataDir, testPrefix+"_unpacked")
	err = Unpack(zip, dstDir)
	if err != nil {
		return fmt.Errorf("unpack archive: %v", err)
	}
	if purgeTestData {
		defer os.RemoveAll(dstDir)
	}

	return nil
}
