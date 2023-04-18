function show_hide_password(target){
	var input = document.getElementById('password');
	if (input.getAttribute('type') == 'password') {
		target.classList.add('view');
		input.setAttribute('type', 'text');
	} else {
		target.classList.remove('view');
		input.setAttribute('type', 'password');
	}
	return false;
}
function show_hide_password1(target){
	var input = document.getElementById('password_repeat');
	if (input.getAttribute('type') == 'password') {
		target.classList.add('view');
		input.setAttribute('type', 'text');
	} else {
		target.classList.remove('view');
		input.setAttribute('type', 'password');
	}
	return false;
}

function toggleButton(){
	var login = document.getElementById('login').value;
    var tel = document.getElementById('tel').value;
	var password = document.getElementById('password').value;
    var password_repeat = document.getElementById('password_repeat').value;
	var surname = document.getElementById('surname').value;
    var name = document.getElementById('name').value;
	var midname = document.getElementById('midname').value;
	var date = document.getElementById('date').value;
 
    if (login && tel && password && password_repeat && surname && name && midname && date ) {
        document.getElementById('submitButton').disabled = false;
		document.getElementById('submitButton').style = 'cursor: pointer;'
		document.getElementById('submitButton').style.background = '#73A2FF';
    } else {
        document.getElementById('submitButton').disabled = true;
		document.getElementById('submitButton').style = 'cursor: default;'
		document.getElementById('submitButton').style.background = '#B5CEFF';
	}
}