package main

var GpxTemplate = `<?xml version="1.0" standalone="yes"?>
{{- $name := .Name }}
{{- $start_time := .StartTime }}
{{- $points := .TrackPoints }}
<gpx xmlns="http://www.topografix.com/GPX/1/1"
  creator="mapsme2gpx"
  version="1.1"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd">
  <metadata>
    <name>{{$name}}</name>
    <time>{{$start_time}}</time>
  </metadata>
  <trk>
    <name>{{$name}}</name>
    <trkseg>
    {{- range $point := $points}}
      <trkpt lat="{{$point.Lat}}" lon="{{$point.Lon}}">
        <ele>{{$point.Elevation}}</ele>
        <time>{{$point.IsoTimeStr}}</time>
      </trkpt>
    {{- end}}
    </trkseg>
  </trk>
</gpx>`
