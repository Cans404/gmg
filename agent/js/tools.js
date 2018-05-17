function displaySW(content) {
	var all = document.getElementById('all');
	var divArray = all.getElementsByClassName('bulk');
	for (var i = 0; i < divArray.length; i++) {
		if (divArray[i].id == content) {
				divArray[i].style.display = '';
			} else {
				divArray[i].style.display = 'none';
			}
		}
	}
