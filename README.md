[![github
release](https://img.shields.io/github/release/learninfun/apidocx.svg?label=current+release)](https://github.com/learninfun/apidocx/releases)

# apidocx
AI-Powered Knowledge Books: Generated from a Tree of Topics

# Translation
[繁體中文](https://github.com/learninfun/apidocx/blob/main/README_zh-TW.md)

# Installing
Download the release from [Github](https://github.com/learninfun/apidocx/releases/)

# Prerequisites
Obtain an API key from OpenAI (Bard still does not provide a web API interface.)

# Basic usage
This command will convert the "**input.md**" and "**config.yaml**" files in the current folder into an "output.epub" file.
```bash
apidocx -t epub ^
        -o output.epub ^
        -apikey-openai "paste-your-api-key" ^
        input.md
```

# input.md
Tree list of knowledge points to learn.
```markdown
- Machine Learning
  - Supervised Learning
    - Regression
      - Linear Regression
      - Polynomial Regression
      - Ridge Regression
      - Lasso Regression
  - Unsupervised Learning
    - Clustering
```

# config.yaml
Tree list of knowledge points to learn.
```yaml
apiProvider: OpenAI
apiModal: gpt-3.5-turbo-0301
initRole: Assuming you are an Machine Learning expert, answer my questions.
questions:
  - key: preview
    desc: Preview
    template: give me five question about {{ .keyword}}
  - key: explain
    desc: Explain
    template: Explain {{ .keyword}} and give an example
  - key: keypoint
    desc: Keypoint
    template: List the key points of {{ .keyword}}
  - key: review
    desc: Review
    template: Give me 5 medium-difficulty questions with answers about {{ .keyword}}
```

## yaml config key
- **apiProvider**: API Provicer, ex: ChatGPT only (Bard in the future)
- **apiModal**: Select the modal of API, ex: gpt-3.5-turbo, gpt-4-32k-0314
- **questions**: Question Array
  - **key**: Key for cache answer
  - **desc**: Question description in output
  - **template**: Question template
