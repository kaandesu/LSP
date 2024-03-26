--- add this to autocmds
local client = vim.lsp.start_client({
	name = "kaandesu/LSP",
	cmd = { "/Users/kaanyilmaz/Desktop/Programming/gits/LSP/main" },
	filetypes = { "markdown" },
})
if not client then
	vim.notify("there is something wrong with the client")
end
vim.api.nvim_create_autocmd("FileType", {
	pattern = { "md", "markdown" },
	callback = function()
		if not client then
			vim.notify("there is something wrong with the client")
			return
		end
		vim.notify("custom lsp attached!")
		vim.lsp.buf_attach_client(0, client)
	end,
})
