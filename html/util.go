package html

import (
	"regexp"
	"strings"
)

/*
 * 安全字符串
 *  */
func HtmlSafeString(source string) string {
	if source == "" {
		return source
	}

	source = HtmlTagFilter(source)
	source = HtmlEncode(source)

	return source
}

/*
 * 安全参数
 *  */
func HtmlSafeParam(source string) string {
	if source == "" {
		return source
	}

	source = HtmlTagFilter(source)

	return source
}

/*
 * html编码
 *  */
func HtmlEncode(source string, args ...bool) string {
	if source == "" {
		return source
	}

	isSpace := false
	if len(args) > 0 {
		isSpace = args[0]
	}

	content := source

	all := make(map[string]string, 0)
	all[">"] = "&lt;"
	all["<"] = "&gt;"
	all["&"] = "&amp;"
	all["\""] = "&quot;"
	all["'"] = "&#39;"

	if isSpace {
		all[" "] = "&nbsp;"
	}

	for k, v := range all {
		content = strings.Replace(content, k, v, -1)
	}

	return content
}

/*
 * html解码
 *  */
func HtmlDecode(source string, args ...bool) string {
	if source == "" {
		return source
	}

	isSpace := false
	if len(args) > 0 {
		isSpace = args[0]
	}

	content := source

	all := make(map[string]string, 0)
	all[">"] = "&lt;"
	all["<"] = "&gt;"
	all["&"] = "&amp;"
	all["\""] = "&quot;"
	all["'"] = "&#39;"

	if isSpace {
		all[" "] = "&nbsp;"
	}

	for k, v := range all {
		content = strings.Replace(content, v, k, -1)
	}

	return content
}

/*
 * 过滤html标签
 *  */
func HtmlTagFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")

	return re.ReplaceAllString(source, "")
}

/*
 * 过滤html超链接标签
 *  */
func HtmlHyperLinkFilter(source string) string {
	if source == "" {
		return source
	}

	//re, _ := regexp.Compile("\\<a[\\s][\\S\\s]+?/\\>|\\<a[\\S\\s]+?\\</a\\>")
	re, _ := regexp.Compile("\\<a[\\s][\\S\\s]+?/\\>|\\<a[\\s][\\S\\s]+?\\</a\\>")

	return re.ReplaceAllString(source, "")
}

/*
 * 过滤html img标签
 *  */
func HtmlImageFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\<img[\\S\\s]+?/\\>")

	return re.ReplaceAllString(source, "")
}

/*
 * 过滤html audio标签
 *  */
func HtmlAudioFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\<audio[\\S\\s]+?/\\>|\\<audio[\\S\\s]+?\\</audio\\>")

	return re.ReplaceAllString(source, "")
}

/*
 * 过滤html video标签
 *  */
func HtmlVideoFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\<video[\\S\\s]+?/\\>|\\<video[\\S\\s]+?\\</video\\>")

	return re.ReplaceAllString(source, "")
}

/*
 * 过滤html css标签
 *  */
func HtmlCssFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\<style[\\S\\s]+?/\\>|\\<style[\\S\\s]+?\\</style\\>")

	return re.ReplaceAllString(source, "")
}

/*
 * 过滤html script标签
 *  */
func HtmlScriptFilter(source string) string {
	if source == "" {
		return source
	}

	//\\<!--[^>]+\\>
	//\\<script[\\S\\s]+?/\\>

	re, _ := regexp.Compile("\\<!--[^>]+\\>|\\<iframe[\\S\\s]+?/\\>|\\<iframe[\\S\\s]+?\\</iframe\\>|\\<script[\\S\\s]+?/\\>|\\<script[\\S\\s]+?\\</script\\>")

	return re.ReplaceAllString(source, "")
}

/*
 * 过滤markdown img标记
 * ![](http://a.b.com/s/1/000002/wKgAA1oAjnCAOrkgAADQl5vsv_s123.jpg)
 *  */
func MarkdownImageFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("!?(\\[.*\\])?\\(.+\\)")

	return re.ReplaceAllString(source, "")
}
