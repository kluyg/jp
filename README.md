Console json prettifier
=================================

While trying different REST APIs I usually use curl.<br/>
Most REST APIs return compact json.<br/>
I just wanted to be able to pipeline curl output to some utility which will return indented json.<br/>
This utility is so simple that it's faster to write one than to google it :)<br/>
So here is my version.

Example:
<pre><code>$ curl "http://finance.yahoo.com/webservice/v1/symbols/GOOG/quote?format=json" | jp
{
  "list": {
    "meta": {
      "count": 1,
      "start": 0,
      "type": "resource-list"
    },
    "resources": [
      {
        "resource": {
          "classname": "Quote",
          "fields": {
            "name": "Google Inc.",
            "price": "733.729980",
            "symbol": "GOOG",
            "ts": "1357753025",
            "type": "equity",
            "volume": "1092418"
          }
        }
      }
    ]
  }
}
</code></pre>

That's how it looks without jp:
<pre><code>$ curl "http://finance.yahoo.com/webservice/v1/symbols/GOOG/quote?format=json"
{
"list" : { 
"meta" : { 
"type" : "resource-list",
"start" : 0,
"count" : 1
},
"resources" : [ 
{
"resource" : { 
"classname" : "Quote",
"fields" : { 
"name" : "Google Inc.",
"price" : "733.700012",
"symbol" : "GOOG",
"ts" : "1357753021",
"type" : "equity",
"volume" : "1091932"
}
}
}

]
}
}
</code></pre>
