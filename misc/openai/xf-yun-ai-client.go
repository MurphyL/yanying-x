package openai

type XfYunAiClient struct {
	options map[string]string
}

type XfYunCompletionResp struct {
}

func (x XfYunCompletionResp) Consume(cb func(ok bool, reasoning string, content string)) {
	panic("implement me")
}

func (c XfYunAiClient) ApplyCompletion(text string, enableWebSearch bool) CompletionResp {
	return XfYunCompletionResp{}
}
