//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package diskstore implements tempStorage interface
// by using disk as a storage
package diskstore

import (
	_gag "github.com/sanhuanshisanshao/unioffice/common/tempstorage"
	_a "io/ioutil"
	_ga "os"
	_g "strings"
)

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage() { _e := diskStorage{}; _gag.SetAsStorage(&_e) }

// Open opens file from disk according to a path
func (_gf diskStorage) Open(path string) (_gag.File, error) { return _ga.Open(path) }

// RemoveAll removes all files in the directory
func (_d diskStorage) RemoveAll(dir string) error {
	if _g.HasPrefix(dir, _ga.TempDir()) {
		return _ga.RemoveAll(dir)
	}
	return nil
}

// Add is not applicable in the diskstore implementation
func (_gd diskStorage) Add(path string) error { return nil }

type diskStorage struct{}

// TempFile creates a new temp directory by calling ioutil TempDir
func (_b diskStorage) TempDir(pattern string) (string, error) { return _a.TempDir("", pattern) }

// TempFile creates a new temp file by calling ioutil TempFile
func (_c diskStorage) TempFile(dir, pattern string) (_gag.File, error) {
	return _a.TempFile(dir, pattern)
}