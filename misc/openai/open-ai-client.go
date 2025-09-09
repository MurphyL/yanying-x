package openai

import (
	"encoding/json"
	"yanying-x/res"
)

var openAiAgents map[string][]string

func init() {
	bytes, _ := res.GetVendorBytes("vendors/ai_agents.json")
	json.Unmarshal(bytes, &openAiAgents)
}

// ApplyCompletion - 执行 AI 对话
func ApplyCompletion(agent CompletionByAI, text string, enableWebSearch bool) CompletionResp {
	return agent.ApplyCompletion(text, enableWebSearch)
}
