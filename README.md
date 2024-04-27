_## 1. Download Golang above 1.20 version
- https://go.dev/doc/install

## 2. Clone this reoisitory
```console
git clone git@github.com:golangkorea/build-with-ai-go-handson.git && cd build-with-ai-go-handson
```

## 3. Init Project
```console
go mod init build-with-ai-go-handson
```


## 4. Install Go SDK for Google Generative AI
```console
go get github.com/google/generative-ai-go
```


## 5. Make your GEMINI API KEY 
- https://makersuite.google.com/app/apikey

### Set ENV from MAC & Linux
```console
export GEMINI_API_KEY=<YOUR_API_KEY>
```
### Set ENV from Windows
```console
setx GEMINI_API_KEY=<YOUR_API_KEY>
```

## 6. Run Go file
```console
go run main.go
```

## 7. Check Results

```
Q1:
너의 이름은?
A:
저는 Gemini입니다. Google에서 개발한 대규모 언어 모델입니다.
==========================================================
Q2:
Gemini의 대단함을 칭송하는 3줄의 짧은 시를 작성해줘
A:
쌍둥이별처럼 영리하고 빠름,
Gemini는 지적 탐구의 표본.
언어와 사고의 달인이라네.
==========================================================
Q3:
내가 가보지 않고 좋아할만한 여행지를 추천해줘
A:
좋아하는 여행지가 이탈리아라고 하셨군요. 그러면 다음과 같은 여행지를 추천해 드릴 수 있습니다.

* **크로아티아:** 아름다운 해안선, 역사적인 도시, 맛있는 요리로 유명합니다.
* **그리스:** 고대 유적, 푸른 바다, 그림 같은 섬으로 유명합니다.
* **포르투갈:** 매력적인 도시, 아름다운 해변, 맛있는 해산물 요리로 유명합니다.
* **스페인:** 활기찬 문화, 맛있는 타파스, 아름다운 건축물로 유명합니다.
* **프랑스:** 낭만적인 분위기, 맛있는 요리, 세계적으로 유명한 와인으로 유명합니다.
==========================================================
```
