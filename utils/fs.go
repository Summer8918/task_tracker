package utils

import (
	"os"
	"fmt"
	"path/filepath"
	"encoding/json"
)

func TasksFilePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: fail to get current work directory:", err)
		return "", fmt.Errorf("getwd: %w", err)
	}
	return filepath.Join(cwd, "tasks.json"), nil
}

func ReadTasksFromFile() ([]Task, error) {
	filePath, err := TasksFilePath()
	if err != nil {
		return nil, err
	}

	// Check if file exists, if doesn't exist, return an empty list and create the file
	_, err = os.Stat(filePath)

	if os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating file...")
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return nil, err
		}

		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())
		defer file.Close()

		return []Task{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Fail to open file:", err)
		return nil, err
	}

	defer file.Close()

	tasks := []Task{}

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return nil, err
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []Task) error {
	filePath, err := TasksFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding file:", err)
		return err
	}

	return nil
}