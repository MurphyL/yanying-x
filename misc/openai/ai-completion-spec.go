package openai

// CompletionResp - 对话响应
type CompletionResp interface {
	Consume(cb func(ok bool, reasoning string, content string))
}

// CompletionByAI - OpenAI 客户端 - 接口定义
type CompletionByAI interface {
	ApplyCompletion(text string, enableWebSearch bool) CompletionResp
}
