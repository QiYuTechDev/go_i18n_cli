
# 安装 gotext
i18n-install-gotext:
	go install golang.org/x/text/cmd/gotext@latest
	# add `~/go/bin` to your PATH env `gotext` command not found

# 更新翻译
i18n-update:
	cd cli && gotext -srclang=en update -out=translations.go -lang=en,zh go_i18n_cli

