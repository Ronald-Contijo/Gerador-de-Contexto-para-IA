package main

func defaultLangByExt() map[string]string {
	return map[string]string{
		".go": "go", ".html": "html", ".htm": "html", ".css": "css",
		".js": "javascript", ".mjs": "javascript", ".cjs": "javascript",
		".ts": "typescript", ".tsx": "tsx", ".jsx": "jsx",
		".py": "python", ".rb": "ruby", ".php": "php",
		".java": "java", ".c": "c", ".h": "c",
		".hpp": "cpp", ".cpp": "cpp", ".cc": "cpp",
		".cs": "csharp", ".rs": "rust", ".kt": "kotlin",
		".swift": "swift", ".sh": "bash", ".bat": "bat", ".ps1": "powershell",
		".sql": "sql", ".json": "json", ".yaml": "yaml", ".yml": "yaml",
		".toml": "toml", ".ini": "ini", ".md": "md", ".tex": "latex",
		".r": "r", ".m": "matlab",
	}
}

func defaultSkipDirs() map[string]bool {
	return map[string]bool{
		".git": true, "node_modules": true, "dist": true, "build": true,
		"out": true, "venv": true, ".venv": true, ".idea": true, ".vscode": true,"":true,
	}
}
