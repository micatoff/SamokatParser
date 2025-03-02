package main

import (
	"SamokatParser/internal/samokat"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	lat := flag.Float64("lat", 0, "Широта")
	lon := flag.Float64("lon", 0, "Долгота")
	targetCategory := flag.String("category", "Комбо-наборы", "Категория дла поиска")

	flag.Parse()

	if *lat == 0 || *lon == 0 {
		fmt.Println("Ошибка: необходимо указать корректные lat и lon")
		flag.Usage()
		os.Exit(1)
	}

	sam := samokat.NewClient()
	OAuthTokenResp, err := sam.GetOauthToken()
	if err != nil {
		log.Fatal("Error while getting OAuth Token:", err.Error())
	}
	sam.SetBearer(OAuthTokenResp.AccessToken)

	showcases, err := sam.GetShowcases(*lat, *lon)
	if err != nil {
		log.Fatal("Error while getting showcases:", err.Error())
	}

	if len(showcases) == 0 {
		log.Fatal("No showcases found for the given location")
	}

	showcase := showcases[0]

	showcaseMain, err := sam.GetShowcaseMain(showcase.ShowcaseID)
	if err != nil {
		log.Fatal("Error while getting showcase main:", err.Error())
	}

	foundCategoryUUID := ""
	for _, category := range showcaseMain.Categories {
		for _, subCategory := range category.Categories {
			if subCategory.Name == *targetCategory {
				foundCategoryUUID = subCategory.UUID
			}
		}
	}

	if foundCategoryUUID == "" {
		log.Fatalf("Category '%s' not found in showcase", *targetCategory)
	}

	getShowcaseCategoryResp, err := sam.GetShowcaseCategoryGoods(showcase.ShowcaseID, foundCategoryUUID)
	if err != nil {
		log.Fatal("Error while getting showcase goods:", err.Error())
	}

	for _, category := range getShowcaseCategoryResp.Categories {
		for _, product := range category.Products {
			fmt.Printf("Название: %s\nКартинки: %s\nЦена: %.1f\nСсылка: %s\n\n",
				product.Name, product.Media, float64(product.Prices.Current)/100.0, "https://samokat.ru/product/"+product.UUID)
		}
	}

}
