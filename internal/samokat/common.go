package samokat

type GetOAuthTokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type Address struct {
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Street      string  `json:"street"`
	PurposeName string  `json:"purposeName"`
	District    string  `json:"district"`
	House       string  `json:"house"`
	City        string  `json:"city"`
	Region      string  `json:"region"`
}

type Showcases []struct {
	ShowcaseID     string `json:"showcaseId"`
	StoreID        string `json:"storeId"`
	Type           string `json:"type"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	SLA            int    `json:"sla"`
	DirectDelivery bool   `json:"directDelivery"`
	LargeImage     struct {
	} `json:"largeImage"`
	SmallImage struct {
	} `json:"smallImage"`
}

type ShowcaseMainRespBody struct {
	HasExpandedSLA bool `json:"hasExpandedSla"`
	Categories     []struct {
		UUID       string `json:"uuid"`
		Name       string `json:"name"`
		Position   int    `json:"position"`
		Featurings []struct {
			UUID        string `json:"uuid"`
			Title       string `json:"title"`
			HideName    bool   `json:"hideName"`
			Type        string `json:"type"`
			DisplayType string `json:"displayType"`
			Image       string `json:"image"`
			PreviewURL  string `json:"previewUrl"`
			VariationID string `json:"variationId,omitempty"`
			Position    int    `json:"position,omitempty"`
			Target      struct {
				Value string `json:"value"`
				Type  string `json:"type"`
			} `json:"target,omitempty"`
			AdvertID string `json:"advertId,omitempty"`
		} `json:"featurings,omitempty"`
		Type        string `json:"type"`
		DisplayType string `json:"displayType"`
		Categories  []struct {
			UUID     string `json:"uuid"`
			Name     string `json:"name"`
			Image    string `json:"image"`
			Position int    `json:"position"`
			Products []struct {
				UUID  string `json:"uuid"`
				Name  string `json:"name"`
				Media []struct {
					Type string `json:"type"`
					URL  string `json:"url"`
				} `json:"media"`
				Prices struct {
					Current  int    `json:"current"`
					Pickup   int    `json:"pickup"`
					MetaInfo string `json:"metaInfo"`
				} `json:"prices"`
				Highlights  []string `json:"highlights,omitempty"`
				Quantity    int      `json:"quantity"`
				IsMatrix    bool     `json:"isMatrix"`
				PromoMarkup struct {
					Badge struct {
						Text  string `json:"text"`
						Color string `json:"color"`
					} `json:"badge"`
				} `json:"promoMarkup,omitempty"`
				Variants struct {
					Items []struct {
						ProductID string `json:"productId"`
						Text      string `json:"text"`
						Quantity  int    `json:"quantity"`
					} `json:"items"`
					AsOneCard bool `json:"asOneCard"`
				} `json:"variants,omitempty"`
				Expires    bool `json:"expires,omitempty"`
				HasCombo   bool `json:"hasCombo,omitempty"`
				Attributes struct {
					Adult string `json:"adult"`
				} `json:"attributes,omitempty"`
			} `json:"products"`
		} `json:"categories,omitempty"`
		HideName bool `json:"hideName,omitempty"`
	} `json:"categories"`
	Featurings []struct {
		UUID        string `json:"uuid"`
		Title       string `json:"title"`
		HideName    bool   `json:"hideName"`
		Type        string `json:"type"`
		DisplayType string `json:"displayType"`
		Position    int    `json:"position"`
		VariationID string `json:"variationId"`
		Categories  []struct {
			UUID        string `json:"uuid"`
			Name        string `json:"name"`
			Position    int    `json:"position"`
			Type        string `json:"type"`
			DisplayType string `json:"displayType"`
			Products    []struct {
				UUID  string `json:"uuid"`
				Name  string `json:"name"`
				Media []struct {
					Type string `json:"type"`
					URL  string `json:"url"`
				} `json:"media"`
				Prices struct {
					Old      int    `json:"old"`
					Current  int    `json:"current"`
					Pickup   int    `json:"pickup"`
					MetaInfo string `json:"metaInfo"`
					PromoID  string `json:"promoId"`
				} `json:"prices"`
				Quantity   int      `json:"quantity"`
				IsMatrix   bool     `json:"isMatrix"`
				Highlights []string `json:"highlights,omitempty"`
				Variants   struct {
					Items []struct {
						ProductID string `json:"productId"`
						Color     string `json:"color"`
						Text      string `json:"text"`
						Quantity  int    `json:"quantity"`
					} `json:"items"`
					AsOneCard bool `json:"asOneCard"`
				} `json:"variants,omitempty"`
				PromoMarkup struct {
					Badge struct {
						Text  string `json:"text"`
						Color string `json:"color"`
					} `json:"badge"`
				} `json:"promoMarkup,omitempty"`
			} `json:"products"`
		} `json:"categories"`
		Image      string `json:"image,omitempty"`
		PreviewURL string `json:"previewUrl,omitempty"`
	} `json:"featurings"`
}

type GetShowcaseCategoryResp struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Position   int    `json:"position"`
	Categories []struct {
		UUID        string `json:"uuid"`
		Name        string `json:"name"`
		Position    int    `json:"position"`
		Type        string `json:"type,omitempty"`
		DisplayType string `json:"displayType"`
		ComboSets   []struct {
			ID             string `json:"id"`
			Title          string `json:"title"`
			Description    string `json:"description"`
			ImageLink      string `json:"imageLink"`
			Discount       int    `json:"discount"`
			PriceFrom      int    `json:"priceFrom"`
			ProductsNumber int    `json:"productsNumber"`
		} `json:"comboSets,omitempty"`
		Products []struct {
			UUID  string `json:"uuid"`
			Name  string `json:"name"`
			Media []struct {
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"media"`
			Prices struct {
				Current  int    `json:"current"`
				Pickup   int    `json:"pickup"`
				MetaInfo string `json:"metaInfo"`
			} `json:"prices"`
			Highlights []string `json:"highlights"`
			Quantity   int      `json:"quantity"`
			IsMatrix   bool     `json:"isMatrix"`
		} `json:"products,omitempty"`
	} `json:"categories"`
}
