package main

var (
	mainPage = `
<html><head><title>Yes Coffee is Brewing!!</title></head>
<body bgcolor={{.Color}}>
<br><br><br><br><br><br><br><br><br>
<center>
  <script type="text/javascript">
	<!--
  google_ad_client = "ca-pub-4109530843345987";
  /* top-coffee-machine */
  google_ad_slot = "2520183523";
  google_ad_width = 728;
  google_ad_height = 90;
  //-->
  </script>
	<script type="text/javascript" src="{{.Schema}}://pagead2.googlesyndication.com/pagead/show_ads.js">
  </script>
	</center></br><center><img src="/img/{{ .Image }}">
  <h1>{{ .Exclamation }}! The Coffee is {{ .Broken }}!</h1></center><center>
  <script type="text/javascript"><!--
  google_ad_client = "ca-pub-4109530843345987";
  /* bottom-coffee */
  google_ad_slot = "2064675845";
  google_ad_width = 728;
  google_ad_height = 90;
  //-->
  </script>
  <script type="text/javascript" src="{{.Schema2}}://pagead2.googlesyndication.com/pagead/show_ads.js">
  </script></center>
  <!--# Brought to you by Fiasco Enterprises, Inc. -->
</body></html>`
)
