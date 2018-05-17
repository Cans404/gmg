function loadXMLDoc() {
	var xmlhttp = new XMLHttpRequest();
	var fd = new FormData();
	var args = document.getElementById("body").value;

	var msg = "the command to be executed:\n" + args
	if (confirm(msg)==false) {
		return;
	}

	fd.append("body", args);

	xmlhttp.onreadystatechange = function() {
		if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
			document.getElementById("output").innerHTML = xmlhttp.responseText;
		}
	}
	
	xmlhttp.open("POST", "/exec/", true);
	xmlhttp.send(fd);
}
