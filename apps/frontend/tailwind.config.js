/** @type {import('tailwindcss').Config} */
export default {
	content: ['./index.html', './src/**/*.{svelte,js,ts,jsx,tsx}'],
	theme: {
		extend: {
			colors: {
				spaceGrey: '#121212',
				partyPurple: '#785BEB',
				brightGrey: '#EBEBF54D',
				darkWhite: '#FFFFFF99',
			},
		},
	},
	plugins: [],
};
