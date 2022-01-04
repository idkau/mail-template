<html lang="en">
<head>
    <title>Migration email output</title>
</head>
<body>
<h1>Migration email processor output</h1>
<p><b>Subject:</b></p>
<p>{{.MgNumber}} - {{.VpsID}} - VPS Migration Plan</p>
<p><b>Email contents:</b></p>
<br>
<p>We are contacting you to inform you of a migration of one of your services with TPP.
    Specifically, VPS {{.VpsID}}. We are moving our clients to the new infrastructure built on AWS.
        This will provide you with the latest hardware along with SSD storage.
        We would like to plan a migration of this VPS with you.
        If the VPS is using external DNS, the IP would need to be updated in the zone files as your IP will change.
        We will answer any questions you may have while planning this migration with us.
</p>
</body>
</html>