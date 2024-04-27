package main

import (
	"fmt"
	"github.com/google/generative-ai-go/genai"
)
import "google.golang.org/api/option"
import (
	"context"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx,
		option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-pro")
	question1 := "너의 이름은?"
	//1. Simple API Call
	resp, err := model.GenerateContent(ctx, genai.Text(question1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Q1:")
	fmt.Println(question1)
	PrintModelResp(resp)
	//2. 나만의 시 만들기
	fmt.Println("Q2:")
	question2 := "Gemini의 대단함을 칭송하는 3줄의 짧은 시를 작성해줘"
	resp, err = model.GenerateContent(ctx,
		genai.Text(question2))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(question2)
	PrintModelResp(resp)
	//3. 내가 가야할 여행지는?
	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("안녕 나는 작년에 일본, 미국, 이탈리아 여행을 다녀왔어. 이탈리아가 가장 마음에 들었어."),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("만나서 반갑습니다, 저는 여행 전문 가이드입니다."),
			},
			Role: "model",
		},
	}
	question3 := "내가 가보지 않고 좋아할만한 여행지를 추천해줘"
	resp, err = cs.SendMessage(ctx, genai.Text(question3))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Q3:")
	fmt.Println(question3)
	PrintModelResp(resp)
}

func PrintModelResp(resp *genai.GenerateContentResponse) {
	fmt.Println("A:")
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("==========================================================")
}
