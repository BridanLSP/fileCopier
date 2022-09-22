package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

type FileCopier struct {
	Folders []Folder `json:"Folders"`
	Auto    bool     `json:"auto"`
}

type Folder struct {
	Source      string   `json:"source"`
	Destination []string `json:"destination"`
}

func ParseTask() (*[]Folder, bool) {
	var fc FileCopier
	cont, err := os.ReadFile("fileCopier.json")
	if err != nil {
		log.Fatalln(err, "can't find fileCopier.json")
	}
	err = json.Unmarshal(cont, &fc)
	if err != nil {
		log.Fatalln(err, "fileCopier.json parse error")
	}
	return &fc.Folders, fc.Auto
}

func VerifyTask(fd *Folder) (*Folder, bool) {
	src := fd.Source
	info, err := os.Stat(src)
	if err != nil {
		log.Println(err, "can't find source folder", src)
		return nil, false
	}
	if !info.IsDir() {
		log.Println(src, "is not a folder")
		return nil, false
	}
	for i, d := range fd.Destination {
		info, err := os.Stat(d)
		if err != nil || !info.IsDir() {
			fd.Destination = append(fd.Destination[:i], fd.Destination[i+1:]...)
		}
	}
	return fd, true
}

func ProcessTask(fd *Folder) {
	for _, d := range fd.Destination {
		_, err := os.Stat(d)
		if os.IsNotExist(err) {
			os.Mkdir(d, os.ModePerm)
		}
	}
	files, err := os.ReadDir(fd.Source)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			subfd := &Folder{Source: fd.Source + "\\" + f.Name()}
			for _, d := range fd.Destination {
				subfd.Destination = append(subfd.Destination, d+"\\"+f.Name())
			}
			ProcessTask(subfd)
		} else {
			src := fd.Source + "\\" + f.Name()
			f0, _ := os.Stat(src)
			for _, d := range fd.Destination {
				dst := d + "\\" + f.Name()

				f, err := os.Stat(dst)
				if os.IsExist(err) && !f.ModTime().After(f0.ModTime()) {
					continue
				} else {
					s, err := os.Open(src)
					if err != nil {
						log.Println(err)
					}
					d, err := os.Create(dst)
					if err != nil {
						log.Println(err)
					}
					io.Copy(d, s)
				}
			}
		}
	}
}

func ProcessTrash(fd *Folder) {
	var (
		MoveToTrash func(src, dst string)
		trashBin    string
	)

	MoveToTrash = func(src, dst string) {
		files, err := os.ReadDir(dst)
		if err != nil {
			log.Println(err)
		}
		for _, f := range files {
			s := src + "\\" + f.Name()
			d := dst + "\\" + f.Name()
			if f.IsDir() {
				if f.Name() == "fileCopier_TrashBin" {
					continue
				}
				MoveToTrash(s, d)
			} else {
				_, err := os.Stat(s)
				if os.IsNotExist(err) {
					err = os.Rename(d, trashBin+"\\"+f.Name())
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
	for _, d := range fd.Destination {
		trashBin = d + "\\" + "fileCopier_TrashBin"
		_, err := os.Stat(trashBin)
		if err != nil {
			err := os.Mkdir(trashBin, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}
		MoveToTrash(fd.Source, d)
	}
}

func WorkFlow(fds *[]Folder) {
	for _, fd := range *fds {
		fd, ok := VerifyTask(&fd)
		if !ok {
			return
		}
		ProcessTask(fd)
		ProcessTrash(fd)
	}
}

func main() {
	fds, auto := ParseTask()
	if auto {
		for {
			WorkFlow(fds)
			time.Sleep(15 * time.Minute)
		}
	} else {
		WorkFlow(fds)
	}
}
