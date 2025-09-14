package main

import "fmt"

/* Делаем апп, который выдает меню:
- 1. Посмотреть закладки
- 2. Добавить закладку
- 3. Удалить закладку
- 4. Выход
При 1- Выводит закладки
При 2 - 2 поля ввода названия и адреса и после добавление
При 3 - Ввод названия и удаление по нему
При 4 - Завершение
*/

type bookmarkMap = map[string]string

func main() {
	m := make(bookmarkMap, 3)
	m["A"] = "1"
	m["B"] = "2"
	m["C"] = "3"
	fmt.Println(len(m))

	bookmarks := make(bookmarkMap)
	fmt.Println("Приложение для закладок")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			bookmarks = addBookmark(bookmarks)
		case 3:
			bookmarks = deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить заметку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
		return
	}

	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	}
}

func addBookmark(bookmarks map[string]string) map[string]string {
	var newBookmarkKey string
	var newBookmarkValue string

	fmt.Print("Введите название: ")
	fmt.Scan(&newBookmarkKey)

	fmt.Print("Введите ссылку: ")
	fmt.Scan(&newBookmarkValue)

	bookmarks[newBookmarkKey] = newBookmarkValue
	return bookmarks
}

func deleteBookmark(bookmarks map[string]string) map[string]string {
	var bookmarkKeyToDelete string
	fmt.Println("Введите название: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
	return bookmarks
}
