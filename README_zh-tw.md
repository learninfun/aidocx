[![github
release](https://img.shields.io/github/release/learninfun/aidocx.svg?label=current+release)](https://github.com/learninfun/aidocx/releases)

# aidocx: 知識擷取小幫手

![alt aidocx flow](https://github.com/learninfun/aidocx/blob/main/static/images/aidocx_flow_zh-tw.png?raw=true)

## 安裝
從[Github](https://github.com/learninfun/aidocx/releases/)下載程式

## 使用前提
經由使用各家廠商所提供的API介面呼叫AI，因此在使用前需要先申請相對應的API Key。
以OpenAI的ChatGPT為例，可使用此網址進行申請: [申請網址](https://openai.com/blog/openai-api)

## 基本用法: epub
這個指令會將當前資料夾中的 input.md 和 config.yaml 檔案轉換成一個名為 output.epub 的檔案
```bash
aidocx -t epub ^
       -o output.epub ^
       -apikey-openai "paste-your-api-key" ^
       input.md
```

## input.md: 準備學習的樹狀知識點
```markdown
- 機器學習
  - 監督式學習
    - 迴歸分析
    - 線性迴歸
    - 多項式迴歸
    - Ridge 迴歸
    - Lasso 迴歸
  - 非監督式學習
    - 分群
```

## config.yaml: 每個知識點想要問的問題
```yaml
apiProvider: OpenAI
apiModal: gpt-3.5-turbo-0301
initRole: 假設你是機器學習專家，回答我問題
questions:
  - key: preview
    desc: 習題預習
    template: 給我5題{{ .keyword}}的中文問題
  - key: explain
    desc: 說明知識
    template: 以中文說明{{ .keyword}}並舉例
  - key: keypoint
    desc: 條列重點
    template: 以中文條列{{ .keyword}}的重點
  - key: test
    desc: 知識測驗
    template: 以中文給我5題{{ .keyword}}的中等難度問題，並在後面列出答案
```

## config.yaml 设定参数
- **apiProvider**: API Provicer, ex: OpenAI
- **apiModal**: 選擇API的模型, ex: gpt-3.5-turbo, gpt-4-32k-0314
- **questions**: 問題清單
  - **key**: 快取答案時使用的key
  - **desc**: 輸出結果時，描述問題的類型
  - **template**: 問題的模板
