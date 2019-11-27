package indexer

const tmpl = `<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">
  <title>{{.Name}} index</title>
  <meta name="description" content="The HTML5 Herald">
  <meta name="author" content="SitePoint">
    <style>
        body { font-size: 80%; font-family: 'Lucida Grande', Verdana, Arial, Sans-Serif; }
        .accordionItem h2 { margin: 0; font-size: 1.1em; padding: 0.4em; color: #fff; background-color: #f52424; border-bottom: 1px solid #66d; }
        .accordionItem h2:hover { cursor: pointer; }
        .accordionItem div { margin: 0; padding: 1em 0.4em; background-color: #eef; border-bottom: 1px solid #66d; }
        .accordionItem.hide h2 { color: #000; background-color: #9a9a9a; }
        .accordionItem.hide div { display: none; }
    </style>
</head>
<body onload="init()">
{{define "directories"}}
  <div class="accordionItem">
    <h2>
      <a href="{{.Path}}">{{.Name}}</a> - Total Size: {{bytesToMegaBytes .TotalSize}}MB~
    </h2>
  <div>
    <ol>
      {{range $key, $value :=  .Files}}
          <li class="file">
              <a href="{{$key}}">{{$value.Name}} - {{bytesToKiloBytes $value.Size}}KB</a>
          </li>
      {{end}}
    </ol>
  </div>
  </div>
  {{if .Directories}}
    <ul>
      {{range $dir := .Directories}}
        <li>
          {{template "directories" $dir}}
        </li>
      {{end}}
    </ul>
  {{end}}
{{end}}

{{template "directories" . }}
<script>
let accordionItems = new Array();
    function init() {
      // Grab the accordion items from the page
      let divs = document.getElementsByTagName( 'div' );
      for ( let i = 0; i < divs.length; i++ ) {
        if ( divs[i].className == 'accordionItem' ) accordionItems.push( divs[i] );
      }
      // Assign onclick events to the accordion item headings
      for ( let i = 0; i < accordionItems.length; i++ ) {
        let h2 = getFirstChildWithTagName( accordionItems[i], 'H2' );
        h2.onclick = toggleItem;
      }
      // Hide all accordion item bodies except the first
      for ( let i = 1; i < accordionItems.length; i++ ) {
        accordionItems[i].className = 'accordionItem hide';
      }
    }
    function toggleItem() {
      let itemClass = this.parentNode.className;
      // Hide all items
      for ( let i = 0; i < accordionItems.length; i++ ) {
        accordionItems[i].className = 'accordionItem hide';
      }
      // Show this item if it was previously hidden
      if ( itemClass == 'accordionItem hide' ) {
        this.parentNode.className = 'accordionItem';
      }
    }
    function getFirstChildWithTagName( element, tagName ) {
      for ( let i = 0; i < element.childNodes.length; i++ ) {
        if ( element.childNodes[i].nodeName == tagName ) return element.childNodes[i];
      }
    }
</script>
</body>
</html>`
