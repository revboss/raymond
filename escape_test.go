package raymond

import "fmt"

func ExampleEscape() {
	tpl := MustParse("{{link url text}}")

	tpl.RegisterHelper("link", func(url string, text string) SafeString {
		return SafeString("<a href='" + Escape(url) + "'>" + Escape(text) + "</a>")
	})

	ctx := map[string]string{
		"url":  "http://www.revboss.com/",
		"text": "This is a <em>cool</em> website",
	}

	result := tpl.MustExec(ctx)
	fmt.Print(result)
	// Output: <a href='http://www.revboss.com/'>This is a &lt;em&gt;cool&lt;/em&gt; website</a>
}
