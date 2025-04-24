package summary

import "os"

func WritePromptToFile(prompt string, filepath string) error {
	return os.WriteFile(filepath, []byte(prompt), 0644)
}
