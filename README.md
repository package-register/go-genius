## go-genius

## trans

- 接口定义

```go
// Translator 翻译接口
type Translator interface {
	Translate(text string) (string, error)
	TranslateWithResult(text string) (*TranslationResult, error)
	Extract(result string) (*TranslationResult, error)
}
```

### 使用

```go
func Test() {
	trans := translator.New(
		translator.WithAppID("1f8c2x2e6"),
		translator.WithSecret("NTllMWQ4NmExxxGVjZDAzYmQ0MTc1N2M3"),
		translator.WithAPIKey("58d5d1sxx47dxxx8f9be721cc18db599e"),
		translator.WithFromLang("cn"),
		translator.WithToLang("en"),
	)
	fmt.Println("原文: ", item)
	result, _ := trans.TranslateWithResult(item)
	fmt.Println(result.Target)
}
```

- 解析翻译

```go
fmt.Println("原文: ", item)

result, _ := trans.TranslateWithResult(item)

fmt.Println(result.Target)
```

- 不解析翻译

```go
// 翻译文本
result, err := trans.Translate("今天天气怎么样")
if err != nil {
	fmt.Println("翻译失败:", err)
	return
}

data, _ := trans.Extract(result)

fmt.Println(data.Target)

fmt.Println("翻译结果:", result)
```

### Option 创建

```go
// Option 方式创建实例
trans := translator.New(
	translator.WithAppID("1f8c2x2e6"),
	translator.WithSecret("NTllMWQ4NmExxxGVjZDAzYmQ0MTc1N2M3"),
	translator.WithAPIKey("58d5d1sxx47dxxx8f9be721cc18db599e"),
	translator.WithFromLang("cn"),
	translator.WithToLang("en"),
)
```

### Struct 创建

```go
trans := translator.NewWithConfig(translator.Config{
	AppID: "xxx",
	Secret: "xxx",
	APIKey: "xxx",
	FromLang: "xxx",
	ToLang: "xxx",
})
```
