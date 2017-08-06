package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/fprimex/zdesk-go"
	"github.com/tidwall/gjson"
)

func main() {
	// must add "os" to imports for this to work
	//client, err := zdesk.NewClient(fmt.Sprintf("%s/token",
	//	os.Getenv("ZDESK_USERNAME")),
	//	os.Getenv("ZDESK_TOKEN"),
	//	os.Getenv("ZDESK_DOMAIN"))

	// The DefaultClient uses all of the above env vars for you.
	client, err := zdesk.DefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	//
	// Create
	//

	// Category
	newcategory := []byte(`{"category": {"position": 0, "locale": "en-us", "name":  "Products", "description": "Product Knowledgebase"}}`)
	resp, err := client.HelpCenterCategoryCreate(&zdesk.HelpCenterCategoryCreateInput{},
		&zdesk.RequestOptions{
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       bytes.NewReader(newcategory),
			BodyLength: int64(len(newcategory))})
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterCategoryCreate: %s", err))
	}
	body, err := zdesk.GetBody(resp)
	if err != nil {
		log.Fatal(err)
	}
	categoryid := gjson.GetBytes(body, "category.id").String()
	fmt.Printf("Category ID: %s\n", categoryid)

	// Section
	newsection := []byte(`{"section": {"position": 0, "locale": "en-us", "name": "Terraform", "description": "Product One Knowledgebase Articles", "manageable_by": "staff"}}`)
	resp, err = client.HelpCenterCategorySectionCreate(&zdesk.HelpCenterCategorySectionCreateInput{ID: categoryid},
		&zdesk.RequestOptions{
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       bytes.NewReader(newsection),
			BodyLength: int64(len(newsection))})
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterCategorySectionCreate: %s", err))
	}
	body, err = zdesk.GetBody(resp)
	if err != nil {
		log.Fatal(err)
	}
	sectionid := gjson.GetBytes(body, "section.id").String()
	fmt.Printf("Section ID: %s\n", sectionid)

	// Article
	newarticle := []byte(`{"article": {"title": "FAQ", "body": "Q. What is the example URL?<br/>A. <a href=\"https://example.com\">https://example.com</a>", "locale": "en-us", "comments_disabled": true, "promoted": false, "position": 0}}`)
	resp, err = client.HelpCenterSectionArticleCreate(&zdesk.HelpCenterSectionArticleCreateInput{ID: sectionid},
		&zdesk.RequestOptions{
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       bytes.NewReader(newarticle),
			BodyLength: int64(len(newarticle))})
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterSectionArticleCreate: %s", err))
	}
	body, err = zdesk.GetBody(resp)
	if err != nil {
		log.Fatal(err)
	}
	articleid := gjson.GetBytes(body, "article.id").String()
	fmt.Printf("Article ID: %s\n", articleid)

	//
	// List
	//

	// Categories
	resp, err = client.HelpCenterCategoriesList(&zdesk.HelpCenterCategoriesListInput{}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterCategory: %s", err))
	}

	// Category Sections
	resp, err = client.HelpCenterCategorySections(&zdesk.HelpCenterCategorySectionsInput{ID: categoryid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterCategorySections: %s", err))
	}

	// List articles in section
	resp, err = client.HelpCenterSectionArticles(&zdesk.HelpCenterSectionArticlesInput{ID: sectionid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterSectionArticles: %s", err))
	}

	// List articles in category
	resp, err = client.HelpCenterCategoryArticles(&zdesk.HelpCenterCategoryArticlesInput{ID: categoryid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterCategoryArticles: %s", err))
	}

	//
	// Retrieve
	//

	// Categories
	resp, err = client.HelpCenterCategoryShow(&zdesk.HelpCenterCategoryShowInput{ID: categoryid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterCategoryShow: %s", err))
	}

	// Section
	resp, err = client.HelpCenterSectionShow(&zdesk.HelpCenterSectionShowInput{ID: sectionid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterCategorySectionShow: %s", err))
	}

	// Section Access Policy
	resp, err = client.HelpCenterSectionAccessPolicy(&zdesk.HelpCenterSectionAccessPolicyInput{SectionID: sectionid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterSectionAccessPolicy: %s", err))
	}

	// Article
	resp, err = client.HelpCenterArticleShow(&zdesk.HelpCenterArticleShowInput{ID: articleid, Locale: "en-us"}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterArticleShow: %s", err))
	}

	//
	// Update
	//

	/*
		// Category

		resp, err = client.HelpCenterCategoryUpdate(&zdesk.HelpCenterCategoryUpdateInput{ID:categoryid}, nil)
		if err != nil {
			log.Fatal(fmt.Sprintf("HelpCenterCategoryUpdate: %s", err))
		}

		// Section
		resp, err = client.HelpCenterSectionUpdate(&zdesk.HelpCenterSectionUpdateInput{ID:sectionid}, nil)
		if err != nil {
			log.Fatal(fmt.Sprintf("HelpCenterCategorySectionUpdate: %s", err))
		}

		// Section Access Policy
		newpolicy := []byte(`{ "access_policy": { "viewable_by": "signed_in_users", "manageable_by": "managers"} }`)
		resp, err = client.HelpCenterSectionAccessPolicyUpdate(&zdesk.HelpCenterSectionAccessPolicyUpdateInput{SectionID: sectionid})
		if err != nil {
			log.Fatal(fmt.Sprintf("HelpCenterSectionAccessPolicyUpdate: %s", err))
		}

		// Article
		resp, err = client.HelpCenterArticleUpdate(&zdesk.HelpCenterArticleUpdateInput{ID:articleid}, nil)
		if err != nil {
			log.Fatal(fmt.Sprintf("HelpCenterArticleUpdate: %s", err))
		}
	*/

	//
	// Delete
	//

	// Article
	resp, err = client.HelpCenterArticleDelete(&zdesk.HelpCenterArticleDeleteInput{ID: articleid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterArticleDelete: %s", err))
	}

	// Section
	resp, err = client.HelpCenterSectionDelete(&zdesk.HelpCenterSectionDeleteInput{ID: sectionid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterCategorySectionDelete: %s", err))
	}

	// Category
	resp, err = client.HelpCenterCategoryDelete(&zdesk.HelpCenterCategoryDeleteInput{ID: categoryid}, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("HelpCenterCategoryDelete: %s", err))
	}
}
