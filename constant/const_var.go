package constant

const HtmlTemplate = `<html>
<head>
	<title>Office Reservation CSV Manual [Hii! :) Checkmarx Assessment BY Shweta Tanwar (: )]</title>
	<style>
		table { border-collapse: collapse; width: 80%; }
		th, td { border: 1px solid #ccc; padding: 8px; text-align: left; }
		th { background-color: #eee; }
		pre { background: #f4f4f4; padding: 10px; }
	</style>
</head>
<body>
	<h1>Office Reservation CSV Data</h1>
	<table>
		<thead>
			<tr>
				<th>Capacity</th>
				<th>Monthly Rate</th>
				<th>Start Date</th>
				<th>End Date</th>
			</tr>
		</thead>
		<tbody>
			{{range .Rows}}
			<tr>
				<td>{{.Capacity}}</td>
				<td>{{printf "%.2f" .MonthlyRate}}</td>
				<td>{{.StartDate}}</td>
				<td>{{.EndDate}}</td>
			</tr>
			{{end}}
		</tbody>
	</table>

	<h2>How to Use the POST /calculate API</h2>
	<p>Send a POST request to <code>/calculate</code> with JSON body specifying the month in <code>YYYY-MM</code> format. For example:</p>
	<pre>{
  "month": "2023-06"
}</pre>

	<p>The API will return JSON with total revenue and unreserved capacity for the given month.</p>
</body>
</html>`
