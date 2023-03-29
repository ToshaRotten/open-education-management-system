var pass = document.querySelector('#password');
var login = document.querySelector('#login');

if(pass){
	pass.addEventListener('input', changeBackground);
}
if(login){
	login.addEventListener('input', changeBackground);
}


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

function changeBackground() {
  if (pass.value !== '' && login.value !== '') {
	setTimeout(function() {
    document.getElementById('submitButton').disabled = false;
	document.getElementById('submitButton').style = 'cursor: pointer;'
	document.getElementById('submitButton').style.background = '#73A2FF';
	}, 100);
    
	
  } else {
    setTimeout(function() {
    document.getElementById('submitButton').disabled = true;
	document.getElementById('submitButton').style = 'cursor: default;'
	document.getElementById('submitButton').style.background = '#B5CEFF';
    }, 100);
    
  }
  
}
export default show_hide_password;