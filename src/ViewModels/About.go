package ViewModels

type About struct{
	Title string
	Active string
}

func GetAbout() About{
	result:= About{
		Title:"About us frm go",
		Active:"About",
	}
	return result
}
