document.addEventListener('DOMContentLoaded', () => {
	const registerForm = document.getElementById('registerForm');
	const responseMessageDiv = document.getElementbyId('responseMessage');

	registerForm.addEventListener('submit', async (event) => {
		event.preventDefault();

		const formData = new FormData(registerForm);
		const data = new URLSearchParams(formData);

		try {
		  const response = await fetch('register', {
		  	method: 'POST',
			body: data,
			});

		 if(!response.ok) {
		 	const errorText = await response.text();
			responseMessageDiv.textContent = 'Error: ${errorText}';
			responseMessageDiv.classList.add('text-red-500');
			responseMessageDiv.classList.remove('text-green-500');
			} else {
			const successText = await response.text();
			responseMessageDiv.textContent = successText;
			responseMessageDiv.classList.add('text-green-500');
			responseMessageDiv.classList.remove('text-red-500');
			}
			} catch (error) {
			responseMessageDiv.textContent = 'Error: ${error.message}';
			responseMessageDiv.classList.add('text-red-500');
			responseMessageDiv.classList.remove('text-green-500');
			}
		});
	});
