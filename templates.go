package ltropnr

const (
	defaultTemplate = `
	<html>
		<head>
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
			{{ if (ne (.GetHeader "Subject") "") }}
				<title>{{ .GetHeader "Subject" }}</title>
			{{ end }}
			<style type="text/css">
			  #container {
				margin: 10px auto;
			  }
			  #message_headers {
				background: #fff;
				font-size: 12px;
				font-family: "Segoe UI", "Helvetica Neue", Arial, sans-serif;
				border-bottom: 1px solid #dedede;
				margin-bottom: 10px;
				overflow: auto;
			  }
			  #message_headers dl {
				float: left;
				line-height: 1.3em;
				padding: 0;
			  }
			  #message_headers dt {
				width: 92px;
				margin: 0;
				float: left;
				text-align: right;
				font-weight: bold;
				color: #7f7f7f;
			  }
			  #message_headers dd {
				margin: 0 0 0 102px;
			  }
			  #message_headers p.alternate {
				float: right;
				margin: 0;
			  }
			  #message_headers p.alternate a {
				color: #09c;
			  }
			  pre#message_body {
				padding: 4px;
				white-space: pre-wrap;
				border: 1px solid #eee;
				background-color: #fcfcfc;
			  }
			  iframe {
				border: 0;
				width: 100%;
				height: 100%;
			  }
			</style>
		</head>
		<body>
			<div id="container">
				<div id="message_headers">
					<dl>
					{{ if (ne (.GetHeader "From") "") }}
						<dt>From:</dt>
						<dd>{{ .GetHeader "From" }}</dd>
					{{ end }}
					{{ if (ne (.GetHeader "Reply To") "") }}
					  <dt>Reply-To:</dt>
					  <dd>{{ .GetHeader "Reply To" }}</dd>
					{{ end }}
					{{ if (ne (.GetHeader "Subject") "") }}
					  <dt>Subject:</dt>
					  <dd><strong>{{ .GetHeader "Subject" }}</strong></dd>
					{{ end }}
					{{ if (ne (.GetHeader "To") "") }}
					  <dt>To:</dt>
					  <dd>{{ .GetHeader "To" }}</dd>
					{{ end }}
					{{ if (ne (.GetHeader "Date") "") }}
					  <dt>To:</dt>
					  <dd>{{ .GetHeader "Date" }}</dd>
					{{ end }}
					{{ if (ne (.GetHeader "Cc") "") }}
					  <dt>CC:</dt>
					  <dd>{{ .GetHeader "Cc" }}</dd>
					{{ end }}
					{{ if (ne (.GetHeader "Bcc") "") }}
					  <dt>BCC:</dt>
					  <dd>{{ .GetHeader "Bcc" }}</dd>
					{{ end }}
				</div>
				{{ .HTML .Body }}
			</div>
		</body>
	</html>`

	lightTemplate = `
	<html>
		<head>
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
			{{ if (ne (.GetHeader "Subject") "") }}
				<title>{{ .GetHeader "Subject" }}</title>
			{{ end }}
		</head>
		<body>
			{{ .HTML .Body }}
		</body>
	</html>`
)
