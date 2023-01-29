package backend

import (
	
	"context"
    "travel-planner/util"
	"travel-planner/model"

	"github.com/google/uuid"
	gogpt "github.com/sashabaranov/go-gpt3"
	//https://pkg.go.dev/github.com/aknuds1/go-gpt3
)

func SearchSitesInChatGPT(query string) ([]model.Site, error){

    config, err := util.LoadApplicationConfig("conf", "chatGPT.yml")


	 c := gogpt.NewClient(config.ChatGPTConfig.Key)
	 ctx := context.Background()

	 req := gogpt.CompletionRequest{
		//Model:  "text-davinci-003",
		Model: gogpt.GPT3Ada,
		MaxTokens:200,
		Prompt: query,
	
	 }

	 resp, err := c.CreateCompletion(ctx,req)
	 if err != nil{
		return nil, err
	 }

    return ReadSitesFromChatGPT(resp)
	 
}

func ReadSitesFromChatGPT(resp  gogpt.CompletionResponse)([]model.Site, error){
	
	var sites []model.Site
    
	var choices []gogpt.CompletionChoice
    choices = resp.Choices
	for _, item := range choices {
		var site model.Site 
		site.Site_name = item.Text
		site.Id = uuid.New().ID()
		sites = append(sites,site )
	}
	
	return sites, nil
}

