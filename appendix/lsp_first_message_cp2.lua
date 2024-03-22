local M = {}

function M.init_lsp()
	local client = vim.lsp.start_client({
		name = "kaandesu/lsp",
		cmd = { "/path/to/the/lsp/project/main" }, -- go build main.go
	})
	if not client then
		vim.notify("something is wrong with custom lsp client")
	end

	vim.api.nvim_create_autocmd("FileType", {
		pattern = "markdown",
		callback = function()
			vim.lsp.buf_attach_client(0, client)
		end,
	})
end

return M
