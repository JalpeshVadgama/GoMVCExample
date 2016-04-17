package ViewModels

type Pricing struct{
	Title string
	Active string
	HostingCategories []HostingCategory
}

type HostingCategory struct{
	Name string
	Description string
	Price float64
}

func GetPricing(){
	result:= Pricing{
		Title:"Pricing",
		Active:"Pricing",
	}

	var sharedHostingCategory:=HostingCategory{
		Name:"Shared Hosting",
		Description: "Description for shared hosting",
		Price:199.00,
	}

	var vpsHostingCategory:=HostingCategory{
		Name : "VPS hosting",
		Description:"Description for VPS hosting",
		Price:299.00,
	}

	var dedicatedHostingCategory:=HostingCategory{
		Name: "Dedicated Hosting",
		Description:"Descriptiong for dedicated hosting",
		Price: 599.00,
	}

	var windowHostingCategory:=HostingCategory{
		Name:"Windows Hosting",
		Description:"Descriptiong for windows hosting",
		Price:699.00,
	}
	result.HostingCategories = []HostingCategory{vpsHostingCategory,sharedHostingCategory,dedicatedHostingCategory,windowHostingCategory}
	return result
}
