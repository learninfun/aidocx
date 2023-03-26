[![github
release](https://img.shields.io/github/release/learninfun/apidocx.svg?label=current+release)](https://github.com/learninfun/apidocx/releases)

# apidocx
知识撷取小帮手

# 安装
从[Github](https://github.com/learninfun/apidocx/releases/)下载程序

# 使用前提
经由使用各家厂商所提供的API接口呼叫AI，因此在使用前需要先申请相对应的API Key。
以OpenAI的ChatGPT为例，可使用此网址进行申请: [申请网址](https://openai.com/blog/openai-api)

# Basic usage
这个指令会将当前文件夹中的 input.md 和 config.yaml 档案转换成一个名为 output.epub 的档案
```bash
aidocx -t epub ^
       -o output.epub ^
       -apikey-openai "paste-your-api-key" ^
       input.md
```

# input.md
准备学习的树状知识点
```markdown
- 机器学习
  - 监督式学习
    - 回归分析
    - 线性回归
    - 多项式回归
    - Ridge 回归
    - Lasso 回归
  - 非监督式学习
    - 分群
```

# config.yaml
每个知识点想要问的问题
```yaml
apiProvider: OpenAI
apiModal: gpt-3.5-turbo-0301
initRole: 假设你是机器学习专家，回答我问题
questions:
  - key: preview
    desc: 习题预习
    template: 给我5题{{ .keyword}}的中文问题
  - key: explain
    desc: 说明知识
    template: 以中文说明{{ .keyword}}并举例
  - key: keypoint
    desc: 条列重点
    template: 以中文条列{{ .keyword}}的重点
  - key: test
    desc: 知识测验
    template: 以中文给我5题{{ .keyword}}的中等难度问题，并在后面列出答案
```

## yaml config key
- **apiProvider**: API Provicer, ex: OpenAI
- **apiModal**: 选择API的模型, ex: gpt-3.5-turbo, gpt-4-32k-0314
- **questions**: 问题清单
  - **key**: 快取答案时使用的key
  - **desc**: 输出结果时，描述问题的类型
  - **template**: 问题的模板
