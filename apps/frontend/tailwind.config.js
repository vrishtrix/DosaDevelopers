/** @type {import('tailwindcss').Config} */
export default {
	content: ['./index.html', './src/**/*.{svelte,js,ts,jsx,tsx}'],
	theme: {
		fontFamily: {
			sans: ['Geist', 'ui-sans-serif', 'system-ui'],
			serif: ['ui-serif', 'Georgia'],
			mono: ['ui-monospace', 'SFMono-Regular'],
		},
		extend: {
			colors: {
				spaceGrey: '#121212',
				partyPurple: '#785BEB',
				brightGrey: '#EBEBF54D',
				darkWhite: '#FFFFFF99',
			},
			fontFamily: {
				ownersTrialWide: [
					'Owners Trial Wide',
					'ui-sans-serif',
					'system-ui',
				],
			},
		},
	},
	plugins: [],
};
