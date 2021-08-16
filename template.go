package main

const templ = `<!DOCTYPE html>
<html lang="en">
	<head>
		<title>Exchanges</title>
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
	</head>
	<body>
		<div class="container">
			<div class="row my-5 justify-content-center">
				<div class="col-12">
					<h1 class="text-center mb-3 text-primary">Cryptocurrencies Exchanges</h1>
				</div>
				<div class="col-12">
					<table class="table table-striped table-bordered">
						<thead>
							<tr>
								<th scope="col">Exchange Id</th>
								<th scope="col">Name</th>
								<th scope="col">Website</th>
							</tr>
						</thead>
						<tbody>
							{{ range . }}
								<tr>
									<th scope="row">{{.ExchangeId}}</th>
									<td>{{.Name}}</td>
									<td>{{.Website}}</td>
								</tr>
							{{ end }}
						</tbody>
					</table>
				</div>
			</div>
		</div>
	</body>
</html>
`