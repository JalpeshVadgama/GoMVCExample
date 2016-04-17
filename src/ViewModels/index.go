package ViewModels

type Index struct{
	Title string
	Active string
}
func GetIndex() Index{
	result:= Index{
		Title:"This is the tile from Go file",
		Active: "Home",
	}
	return result
}