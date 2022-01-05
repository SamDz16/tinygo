const go = new Go(); // Defined in wasm_exec.js
const WASM_URL = "main.wasm"
var wasm;

if ('instantiateStreaming' in WebAssembly) {
	WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
		wasm = obj.instance;
		go.run(wasm);
	})
} else {
	fetch(WASM_URL).then(resp =>
		resp.arrayBuffer()
	).then(bytes =>
		WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
			wasm = obj.instance;
			go.run(wasm);
		})
	)
}

// EXEMPLE 1
const number1 = document.querySelector("#number1");
const number2 = document.querySelector("#number2");
const btn = document.querySelector("#btn");
const form1 = document.querySelector("#ex1");
const span = document.querySelector("#result");

form1.addEventListener("submit", (e) => {
	e.preventDefault();

    const result = wasm.exports.add(+number1.value, +number2.value)
    span.textContent = result;

	number1.value = number2.value = ''
})

// EXEMPLE 2
const userInput = document.querySelector("#userInput");
const form2 = document.querySelector("#ex2");
const capitalized = document.querySelector("#capitalized")

form2.addEventListener("submit", (e) => {
	e.preventDefault();

	console.log(userInput.value)
	console.log(stringToIntegerArray(userInput.value))

	const arr = wasm.exports.hello(stringToIntegerArray(userInput.value))
	
	// console.log(arr)
	const str = ''

	for (let i = 0; i < arr.length; i++) {
		str += integerArrayToString()
		
	}

	console.log("le résultat:", wasm.exports.hello([1, 2, 3, 4]))

	// console.log(integerArrayToString(arr))

	userInput.value = ''
})


// ALPHABET
const alphabet = "abcdefghijklmnopqrstuvwxyz,! ";


// ENCODE & DECODE functions
const stringToIntegerArray = (string) => {
	const array = []
	for (let i = 0; i < string.length; i++) {
	  array[i] = alphabet.indexOf(string[i]);
	}

	return array
  };

  const integerArrayToString = (array) => {
	let string = "";
	for (let i = 0; i < array.length; i++) {
	  string += alphabet[array[i]];
	}

	return string;
  };


