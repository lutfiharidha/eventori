package database

import (
	"eventori/app/models"
	"strconv"
)

func ModelCatalogueSeeder() {
	for i := 1; i <= 10; i++ {
		db.Create(&models.ModelCatalogue{
			ModelName:   "Model " + strconv.Itoa(i),
			ImageUrl:    "https://images.unsplash.com/photo-1633157953349-75c66213ca2f",
			Description: "Sebastian is a guru at all things UX design, consistently producing intuitive, modern, and bold designs. He is a natural problem solver in the design space so everyone went to him if they were stuck on a feature or flow. He has a arsenal of design skills including animation and graphic design. Not to mention, Seb was a design team lead, which means he managed a full team of both onshore and offshore designs reviewing and providing feedback on all of their designs... Any business would be lucky to have him.",
		})

	}

}
