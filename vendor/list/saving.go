package list

// func parseList(text string) (TodoList, error) {
// }

// func listToText(list TodoList) (string, error) {
// 	out := ""
// 	for i, task := range list.Items {
// 		out += fmt.Sprintf("%s\n", task.Name)
// 	}
// }

// func LoadList(listName string) error {
// 	// Load list from file listName
// 	// Set CurrentList to loaded list

// 	return nil

// }

// func SaveList(listName string) error {
// 	// Save CurrentList to file listName
// 	file, err := os.Create(fmt.Sprintf("%APPDATA%\\domorg\\lists\\%s.txt", listName))
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	txt, err := listToText(CurrentList)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = file.WriteString(txt)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
