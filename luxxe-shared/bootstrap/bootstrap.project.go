package bootstrap

import (
	"fmt"
	"os"
	"path/filepath"
)

type Folder struct {
	Name     string
	Children []Item
}

type File struct {
	Name string
}

type Item interface{}

var baseFolderStructure = []Item{
	Folder{
		Name: "api",
		Children: []Item{
			File{
				Name: "controllers.go",
			},
			File{
				Name: "router.go",
			},
		},
	},
	Folder{
		Name: "dtos",
	},
	Folder{
		Name: "messages",
	},
	Folder{
		Name: "pipes",
	},
}

func _bootstrapProject(projectName string, structure []Item) {
	for _, item := range structure {
		switch v := item.(type) {
		case Folder:
			folderPath := filepath.Join(projectName, v.Name)
			if _, err := os.Stat(folderPath); os.IsNotExist(err) {
				err := os.Mkdir(folderPath, os.ModePerm)
				if err != nil {
					fmt.Printf("error creating root folder %s: %s\n", folderPath, err)
					continue
				}
			}

			if len(v.Children) > 0 {
				_bootstrapProject(folderPath, v.Children)
			}
		case File:
			filePath := filepath.Join(projectName, v.Name)
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Printf("error creating root folder %s: %s\n", filePath, err)
				continue
			}
			file.Close()
		}
	}
}

func BootstrapProject(rootFolderName string) {
	rootFolder := filepath.Join(rootFolderName)
	fmt.Printf("bootstrapping project for %s\n", rootFolderName)

	if _, err := os.Stat(rootFolder); os.IsNotExist(err) {
		err := os.Mkdir(rootFolder, os.ModePerm)
		if err != nil {
			fmt.Printf("error creating root folder %s: %s\n", rootFolder, err)
			return
		}
	}
	_bootstrapProject(rootFolder, baseFolderStructure)
	fmt.Printf("Project %s fully set up.\n", rootFolderName)
}
