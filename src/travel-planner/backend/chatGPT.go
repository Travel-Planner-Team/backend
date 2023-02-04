package backend

import (
	"context"
	"fmt"
	"strings"

	"travel-planner/model"
	"travel-planner/util"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func SearchSitesInChatGPT(query string) ([]model.Site, error) {
	config, _ := util.LoadApplicationConfig("conf", "chatGPT.yml")
	c := gogpt.NewClient(config.ChatGPTConfig.Key)
	ctx := context.Background()
	// get the config from config file, print client config
	fmt.Println(query)

	req := gogpt.CompletionRequest{
		Model:       "text-davinci-003",
		MaxTokens:   200,
		Prompt:      query,
		Temperature: 0,
	}

	resp, err := c.CreateCompletion(ctx, req)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	reply := resp.Choices[0].Text
	var sites []model.Site

	// parse rsp_text return a list of string
	rsp_list := strings.Split(reply, "\n")
	// print each item in rsp_list
	for _, item := range rsp_list {
		// print item if not empty
		if item != "" {
			i := strings.Index(item, ".")
			item_clean := item[i+2:]
			var site model.Site
			site.SiteName = item_clean
			sites = append(sites, site)
		}
	}

	return sites, nil
}

func ReadSitesFromChatGPT(resp gogpt.CompletionResponse) ([]model.Site, error) {
	fmt.Printf("ReadSitesFromCHatGPT%v\n", resp)
	var sites []model.Site

	choices := resp.Choices
	for _, item := range choices {
		var site model.Site
		site.SiteName = item.Text
		sites = append(sites, site)
	}

	return sites, nil
}
