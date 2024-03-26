package lsp

type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageID string `json:"languageId"`
	Text       string `json:"text"`
	Version    int    `json:"version"`
}
